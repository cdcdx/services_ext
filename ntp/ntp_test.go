package ntp

import (
	"fmt"
	"testing"
)

func TestNTPSync(t *testing.T) {
	o, err := DefaultSync()
	if err != nil {
		t.Log(err)
	}
	fmt.Println(o)
}
