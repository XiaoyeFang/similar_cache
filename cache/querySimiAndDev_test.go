package cache

import (
	"fmt"
	"testing"
)

func TestQuerySimiAndDev(t *testing.T) {
	_, err := QuerySimiAndDev("com.whatsapp", "WhatsApp Inc.", 0, 0, 1000000, 0,
		"", "WEBSITE", "1", "1", "", "com.whatsapp", true, true)
	//fmt.Println("reply.Reply====", reply.Reply)
	if err != nil {
		fmt.Errorf("QuerySimiAndDev %v", err)
	}
	//language_recommend !!!
}

func JudegePackage(old, new []string) {
	old = append(old, new...)
	nums := make(map[string]int, 0)

	for k, v := range old {

		if _, ok := nums[v]; ok {
			delete(nums, v)
		} else {
			nums[v] = k
		}

	}

	//glog.V(0).Infoln(nums)
}
