package dingtalk

import (
	"fmt"
	"os"

	"github.com/wanghuiyt/ding"
)

func sendMsgtoDingTalk(msg string, is_all bool) error {
	if os.Getenv("DINGTALK_TOKEN") != "" && os.Getenv("DINGTALK_SECRET") != "" {
		var ding_token = os.Getenv("DINGTALK_TOKEN")
		var ding_secret = os.Getenv("DINGTALK_SECRET")
		d := ding.Webhook{
			AccessToken: ding_token,
			Secret:      ding_secret,
		}
		d.EnableAt = is_all
		d.AtAll = is_all
		err := d.SendMessage(msg)
		if err != nil {
			// fmt.Printf("send dt err: %v \n", err)
			return err
		} else {
			fmt.Printf("send dt: %s \n", msg)
		}
	}
	return nil
}

func SendMessage(msg string, is_send bool, is_all bool) error {
	if is_send {
		// DingTalk
		return sendMsgtoDingTalk(msg, is_all)
	}
	return nil
}
