package main

import (
	"github.com/ghiac/go-smpp"
	"time"
)

func main() {
	token := "token"
	temp := go_smpp.NewClient("localhost:7878", "admin", "admin")
	temp.Start()
	des := "898625337"
	var parameters []go_smpp.Parameter
	parameters = append(parameters, go_smpp.NewParameter(1, []byte(token)))
	time.Sleep(1000 * time.Millisecond)
	temp.Send("userId", des, "msg1", parameters)
	time.Sleep(1000 * time.Millisecond)
	temp.Send("aa", des, "msg2", parameters)
	time.Sleep(1000 * time.Millisecond)
	temp.Send("aa", des, "msg3", parameters)
	//<-signal.Wait()
}
