package structs

import (
	"fmt"
	"sync/atomic"
)

type Lock struct {
	locked int32
}

var (
	ErrAlreadyLocked error = fmt.Errorf("lock is already locked")
	ErrNotLocked     error = fmt.Errorf("lock is already unlocked")
)

func NewLock() *Lock {
	return &Lock{
		locked: 0,
	}
}

func (l *Lock) Lock() error {
	free := atomic.CompareAndSwapInt32(&l.locked, 0, 1)

	if !free {
		return ErrAlreadyLocked
	}

	return nil
}

func (l *Lock) Unlock() error {
	free := atomic.CompareAndSwapInt32(&l.locked, 1, 0)

	if free {
		return ErrNotLocked
	}

	return nil
}

func (l *Lock) IsLocked() bool {
	return atomic.LoadInt32(&l.locked) == 1
}

func (l *Lock) IsUnlocked() bool {
	return atomic.LoadInt32(&l.locked) == 0
}
