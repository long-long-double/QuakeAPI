package db

import (
	"QuakeAPI/log"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func SaveDataToMySQL(dataList []string, Config MySQLConfig) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",
		Config.Username,
		Config.Password,
		Config.Network,
		Config.Server,
		Config.Port,
		Config.Database)
	log.Log("Connecting MySQL...", log.INFO)
	DB, err := sql.Open("mysql", dsn)
	if DB != nil {
		err = DB.Ping()
		if err != nil {
			fmt.Println(err)
			log.Log("MySQL Connect Error", log.ERROR)
			return
		}
	}
	defer func() {
		if DB != nil {
			_ = DB.Close()
		}
	}()
	if err != nil {
		log.Log("MySQL Connect Error", log.ERROR)
		return
	}
	log.Log("Drop Old Database", log.INFO)
	_, _ = DB.Exec("DROP DATABASE IF EXISTS `quakeapi`;")
	log.Log("Create New Database", log.INFO)
	_, _ = DB.Exec("CREATE DATABASE `quakeapi`;")
	log.Log("Use Database", log.INFO)
	_, _ = DB.Exec("USE `quakeapi`;")
	log.Log("Create Table", log.INFO)
	_, _ = DB.Exec("CREATE TABLE IF NOT EXISTS `result`(" +
		"`result_id` INT UNSIGNED AUTO_INCREMENT," +
		"`result` VARCHAR(100) NOT NULL," +
		"PRIMARY KEY (`result_id`))" +
		"ENGINE=InnoDB DEFAULT CHARSET=utf8;")
	log.Log("Saving Data...", log.INFO)
	for i := 0; i < len(dataList); i++ {
		_, _ = DB.Exec("INSERT INTO `result` (result) VALUE ('" + dataList[i] + "');")
	}
	var res *sql.Rows
	res, err = DB.Query("SELECT COUNT(1) AS result FROM `result`;")
	defer func() {
		if res != nil {
			_ = res.Close()
		}
	}()
	if err != nil {
		log.Log("Saving Data Error", log.ERROR)
		return
	}
	var result string
	for res.Next() {
		err = res.Scan(&result)
		if err != nil || result != strconv.Itoa(len(dataList)) {
			log.Log("Saving Data Error", log.ERROR)
			return
		}
	}
	log.Log("Saving Data Successfully", log.INFO)
}
