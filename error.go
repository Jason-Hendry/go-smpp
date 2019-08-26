package go_smpp

import "fmt"

func HandleError(msg string, err error) bool {
	if err != nil {
		fmt.Println(msg, " : ", err)
		return false
	}
	return true
}
