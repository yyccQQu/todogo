package morefun

import (
	"net/http"
	"morefun/defs"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
					"io"
	"io/ioutil"
	"fmt"
	"morefun/dbops"
	"strconv"
)

func findusers() []*(defs.SalesCombo){
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:33061)/agency_dev?charset=utf8")

	query := "SELECT `id`,`combo_name`,`site_num`,`platform_num`,`road_fee`,`status`,`create_time`,`delete_time` FROM `sales_combo`"

	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	var users = make([]*(defs.SalesCombo),0)

	for rows.Next() {
		scombo := new(defs.SalesCombo)
		err := rows.Scan(
			&scombo.Id,
			&scombo.ComboName,
			&scombo.SiteNum,
			&scombo.PlatformNum,
			&scombo.RoadFee,
			&scombo.Status,
			&scombo.CreateTime,
			&scombo.DeleteTime)

		if err != nil {
			panic(err)
		}
		users = append(users, scombo)
	}
	return users
	
}

func Register(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	scombo := findusers()

	data, err := json.Marshal(scombo)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(200)
	w.Write(data)

}

func findsimple() []*(defs.Userstable) {
	db, err := sql.Open("mysql","root:root@tcp(127.0.0.1:33061)/agency_dev?charset=utf8")
	query := "SELECT `id`,`login_name`,`pwd` FROM `userstable`"

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	var simpleUser = make([]*(defs.Userstable),0)

	for rows.Next(){
		simUser := new(defs.Userstable)
		err := rows.Scan(&simUser.Id,&simUser.LoginName,&simUser.Pwd)
		if err!=nil {
			panic(err)
		}
		simpleUser = append(simpleUser, simUser)
	}
	return simpleUser
}

func Hasagain(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	someList := findsimple()

	data, err := json.Marshal(someList)

	if err!= nil{
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(200)
	w.Write(data)
}

//post
func Resign(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {

	name := p.ByName("fun")

	res, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(res),"res",name)

	ubody := &defs.UserCredential{}
	//默认username 为路由的子级
	if len(ubody.Username)==0 {
		ubody.Username = name
	}

	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println("---->err",err)
		panic(err)
		//io.WriteString(w, string(resStr))
		return
	}

	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		fmt.Println("---->",err)
		//报错方法
		dbops.SendErrorResponse(w, defs.ErrorDBError)
		return
	}
	w.WriteHeader(200)
	io.WriteString(w, "ok "+name)
}

func Delfun(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {

	inum := p.ByName("fun")
	ins, err:= strconv.Atoi(inum)
	if err != nil {
		panic(err)
	}
	//ins,_ := strconv.ParseFloat(inum,64)
	dbops.DeleteUser(ins)
	w.WriteHeader(200)
	io.WriteString(w, "ok "+inum)
}






