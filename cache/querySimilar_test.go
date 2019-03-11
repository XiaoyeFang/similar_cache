package cache

import (
	"fmt"
	"testing"
)

func TestQuerSimilar(t *testing.T) {
	reply, err := QuerySimilar("com.lofter.android", 1, 1, "ddd", "0", "id.co.paytren.user", false)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(reply)

}
