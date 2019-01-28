package defs

type SalesCombo struct {
	Id int `json:"id"`
	ComboName string `json:"combo_name"`
	SiteNum int `json:"site_num"`
	PlatformNum int `json:"platform_num"`
	RoadFee int `json:"road_fee"`
	Status int `json:"status"`
	CreateTime int `json:"create_time"`
	DeleteTime int `json:"delete_time"`
}

type Userstable struct {
	Id int `json:"id"`
	LoginName string `json:"loginName"`
	Pwd string `json:"pwd"`
}

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}




// error
type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrResponse struct {
	HttpSC int
	Error Err
}

var (
	ErrorRequestBodyParseFailed = ErrResponse{HttpSC: 400, Error: Err{Error: "Request body is not correct", ErrorCode: "001"}}
	ErrorNotAuthUser = ErrResponse{HttpSC: 401, Error: Err{Error: "User authentication failed.", ErrorCode: "002"}}
	ErrorDBError = ErrResponse{HttpSC: 500, Error: Err{Error: "DB ops failed", ErrorCode: "003"}}
	ErrorInternalFaults = ErrResponse{HttpSC: 500, Error: Err{Error: "Internal service error", ErrorCode: "004"}}
)