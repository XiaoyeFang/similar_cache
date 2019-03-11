package cache

import (
	"testing"
)

func TestQueryDeveloper(t *testing.T) {
	reply, err := QueryDeveloper("WhatsApp Inc.", 0, 0, "", "0", "com.whatsapp", true)
	if err != nil {
		t.Error(err)
	}
	t.Log(reply)
}
