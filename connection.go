package connection

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func connect() {
	fmt.Println("connection.go is running")
	//removeDb()
	//createDb()
	sqliteDatabase, _ := sql.Open("sqlite3", "./users.db")
	defer sqliteDatabase.Close()
	// createTable(sqliteDatabase)
	insertUser(sqliteDatabase, "Ali", "5636532")
	// listUsers(sqliteDatabase)
}
func createTable(db *sql.DB) {
	createUsersTableSQL := `CREATE TABLE users (
		"idUser" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"username" TEXT,
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

func insertUser(db *sql.DB, username string, password string) {
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

func listUsers(db *sql.DB) {
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

func createDb() {
	log.Println("Creating users.db...")
	file, err := os.Create("users.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("users.db created")
}

func removeDb() {
	os.Remove("users.db")
}
