package connection

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//DbConn var connection
var DbConn *sql.DB

//InitDal connection BD
func InitDal() {
	//connString := os.Getenv("MY_SQL_URL")
	connString := "root:root@tcp(127.0.0.1:3306)/db_golang_echo_jwt?charset=utf8mb4"
	db, err := sql.Open("mysql", connString)
	CheckErr(err)
	DbConn = db
}

//CheckErr Error utils
func CheckErr(err error) {
	if err != nil {
		panic(err)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}