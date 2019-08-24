package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func SendMessage(phone string, text string, token string) {
	url1 := "https://tapi.bale.ai/" + token + "/SendMessage"
	resp, err := http.PostForm(url1, url.Values{"chat_id": {phone}, "text": {text}})
	if nil != err {
		fmt.Println("HttpRequest Error")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		fmt.Println("HttpRequest Error")
		return
	}
	fmt.Println("Body: ", string(body))
}
