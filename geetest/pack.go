package geetest

import (
	"./lib"
	"./lib/MHttp"
	"encoding/json"
	"github.com/Mas0nShi/goConsole/console"
	"github.com/tidwall/gjson"
	"regexp"
	"strconv"
	"time"
)
const errorMsg = `{"success": 0, "message": "you should check error."}`
// JsonpParse parse jsonp
func JsonpParse(respStr string) string {
	reg, _ := regexp.Compile(`^.*\((.*?)\)$`)
	return reg.ReplaceAllString(respStr, "$1")
}
// TimeStamp 13 timestamp
func TimeStamp() string {
	return strconv.FormatInt(time.Now().UnixNano() / 1e6, 10)
}

func EmulatorParams(gt string, challenge string, x int, s string) string {
	Random := lib.Random{}
	libp := lib.Params{}
	Crypto := lib.Crypto{}

	tracks, slideX ,times := Random.MoveSlide(x - 9, 5) // TODO: success rate and accuracy. 60~100%
	structBody := struct {
		Userresponse string `json:"userresponse"`
		Passtime int `json:"passtime"`
		Imgload int `json:"imgload"`
		Aa string `json:"aa"`
		Ep struct {
			V string `json:"v"`
		} `json:"ep"`
		Rp string `json:"rp"`
	}{
		Userresponse: libp.Userresponse(slideX, challenge),
		Passtime: times,
		Imgload:  Random.Range(100, 300),
		Aa:       libp.Aa(libp.EncTrack(tracks),[]int{12, 58, 98, 36, 43, 95, 62, 15, 12}, s),
		Ep:		  struct {
			V string `json:"v"`
		}{V: "8.7.8"},
	}
	structBody.Rp = Crypto.Md5(gt + challenge[0:32] + strconv.FormatInt(int64(structBody.Passtime), 10))
	data, _ := json.Marshal(structBody)
	aesKey := Random.AesKey()
	rsaData := Crypto.RsaEncrypt(aesKey, "00c1e3934d1614465b33053e7f48ee4ec87b14b95ef88947713d25eecbff7e74c7977d02dc1d9451f79dd5d1c10c29acb6a9b4d6fb7d0a0279b6719e1772565f09af627715919221aef91899cae08c0d686d748b20a3603be2318ca6bc2b59706592a9219d0bf05c9f65023a21d2330807252ae0066d59ceefa5f2748ea80bab81")
	aesData := Crypto.AesEncrypt(data, aesKey)
	return aesData + rsaData
}

func GetPass(gt string, challenge string) string {
	var (
		url     = ""
		ret     = ""
		headers = map[string]string {
			"User-agent": "Mozilla/5.0 (Linux; U; Android 8.1.0; zh-cn; BLA-AL00 Build/HUAWEIBLA-AL00) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/57.0.2987.132 MQQBrowser/8.9 Mobile Safari/537.36",
			"Content-Type": "text/javascript;charset=UTF-8",
		}
	)
	http := &MHttp.MHttp{}

	url = "https://api.geetest.com/gettype.php?gt=" + gt + "&lang=zh-cn&pt=3&client_type=web_mobile&w=&callback=geetest_" + TimeStamp()
	ret = http.Get(url,&headers,nil).GetResponseText()

	url = "https://api.geetest.com/ajax.php?gt=" + gt + "&challenge="+ challenge + "&lang=zh-cn&pt=3&client_type=web_mobile&w=&callback=geetest_" + TimeStamp() // TODO: the init param w: coming soon 🙂
	ret = http.Get(url,&headers,nil).GetResponseText()

	ret = JsonpParse(ret)
	if gjson.Get(ret, "data.result").String() != "slide" {
		console.Warn(ret)
		return errorMsg
	}

	url = "https://api.geetest.com/get.php?is_next=true&type=slide3&gt=" + gt + "&challenge=" + challenge + "&lang=zh-cn&https=false&protocol=https%3A%2F%2F&offline=false&product=embed&api_server=api.geetest.com&isPC=true&autoReset=true&width=100%25&callback=geetest_" + TimeStamp()
	ret = http.Get(url,&headers,nil).GetResponseText()

	jp := gjson.Parse(JsonpParse(ret))
	s := jp.Get("s").String()
	challenge = jp.Get("challenge").String()

	fullbgUrl := "https://static.geetest.com/"  + jp.Get("fullbg").String()
	bgUrl := "https://static.geetest.com/"  + jp.Get("bg").String()

	fullbgBytes := http.Get(fullbgUrl, &headers, nil).GetResponseBody()
	bgBytes := http.Get(bgUrl, &headers, nil).GetResponseBody()

	fullbg, err := lib.RefuseImage(fullbgBytes) // refuse picture
	bg, err := lib.RefuseImage(bgBytes) // refuse picture
	if fullbg == nil || bg == nil {
		console.Error(err)
		return errorMsg
	}

	x, _ := lib.CompareOcr(&bg, &fullbg) // ocr x

	w := EmulatorParams(gt, challenge, x, s)
	url = "https://api.geetest.com/ajax.php?gt=" + gt + "&challenge=" + challenge + "&lang=zh-cn&pt=3&client_type=web_mobile&w=" + w + "&callback=geetest_" + TimeStamp()

	ret = http.Get(url, &headers, nil).GetResponseText()
	return JsonpParse(ret)
}