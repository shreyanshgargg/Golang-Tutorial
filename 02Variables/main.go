package main

import "fmt"

const LoginToken string = "Shreyansh" //public

func main() {
	var username string = "shreyansh"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("variable is of type: %T \n", smallval)

	var smallFloat float32 = 255.32432423
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	//default valuses and some aliases
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("variable is of type: %T \n", anothervariable)

	//implicit type

	var website = "LEARNCODE.IN"
	fmt.Println(website)

	//no var style

	numberofUser := 3004324.4234
	fmt.Println(numberofUser)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)
}
