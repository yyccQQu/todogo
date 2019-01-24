package api

import (
	"net/http"
	"todogo/pkg/api/session"
	)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"

// 如果没有过期就把 未过期的 用户名保存到 header头（ X-User-Name ）中
func ValidateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid)==0 {
		return false
	}
	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}

	r.Header.Add(HEADER_FIELD_UNAME, uname)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0 {
		//pkg.SendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	return true
}

