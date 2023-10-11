package DBConn

import "github.com/go-sql-driver/mysql"

var Cfg = mysql.Config{
	User:   "root",
	Passwd: "astropass",
	Net:    "tcp",
	Addr:   "127.0.0.1:3306",
	DBName: "AstroRankings",
}
