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
