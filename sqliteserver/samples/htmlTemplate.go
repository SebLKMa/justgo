package samples

import (
	"database/sql" // sql database interface
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3" // the db driver supports database/sql
)

/*
iox@ubuntu:~/dev/iox-go$ cd sqliteclient/
iox@ubuntu:~/dev/iox-go/sqliteclient$ sqlite3 htmlTemplate.db
SQLite version 3.22.0 2018-01-22 18:45:57
Enter ".help" for usage hints.
DROP TABLE IF EXISTS person;
CREATE TABLE person (
	email VARCHAR(250) NOT NULL PRIMARY KEY CHECK(email<>''),
	name VARCHAR(250),
	age_group TINYINT);
iox@ubuntu:~/dev/iox-go/sqliteclient$ ls -l htmlTemplate.db
-rw-r--r-- 1 iox iox 8192 Jul  5 00:49 htmlTemplate.db
*/

// DATABASE is the name of the database
const DATABASE = "htmlTemplate.db"

// Person definition for the table named person
type Person struct {
	Email    string
	Name     string
	AgeGroup int
}

/* The database table is mapped to this slice:
iox@ubuntu:~/dev/iox-go$ cd sqliteclient/
iox@ubuntu:~/dev/iox-go/sqliteclient$ sqlite3 htmlTemplate.db
SQLite version 3.22.0 2018-01-22 18:45:57
Enter ".help" for usage hints.
sqlite> CREATE TABLE data (
   ...> number INTEGER PRIMARY KEY,
   ...> double INTEGER,
   ...> square INTEGER );
sqlite>
iox@ubuntu:~/dev/iox-go/sqliteclient$ ls -l htmlTemplate.db
-rw-r--r-- 1 iox iox 8192 Jul  5 00:49 htmlTemplate.db
*/

// Entry definition for the table named data
type Entry struct {
	Number int
	Double int
	Square int
}

// DATA is a slice of Entries to be passed to template file
var DATA []Entry

var templateFile string // name of template file
var templateSurveyFile string

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
	log.Printf("Host: %s Path: %s\n", r.Host, r.URL.Path)
	myTemplate := template.Must(template.ParseGlob(templateFile))
	myTemplate.ExecuteTemplate(w, templateFile, DATA)
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
	db := dbConn(DATABASE)
	defer db.Close()
	if r.Method == "POST" {
		email := r.FormValue("user_email")
		name := r.FormValue("user_name")
		ageGroup := getAgeGroup(r)
		ageGroupNum, err := strconv.Atoi(ageGroup)
		if err != nil {
			fmt.Printf("Error getting age group: %v\n", err)
		}
		log.Println("Person:" + email)
		if email != "" {
			// the Rows object has no count/len, using QueryRow instead
			//rows, err := db.Query("SELECT * FROM person WHERE email=?", email)
			var count int
			db.QueryRow("SELECT COUNT(*) FROM person WHERE email=?", email).Scan(&count)

			if count == 0 {
				stmt, err := db.Prepare("INSERT INTO person(email, name, age_group) VALUES(?,?,?)")
				if err != nil {
					panic(err.Error())
				}
				_, err = stmt.Exec(email, name, ageGroupNum)
				if err != nil {
					log.Println(err)
				} else {
					log.Println("Person created")
				}
			} else {
				// db.Prepare("UPDATE Employee SET name=?, city=? WHERE id=?")
				stmt, err := db.Prepare("UPDATE person SET name=?, age_group=? where email=?")
				if err != nil {
					panic(err.Error())
				}
				_, err = stmt.Exec(name, ageGroupNum, email)
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

// getAgeGroup determine the age group number from radio buttons
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

func populateSampleData(dbName string) error {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return err
	}
	defer db.Close()

	log.Println("Emptying database table")
	_, err = db.Exec("DELETE FROM data")
	if err != nil {
		return err
	}

	log.Println("Populating database table")
	stmt, _ := db.Prepare("INSERT INTO data(number, double, square) VALUES(?,?,?)")
	for i := 20; i < 50; i++ {
		_, err = stmt.Exec(i, 2*i, i*i)
		if err != nil {
			return err
		}
	}

	rows, err := db.Query("SELECT * FROM data")
	if err != nil {
		return err
	}

	// fill up DATA from db table
	var n int
	var d int
	var s int
	for rows.Next() {
		err = rows.Scan(&n, &d, &s)
		tmp := Entry{Number: n, Double: d, Square: s}
		DATA = append(DATA, tmp)
	}

	return nil
}

// go run htmlTemplate.go htmlTemplate.db html.gohtml
func main() {
	args := os.Args
	if len(args) != 3 {
		log.Println("arguments expected: sqlite db file name and template file name")
		return
	}

	dbName := args[1]
	templateFile = args[2]
	templateSurveyFile = "survey.html"

	err := populateSampleData(dbName)
	if err != nil {
		log.Println(err)
		return
	}

	// expose REST to view DATA
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/survey", surveyHandler)
	http.HandleFunc("/insertOrUpdate", insertOrUpdateHandler)

	log.Println("Listening 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
		return
	}

}
