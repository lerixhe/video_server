package main

import (
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

}
