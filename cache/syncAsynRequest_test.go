package cache

import (
	"testing"
)

func TestSyncHttpRequest(t *testing.T) {
	similar, err := SyncHttpRequest("com.appybuilder.santoshsubudhijio.Board_Results")

	if err != nil {
		t.Error(err)
	}
	t.Log("similar=", similar)

	//AsynHttpRequest("com.appybuilder.santoshsubudhijio.Board_Results")
}
