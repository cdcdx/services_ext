package mail

import (
	"testing"
)

func TestSendMail(t *testing.T) {
	err := SendMessage(nil, "alarm", "msg", true)
	if err != nil {
		t.Log(err)
	}
}
