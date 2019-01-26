package src

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"todogo/pkg/scheduler/src/dbops"
)

func VidDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	vid := p.ByName("vid-id")

	if len(vid) == 0 {
		sendResponse(w, 400, "videoId should not be empty")
	}

	err := dbops.AddVideoDeletionRecord(vid)
	if err != nil{
		sendResponse(w, 500, "Internal server error")
		return
	}

	sendResponse(w, 200, "")
	return
}
