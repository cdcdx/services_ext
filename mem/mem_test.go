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
	DropPageCacheService(true, 8)
}
