# services_ext

## ntp time service
# from https://github.com/etcd-io/dbtester/tree/master/pkg/ntp
```golang
import "github.com/cdcdx/services_ext/ntp"
func xxx {
	...
	# DefaultSync
	out, err := ntp.DefaultSync()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)
	...
}
```

## drop PageCache service
# from https://github.com/etcd-io/etcd/blob/main/tests/functional/agent/utils.go
```golang
import "github.com/cdcdx/services_ext/mem"
func xxx {
	...
	# DropPageCache
	err := mem.DropPageCache()
	if err != nil {
		fmt.Println(err)
	}
	...
}
```

## Send message to DingTalk
```golang
import "github.com/cdcdx/services_ext/dingtalk"
func xxx {
	...
	# DingTalk
	os.Setenv("DINGTALK_TOKEN", "0000000000xxxxxxxxxxx")
	os.Setenv("DINGTALK_SECRET", "SECxxxxxxxxxxx")
	
	err := dingtalk.SendMessage("alarm - msg", true, true)
	if err != nil {
		fmt.Println(err)
	}
	...
}
```

## Send message to WxPusher
```golang
import "github.com/cdcdx/services_ext/wxpusher"
func xxx {
	...
	# WxPusher
	os.Setenv("WXPUSHER_TOKEN", "AT_xxxxxxxxxxx")
	os.Setenv("WXPUSHER_UID", "UID_xxxxxxxxxxx")
	
	err := wxpusher.SendMessage("alarm - msg", true)
	if err != nil {
		fmt.Println(err)
	}
	...
}
```

## Send message to Telegram
```golang
import "github.com/cdcdx/services_ext/telegram"
func xxx {
	...
	# Telegram
	os.Setenv("TELEGRAM_TOKEN", "0000000000:xxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	os.Setenv("TELEGRAM_CHATID", "111111111")
	
	err := telegram.SendMessage("alarm - msg", true)
	if err != nil {
		fmt.Println(err)
	}
	...
}
```

## Send message to Mail
```golang
import "github.com/cdcdx/services_ext/mail"
func xxx {
	...
	# Mail
	os.Setenv("SMTP_HOST", "smtp.163.com")
	os.Setenv("SMTP_PORT", "465")
	os.Setenv("SMTP_USERNAME", "username@163.com")
	os.Setenv("SMTP_PASSWORD", "passwd")
	
	sendTo := make([]string, 0)
	sendTo = append(sendTo, "cdcdx888@gmail.com")
	err := mail.SendMessage(sendTo, "alarm", "msg", true)
	if err != nil {
		fmt.Println(err)
	}
	...
}
```
