package dbops

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ReadVideoDeletionRecord(count int) ([]string, error) {
	var ids []string
	stmtOut, err := db.Prepare("select video_id from video_del_rec limit ?")
	if err != nil {
		return nil, err
	}
	rows, err := stmtOut.Query(count)
	if err != nil {
		log.Println("Query videodeletionRecord error:", err)
		return ids, err
	}
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}
	defer stmtOut.Close()
	return ids, nil
}
func DelVideoDeletionRecord(vid string) error {
	stmtDel, err := db.Prepare("delete from video_del_rec where video_id = ?")
	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(vid)
	if err != nil {
		log.Println("Deleting VideoDeletionRecord error:", err)
		return err
	}
	defer stmtDel.Close()
	return nil
}
