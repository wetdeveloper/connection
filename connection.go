package connection

import (
	// "database/sql"

	"fmt"
	"log"

	"github.com/wetdeveloper/crud-api-config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// _ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID int `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	gorm.Model
	Username string `gorm:UNIQUE`
	Password string
}

func Connect() (db *gorm.DB, err error) {
	dbAdd := config.AppConfig()["database-address"]
	fmt.Println("database-address:", dbAdd)
	mydb, err := gorm.Open(sqlite.Open(dbAdd), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	mydb.AutoMigrate(&User{})
	return mydb, err
}

func InsertUser(mydb *gorm.DB, username string, password string) bool {
	if SelectUser(mydb, username) == false {
		fmt.Println("Inserting user record ...")
		mydb.Create(&User{Username: username, Password: password})
		return true
	}
	fmt.Println("Error:This username already exists")
	return false
}

func UpdateUser(mydb *gorm.DB, username string, password string) bool {
	if SelectUser(mydb, username) {
		var user User
		res := mydb.Model(&user).Where("username= ?", username).Updates(map[string]interface{}{"Username": username, "Password": password})
		if res.Error != nil {
			fmt.Println(res.Error)
			return false
		}
		log.Println("Updating user record ...")
		return true
	}
	log.Println("Error:user not found to update")
	return false
}

func DeleteUser(mydb *gorm.DB, username string) bool {
	if SelectUser(mydb, username) {
		var user User
		res := mydb.Model(&user).Where("username= ?", username).Delete(&user)
		if res.Error != nil {
			fmt.Println(res.Error)
			return false
		}
		log.Println("Deleting user record ...")
		return true
	}
	log.Println("Error:user not found to delete")
	return false
}

func SelectUser(mydb *gorm.DB, username string) bool {
	var user User
	result := mydb.First(&user, "username= ?", username)
	if result.Error != nil {
		fmt.Println(result.Error, "------ignore this message if you're inserting!")
		return false
	}
	return true
}

func ListUsers(mydb *gorm.DB) map[string]interface{} {
	userslist := map[string]interface{}{}
	var user User
	rows, _ := mydb.Model(&User{}).Rows()
	for rows.Next() {
		mydb.ScanRows(rows, &user)
		// fmt.Println(user.Username)
		userslist[user.Username] = user.Password
	}
	return userslist

}
