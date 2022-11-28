package dingtalk

import (
	"testing"
)

func TestSendDingTalk(t *testing.T) {
	err := SendMessage("alarm - msg", true, true)
	if err != nil {
		t.Log(err)
	}
}
