package src

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"os"
	"log"
	"time"
)

func StreamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR + vid

	video, err := os.Open(vl)

	if err != nil {
		log.Printf("Error when try to open file: %v", err)
		SendErrorResponse(w, http.StatusInternalServerError, "Internal Error") //写一个报错数据
		return
	}

	w.Header().Set("Content-Type", "video/mp4") //让浏览器默认按照mp4格式来解析视屏
	//Ctrl + j 查看方法详情
	http.ServeContent(w, r, "", time.Now(), video) //服务端返回，保存 视屏内容

	defer video.Close()
}


