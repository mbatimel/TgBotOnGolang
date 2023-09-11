package database

import (
	//"time"
	"fmt"
	"database/sql"
	//"net/http"
	"log"
	//_ "github.com/lib/pq"
)
type DataBAseInterfaser interface { 
	 ConnectedForDB()

}
// type Database struct {
// 	datetime time.Time
// 	URL string
// }
const (
    host     = "localhost"
    port     = 5433
    user     = "postgres"
    password = "tatarin17"
    dbname   = "postgres"
)
 
func ConnectedForDB() {
        // connection string
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
        // open database
    db, err := sql.Open("postgres", psqlconn)
    CheckError(err)
        // close database
    defer db.Close()
    log.Println("Connected!")
	    // insert
    // dynamic
    insertDynStmt := `insert into "Students"("Name", "Roll_Number") values($1, $2)`
    _, err = db.Exec(insertDynStmt, "Jack", 21)
    CheckError(err)

		// update
	updateStmt := `update "Students" set "Name"=$1, "Roll_Number"=$2 where "id"=$3`
	_, err = db.Exec(updateStmt, "Rachel", 24, 8)
	CheckError(err)

		// Delete
	deleteStmt := `delete from "Students" where id=$1`
	_, err = db.Exec(deleteStmt, 1)
	CheckError(err)
//возвразение значений из бд
		rows, err := db.Query(`SELECT "Name", "Roll_Number" FROM "Students"`)
	CheckError(err)
	
	defer rows.Close()
	for rows.Next() {
		var name string
		var roll_number int
	
		err = rows.Scan(&name, &roll_number)
		CheckError(err)
	
		fmt.Println(name, roll_number)
	}
	
	CheckError(err)
}
 
func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}