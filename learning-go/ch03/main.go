package main

import "fmt"

func exercise1() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	greetings2 := greetings[:2]
	greetings3 := greetings[1:4]
	greetings4 := greetings[3:]

	fmt.Println(greetings)
	fmt.Println(greetings2)
	fmt.Println(greetings3)
	fmt.Println(greetings4)
}

func exercise2() {
	// var message string
	message := "Hi 👨 and 👩"
	// emo := "👩"
	char := message[3]
	fmt.Println(string(char))
}

func exercise3() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}
	person1 := Employee{"Funke", "Akindele", 1}
	person2 := Employee{
		firstName: "Odunlade",
		lastName:  "Adekola",
		id:        2,
	}
	var person3 struct {
		firstName string
		lastName  string
		id        int
	}
	person3.firstName = "Bola"
	person3.lastName = "Are"
	person3.id = 3

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)
}

func main() {
	exercise3()
}
// redo the notification button compoent so i can use as a feature in dashboard tabs. I want the number of fetched data be shown at the top of the tabs just like number 4 on the notification button component