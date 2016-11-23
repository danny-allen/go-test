package third

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	//"strconv"
	"test/demo/Exercise/Third/model"
)

// This packages function that is called from exercise.
func Run(){

	// Register models.
	orm.RegisterModel(new(model.User))
	orm.RegisterModel(new(model.Role))
	orm.RegisterModel(new(model.UserRole))

	// Run tests
	Test();
}



// Where the magic is used.
func Test() {

	// Get user.
	user := GetUserById(2)

	// Say what the user is.
	//fmt.Println("The user with ID " + strconv.Itoa(user.Id) + " is " + user.Name + ".")

	// Loop the users roles and print the role names.
	for _, v := range user.UserRoles {
		fmt.Println(v.Role)
	}

	fmt.Println(user)
}


// Get a user by their ID.
func GetUserById(id int) *model.User{

	// Show queries.
	orm.Debug = true

	// Set up ORM.
	o := orm.NewOrm()

	// Define results as Users collection.
	result := &model.User{}

	// The users are on Blaze DB so lets start using that one.
	o.Using("people")

	// Prepare the SQL.
	err := o.QueryTable("user").Filter("id", id).RelatedSel("user_role").One(result)

	if err != nil {
		fmt.Println("Something went wrong")
	}

	// Return the result
	return result
}