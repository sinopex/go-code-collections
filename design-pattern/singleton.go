package design_pattern

import (
	"sync"
)

// 通过sync.Once指令比sync.Mutex效率要搞
// var lock *sync.Mutex = &sync.Mutex {}
//
// func GetInstance() *Manager {
// 	if m == nil {
// 		lock.Lock()
// 		defer lock.Unlock()
// 		m = &Manager {}
// 	}
// 	return m
// }

var singleton *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(
		func() {
			singleton = &Singleton{
				Mutex: sync.Mutex{},
			}
		})
	return singleton
}

type Singleton struct {
	sync.Mutex
	number int
}

func (s *Singleton) Set(number int) {
	s.Lock()
	defer s.Unlock()
	s.number = number
}

func (s *Singleton) Get() int {
	return s.number
}
