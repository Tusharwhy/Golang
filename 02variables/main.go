package main

import "fmt"

// we use initial capital which denotes this variable is publically availaible
const LoginToken string = "nbhegdbdgeirgbek"

func main() {
	fmt.Println("--variables--")

	var username string = "Tushar"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var userID string
	fmt.Println(userID)

	//implicit type
	var web = "hello.com"
	fmt.Println(web)

	//no var style
	//not allowed to use this walrus operator outside the method. like global
	numberOfUsers := 5000
	fmt.Println(numberOfUsers)

	//getting global variable
	fmt.Println(LoginToken)

}
