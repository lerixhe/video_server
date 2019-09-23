package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	videoID := p.ByName("vid-id")
	videoLink := VIDEO_DIR + videoID
	video, err := os.Open(videoLink)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal err")
		return
	}
	w.Header().Set("Conten-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
	defer video.Close()
}

func uploaderHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	err := r.ParseMultipartForm(MAX_UPLOAD_SIZE)
	if err != nil {
		log.Println("Erroe when try to open file:", err)
		sendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return
	}
	file, _, err := r.FormFile("file") //获取name="file"的form
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal err")
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Read file error:%v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal err")
		return
	}
	fileName := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+fileName, data, 0666)
	if err != nil {
		log.Println("Write file error:", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal err")
		return
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Upload file successful!")
	log.Println("Upload file successful!")
}
