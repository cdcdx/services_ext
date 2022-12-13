package mem

import (
	"testing"
)

func TestMemSync(t *testing.T) {
	err := DropPageCache()
	if err != nil {
		t.Log(err)
	}
}

func TestMemSyncService(t *testing.T) {
	go DropPageCacheService(true, 8)
}
