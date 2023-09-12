package database

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	//	"net/http"
	"log"

	_ "github.com/lib/pq"
)


const (
    host     = "localhost"
    port     = 5433
    user     = "postgres"
    password = "tatarin17"
    dbname   = "postgres"
)
 func inserturl(url string, db *sql.DB) bool{
	// insert
    // dynamic
	log.Println("\n"+url+"\n")
    insertDynStmt := `insert into "linkforbyobject"("datetyme", "linkforcite") values($1, $2)`
    _, err := db.Exec(insertDynStmt, time.Now(), url)
	if err != nil {
        log.Print("\n",err)
		return false
    }
	return true
 }
func ConnectedForDB(url string ) bool {
	
       
	// connection string
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
       
	// open database
    db, err := sql.Open("postgres", psqlconn)
    if err != nil {
        log.Print("\n",err)
		return  false
    }
        
	// close database
    defer db.Close()
    log.Println("Connected!")
	
	// insert
    // dynamic
	return inserturl(url, db)



	// 	// update
	// updateStmt := `update "Students" set "Name"=$1, "Roll_Number"=$2 where "id"=$3`
	// _, err = db.Exec(updateStmt, "Rachel", 24, 8)
//  panic = CheckError(err)
	// 	// Delete
	// deleteStmt := `delete from "Students" where id=$1`
	// _, err = db.Exec(deleteStmt, 1)
	// panic = CheckError(err)
	// //возвразение значений из бд
// 		rows, err := db.Query(`SELECT "Name", "Roll_Number" FROM "Students"`)
// panic = CheckError(err)	
// 	defer rows.Close()
// 	for rows.Next() {
// 		var name string
// 		var roll_number int
	
// 		err = rows.Scan(&name, &roll_number)
// 		CheckError(err)
	
// 		fmt.Println(name, roll_number)
// 	}
	
// panic = CheckError(err)
	
}

func IsAccessibleURL(url string) bool {
    timeout := time.Duration(5 * time.Second)
    client := http.Client{
        Timeout: timeout,
    }

    resp, err := client.Get(url)
    if err != nil {
		log.Println(err)
        return false
    }
    defer resp.Body.Close()

    return resp.StatusCode == 200
}