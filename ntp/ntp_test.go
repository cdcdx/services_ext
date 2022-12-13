package ntp

import (
	"fmt"
	"testing"
)

func TestNTPSync(t *testing.T) {
	o, err := DefaultSync()
	if err != nil {
		t.Log(err)
	}
	fmt.Println(o)
}

func TestNTPSyncService(t *testing.T) {
	go DefaultSyncService(true, 8)
}
