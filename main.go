package main

import (
	"fmt"
	"os"
)

type Todo struct {
	id          int32  `json:"id"`
	name        string `json:"name"`
	isCompleted bool   `json:"isCompleted"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo [add|list|complete]")
		return
	}

}
