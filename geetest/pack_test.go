package geetest

import (
	"./lib/MHttp"
	"github.com/Mas0nShi/goConsole/console"
	"github.com/tidwall/gjson"
)
import "testing"
func getpass(gt string, challenge string) string {
	return GetPass(gt, challenge)
}

func BenchmarkGetPass(b *testing.B) {
	http := &MHttp.MHttp{}
	headers := map[string]string{
		"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36 Edg/91.0.864.54",
	}
	url := "https://passport.bilibili.com/x/passport-login/captcha?source=main_web"

	ret := http.Get(url, &headers, nil).GetResponseText()
	//console.Log(ret)
	gt := "7fa7d480550df273db851dcb2b04babf"
	challenge := gjson.Get(ret, "data.geetest.challenge").String()
	ret = getpass(gt, challenge)
	console.Log("ajax => ", ret)
}


func Test(t *testing.T) {
	http := &MHttp.MHttp{}
	headers := map[string]string{
		"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36 Edg/91.0.864.54",
	}
	url := "https://passport.bilibili.com/x/passport-login/captcha?source=main_web"

	ret := http.Get(url, &headers, nil).GetResponseText()
	//console.Log(ret)
	gt := "7fa7d480550df273db851dcb2b04babf"
	challenge := gjson.Get(ret, "data.geetest.challenge").String()
	ret = getpass(gt, challenge)
	console.Log("ajax => ", ret)
}