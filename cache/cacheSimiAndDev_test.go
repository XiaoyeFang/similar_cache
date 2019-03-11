package cache

import (
	"similar_cache/config"
	"testing"
)

func TestHandlerCache(t *testing.T) {

	//HandlerCache(config.CacheConfig.DeveloperPrefix, "11111111111111")
	for i := 0; i < 30; i++ {
		go func() {
			Getcache(config.CacheConfig.DeveloperPrefix + "zozo.android.crosswords")
		}()

	}

	select {}

	//client, err := models.ConnSSDB()
	//if err != nil {
	//	glog.V(0).Infoln(err)
	//}
	//err = client.Del(config.CacheConfig.DeveloperPrefix + "zozo.android.crosswords")
	//glog.V(0).Infoln("client.Del==", err)

}
