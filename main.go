package main

import (
	"fmt"
	"task_manager/config"
)

func main() {
	fmt.Println("hello world")
	config.ConnectDatabase()
}

