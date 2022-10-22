package main

import "github.com/aliicoban/store"

func main() {

	//insert user

	u := store.User{
		Name: "Ali Coban",
		Age:  "27",
		Job:  "Software Developer",
	}

	store.InsertUser(u)

	//get users
	//store.GetUsers()

	//get user by id
	//store.GetUserByID(2)

	//update user
	/*u := store.User{
		ID:   2,
		Name: "Ali",
		Age:  "30",
		Job:  "Computer Engineer",
	}
	store.UpdateUser(u)
	*/

	// Delete user
	//store.DeleteUser(6)
}
