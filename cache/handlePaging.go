package cache

import (
	"fmt"
	"similar_cache/models"
	"github.com/golang/glog"
)

func HandlerPaging(list []models.ApkDetail, skip, limit int32, disable string, enable_google_adsense string, ne_package_name string, notNexting bool) []models.ApkDetail {

	//disable 禁用的包名数组
	//enable_google_adsense 是否支持谷歌广告，如果为1必须支持广告，0的话支不支持广告都行
	//ne_package_name 要排除的包名，不能在特殊页显示

	if limit == 0 {
		limit = int32(len(list))
	}
	var newList = make([]models.ApkDetail, 0, 0)
	var total int32 = 0
	var temp_limit int32 = 0

	for i := 0; i < len(list); i++ {

		if len(list[i].Disable) > 0 && JudgeSubset(list[i].Disable, disable) {
			glog.V(0).Infoln("continue  disable")
			continue
		}

		//glog.V(0).Infoln("con=", enable_google_adsense == 1 && !list[i].EnableGoogleAdsense)
		if enable_google_adsense == "1" && !list[i].EnableGoogleAdsense {
			glog.V(0).Infoln("continue adsense")
			continue
		}
		//glog.V(0).Infoln("ne_package_name==",ne_package_name,"list[i].PackageName==",list[i].PackageName)
		//fmt.Println("ne_package_name===", ne_package_name)
		if ne_package_name == list[i].PackageName {
			//fmt.Println(list[i].PackageName)
			glog.V(0).Infoln("ne_package_name continue2", list[i].PackageName)
			continue
		}
		//glog.V(0).Infoln("list[i].Enable===",list[i].Enable,"list[i].Active======",list[i].Active)
		if list[i].Enable && list[i].Active {
			total++
			if skip <= (total-1) && temp_limit < limit {
				temp_limit++
				newList = append(newList, list[i])
				if notNexting && temp_limit == limit {
					break
				}
			}
		}
	}

	return newList

}

func HandSimilar(list []models.ApkDetail, skip, limit int32, disable string, enable_google_adsense string, ne_package_name string, notNexting bool) []models.ApkDetail {
	newList := []models.ApkDetail{}
	for i := 0; i < len(list); i++ {

		if len(list[i].Disable) > 0 && JudgeSubset(list[i].Disable, disable) {
			glog.V(0).Infoln("continue  disable")
			continue
		}

		//glog.V(0).Infoln("con=", enable_google_adsense == 1 && !list[i].EnableGoogleAdsense)
		if enable_google_adsense == "1" && !list[i].EnableGoogleAdsense {
			glog.V(0).Infoln("continue adsense")
			continue
		}
		//glog.V(0).Infoln("ne_package_name==",ne_package_name,"list[i].PackageName==",list[i].PackageName)
		//fmt.Println("ne_package_name===", ne_package_name)
		if ne_package_name == list[i].PackageName {
			//fmt.Println(list[i].PackageName)
			glog.V(0).Infoln("ne_package_name continue2", list[i].PackageName)
			continue
		}
		newList = append(newList, list[i])
	}
	if limit == 0 {
		limit = int32(len(newList))
	}
	if skip > int32(len(newList)) {
		fmt.Println("处理完成 similar长度 ", 0)
		return []models.ApkDetail{}
	}
	if skip <= int32(len(newList)) && skip+limit > int32(len(newList)) {
		fmt.Println("处理完成 similar长度 ", len(newList[skip:]))
		return newList[skip:]
	}
	newList = newList[skip : skip+limit]
	fmt.Println("处理完成 similar长度 ", len(newList))

	return newList
}

//判断list[i].Disable是否包含disable
func JudgeSubset(listDisable []string, disable string) bool {

	for i := 0; i < len(listDisable); i++ {
		if listDisable[i] == disable {
			return true
		}

	}

	return false
}
