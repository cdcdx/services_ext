package mem

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("mem")

func DropPageCache() error {
	log.Info("start drop cache")
	cmd := exec.Command("sudo", "/bin/sh", "-c", "sync; echo 1 > /proc/sys/vm/drop_caches")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Errorf("Failed to drop caches: %v", err)
		return err
	}
	return nil
}

func DropPageCacheService(switch_timer bool, hours int) {
	DropPageCache()

	if switch_timer {
		if hours < 1 {
			hours = 240
		}
		log.Info("start drop cache service")
		ticker := time.NewTicker(time.Duration(hours) * time.Hour)
		for {
			select {
			case <-ticker.C:
				DropPageCache()
			}
		}
	}
}
