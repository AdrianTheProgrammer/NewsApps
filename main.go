package main

import (
	"fmt"
	"newsapps/configs"
	a_rep "newsapps/internal/features/articles/repositories"
	c_rep "newsapps/internal/features/comments/repositories"
	u_rep "newsapps/internal/features/users/repositories"
)

func main() {
	db := configs.ConnectDB()

	var input int
	fmt.Print("Masukkan '1' untuk Migrasi Database: ")
	fmt.Scan(&input)

	if input == 1 {
		err := db.AutoMigrate(&u_rep.Users{}, &c_rep.Comments{}, &a_rep.Articles{})
		if err != nil {
			fmt.Println(err)
		}
	}
}
