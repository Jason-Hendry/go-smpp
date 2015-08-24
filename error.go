package go_smpp

import "fmt"

func HandleError(msg string, err error){
	if err != nil {
		fmt.Println("%s: %v", msg, err)
	}
}