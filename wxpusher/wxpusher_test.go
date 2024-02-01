package wxpusher

import (
	"testing"
)

func TestSendWxPusher(t *testing.T) {
	err := SendMessage("alarm - msg", true)
	if err != nil {
		t.Log(err)
	}
}
