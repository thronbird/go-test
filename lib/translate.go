package main

import (
	"fmt"
	"github.com/jinzhongmin/gtra"
	"log"
	"sync"
)

func main() {
	var langs = []string{"auto", "af", "sq", "ar", "hy", "az", "eu", "be", "bn", "bs", "bg", "ca", "ceb", "ny", "zh-cn", "zh-tw", "co", "hr", "cs", "da", "nl", "en", "eo", "et", "tl", "fi", "fr", "fy", "gl", "ka", "de", "el", "gu", "ht", "ha", "haw", "iw", "hi", "hmn", "hu", "is", "ig", "id", "ga", "it", "ja", "jw", "kn", "kk", "km", "ko", "ku", "ky", "lo", "la", "lv", "lt", "lb", "mk", "mg", "ms", "ml", "mt", "mi", "mr", "mn", "my", "ne", "no", "ps", "fa", "pl", "pt", "ma", "ro", "ru", "sm", "gd", "sr", "st", "sn", "sd", "si", "sk", "sl", "so", "es", "su", "sw", "sv", "tg", "ta", "te", "th", "tr", "uk", "ur", "uz", "vi", "cy", "xh", "yi", "yo", "zu"}
	var wg sync.WaitGroup
	for _, lang := range langs {
		wg.Add(1)
		lang := lang
		go func() {
			defer func() {
				if err := recover(); err != nil {
					log.Println("panic occurred:", err)
				}
				wg.Done()
			}()
			//var _, res = gtra.Translate("fire", lang)
			var _, res = gtra.Translate("fire", lang, "en")
			//var _, res = gtra.Translate("人", lang,"zh-cn")
			fmt.Println(lang, " : ", res)
		}()
	}
	wg.Wait()

	//t := gtra.NewTranslater()
	//fmt.Println(t.Translate("hello"))
	//fmt.Println(t.Vector(lang.ZHCN, lang.EN).Translate("你好世界"))
	//fmt.Println(t.To(lang.JA).Translate("你好世界"))
	//
	//_, j := t.Vector(lang.EN, lang.ZHCN).Dt("like", "t", "at")
	//fmt.Println(j.String())
}
