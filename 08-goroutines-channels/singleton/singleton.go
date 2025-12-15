package singleton

import (
	"sync"
)

import "sync/atomic"

type Singleton struct{}

var (
	instance *Singleton
	doOnce   sync.Once

	//checked
	instance_checked *Singleton
	lock             sync.Mutex

	//checked_atomic
	instance_checked_atomic *Singleton
	lock_checked_atomic     sync.Mutex
	initialized             uint32
)

func GetInstance_serial() *Singleton {
	if instance_checked == nil {
		instance_checked = &Singleton{}
	}
	return instance_checked
}

func GetInstance_DoOnce() *Singleton {
	if instance == nil {
		doOnce.Do(func() {
			instance = &Singleton{}
		})
	}
	return instance
}

func GetInstance_checked() *Singleton {
	if instance_checked == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance_checked == nil {
			instance_checked = &Singleton{}
		}
	}

	return instance_checked
}

func GetInstance_checked_atomic() *Singleton {
	if atomic.LoadUint32(&initialized) == 0 {
		lock_checked_atomic.Lock()
		defer lock_checked_atomic.Unlock()

		if initialized == 0 {
			instance = &Singleton{}
			atomic.StoreUint32(&initialized, 1)
		}
	}
	return instance_checked_atomic
}

func GetInstance_locking_always() *Singleton {
	lock_checked_atomic.Lock()
	defer lock_checked_atomic.Unlock()

	if initialized == 0 {
		instance = &Singleton{}

	}
	return instance_checked_atomic
}
