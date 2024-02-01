package wxpusher

import (
	"fmt"
	"os"

	"github.com/wxpusher/wxpusher-sdk-go"
	"github.com/wxpusher/wxpusher-sdk-go/model"
	"golang.org/x/xerrors"
)

func sendMsgtoWxPusher(msg string) error {
	if os.Getenv("WXPUSHER_TOKEN") != "" && os.Getenv("WXPUSHER_UID") != "" {
		var wxpusher_token = os.Getenv("WXPUSHER_TOKEN")
		var wxpusher_uid = os.Getenv("WXPUSHER_UID")
		model_msg := model.NewMessage(wxpusher_token).SetContent(msg).AddUId(wxpusher_uid)
		msgArr, err := wxpusher.SendMessage(model_msg)
		fmt.Println(msgArr, err)

		if err != nil {
			// fmt.Printf("send wxpusher err: %v \n", err)
			return err
		} else {
			fmt.Printf("send wxpusher: %s \n", msg)
		}
		return nil
	}
	return xerrors.Errorf("env WXPUSHER_TOKEN is empty")
}

func SendMessage(msg string, is_send bool) error {
	if is_send {
		// WxPusher
		return sendMsgtoWxPusher(msg)
	}
	return nil
}
