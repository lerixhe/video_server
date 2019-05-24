//限制保护，防止系统资源被恶意消耗完，导致系统崩溃
package main

import "log"

type ConnLimiter struct {
	concurrentConn int
	bucket         chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}

// 链接限制器的方法，判断当前连接数是否已满，未满才可接受token
func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reached the rate limitation")
		return false
	}
	cl.bucket <- 1
	return true
}

// 释放链接.只要能把token拿出来，即可让其他连接进入
func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket
	log.Printf("New connect coming:%d", c)
}
