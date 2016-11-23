package first

import (
	"fmt"
	"test/demo/Exercise/First/Pet"
	"test/demo/Exercise/First/Person"
)

// Our main function.
func Run() {

	// Set the power
	power := 1200;

	// Create a new person, with the pet!
	pet1 := pet.New("Garfield", "cat", 120);
	person1 := person.New("Jon", power, pet1, true);

	// Create a another person, with the pet!
	pet2 := pet.New("Tom", "cat", 160);
	person2 := person.New("Jill", power*2, pet2, false);

	// Create people slice.
	people := []*person.Person{person1, person2}

	// Display the people details!
	displayPeople(people)
}


// Loop all the people.
func displayPeople(people []*person.Person){

	// Loop over the people.
	for _, value := range people {

		// Declare the person.
		fmt.Printf("\nPerson: %s.\n", value.Name);

		// Check power up enabled.
		if value.PowerUp {

			// Super power this person!
			super(value, 500);
			fmt.Println("\nPOWER UP!");
		} else {

			// No power up :(
			fmt.Println("\nNo power up :(");
		}

		// Whats happened?
		fmt.Printf("\n%s's power is over %d\n", value.Name, value.Power);
		fmt.Printf("%s has a pet %s called %s.\n", value.Name, value.Pet.Type, value.Pet.Name);
		fmt.Printf("%s the cat has a power value of %d.\n\n", value.Pet.Name, value.Pet.Power);
	}
}


// Make person super powered.
func super(p *person.Person, number int) {

	// Add 1000 power.
	p.Power += number

	// Increase pet power, as long as it is already 1000 less.
	if p.Pet.Power < (p.Power - number) {
		p.Pet.Power += number
	}
}