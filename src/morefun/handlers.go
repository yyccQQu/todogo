package morefun

import (
	"net/http"
	"morefun/defs"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
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











