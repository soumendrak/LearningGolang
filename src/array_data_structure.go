package main

import "fmt"

// func main() {
// fruit := [5]string{}
// fruit[0] = "APPLE"
// fmt.Printf("fruit: %v address: %p\n", fruit[0], &fruit[0])
// fruit[1] = "ORANGE"
// fmt.Printf("fruit: %v address: %p\n", fruit[1], &fruit[1])
// fruit[2] = "KIWI"
// fmt.Printf("fruit: %v address: %p\n", fruit[2], &fruit[2])
// fmt.Printf("fruit: %v address: %p\n", fruit[3], &fruit[3])
// fmt.Printf("fruit: %v address: %p\n", fruit[4], &fruit[4])
// fmt.Printf("fruit: %T address: %p\n", fruit, &fruit)
// }

type user struct {
	likes int
}

func main() {
	users := make([]user, 3)
	shareUser := &users[1]
	shareUser.likes++
	for i := range users {
		fmt.Printf("User: %d Likes: %d\n", i, users[i].likes)
	}
	users = append(users, user{})
	shareUser.likes++
	for i := range users {
		fmt.Printf("User: %d Likes: %d\n", i, users[i].likes)
	}
}
