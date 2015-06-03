package main

import (
	"fmt"
	"github.com/ibigbug/weibo2"
	"net/url"
)

func main() {
	c := &weibo.Client{}
	c.ApiKey = "YOUR KEY"
	c.ApiSecret = "YOUR SECRET"
	c.RedirectUri = "FIND IT IN YOUR APP SETTINGS"

	fmt.Println(c.GetAuthUrl())

	code := ""

	fmt.Println("Enter code:")
	fmt.Scanln(&code)

	fmt.Println(c.GetAccessToken(code))

	fmt.Println(c.Get("statuses/public_timeline", url.Values{}))

	fmt.Println(c.Post("statuses/update", url.Values{"status": []string{"test"}}))
}
