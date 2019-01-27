package dbops

import (
	"log"
	_ "github.com/go-sql-driver/mysql"
)

// api -> video_id -> mysql
// dispatcher -> mysql -> video_id -> datachannel  behere
// executor -> datachannel-> video_id -> delete videos

func ReadVideoDeletionRecord(count int) ([]string, error) {
	stmtOut, err := dbConn.Prepare("SELECT video_id FROM video_del_rec LIMIT ?") //搜索 将要删除的video的video_id

	var ids []string

	if err != nil {
		return ids, err
	}

	rows, err := stmtOut.Query(count)
	if err != nil {
		log.Printf("Query VideoDeletionRecord error: %v", err)
		return ids, err
	}

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	defer stmtOut.Close()
	return ids, nil
}

func DelVideoDeletionRecord(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_del_rec WHERE video_id=?") //将搜索到的id匹配进行删除操作
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		log.Printf("Deleting VideoDeletionRecord error: %v", err)
		return err
	}

	defer stmtDel.Close()
	return nil
}






