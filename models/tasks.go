package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Task is a struct containing Task data
type Task struct {
	UserId       int    `json:"userId"`
	UserName     string `json:"userName"`
	UserPassword string `json:"userPassword"`
}

// TaskCollection is collection of Tasks
type TaskCollection struct {
	Tasks []Task `json:"items"`
}

// GetTasks from the DB
func GetTasks(db *sql.DB) TaskCollection {
	sqlStatement := "SELECT * FROM ev_useraccount"
	rows, err := db.Query(sqlStatement)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		panic(err)
	}
	// make sure to cleanup when the program exits
	defer rows.Close()

	result := TaskCollection{}
	for rows.Next() {
		task := Task{}
		err2 := rows.Scan(&task.UserId, &task.UserName, &task.UserPassword)
		// Exit if we get an error
		if err2 != nil {
			panic(err2)
		}
		result.Tasks = append(result.Tasks, task)
	}
	return result
}

// PutTask into DB
func PutTask(db *sql.DB, userName, userPassword string) (int, error) {
	sqlStatement := `INSERT INTO ev_useraccount(username, userpassword) 
			VALUES($1, $2)
			RETURNING userid`

	// Replace the '?' in our prepared statement with 'name'
	result := 0
	err := db.QueryRow(sqlStatement, userName, userPassword).Scan(&result)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}

	return result, err
}

// DeleteTask from DB
func DeleteTask(db *sql.DB, userId int) (int64, error) {
	sqlStatement := `DELETE FROM ev_useraccount
					 WHERE userid = $1`
	fmt.Println(userId)
	res, err := db.Exec(sqlStatement, userId)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	return count, err
}
