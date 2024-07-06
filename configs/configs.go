package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type connection struct {
	Host, User, Password, DBName, Port string
}

func ImportPasskey() string {
	var passkey string
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println(err)
		return passkey
	}

	passkey = os.Getenv("passkey")
	return passkey
}

func ImportDB() connection {
	var result connection
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println(err)
		return result
	}

	result.Host = os.Getenv("host")
	result.User = os.Getenv("user")
	result.Password = os.Getenv("password")
	result.DBName = os.Getenv("dbname")
	result.Port = os.Getenv("port")

	return result
}

func ConnectDB() *gorm.DB {
	s := ImportDB()

	conn_string := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", s.Host, s.User, s.Password, s.DBName, s.Port)
	db, err := gorm.Open(postgres.Open(conn_string), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	return db
}
