package dbops

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strconv"
	"io/ioutil"
	"morefun/defs"
	"encoding/json"
	"fmt"
	"io"
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


func DeleteUser(sid int) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM userstable WHERE id = ?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	if _, err := stmtOut.Query(sid); err != nil {
		return err
	}

	return nil
}

func Putfun(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {

	inum ,err := strconv.Atoi(p.ByName("fun"))
	if err != nil {
		log.Println(err,"-----<fun error>")
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println("---->err",err)
		panic(err)
		return
	}

	// 方法1
	stmtIn, err := dbConn.Prepare("UPDATE `userstable` SET login_name=?,pwd=? WHERE id= ? ")
	if err != nil {
		log.Println(err,"-----<error>")
		return
	}
	_, err = stmtIn.Exec(ubody.Username,ubody.Pwd,inum)
	if err != nil {
		panic(err)
		return
	}

	// 方法2
	//update := fmt.Sprintf("UPDATE `userstable` SET login_name='%s',pwd='%s' WHERE id='%v' ",ubody.Username,ubody.Pwd,inum)
	//fmt.Println(update)
	////dbConn.QueryRow()
	//result, err := dbConn.Exec(update)
	//if err != nil {
	//	log.Println("exec failed:", err, ", sql:", update)
	//	return
	//}
	//fmt.Println(result)

	w.WriteHeader(200)
	io.WriteString(w, "ok update")
	//w.Write([]byte("hello world"))
}