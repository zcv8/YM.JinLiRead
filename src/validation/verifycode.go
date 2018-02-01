package validation

import (
	"encoding/json"
	"github.com/mojocn/base64Captcha"
	"log"
	"net/http"
)

type ConfigJsonBody struct {
	Id              string
	CaptchaType     string
	VerifyValue     string
	ConfigAudio     base64Captcha.ConfigAudio
	ConfigCharacter base64Captcha.ConfigCharacter
	ConfigDigit     base64Captcha.ConfigDigit
}

func GenerateCaptchaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var postParameters = ConfigJsonBody{
		CaptchaType: "character",
		ConfigCharacter: base64Captcha.ConfigCharacter{
			Height: 50,
			Width:  138,
			//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
			Mode:               base64Captcha.CaptchaModeNumber,
			ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
			ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
			IsShowHollowLine:   false,
			IsShowNoiseDot:     false,
			IsShowNoiseText:    false,
			IsShowSlimeLine:    false,
			IsShowSineLine:     false,
			CaptchaLen:         4,
		},
	}
	//create base64 encoding captcha
	//创建base64图像验证码

	var config interface{}
	switch postParameters.CaptchaType {
	case "audio":
		config = postParameters.ConfigAudio
	case "character":
		config = postParameters.ConfigCharacter
	default:
		config = postParameters.ConfigDigit
	}
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	captchaId, digitCap := base64Captcha.GenerateCaptcha("", config)
	base64Png := base64Captcha.CaptchaWriteToBase64Encoding(digitCap)

	//or you can do this
	//你也可以是用默认参数 生成图像验证码
	//base64Png := captcha.GenerateCaptchaPngBase64StringDefault(captchaId)

	//set json response
	//设置json响应

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body := map[string]interface{}{"code": 1, "data": base64Png, "captchaId": captchaId, "msg": "success"}
	json.NewEncoder(w).Encode(body)

}

// base64Captcha verify http handler
func CaptchaVerifyHandler(w http.ResponseWriter, r *http.Request) {

	//parse request parameters
	//接收客户端发送来的请求参数
	decoder := json.NewDecoder(r.Body)
	var postParameters ConfigJsonBody
	err := decoder.Decode(&postParameters)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	//verify the captcha
	//比较图像验证码
	verifyResult := base64Captcha.VerifyCaptcha(postParameters.Id, postParameters.VerifyValue)

	//set json response
	//设置json响应
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body := map[string]interface{}{"code": "error", "data": "验证失败", "msg": "captcha failed"}
	if verifyResult {
		body = map[string]interface{}{"code": "success", "data": "验证通过", "msg": "captcha verified"}
	}
	json.NewEncoder(w).Encode(body)
}
