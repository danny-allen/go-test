package main

import (
	"os"
	"test/demo/Exercise/Second"
	"test/demo/Exercise/Third"
	"test/demo/Exercise/First"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"test/demo/Exercise/Fourth"
	"test/demo/Exercise/Fifth"
)

func init() {

	//// Setting up ORM.
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	//
	//// Set default database.
	//err := orm.RegisterDataBase("default", "mysql", "root:password@tcp(127.0.0.1:3306)/people?charset=utf8")
	//
	//// Throw error.
	//if(err != nil) {
	//	panic(err)
	//}
}

// Our main function.
func main() {

	// Check we have a second argument (cli).
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	// Get the package name to run.
	ex := os.Args[1];

	// Check the map has the value from command.
	if val, ok := dataStore()[ex]; ok {

		// Use switch to figure out how to use the value.
		switch val {
			case "First":
				first.Run();
				break
			case "Second":
				second.Run()
				break
			case "Third":
				third.Run()
				break
			case "Fourth":
				fourth.Run()
				break
			case "Fifth":
				fifth.Run()
				break
			default:
				fmt.Println(val)
		}

	} else {

		// If this happens, see message below.
		panic("Your command sucks.")
	}

}

// Get the package data
func dataStore()(map[string]string){

	// Create the exercises map.
	exercises := make(map[string]string)

	// Add to it.
	exercises["first"] = "First"
	exercises["second"] = "Second"
	exercises["third"] = "Third"
	exercises["fourth"] = "Fourth"
	exercises["fifth"] = "Fifth"

	// Return the exercises.
	return exercises
}
