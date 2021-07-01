package main

import (
	"./geetest"
	"./geetest/lib/MHttp"
	"github.com/Mas0nShi/goConsole/console"
	"github.com/tidwall/gjson"
)

func main() {
	http := &MHttp.MHttp{}
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36 Edg/91.0.864.54",
	}

	for i := 0; i < 10; i++ {
		url := "https://www.geetest.com/demo/gt/register-slide?t=" + geetest.TimeStamp()
		ret := http.Get(url, &headers, nil).GetResponseText()
		t := gjson.GetMany(ret, "gt", "challenge")
		gt, challenge:= t[0].String(), t[1].String()
		ret = geetest.GetPass(gt, challenge)
		console.Log("ajax => ", ret)
	}
}