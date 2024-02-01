package main

import (
	"fmt"
	"os"

	"github.com/cdcdx/services_ext/dingtalk"
	"github.com/cdcdx/services_ext/mail"
	"github.com/cdcdx/services_ext/mem"
	"github.com/cdcdx/services_ext/ntp"
	"github.com/cdcdx/services_ext/telegram"
	"github.com/cdcdx/services_ext/wxpusher"
)

func main() {
	for i, v := range os.Args {
		fmt.Println("args:", i, v)
	}

	// sync ntptime
	fmt.Println("\n sync ntptime")
	out, err := ntp.DefaultSync()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)

	// sync pagecache
	fmt.Println("\n sync pagecache")
	err = mem.DropPageCache()
	if err != nil {
		fmt.Println(err)
	}

	// send dingtalk
	fmt.Println("\n send dingtalk")
	err = dingtalk.SendMessage("alarm - msg", true, false)
	if err != nil {
		fmt.Println(err)
	}

	// send wxpusher
	fmt.Println("\n send wxpusher")
	err = wxpusher.SendMessage("alarm - msg", true)
	if err != nil {
		fmt.Println(err)
	}

	// send telegram
	fmt.Println("\n send telegram")
	err = telegram.SendMessage("alarm - msg", true)
	if err != nil {
		fmt.Println(err)
	}

	// send mail
	fmt.Println("\n send mail")
	sendTo := make([]string, 0)
	sendTo = append(sendTo, "cdcdx888@gmail.com")
	err = mail.SendMessage(sendTo, "alarm", "msg", true)
	if err != nil {
		fmt.Println(err)
	}

}
