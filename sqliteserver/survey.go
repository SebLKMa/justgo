package main

import (
	"database/sql" // sql database interface
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3" // the db driver supports database/sql
)

/*
iox@ubuntu:~/dev/iox-go$ cd sqliteclient/
iox@ubuntu:~/dev/iox-go/sqliteclient$ sqlite3 survey.db
SQLite version 3.22.0 2018-01-22 18:45:57
Enter ".help" for usage hints.
DROP TABLE IF EXISTS person;
CREATE TABLE person (
	email VARCHAR(250) NOT NULL PRIMARY KEY CHECK(email<>''),
	name VARCHAR(250),
	age_group TINYINT,
	bio VARCHAR(1000),
	job VARCHAR (250),
	interests VARCHAR(1000)
	);
iox@ubuntu:~/dev/iox-go/sqliteclient$ ls -l htmlTemplate.db
-rw-r--r-- 1 iox iox 8192 Jul  5 00:49 htmlTemplate.db
*/

// DBname is name of the database
var DBname string

// Person definition for the table named person
type Person struct {
	Email       string
	Name        string
	AgeGroupNum int
	Bio         string
	Job         string
	Interests   string
}

var (
	templates = template.Must(template.ParseGlob("templates/*.html"))
)

// dbConn opens and return the db object. Caller must close db when done with it.
func dbConn(dbName string) (db *sql.DB) {
	// mysql
	/*
		dbDriver := "mysql"
		dbUser := "root"
		dbPass := "root"
		dbName := "goblog"
		db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	*/
	// sqlite
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// defaultHandler handles / route, renders a simple page with a table
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "guten tag!")
}

// surveyHandler handles /survey route, renders go html templates in templates directory
func surveyHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Host: %s Path: %s\n", r.Host, r.URL.Path)

	if err := templates.ExecuteTemplate(w, "survey", nil); err != nil {
		log.Println(err)
	}
}

// insertOrUpdateHandler handles /insertOrUpdate route.
// It inserts or update the person record depending on person PK is found.
// See https://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html
func insertOrUpdateHandler(w http.ResponseWriter, r *http.Request) {
	db := dbConn(DBname)
	defer db.Close()
	if r.Method == "POST" {
		person := Person{}
		person.Email = r.FormValue("user_email")
		person.Name = r.FormValue("user_name")
		ageGroup := getAgeGroup(r)
		num, err := strconv.Atoi(ageGroup)
		if err != nil {
			fmt.Printf("Error getting age group: %v\n", err)
		}
		person.AgeGroupNum = num
		person.Bio = r.FormValue("user_bio")
		// Form returns a list
		// FormValue returns the first value from the list
		person.Job = r.FormValue("user_job")
		person.Interests = getInterests(r)
		log.Printf("Person: %s, %s, %v \n", person.Email, person.Job, person.Interests)

		if person.Email != "" {
			// the Rows object has no count/len, using QueryRow instead
			//rows, err := db.Query("SELECT * FROM person WHERE email=?", email)
			var count int
			db.QueryRow("SELECT COUNT(*) FROM person WHERE email=?", person.Email).Scan(&count)

			if count == 0 {
				stmt, err := db.Prepare("INSERT INTO person(email, name, age_group, bio, job, interests) VALUES(?,?,?,?,?,?)")
				if err != nil {
					panic(err.Error())
				}
				_, err = stmt.Exec(person.Email, person.Name, person.AgeGroupNum, person.Bio, person.Job, person.Interests)
				if err != nil {
					log.Println(err)
				} else {
					log.Println("Person created")
				}
			} else {
				// db.Prepare("UPDATE Employee SET name=?, city=? WHERE id=?")
				stmt, err := db.Prepare("UPDATE person SET name=?, age_group=?, bio=?, job=?, interests=? where email=?")
				if err != nil {
					panic(err.Error())
				}
				_, err = stmt.Exec(person.Name, person.AgeGroupNum, person.Bio, person.Job, person.Interests, person.Email)
				if err != nil {
					log.Println(err)
				} else {
					log.Println("Person updated")
				}
			}

		}
	}
	// redirect to transaction completed page
	http.Redirect(w, r, "/", 301)
}

// getAgeGroup determine the age group number from radio buttons.
// It returns nil if not found.
// https://astaxie.gitbooks.io/build-web-application-with-golang/en/04.2.html
func getAgeGroup(r *http.Request) string {
	slice := []string{"1", "2", "3"}
	for _, v := range slice {
		log.Println(r.Form.Get("user_age_group"))
		if v == r.Form.Get("user_age_group") {
			return v
		}
	}
	return ""
}

// getInterests from checkboxes
func getInterests(r *http.Request) string {
	const delimiter = ","
	result := ""
	var interests []string
	if len(r.Form["user_interest"]) > 0 {
		interests = r.Form["user_interest"]
		for _, v := range interests {
			result += v
			result += delimiter
		}
	}
	result = strings.TrimRight(result, delimiter) // remove last delimiter
	return result
}

// go run survey.go survey.db
func main() {
	args := os.Args
	if len(args) != 2 {
		log.Println("arguments expected: sqlite db file name")
		return
	}

	DBname = args[1]

	// expose REST to view DATA
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/survey", surveyHandler)
	http.HandleFunc("/insertOrUpdate", insertOrUpdateHandler)

	log.Println("Listening at 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
		return
	}

}
