package main

import "fmt"

func main() {

	originalString := "hello world"

	smth := []rune(originalString)
	var rev []rune

	for i := len(smth) - 1; i >= 0; i-- {
		rev = append(rev, smth[i])
	}

	fmt.Println(string(rev))
}
