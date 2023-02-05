package gomutextreentrancycontext

import (
	"sync"
)

type MutexReentrancyContext struct {
	OwnLock     sync.Mutex
	LockerItems []*MutexReentrancyContextLockerItem
}

type MutexReentrancyContextLockerItem struct {
	Subject sync.Locker
	Counter int
}

func (self *MutexReentrancyContext) LockMutex(subject *sync.Mutex) {
	self.OwnLock.Lock()
	defer self.OwnLock.Unlock()

	var item *MutexReentrancyContextLockerItem

	for _, i := range self.LockerItems {
		if i.Subject == subject {
			item = i
			break
		}
	}

	if item == nil {
		item = new(MutexReentrancyContextLockerItem)
		item.Subject = subject
		item.Counter = 0
		self.LockerItems = append(self.LockerItems, item)
	}

	if item.Counter == 0 {
		item.Subject.Lock()
	}

	item.Counter++
}

func (self *MutexReentrancyContext) UnlockMutex(subject *sync.Mutex) {
	self.OwnLock.Lock()
	defer self.OwnLock.Unlock()

	var item *MutexReentrancyContextLockerItem

	for _, i := range self.LockerItems {
		if i.Subject == subject {
			item = i
			break
		}
	}

	if item == nil {
		panic("trying to unlock not Locked item")
	}

	if item.Counter == 0 {
		panic("trying to unlock item, which have counter == 0")
	}

	item.Counter--

	if item.Counter == 0 {
		for i := len(self.LockerItems) - 1; i != -1; i-- {
			if self.LockerItems[i] == item {
				self.LockerItems =
					append(
						self.LockerItems[:i],
						self.LockerItems[i+1:]...,
					)
			}
		}
		item.Subject.Unlock()
	}
}
