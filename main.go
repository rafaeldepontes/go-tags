package main

type User struct {
	Name string `validate:"min=3,max=32"`

	// FIXME: I don't personally like email validation by some regex...
	//
	// Not only its a bad thing, but it can also break my whole system
	// with a big enough regex! So I'm not gonna use it here.
	Email string `validate:"required"`
}

func main() {

}
