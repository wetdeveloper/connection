package connection

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() (db *sql.DB, err error) {
	fmt.Println("connection.go is running")
	// RemoveDb()
	// CreateDb()
	sqliteDatabase, _ := sql.Open("sqlite3", "./users.db")
	// defer sqliteDatabase.Close()
	// CreateTable(sqliteDatabase)
	// InsertUser(sqliteDatabase, "Ehsan", "5636532")
	// DeleteUser(sqliteDatabase, "Ehsan", "5636532")
	// ListUsers(sqliteDatabase)
	return sqliteDatabase, err
}
func CreateTable(db *sql.DB) {
	createUsersTableSQL := `CREATE TABLE users (
		"idUser" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"username" TEXT NOT NULL UNIQUE,
		"password" TEXT
	  );` 

	log.Println("Create user table...")
	statement, err := db.Prepare(createUsersTableSQL) 
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("user table created")
}

func InsertUser(db *sql.DB, username string, password string) bool {
	if SelectUser(db, username) == false {
		log.Println("Inserting user record ...")
		insertUserSQL := `INSERT INTO users(username,password) VALUES (?, ?)`
		statement, err := db.Prepare(insertUserSQL)

		if err != nil {
			log.Fatalln(err.Error())
			return false
		}
		_, err = statement.Exec(username, password)
		if err != nil {
			log.Fatalln(err.Error())
			return false
		}
		return true
	}
	return false
}

func UpdateUser(db *sql.DB, username string, password string) bool {
	if SelectUser(db, username) {
		log.Println("Updating user record ...")
		UpdateUserSQL := `UPDATE users SET password=? where username=?`
		statement, err := db.Prepare(UpdateUserSQL)

		if err != nil {
			log.Fatalln(err.Error())
			return false
		}
		_, err = statement.Exec(password, username)
		if err != nil {
			log.Fatalln(err.Error())
			return false
		}
		return true
	}
	return false
}

func DeleteUser(db *sql.DB, username string) bool {
	if SelectUser(db, username) {
		log.Println("Deleting user record ...")
		UpdateUserSQL := `DELETE FROM  users WHERE username=?`
		statement, err := db.Prepare(UpdateUserSQL)

		if err != nil {
			log.Fatalln(err.Error())
			return false
		}
		_, err = statement.Exec(username)
		if err != nil {
			log.Fatalln(err.Error())
			return false
		}
		return true
	}
	return false
}

func SelectUser(db *sql.DB, Username string) bool {
	sqlStatement := `SELECT username FROM users WHERE username=$1`
	var username string
	row := db.QueryRow(sqlStatement, Username)
	switch err := row.Scan(&username); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return false
	case nil:
		fmt.Println(username)
		return true
	default:
		panic(err)
	}
}

func ListUsers(db *sql.DB) map[string]interface{} {
	row, err := db.Query("SELECT * FROM users ORDER BY username")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	userslist := map[string]interface{}{}
	for row.Next() {
		var id int
		var username string
		var password string
		row.Scan(&id, &username, &password)
		userslist[username] = password
		log.Println("User: ", username, " ", password, " ")
	}
	return userslist
}

func CreateDb() {
	log.Println("Creating users.db...")
	file, err := os.Create("users.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("users.db created")
}

func RemoveDb() {
	os.Remove("users.db")
}
