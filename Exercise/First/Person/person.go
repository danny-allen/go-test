package person

import(
	"test/demo/Exercise/First/Pet"
)


// Declare the Person structure.
type Person struct {
	Name string
	Power int
	Pet *pet.Pet
	PowerUp bool
}

// Create a new Person.
func New(n string, p int, pt *pet.Pet, pu bool) *Person {
	return &Person {
		Name: n,
		Power: p,
		Pet: pt,
		PowerUp: pu,
	}
}