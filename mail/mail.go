package mail

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-gomail/gomail"
	logging "github.com/ipfs/go-log/v2"
	"golang.org/x/xerrors"
)

var log = logging.Logger("mail")

func sendMail(to []string, title string, context string) error {
	if os.Getenv("SMTP_HOST") != "" && os.Getenv("SMTP_USERNAME") != "" {
		log.Infof("Mail to num %d", len(to))

		send_to := []string{}
		for _, rec := range to {
			send_to = append(send_to, strings.TrimSpace(rec))
		}

		m := gomail.NewMessage()

		// env
		var smtp_host = os.Getenv("SMTP_HOST")
		var smtp_port, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
		var smtp_username = os.Getenv("SMTP_USERNAME")
		var smtp_password = os.Getenv("SMTP_PASSWORD")
		// 发送邮件服务器/端口/账号/密码
		d := gomail.NewPlainDialer(smtp_host, smtp_port, smtp_username, smtp_password)
		// 发件人
		m.SetAddressHeader("From", smtp_username, "server")

		// 收件人
		m.SetHeader("To", send_to...)
		// 主题
		m.SetHeader("Subject", title)
		// 正文
		m.SetBody("text/html", context)

		// // 抄送
		// m.SetHeader("Cc", send_cc)
		// // 暗送
		// m.SetHeader("Bcc", rec)
		// // 附件
		// m.Attach("filename")

		// 发送
		err := d.DialAndSend(m)
		if err != nil {
			// fmt.Printf("send Mail err: %v \n", err)
			return err
		} else {
			fmt.Printf("send Mail %s, %s - %s, %v \n", send_to, title, context, err)
		}
		return nil
	}
	return xerrors.Errorf("env SMTP_HOST is empty")
}

func SendMessage(to []string, title string, context string, is_send bool) error {
	if is_send {
		if len(to) > 0 {
			// Mail
			return sendMail(to, title, context)
		} else {
			if os.Getenv("SMTP_TO") != "" {
				to_str := os.Getenv("SMTP_TO")
				to_list := strings.Split(to_str, ",")
				// Mail
				return sendMail(to_list, title, context)
			} else {
				log.Infof("env SMTP_TO is empty")
			}
		}
	}
	return nil
}
