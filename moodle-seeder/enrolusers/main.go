package main

import (
	"fmt"
	"os"
)

func main() {

	file, _ := os.Create("./enrollments.csv")

	file, err := os.OpenFile("enrollments.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
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
	line := fmt.Sprintf("%s,%s", "username", "course1")
	if _, err := file.WriteString(line); err != nil {
		panic(err)
	}

	for i := 0; i < 1000; i++ {
		line := fmt.Sprintf("\n%s%v,%s", "student", i, "docker101")
		if _, err := file.WriteString(line); err != nil {
			panic(err)
		}
	}

}
