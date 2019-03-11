package cache

import (
	"encoding/json"
	"fmt"
	"similar_cache/config"
	"similar_cache/models"
	"similar_cache/protos"
	"github.com/golang/glog"
)

func QuerySimiAndDev(packageName, developerId string, skip1, skip2, limit1, limit2 int32, disable1, disable2 string, enable_google_adsense1, enable_google_adsense2 string, ne_package_name1, ne_package_name2 string, notNexting1, notNexting2 bool) (reply *protos.GeneralReply, err error) {
	allresult := models.SimDevResult{}
	reply = &protos.GeneralReply{}
	//从缓存中查询，缓存中无数据再去mongoDB中查询
	cache, err := Getcache(config.CacheConfig.DeveloperPrefix + packageName)
	if err != nil {
		glog.V(0).Infoln("Getcache", err)
	}
	if len(cache) != 0 {
		glog.V(0).Infoln("缓存中读取")

		err := json.Unmarshal(cache, &allresult)
		if err != nil {
			glog.V(0).Infof("json err %v", err)
		}
		allresult.Similar.List = HandSimilar(allresult.Similar.List, skip1, limit1, disable1, enable_google_adsense1, ne_package_name1, notNexting1)
		allresult.Similar.Total = len(allresult.Similar.List)

		//处理developer
		//fmt.Println("===缓存 Developer11111", len(allresult.Developer.List))
		allresult.Developer.List = HandlerPaging(allresult.Developer.List, skip2, limit2, disable2, enable_google_adsense2, ne_package_name2, notNexting2)
		allresult.Developer.Total = len(allresult.Developer.List)
		//fmt.Println("===缓存 Developer2222", len(allresult.Developer.List))
		jsonStr, err := json.Marshal(allresult)
		reply.Reply = string(jsonStr)

		return reply, err
	}

	//无缓存时拼接similar和developer结果并缓存
	simresult, err := QuerySimilar(packageName, skip1, limit1, disable1, enable_google_adsense1, ne_package_name1, notNexting1)
	if err != nil {
		fmt.Printf("simresult err %v", err)
	}
	devresult, err := QueryDeveloper(developerId, skip2, limit2, disable2, enable_google_adsense2, ne_package_name2, notNexting2)
	if err != nil {
		fmt.Printf("devresult err %v", err)
	}
	fmt.Println()
	allresult.Similar = simresult
	allresult.Developer = devresult

	//fmt.Println("allresult===",allresult)
	jsonStr, err := json.Marshal(allresult)

	//更新缓存
	err = HandlerCache(config.CacheConfig.DeveloperPrefix+packageName, jsonStr)
	if err != nil {
		fmt.Println("HandlerCache", err)
	}
	//根据客户端参数修改最终结果
	//fmt.Println("===mongodb 处理前Similar", allresult.Similar.Total)
	allresult.Similar.List = HandSimilar(allresult.Similar.List, skip1, limit1, disable1, enable_google_adsense1, ne_package_name1, notNexting1)
	allresult.Similar.Total = len(allresult.Similar.List)
	//fmt.Println("===mongodb 处理后Similar", allresult.Similar.Total)
	//处理developer
	allresult.Developer.List = HandlerPaging(allresult.Developer.List, skip2, limit2, disable2, enable_google_adsense2, ne_package_name2, notNexting2)
	//allresult.Developer.Total = len(allresult.Developer.List)
	//fmt.Println("=== Similar", allresult.Similar.Total)
	//fmt.Println("=== Developer", allresult.Developer.Total)
	jsonStr, err = json.Marshal(allresult)
	reply.Reply = string(jsonStr)
	return reply, err
}
