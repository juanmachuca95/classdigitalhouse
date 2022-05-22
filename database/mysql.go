package database

import (
	"database/sql"

	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

/*
	Driver MySQL GoDriver
	https://github.com/go-sql-driver/mysql

	Example and uses
	http://go-database-sql.org/index.html
*/

type MySQLClient struct {
	*sql.DB
}

func NewMySQLClient() *MySQLClient {
	/* Seguir el siguiente formato user:password@tcp(127.0.0.1:3306)/hello */
	user := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	hostname := os.Getenv("HOSTNAME")
	port := os.Getenv("PORT")
	database := os.Getenv("DATABASE")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, hostname, port, database)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("No fue posible conectar con la base de datos ", connectionString)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("La conexion a la base de datos no esta disponible")
	}

	return &MySQLClient{db}
}
