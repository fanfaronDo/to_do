package main

import (
	"fmt"
	"os"
)

func main() {
	//cfg, _ := config.ConfigLoad()
	fmt.Println(os.Getenv("POSTGRES_PASSWORD"))

	//for _, env := range os.Environ() {
	//	fmt.Println(env)
	//}
}
