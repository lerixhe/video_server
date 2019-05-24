package dbops

import "log"

func AddVideoDeletionRecord(vid string) error {
	stmtIns, err := db.Prepare("insert into video_del_rec (video_id) values(?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(vid)
	if err != nil {
		log.Printf("addvideodelectionrecord error:", err)
		return err
	}
	defer stmtIns.Close()
	return nil

}
