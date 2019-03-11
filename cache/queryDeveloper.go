package cache

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"similar_cache/config"
	"similar_cache/models"
)

func QueryDeveloper(developerId string, skip, limit int32, disable string, enable_google_adsense string, ne_package_name string, notNexting bool) (reply models.ApkResult, err error) {
	devResult := []models.ApkDetail{}
	reply = models.ApkResult{}

	db, err := config.CreatDatabase()
	if err != nil {
		panic(err)
	}
	dcol := db.C(models.APKDETAIL)
	err = dcol.Find(bson.M{"developer": developerId}).All(&devResult)
	if err != nil {
		fmt.Printf("dcol.Find developer_id err %v", err)
	}
	reply.List = devResult
	reply.Total = len(reply.List)
	//fmt.Println("===", reply.Total)
	return reply, err
}
