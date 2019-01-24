package defs


//requests
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}

// Data model
type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}

type Comment struct {
	Id string
	VideoId string
	Author string
	Content string
}

type SimpleSession struct {
	Username string //loginname
	TTL int64 //用来检查用户是否登录过期 当用户返回的时间与ttl不匹配的时候= 就返回登录过期
}

type SignedUp struct {
	Success bool
	SessionId string
}