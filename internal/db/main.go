package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS folders (
	password TEXT,
	creation_date INT DEFAULT CURRENT_TIMESTAMP,
	ip TEXT
)`

type DB struct {
	*sqlx.DB
}

type Folder struct {
	Password     string `db:"password" json:"password"`
	CreationDate int    `db:"creation_date" json:"creation_date"`
	Ip           string `db:"ip" json:"ip"`
}

func Connect(DBFile string) *DB {
	db, err := sqlx.Connect("sqlite3", DBFile)
	if err != nil {
		log.Fatalln(fmt.Errorf("connect to sqlite db: %w", err))
	}

	db.MustExec(schema)

	return &DB{db}
}

func (db *DB) CreateFolder(c context.Context, pw string, ip string) *Folder {
	folder := &Folder{
		Password:     pw,
		CreationDate: int(time.Now().Unix()),
		Ip:           ip,
	}

	_, err := db.NamedExecContext(c, "INSERT INTO folders VALUES (:password, :creation_date, :ip)", folder)
	if err != nil {
		panic(fmt.Errorf("create folder: %w", err))
	}

	return folder
}

func (db *DB) TestFolder(c context.Context, pw string) bool {
	var res bool

	err := db.GetContext(c, &res, "SELECT true FROM folders f WHERE password = $1", pw)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false
		}

		panic(fmt.Errorf("test folder: %w", err))
	}

	return true
}
