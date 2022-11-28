package telegram

import (
	"testing"
)

func TestSendTelegram(t *testing.T) {
	err := SendMessage("alarm - msg", true)
	if err != nil {
		t.Log(err)
	}
}
