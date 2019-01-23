package dbops

import (
	"log"
		_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"time"
	"todogo/pkg/api/utils"
	"todogo/pkg/api/defs"
	)

func AddUserCredential(loginName string, pwd string) error{

	stmtIns, err := dbConn.Prepare("INSERT INTO userstable (login_name, pwd) VALUES (?, ?)")//预编译
	if err != nil {
		return nil
	}
	_, err = stmtIns.Exec(loginName, pwd)//执行 prepare 之后的query
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}
//
func GetUserCredential(loginName string)(string, error) {

	stmtOut, err := dbConn.Prepare("SELECT pwd FROM userstable WHERE	login_name = ?")

	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd) // query 一整行 将搜索到的数据 写进pwd
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()

	return pwd, nil
}
//
//
func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM userstable WHERE login_name =? AND pwd=?")

	if err != nil {
		log.Printf("DeleteUser %s", err)
		return err
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error)  {

	// create uuid
	vid, err := utils.NewUUID()

	if err != nil {
		return nil, err
	}
	// createtime -> db ->

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")//M D y, HH:MM:SS 时间原点

	stmtIns, err := dbConn.Prepare(`INSERT INTO video_info 
			(id, author_id, name, display_ctime) VALUES (?, ?, ?, ?)`)

	if err != nil {
		return nil, err
	}

	_, err = stmtIns.Exec(vid, aid, name, ctime) //执行 prepare 之后的query

	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}

	defer stmtIns.Close()

	return res, nil
}

func GetVideoInfo(vid string)(*defs.VideoInfo, error)  {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE id = ?")

	var aid int
	var dct string
	var name string

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)

	if err != nil && err != sql.ErrNoRows{
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}

	return res, nil
}

func DeleteVideoInfo(vid string) error {

	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id=?")

	if err != nil {
		return err
	}

	_,err = stmtDel.Exec(vid)
	if err != nil{
		return err
	}

	defer stmtDel.Close()

	return nil
}

func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) values(?, ?, ?, ?)")
	if err != nil{
		return err
	}

	_,err = stmtIns.Exec(id, vid, aid, content)

	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	//http://www.ruanyifeng.com/blog/2019/01/table-join.html
	stmtOut, err := dbConn.Prepare(` SELECT comments.id, userstable.Login_name, comments.content FROM comments
		INNER JOIN userstable ON comments.author_id = userstable.id
		WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)

	var res []*defs.Comment
	rows, err := stmtOut.Query(vid, from, to)
	//fmt.Println(rows,"------->>>")
	if err != nil {
		log.Println(err.Error())
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}

		c := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, c)
	}
	defer stmtOut.Close()

	return res, nil
}









