package postgresql

import (
	"encoding/xml"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	var dsnXML dsn
	if err := dsnXML.Configure(); err != nil {
		return nil, err
	}
	db, err := gorm.Open(postgres.Open(dsnXML.toString()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = Ping(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Ping(db *gorm.DB) error {
	dbSql, err := db.DB()
	if err != nil {
		return err
	}
	if err = dbSql.Ping(); err != nil {
		return err
	}
	return nil
}

func TxSaveExecutor(db *gorm.DB, fn func(*gorm.DB) error) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	dbSql, err := db.DB()
	if err != nil {
		tx.Rollback()
		return err
	}
	if err = dbSql.Ping(); err != nil {
		tx.Rollback()
		return err
	}
	if err = fn(db); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

type dsn struct {
	Host   string `xml:"host"`
	Port   string `xml:"portdb"`
	Dbname string `xml:"dbname"`

	User string `xml:"user"`
	Pass string `xml:"pass"`

	Ssl string `xml:"ssl"`
}

func (r *dsn) Configure() error {
	xmlData, err := os.ReadFile("../server/cfg.xml")
	if err != nil {
		return err
	}

	err = xml.Unmarshal(xmlData, &r)
	if err != nil {
		return err
	}
	return nil
}

func (r *dsn) toString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		r.Host,
		r.Port,
		r.User,
		r.Pass,
		r.Dbname,
		r.Ssl)
}