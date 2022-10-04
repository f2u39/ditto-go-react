package mysql

import (
	_ "github.com/go-sql-driver/mysql"
)

// func New() *sql.DB {
// 	cfg := core.AppCfg()
// 	db, err := sql.Open("mysql", cfg.Database.MySQL.Username+":"+cfg.Database.MySQL.Password+"@/"+cfg.Database.MySQL.Database)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return db
// }
