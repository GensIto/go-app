package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"todo-app/config"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser     = "users"
	tableNameTodo     = "todos"
	tableNameSessions = "sessions"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := `CREATE TABLE IF NOT EXISTS ` + tableNameUser + `(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`
	Db.Exec(cmdU)

	cmdT := `CREATE TABLE IF NOT EXISTS ` + tableNameTodo + `(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME)`
	Db.Exec(cmdT)

	cmdS := `CREATE TABLE IF NOT EXISTS ` + tableNameSessions + `(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		user_id INTEGER,
		created_at DATETIME)`
	Db.Exec(cmdS)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
