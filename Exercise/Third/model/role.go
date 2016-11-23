package model


// Role model.
type Role struct {
	// Fields.
	Id 				int
	Name 			string
	Permissions 	string
	UserRoles		[]*UserRole `orm:"reverse(many)"`

}

// Collection
type Roles []Role

//
//
//type Persons struct {
//	Id int64 orm:"column(id);auto;pk"
//FullName string orm:"column(full_name);size(512);index"
//Addresses []*ContactAddresses orm:rel(m2m)
//}
//
//type ContactAddresses struct {
//	Id int64 orm:"column(id);auto"
//Street1 string orm:"column(street1);size(256)"
//Street2 string orm:"column(street2);size(256)"
//area *Areas orm:"column(area_id);rel(fk)"
//}
//type Areas struct {
//	Id int64 orm:"column(id);auto;pk"
//Zip string orm:"column(pincode);size(128);index"
//City string orm:"column(taluka);size(128);index"
//County string orm:"column(district);size(128);index"
//State string orm:"column(state);size(128);index"
//}
/////// dbsync.go
//package main
//import (
//"github.com/astaxie/beego/orm"
//_ "github.com/go-sql-driver/mysql"
//)
//
//func main() {
//	connStr := "root:@tcp(localhost:3306)/testdb?charset=utf8"
//	orm.RegisterDriver("mysql", orm.DR_MySQL)
//	orm.RegisterDataBase("default", "mysql", connStr)
//	orm.RegisterModel(new(models.Areas), new(models.ContactAddresses), new(models.Persons))
//	orm.RunCommand()
//}