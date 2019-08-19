package main

import (
	"ghiac/sample-smpp/signal"
	"time"
)

//	host string
//	username string
//	password string
//	conn *net.TCPConn
//	OnBind OnPduCallback
//	OnSubmit OnPduCallback

func main() {
	temp := NewClient("localhost:7878", "<username>", "<password>")
	temp.Start()
	temp.Send("ss33s", "ss333ss", "sss213213s")
	time.Sleep(2 * time.Millisecond)
	temp.Send("s3123ss22", "ss3213123ss", "ss321312332ss")
	time.Sleep(2 * time.Millisecond)
	temp.Send("ss33s", "s1321sss", "sss312s")
	<-signal.Wait()
}
