package src

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"os"
	"log"
	"time"
	"fmt"
	"io/ioutil"
	"io"
	"html/template"
)

func TestPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	t, _ := template.ParseFiles("./videos/upload.html")

	t.Execute(w, nil)
}

func StreamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	//vid := p.ByName("vid-id")
	//vl := VIDEO_DIR + vid + ".mp4"
	v1 := "/Users/.../gopath1/src/todogo/pkg/streamserver/src/videos/testvideo.mp4"
	fmt.Println(v1,"v1")
	video, err := os.Open(v1)

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

func UploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	r.Body = http.MaxBytesReader(w, r.Body, int64(MAX_UPLOAD_SIZE))
	if err := r.ParseMultipartForm(int64(MAX_UPLOAD_SIZE)); err != nil{
		SendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return
	}

	file, _, err := r.FormFile("") // <form name="file"的值

	if err != nil{
		SendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil{
		log.Printf("Read file error: %v", err)
		SendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	}

	fn := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+fn, data, 0666)
	if err !=nil{
		log.Printf("Write file error: %v", err)
		SendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully")
}

