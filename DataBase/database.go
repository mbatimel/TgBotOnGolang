package database

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
    constants "example/main/Constants"
	//	"net/http"
	"log"

	_ "github.com/lib/pq"
)



 func inserturl(url string, db *sql.DB) int{

	log.Println("\n"+url+"\n")
    insertDynStmt := `insert into "linkforbyobject"("datetyme", "linkforcite") values($1, $2)`
    _, err := db.Exec(insertDynStmt, time.Now(), url)
	if err != nil {
        log.Print("\n",err)
		return constants.ErrorFromDB
    }
	return constants.OperationDone
 }
func ConnectedForDB(url string ) int {
	
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", constants.Host, constants.Port, constants.User, constants.Password, constants.Dbname)
       
    db, err := sql.Open("postgres", psqlconn)
    if err != nil {
        log.Print("\n",err)
		return  constants.ErrorFromDB
    }  
	
    defer db.Close()
    log.Println("Connected!")
	
    checkResult := CheckUrlInDB(url, db)
    switch checkResult {
    case constants.ThereIsNoLink:
        return inserturl(url, db)
    case constants.LinkIsAlreadyThere:
        return constants.LinkIsAlreadyThere
    case constants.ErrorFromDB:
        return constants.ErrorFromDB
    }

	




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
	return 1
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
func CheckUrlInDB(url string, db *sql.DB ) int {
    sqlStmt := `SELECT linkforcite FROM linkforbyobject WHERE linkforcite = $1`
    var foundURL string
    err := db.QueryRow(sqlStmt, url).Scan(&foundURL)

    if err != nil {
        if err == sql.ErrNoRows {
            // URL не найден в базе данных
            return constants.ThereIsNoLink
        }
        // Произошла ошибка при выполнении запроса
        log.Print(err)
        return constants.ErrorFromDB
    }

    // URL найден в базе данных
    return constants.LinkIsAlreadyThere
}
