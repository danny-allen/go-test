package pet

// Declare the Person structure.
type Pet struct {
	Name string
	Type string
	Power int
}

// Create a new Pet.
func New(n string, t string, p int) *Pet {
	return &Pet {
		Name: n,
		Type: t,
		Power: p,
	}
}