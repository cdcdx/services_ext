package mem

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func DropPageCache() error {
	cmd := exec.Command("sudo", "/bin/sh", "-c", "sync; echo 1 > /proc/sys/vm/drop_caches")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to drop caches: %v", err)
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
		ticker := time.NewTicker(time.Duration(hours) * time.Hour)
		for {
			select {
			case <-ticker.C:
				DropPageCache()
			}
		}
	}
}
