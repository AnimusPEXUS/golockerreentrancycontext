package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/AnimusPEXUS/gomutextreentrancycontext"
)

type C1 struct {
	m sync.Mutex
}

func (self *C1) f1(rc *gomutextreentrancycontext.MutexReentrancyContext) {
	if rc == nil {
		rc = new(gomutextreentrancycontext.MutexReentrancyContext)
	}
	rc.LockMutex(&self.m)
	defer rc.UnlockMutex(&self.m)

	fmt.Println("f1 start")
	defer fmt.Println("f1 exiting")

	go self.f2("started by f1 through go", rc)
	self.f2("started directly by f1", rc)

}

func (self *C1) f2(
	msg string,
	rc *gomutextreentrancycontext.MutexReentrancyContext,
) {
	if rc == nil {
		rc = new(gomutextreentrancycontext.MutexReentrancyContext)
	}

	rc.LockMutex(&self.m)
	defer rc.UnlockMutex(&self.m)

	countdown := 10

	for countdown >= 0 {
		fmt.Println("f2:", msg, countdown)
		time.Sleep(time.Second)
		countdown--
	}
}

func main() {
	c1 := new(C1)
	c1.f1(nil)
}
