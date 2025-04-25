package main

import "fmt"

func main() {
	for abUpp := 'A'; abUpp <= 'z'; abUpp++ {
		fmt.Printf("%c and type:%T\n", abUpp, abUpp)
	}
}
