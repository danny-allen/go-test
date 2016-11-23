package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateRoleTable_20161110_134040 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateRoleTable_20161110_134040{}
	m.Created = "20161110_134040"
	migration.Register("CreateRoleTable_20161110_134040", m)
}

// Run the migrations
func (m *CreateRoleTable_20161110_134040) Up() {
	m.SQL("CREATE TABLE people.role (" +
		"`id` int(11) unsigned NOT NULL AUTO_INCREMENT," +
		"`name` varchar(100) NOT NULL DEFAULT ''," +
		"`permissions` text COLLATE utf8_unicode_ci," +
		"PRIMARY KEY (`id`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;")

}

// Reverse the migrations
func (m *CreateRoleTable_20161110_134040) Down() {
	m.SQL("DROP TABLE people.role")

}
