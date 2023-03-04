package singleton

import (
	"fmt"
	"sync"
)

type singleton struct{}

func (s *singleton) DoA() { fmt.Printf("Do A, my address is %p\n", s) }
func (s *singleton) DoB() { fmt.Printf("Do B, my address is %p\n", s) }
func (s *singleton) DoC() { fmt.Printf("Do C, my address is %p\n", s) }
func (s *singleton) DoD() { fmt.Printf("Do D, my address is %p\n", s) }

var (
	once     sync.Once
	instance *singleton
)

func GetSingleton() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}
