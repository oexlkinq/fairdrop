package main

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/glebarez/go-sqlite"
	"github.com/jmoiron/sqlx"
)

var schema = `
CREATE TABLE folders (
	password text,
	creation_date int default current_timestamp,
	ip text
)`

type DB struct {
	*sqlx.DB
}

type Folder struct {
	Password     string `json:"password"`
	CreationDate int    `db:"creation_date" json:"creation_date"`
	Ip           string `json:"ip"`
}

func Connect(DBFile string) *DB {
	db, err := sqlx.Connect("sqlite", DBFile)
	if err != nil {
		log.Fatalln(fmt.Errorf("connect to sqlite db: %w", err))
	}

	db.MustExec(schema)

	return &DB{db}
}

func (db *DB) CreateFolder(ctx context.Context, passwordComplexity PasswordComplexity, ip string) (*Folder, error) {
	folder := &Folder{
		Password:     generatePassword(passwordComplexity),
		CreationDate: int(time.Now().Unix()),
		Ip:           ip,
	}

	_, err := db.NamedExecContext(ctx, "INSERT INTO folders VALUES (:password, :creation_date, :ip)", folder)
	if err != nil {
		return nil, fmt.Errorf("create folder: %w", err)
	}

	return folder, nil
}
