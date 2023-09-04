package perror

import "fmt"

// function to handle error
func ErrorLog(line int, message string) {
	Report(line, "", message)
}

// format the error to report to user
func Report(line int, where string, message string) {
	fmt.Printf("Line %d] Error %s : %s", line, where, message)
}

// utils to panic if you find that err is not nill
// so the function that you call prev had an error
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
