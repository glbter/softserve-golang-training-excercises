//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
// Hint: time.Ticker can be used
// Hint 2: to calculate timediff for Advanced lvl use:
//
//  start := time.Now()
//	// your work
//	t := time.Now()
//	elapsed := t.Sub(start) // 1s or whatever time has passed

package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	timeUsed  int32 // in seconds
}

func (u User) TimeUsed() int32 {
	return u.timeUsed
}

func (u *User) IncrementTime() {
	atomic.AddInt32(&u.timeUsed, 1)
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) (ok bool) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	done := make(chan bool)
	go func() {
		process()
		done <- true
	}()

	for {
		select {
		case <-ticker.C:
			u.IncrementTime()
			if !u.IsPremium && u.TimeUsed() >= 10 {
				fmt.Println(u.TimeUsed())
				return false
			}
		case <-done:
			return true
		}
	}
}

func main() {
	RunMockServer()
}
