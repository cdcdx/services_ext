package mail

import (
	"testing"
)

func TestSendMail(t *testing.T) {
	sendTo := make([]string, 0)
	sendTo = append(sendTo, "cdcdx888@gmail.com")
	err := SendMessage(sendTo, "alarm", "msg", true)
	if err != nil {
		t.Log(err)
	}
}
