package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host      = "localhost"
	port      = 5432
	user      = "postgres"
	password  = "1234"
	dbname    = "crudbase"
	tablename = "todo"
)

func main() {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname) //io.WriteString(os.Stdout, conn) to print the sprintf(the result is in the format of string) && sslmode is secure socket layer mode is to verify the encripting and libpq will verify the server host name matches the cerificate

	db, err := sql.Open("postgres", conn) //func sql.Open(driverName string, dataSourceName string) (*sql.DB, error)
	CheckError(err)                       //if error is returned we should check the error and the function is written below the main

	err = db.Ping() //this ping() function is used to check whether the database connection is still alive
	CheckError(err) // to check if there is error in pinging

	fmt.Println("Database ", dbname, " is connected")

	//Creating table
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (TASK VARCHAR(40) PRIMARY KEY, DUE VARCHAR(20));", tablename)
	_, e := db.Exec(query) //Executes the query
	CheckError(e)
	for {
		fmt.Println("Enter 1 to Create a Task")
		fmt.Println("Enter 2 to Update a Task")
		fmt.Println("Enter 3 to Get all a Task")
		fmt.Println("Enter 4 to Delete Task")
		fmt.Println("Enter 5 to Exit")
		var choice int
		fmt.Scan(&choice)
		if choice == 5 {
			break
		}
		switch choice {
		case 1:
			CreateTask(db)
		case 2:
			UpdateTask(db)
		case 3:
			GetTask(db)
		case 4:
			DeleteTask(db)
		default:
			fmt.Println("Wrong Choice")
		}
	}

}
func CreateTask(db *sql.DB) {
	var task string
	var due string
	fmt.Println("Enter the Task: ")
	fmt.Scan(&task)
	fmt.Println("Enter the Due Day")
	fmt.Scan(&due)
	insert := fmt.Sprintf("INSERT INTO %s VALUES('%s', '%s');", tablename, task, due)
	_, e := db.Exec(insert)
	CheckError(e)
}
func UpdateTask(db *sql.DB) {
	var task string
	var newtask string
	var due string
	fmt.Println("Enter the Task you want to update: ")
	fmt.Scan(&task)
	fmt.Println("Enter the New Task: ")
	fmt.Scan(&newtask)
	fmt.Println("Enter the new Due Date")
	fmt.Scan(&due)
	update := fmt.Sprintf("UPDATE %s SET TASK = '%s', DUE = '%s' WHERE TASK = '%s';", tablename, newtask, due, task)
	_, e := db.Exec(update)
	CheckError(e)
}
func GetTask(db *sql.DB) {
	get := fmt.Sprintf("SELECT * FROM %s", tablename)
	rows, err := db.Query(get)
	CheckError(err)
	defer rows.Close()
	for rows.Next() {
		var task string
		var due string

		err = rows.Scan(&task, &due)
		CheckError(err)

		fmt.Println(task, due)
	}
	CheckError(err)

}
func DeleteTask(db *sql.DB) {
	var task string
	fmt.Println("Enter the Task that You Want to Delete: ")
	fmt.Scan(&task)
	del := fmt.Sprintf("DELETE FROM %s WHERE TASK = '%s'", tablename, task)
	_, err := db.Exec(del)
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}

}
