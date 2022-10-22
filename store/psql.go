package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type User struct {
	ID   int
	Name string
	Age  string
	Job  string
}

// postgres informations
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "usersdb"
)

var db *sql.DB

// call postgres connection in init
func init() {
	var err error
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
}

//Insert User
func InsertUser(u User) {
	res, err := db.Exec("INSERT INTO users(name,age,job) VALUES($1,$2,$3)", u.Name, u.Age, u.Job)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Insert Response is: ", res)
}

func UpdateUser(u User) {
	res, err := db.Exec("UPDATE users SET name=$2, age=$3, job=$4 WHERE id=$1", u.ID, u.Name, u.Age, u.Job)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Update Response is: ", res)
}

// Get All Users from Db
func GetUsers() {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Records Found")
			return
		}
		log.Fatal(err)
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		usr := &User{}
		err := rows.Scan(&usr.ID, &usr.Name, &usr.Age, &usr.Job)

		if err != nil {
			log.Fatal(err)
		}
		users = append(users, usr)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, u := range users {
		fmt.Printf("%d- %s, %s, %s, %s", u.ID, u.Name, u.Age, u.Job)
	}
}

//Get User by Id from DB
func GetUserByID(id int) {
	var user User
	err := db.QueryRow("SELECT * FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Age, &user.Job)

	switch {
	case err == sql.ErrNoRows:
		log.Fatal("No user with that id")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Println("User is name is ", user.Name)
	}
}
