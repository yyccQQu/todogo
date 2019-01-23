package session

import (
	"sync"
	"todogo/pkg/api/dbops"
	"todogo/pkg/api/defs"
	"todogo/pkg/api/utils"
	"time"
)

// session?
// 一种中间态 当我们在客户端与服务器端交互的时候，用的是http，严格来说是restfulAPI，
// restfulAPI 有一个非常大的特点就是状态不会保持的，也就是stateless ，意思就是状态是不会保持的
// 为了记录用户在服务器端的状态，必须用一个东西来保存它以完成一些需求，这个时候就需要session，
//

// 为什要用session
// 我们在保存状态的时候，我们不存状态，那么一些用户操作的东西，需要记录的东西就会丢掉，那么我们必须要用东西来保存他，否则用户就会从头到尾再操作一遍
// 举一个很简单的例子， 我们如果登录的时候，登录到第二个页面，检查用户是否登录，这个就是他的一个状态，如果我们每个页面都输入一下用户名和密码，这样网页是根本不可用的

// session 和cookie 的区别
// session是一种机制，在服务端为用户保存状态的一种机制，而cookie是在客户端为用户保存状态的一种机制，
// 我们经常会把cookie和session放在一块儿说是因为，我们使用session的时候肯定需要一个sessionID，当用户端为了方便访问这个session的时候，会把sessionID放在cookie里面
// 这就是session 和cookie之间的关系

var sessionMap *sync.Map //1.9+ 实现线程安全的map机制，特别是在并发读上面 leggency都很平稳. 但是在并发写的时候，都会加一个全局锁



func init()  {
	sessionMap = &sync.Map{}
}

func LoadSessionsFromDB()  {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}

	//将数据一条条的写入cache里面
	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k,ss)
		return true
	})

}

func nowInMilli() int64{
	return time.Now().UnixNano()/1000000
}

func GenerateNewSessionId(un string) string {

	id, _ := utils.NewUUID()
	ct := nowInMilli() //毫秒
	ttl := ct + 30*60*1000 // severside session valid time: 30 min  --sessionid 在本地的过期时间

	ss := &defs.SimpleSession{Username:un, TTL:ttl} //将获得的值赋值给新的 session对象
	sessionMap.Store(id, ss) //用 Store设置键的值。
	dbops.InserSession(id, ttl, un) //将值插入sessions表
	return id
}

func deleteExpiredSession(sid string)  {
	sessionMap.Delete(sid) //从cache里面删掉
	dbops.DeleteSession(sid)
}

func IsSessionExpired(sid string) (string, bool) {
	// 如果没有sessionExpired，就返回一个username，第二个返回值就是false
	// 当是sessionExired的时候，username就是一个空值，第二个返回true

	// Load返回键在映射中存储的值，如果没有值，则返回nil ---> ok结果指示是否在映射中找到值。
	ss, ok := sessionMap.Load(sid)

	if ok {
		ct := nowInMilli()
		if ss.(*defs.SimpleSession).TTL < ct { //如果过期了就从数据库中删掉对应的sessionid
			//delete expired session
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	}
	return "", true

}









