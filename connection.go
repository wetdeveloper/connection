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
	//RemoveDb()
	// CreateDb()
	sqliteDatabase, _ := sql.Open("sqlite3", "./users.db")
	// defer sqliteDatabase.Close()
	// CreateTable(sqliteDatabase)
	// InsertUser(sqliteDatabase, "Ali", "5636532")
	// ListUsers(sqliteDatabase)
	return sqliteDatabase, err
}
func CreateTable(db *sql.DB) {
	createUsersTableSQL := `CREATE TABLE users (
		"idUser" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"username" TEXT NOT NULL UNIQUE,
		"password" TEXT
	  );` // SQL Statement for Create Table

	log.Println("Create user table...")
	statement, err := db.Prepare(createUsersTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("user table created")
}

func InsertUser(db *sql.DB, username string, password string) {
	log.Println("Inserting user record ...")
	insertUserSQL := `INSERT INTO users(username,password) VALUES (?, ?)`
	statement, err := db.Prepare(insertUserSQL)

	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(username, password)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func ListUsers(db *sql.DB) {
	row, err := db.Query("SELECT * FROM users ORDER BY username")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		var id int
		var username string
		var password string
		row.Scan(&id, &username, &password)
		log.Println("User: ", username, " ", password, " ")
	}
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
