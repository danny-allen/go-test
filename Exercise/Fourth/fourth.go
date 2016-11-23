package fourth

import (
	"github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/go-xweb/log"
	"fmt"
	"test/demo/Struqt"
)

// Declare the DB var.
var db *sqlx.DB


// User model.
type User struct {
	Id 			int
	Name 		string
	Email 		string
	Password	string
}
type Users []User


// Initial functionality - Setup DB connection.
func init()  {

	// Make and set the DB connection.
	db, _ = sqlx.Open("mysql", "root:password@tcp(localhost:3306)/people")
}


// This packages function that is called from exercise.
func Run(){

	//// Run squirrel test.
	//SquirrelTest();
	//
	//// Run SQLX test.
	SQLXTest();

	// Populate struct from abstract method.
	//AbstractEntityTest();
}


// Squirrel query builder usage.
func SquirrelTest() {

	// Define new User.
	user := &User{}

	// Make query for particular user.
	query := squirrel.Select("*").From("user").Where("id = 1")

	// Make SQL from the query.
	sql, _, _ := query.ToSql()

	// Run the query.
	rows, _ := db.Query(sql)

	// Loop the results rows.
	for rows.Next() {

		// Populate the struct.
		err := rows.Scan(&user.Id, &user.Name, &user.Email)

		// Check for struct scan error.
		if err != nil {
			log.Fatal(err)
		}
	}

	// Tell out put the test type.
	fmt.Println("SQUIRREL TEST:")

	// Print the user information.
	fmt.Println(user)
}


// Squirrel usage with Scan Struct from SQLX.
func SQLXTest() {

	// Define new User.
	user := &User{}

	// Make query for particular user.
	query := squirrel.Select("*").From("user").Where("id = 2")

	// Make SQL from the query.
	sql, _, _ := query.ToSql()

	// Run the query.
	rows, _ := db.Queryx(sql) // note the X!

	// Loop the results rows.
	for rows.Next() {

		// Populate the struct.
		err := rows.StructScan(&user) // SQLX method.

		// Check for struct scan error.
		if err != nil {
			log.Fatal(err)
		}
	}

	// Tell out put the test type.
	fmt.Println("SQLX TEST:")

	// Print the user information.
	fmt.Println(user)
}

func AbstractEntityTest() {

	fmt.Println("hello!")

	//Define new User.
	user := &User{}

	// Make query for particular user.
	query := squirrel.Select("*").From("user").Where("id = 3")

	// Make SQL from the query.
	sql, _, _ := query.ToSql()

	// Abstraction.
	struqt.Bind(sql, &user)

	// Tell out put the test type.
	fmt.Println("SQLX TEST:")

	// Print the user information.
	fmt.Println(user)

}