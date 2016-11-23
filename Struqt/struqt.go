package struqt

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"log"
)

func Bind(sql string, entity interface{}) interface{} {


	// DB connection.
	db, _ := sqlx.Open("mysql", "root:password@tcp(localhost:3306)/people")

	// Run the query.
	rows, _ := db.Queryx(sql)

	fmt.Println(rows)

	// Loop the results rows
	for rows.Next() {

		// Populate the struct.
		err := rows.StructScan(&entity)

		// Check for struct scan error.
		if err != nil {
			log.Fatal(err)
		}
	}
	err := rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return entity
}