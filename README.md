# services_ext

## drop PageCache service
from https://github.com/etcd-io/etcd/blob/main/tests/functional/agent/utils.go
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

## ntp time service
from https://github.com/etcd-io/dbtester/tree/master/pkg/ntp
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