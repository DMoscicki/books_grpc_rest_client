package connector

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB // объект бд

var once sync.Once

var MysqlInfo string

// функция чтения параметров из переменного окружения (в данном случае из .env файла)
func Getenv(key string) string {

	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv(key)
}

// Паттерн синглтон для возврата коннекта к БД
func Singleton() *sql.DB {

	once.Do(func() {
			var err error
			host := Getenv("host")
			user := Getenv("user")
			password := Getenv("password")
			protocol := Getenv("protocol")
			dbname := Getenv("dbname")
			port := Getenv("port")
		
			MysqlInfo = fmt.Sprintf("%s:%s@%s(%s:%s)/%s", user, password, protocol, host, port, dbname)
		
			log.Println(MysqlInfo)
		
			db, err = sql.Open(Getenv("driver"), MysqlInfo)
			if err != nil {
				panic(err)
			}

			err = db.Ping()
			if err != nil {
				db, _ = sql.Open(Getenv("driver"), MysqlInfo)
			}

		},
	)

	return db
}