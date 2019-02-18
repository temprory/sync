package sync

import (
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	SetDebug(true, time.Second*3)

	mtx := Mutex{}
	mtx.Lock()
	go func() {
		mtx.Lock()
	}()
	time.Sleep(time.Second * 4)

	rwmtx := RWMutex{}
	rwmtx.Lock()
	go func() {
		rwmtx.Lock()
	}()

	time.Sleep(time.Second * 4)
}
