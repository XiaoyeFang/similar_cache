package cache

import (
	"fmt"
	"similar_cache/config"
	"similar_cache/models"
	"github.com/golang/glog"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//查询 mongodb(apk_similar,apk) ，查询包名排序，集群配置，合并管理员编辑数据(concat(similar.editor_result, similar.result)
//)
func QuerySimilar(packageName string, skip, limit int32, disable string, enable_google_adsense string, ne_package_name string, notNexting bool) (reply models.ApkResult, err error) {
	similarRes := &models.Similar{}
	db, err := models.CreatDatabase()
	if err != nil {
		panic(err)
	}

	scol := db.C(models.APKSIMILAR)
	err = scol.Find(bson.M{"package_name": packageName}).Select(bson.M{"_id": 0, "result": 1, "editor_result": 1}).One(similarRes)
	if err != nil {
		glog.V(0).Infof("scol.Find %v", err)
		//return reply, err
	}

	//如果similarRes.Result为空或超时，请求爬虫接口
	//调用爬虫接口，第一次无数据调用同步爬虫接口，超时240h调用异步爬虫接口
	sumH := time.Now().Sub(similarRes.UpdateDate)

	if len(similarRes.Result) == 0 {
		fmt.Println("无数据，调用同步爬虫接口")
		glog.V(0).Infoln("无数据，调用同步爬虫接口")
		//接受返回的包名列表，根据列表去查询
		similar, err := SyncHttpRequest(packageName)

		reply, err = CheckSliceDetail(db, similar, skip, limit, disable, enable_google_adsense, ne_package_name, notNexting)
		if err != nil {
			fmt.Errorf("CheckSliceDetail err %v", err)
		}
		return reply, err
	}

	if sumH.Hours() > config.SIMILARTIMEOUT {
		//glog.V(0).Infoln("similar数据超时，调用异步爬虫接口")
		glog.V(0).Infoln("similar数据超时，调用异步爬虫接口")
		AsynHttpRequest(packageName)
		//无需等待结果，再去查询一遍similar表
		err = scol.Find(bson.M{"package_name": packageName}).Select(bson.M{"_id": 0, "result": 1, "editor_result": 1}).One(similarRes)
		if err != nil {
			glog.V(0).Infof("scol.Find %v", err)
			return reply, err
		}

	}

	similar := append(similarRes.EditorResult, similarRes.Result...)
	//glog.V(0).Infoln(similar)
	reply, err = CheckSliceDetail(db, similar, skip, limit, disable, enable_google_adsense, ne_package_name, notNexting)
	if err != nil {
		fmt.Errorf("CheckSliceDetail err %v", err)
	}

	return reply, err
}

func CheckSliceDetail(db *mgo.Database, similar []string, skip, limit int32, disable string, enable_google_adsense string, ne_package_name string, notNexting bool) (models.ApkResult, error) {
	reply := models.ApkResult{}
	detailRes := []models.ApkDetail{}
	//d :=models.ApkDetail{}

	dcol := db.C(models.APKDETAIL)
	err := dcol.Find(bson.M{"package_name": bson.M{"$in": similar}}).All(&detailRes)

	if err != nil {
		glog.V(0).Infof("dcol.Find %v", err)
		return reply, err
	}
	orderRes := make([]models.ApkDetail, 0, 0)
	for _, v := range similar {

		for _, r := range detailRes {
			if r.PackageName == v {

				orderRes = append(orderRes, r)
			}

		}

	}
	//fmt.Println("orderRes===",orderRes)

	reply.List = orderRes
	reply.Total = len(orderRes)

	//fmt.Println(Result)
	//jsonStr, err := json.Marshal(Result)
	//if err != nil {
	//	fmt.Println("Marshal==",err)
	//}
	//reply.Reply = string(jsonStr)
	return reply, err
}
