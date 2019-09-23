package taskrunner

import (
	"errors"
	"log"
	"os"
	"sync"
	"video_server/scheduler/dbops"
)

func deleteVideo(vid string) error {
	err := os.Remove(VIDEO_PATH + vid)
	if err != nil && !os.IsNotExist(err) {
		log.Println("deleting video error:", err)
		return err
	}
	return nil
}
func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Println("cideo clear dispatcher error:", err)
		return err
	}
	if len(res) == 0 {
		return errors.New("all task finished")
	}
	for _, id := range res {
		dc <- id
	}
	return nil
}
func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error
forloop:
	for {
		select {
		case vid := <-dc:
			go func(id interface{}) {
				err := deleteVideo(id.(string))
				if err != nil {
					errMap.Store(id, err)
					return
				}
				err = dbops.DelVideoDeletionRecord(id.(string))
				if err != nil {
					errMap.Store(id, err)
					return
				}
			}(vid)
		default:
			break forloop
		}
	}
	errMap.Range(func(k, v interface{}) bool {
		err = v.(error)
		if err != nil {
			return false
		}
		return true
	})
	return err
}
