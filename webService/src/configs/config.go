package configs

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

type Config struct {
	AdminUsername string
	AdminPassword string

	DSN *Dsn
}

type Dsn struct {
	Host   string
	Port   string
	Dbname string

	User string
	Pass string
}

func (r *Dsn) ToString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		r.Host,
		r.Port,
		r.User,
		r.Pass,
		r.Dbname)
}

func GetConfig(log *slog.Logger) *Config {

	file, err := os.Open("../../configs/admin_credentials.txt")
	if err != nil {
		if log != nil {
			log.Error("Error opening file", "error", err.Error())
		}
		return nil
	}
	defer file.Close()

	r := &Config{
		DSN: &Dsn{},
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ": ", 2) // Разделяем строку по ": "
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]

			switch key {
			case "AdminUsername":
				r.AdminUsername = value
			case "AdminPassword":
				r.AdminPassword = value
			case "Host":
				r.DSN.Host = value
			case "Port":
				r.DSN.Port = value
			case "DBname":
				r.DSN.Dbname = value
			case "User":
				r.DSN.User = value
			case "Pass":
				r.DSN.Pass = value
			}
		}
	}
	if scanner.Err() != nil {
		if log != nil {
			log.Error("Err reading file", "err", scanner.Err().Error())
		}
		return nil
	}
	if log != nil {
		log.Debug("Successfully parsed txt", "msg", r)
	}
	return r
}
