package src

import "log"

type ConnLimiter struct {
	concurrentConn int //连接数
	bucket chan int //channel 个数
}

func NewConnLimiter(cc int) *ConnLimiter{ //自己申明一个结构体，通过方法调用返回一个全新的结构体
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:make(chan int, cc),
	}

}

func (cl *ConnLimiter) GetConn() bool {
	//到达最大连接数为false，否则返回true，并用channel接受
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reached the rate limitation.")
		return false
	}
	cl.bucket <- 1
	return true
}

func (cl *ConnLimiter) ReleaseConn()  {
	c := <- cl.bucket //打印已接受的最新连接数
	log.Printf("New connction coming: %d", c)
}


