package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateUserRoleTable_20161110_134051 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUserRoleTable_20161110_134051{}
	m.Created = "20161110_134051"
	migration.Register("CreateUserRoleTable_20161110_134051", m)
}

// Run the migrations
func (m *CreateUserRoleTable_20161110_134051) Up() {
	m.SQL("CREATE TABLE people.user_role (" +
		"`id` int(11) unsigned NOT NULL AUTO_INCREMENT," +
		"`user_id` int(11) NOT NULL," +
		"`role_id` int(11) NOT NULL," +
		"PRIMARY KEY (`id`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;")

}

// Reverse the migrations
func (m *CreateUserRoleTable_20161110_134051) Down() {
	m.SQL("DROP TABLE people.user_role")

}
