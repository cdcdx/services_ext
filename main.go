package main

import (
	"fmt"
	"os"

	"github.com/cdcdx/services_ext/dingtalk"
	"github.com/cdcdx/services_ext/mem"
	"github.com/cdcdx/services_ext/ntp"
	"github.com/cdcdx/services_ext/telegram"
)

func main() {
	for i, v := range os.Args {
		fmt.Println("args:", i, v)
	}

	fmt.Println("\n sync pagecache")
	// sync pagecache
	err1 := mem.DropPageCache()
	if err1 != nil {
		fmt.Println(err1)
	}

	fmt.Println("\n sync ntptime")
	// sync ntptime
	out2, err2 := ntp.DefaultSync()
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(out2)

	fmt.Println("\n send dingtalk")
	// send dingtalk
	err3 := dingtalk.SendMessage("alarm - msg", true, false)
	if err3 != nil {
		fmt.Println(err1)
	}

	fmt.Println("\n send telegram")
	// send telegram
	err4 := telegram.SendMessage("alarm - msg", true)
	if err4 != nil {
		fmt.Println(err1)
	}
}
