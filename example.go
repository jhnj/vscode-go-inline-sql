package main

import "fmt"

func main() {
	sql := `--sql
		SELECT * FROM users;
	`
	fmt.Println(sql)
}
