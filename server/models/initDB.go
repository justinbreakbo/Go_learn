package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE person (
    first_name text,
    last_name text,
    email text
);

CREATE TABLE place (
    country text,
    city text NULL,
    telcode integer
)`

// 通信的时候，转成json数据 first_name: "xxx"
type Person struct {
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	Email     string `json:"email"`
}

type Place struct {
	Country string         `json:"country"`
	City    sql.NullString `json:"city"`
	TelCode int            `json:"tel_code"`
}

var DB *sqlx.DB

const (
	//host = "106.14.171.4"
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbName   = "postgres"
	port     = 5432
	mode     = "disable"
)

func init() {
	logs.Debug("start to init")

	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbName, mode)

	var err error // open ping
	DB, err = sqlx.Connect("postgres", dbInfo)
	if err != nil {
		logs.Error("Connect Error:\n", err)
		return
	}

	result := DB.MustExec(schema)
	logs.Debug(result)

	tx := DB.MustBegin()
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "John", "Doe", "johndoeDNE@gmail.net")
	tx.MustExec("INSERT INTO place (country, city, telcode) VALUES ($1, $2, $3)", "United States", "New York", "1")
	tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Hong Kong", "852")
	tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Singapore", "65")
	tx.NamedExec("INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", &Person{"Jane", "Citizen", "jane.citzen@example.com"})
	err = tx.Commit()
	if err != nil {
		logs.Error(err)
		return
	}

	logs.Debug("init DB success!")
}
