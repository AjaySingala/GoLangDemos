package main

import "fmt"

type Person struct {
	Name string
}

// This defines that Android has a Person.
type AndroidFake struct {
	Person Person
	Model  string
}

/*
	This defines that Andriod is a Person.
	Just omit the "name" for Person. Then we can say:
	a := new(Android)
	a.Person.Talk()

	And also: a.Talk()
	Any method of Person can be called on an object of type Android
	because it is an "embedded" or "anonymous" type.
*/
type Android struct {
	Person
	Model string
}

func main() {
	p := Person{"John"}
	p.Talk()

	a := new(Android)
	a.Name = "Mary"
	a.Model = "Robot"
	a.Person.Talk()
	a.Talk()
	fmt.Println("And I am a", a.Model)
}

func (p *Person) Talk() {
	fmt.Println("Hi! My name is", p.Name)
}
