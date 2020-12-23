package main

import (
	"fmt"
	"os"
)

func main() {

	file, _ := os.Create("./users.csv")

	file, err := os.OpenFile("users.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	// close fi on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	// write header
	line := fmt.Sprintf("%s,%s,%s,%s,%s", "username", "firstname", "lastname", "password", "email")
	if _, err := file.WriteString(line); err != nil {
		panic(err)
	}

	for i := 0; i < 1000; i++ {
		line := fmt.Sprintf("\n%s%v,%s%v,%s,%s,%s%v@mdlcontainer.must.edu.eg", "student", i, "student", i, ".", "123", "student", i)
		if _, err := file.WriteString(line); err != nil {
			panic(err)
		}
	}

}
