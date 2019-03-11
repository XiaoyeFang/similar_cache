package cache

import (
	"encoding/json"
	"fmt"
	"similar_cache/config"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Resp struct {
	StatusCode int
	Result     struct {
		List []string
	}
}

//如果similarRes.Result为空或超时，请求爬虫接口
//调用爬虫接口，第一次无数据调用同步爬虫接口，超时240h调用异步爬虫接口
//只缓存 similar_developer 接口
func SyncHttpRequest(packageName string) (names []string, err error) {
	var response Resp
	client := &http.Client{}
	request, err := MakeCheckUrl(config.CacheConfig.SyncHttpReq, packageName)
	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("client.Do  %v", err)
		return names, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &response)
	if err != nil {
		glog.V(0).Infoln("json.Unmarshal==", err)
	}

	names = response.Result.List
	return names, err
}

func AsynHttpRequest(packageName string) {
	client := &http.Client{}
	request, err := MakeCheckUrl(config.CacheConfig.AsynHttpReq, packageName)
	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("client.Do  %v", err)
	}
	//body, _ := ioutil.ReadAll(resp.Body)
	//glog.V(0).Infoln(string(body))
	if resp.StatusCode != 200 {
		glog.V(0).Infoln(resp.StatusCode)
	}
}

func MakeCheckUrl(link, name string) (*http.Request, error) {

	r, err := url.Parse(link)
	if err != nil {
		glog.V(0).Infoln(err)
	}
	// {type: 'apk_similar',par: {package_name: package_name}}

	url := `{"type": "apk_similar","par": {"package_name": "` + name + `"}}`

	//req, err := http.NewRequest(http.MethodPost, r.String(), strings.NewReader(url))
	req, err := http.NewRequest(http.MethodPost, r.String(), strings.NewReader(url))
	//body, _ := ioutil.ReadAll(req.Body)
	//glog.V(0).Infoln(string(body))
	//defer req.Body.Close()
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", config.HTTP_ACCEPT)
	req.Header.Set("User-Agent", config.HTTP_USER_AGENT)

	return req, err

}
