package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateUserTable_20161110_134046 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUserTable_20161110_134046{}
	m.Created = "20161110_134046"
	migration.Register("CreateUserTable_20161110_134046", m)
}

// Run the migrations
func (m *CreateUserTable_20161110_134046) Up() {
	m.SQL("CREATE TABLE people.user (" +
		"`id` int(11) NOT NULL AUTO_INCREMENT," +
		" `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL," +
		"`email` varchar(255) COLLATE utf8_unicode_ci NOT NULL," +
		"`password` varchar(255) COLLATE utf8_unicode_ci NOT NULL," +
		"PRIMARY KEY (`id`)" +
		") ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;")

}

// Reverse the migrations
func (m *CreateUserTable_20161110_134046) Down() {
	m.SQL("DROP TABLE people.user")

}
