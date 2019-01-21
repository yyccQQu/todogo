package dbops

import (
	"log"
		_ "github.com/go-sql-driver/mysql"
	"database/sql"
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

