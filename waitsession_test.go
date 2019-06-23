package sync

import (
	"fmt"
	"testing"
	"time"
)

func TestWaitSession(t *testing.T) {
	ws := WaitSession{}

	ws.Add(1)

	go func() {
		time.Sleep(time.Second / 1000)
		ws.Done(1, 2)
	}()

	data, err := ws.Wait(1, time.Second*3)
	fmt.Println("---- ", data, err, ws.Len())

	err = ws.Add(1)
	fmt.Println("--- 111:", err)

	err = ws.Add(1)
	fmt.Println("--- 222:", err, ws.Len())

	err = ws.Add(2)
	fmt.Println("--- 333:", err, ws.Len())

	data, err = ws.Wait(1, time.Second/10)
	fmt.Println("--- 444:", data, err, ws.Len())

	data, err = ws.Wait(2, time.Second/10)
	fmt.Println("--- 555:", data, err, ws.Len())
}
