package main

import (
	"fmt"
	"github.com/jinzhongmin/gtra"
)

func main() {
	var langs = []string{"auto", "af", "sq", "ar", "hy", "az", "eu", "be", "bn", "bs", "bg", "ca", "ceb", "ny", "zh-cn", "zh-tw", "co", "hr", "cs", "da", "nl", "en", "eo", "et", "tl", "fi", "fr", "fy", "gl", "ka", "de", "el", "gu", "ht", "ha", "haw", "iw", "hi", "hmn", "hu", "is", "ig", "id", "ga", "it", "ja", "jw", "kn", "kk", "km", "ko", "ku", "ky", "lo", "la", "lv", "lt", "lb", "mk", "mg", "ms", "ml", "mt", "mi", "mr", "mn", "my", "ne", "no", "ps", "fa", "pl", "pt", "ma", "ro", "ru", "sm", "gd", "sr", "st", "sn", "sd", "si", "sk", "sl", "so", "es", "su", "sw", "sv", "tg", "ta", "te", "th", "tr", "uk", "ur", "uz", "vi", "cy", "xh", "yi", "yo", "zu"}
	for _, lang := range langs {
		var _, res = gtra.Translate("to be killed", lang)
		fmt.Println(res)
	}
	//t := gtra.NewTranslater()
	//fmt.Println(t.Translate("hello"))
	//fmt.Println(t.Vector(lang.ZHCN, lang.EN).Translate("你好世界"))
	//fmt.Println(t.To(lang.JA).Translate("你好世界"))
	//
	//_, j := t.Vector(lang.EN, lang.ZHCN).Dt("like", "t", "at")
	//fmt.Println(j.String())
}
