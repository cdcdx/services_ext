package telegram

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Webhook struct {
	AccessToken string
	ChatID      string
}

// SendMessage Function to send message
//goland:noinspection GoUnhandledErrorResult
func (t *Webhook) sendMessage(s string) error {
	data := []byte(fmt.Sprintf(`{"chat_id":%s, "text":"%s"}`, t.ChatID, s))
	body := bytes.NewReader(data)
	resp, err := http.Post(t.getURL(), "application/json", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func (t *Webhook) getURL() string {
	url := "https://api.telegram.org/bot" + t.AccessToken + "/sendMessage"
	return url
}

func sendMsgtoTelegram(msg string) error {
	if os.Getenv("TELEGRAM_TOKEN") != "" && os.Getenv("TELEGRAM_CHATID") != "" {
		var tele_token = os.Getenv("TELEGRAM_TOKEN")
		var tele_chatid = os.Getenv("TELEGRAM_CHATID")
		d := Webhook{
			AccessToken: tele_token,
			ChatID:      tele_chatid,
		}
		err := d.sendMessage(msg)
		if err != nil {
			// fmt.Printf("send tg err: %v \n", err)
			return err
		} else {
			fmt.Printf("send tg: %s \n", msg)
		}
	}
	return nil
}

func SendMessage(msg string, is_send bool) error {
	if is_send {
		// Telegram
		return sendMsgtoTelegram(msg)
	}
	return nil
}
