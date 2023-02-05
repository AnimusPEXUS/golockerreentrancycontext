package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/AnimusPEXUS/gomutextreentrancycontext"
)

type C1 struct {
	m  sync.Mutex
	wg *sync.WaitGroup
}

func (self *C1) f1(rc *gomutextreentrancycontext.MutexReentrancyContext) {
	if rc == nil {
		rc = new(gomutextreentrancycontext.MutexReentrancyContext)
	}
	rc.LockMutex(&self.m)
	defer rc.UnlockMutex(&self.m)

	fmt.Println("f1 start")
	defer fmt.Println("f1 exiting")

	// emulating call from some different place which doesn't
	// have value of reentrancy context (rc value)
	go func() {
		self.f2("started by f1 through go", nil)
		self.wg.Done()
	}()

	// f1 started throug go keyword can't actually continue,
	// because self.m already locked by this time
	time.Sleep(2 * time.Second)

	self.f2("started directly by f1", rc)

	self.wg.Done()
}

func (self *C1) f2(
	msg string,
	rc *gomutextreentrancycontext.MutexReentrancyContext,
) {
	if rc == nil {
		rc = new(gomutextreentrancycontext.MutexReentrancyContext)
	}

	fmt.Println("f2 - before Lock():", fmt.Sprintf("(%s)", msg))
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
	c1.wg = new(sync.WaitGroup)
	c1.wg.Add(2)
	c1.f1(nil)
	c1.wg.Wait()
}
