package main

import (
	"flag"
	"fmt"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
)

var universalTranslator *ut.UniversalTranslator

func main() {
	acceptLanguage := flag.String("language", "zh", "语言")
	flag.Parse()

	e := en.New()
	universalTranslator = ut.New(e, e, zh.New(), zh_Hant_TW.New())

	translator, _ := universalTranslator.GetTranslator(*acceptLanguage)

	switch *acceptLanguage {
	case "zh":
		translator.Add("welcome", "欢迎{0}来到gebitang.com", false)
		translator.AddCardinal("days", "你只剩{0}天时间可以生存", locales.PluralRuleOther, false)
		translator.AddOrdinal("day-of-month", "第{0}天", locales.PluralRuleOther, false)
		translator.AddRange("between", "距离{0}-{1}天", locales.PluralRuleOther, false)
	case "en":
		translator.Add("welcome", "Welcome {0} to gebitang.com", false)
		translator.AddCardinal("days", "You have {0} day left to live", locales.PluralRuleOne, false)
		translator.AddOrdinal("day-of-month", "{0}st", locales.PluralRuleOne, false)
		translator.AddRange("between", "It's {0}-{1} days in range", locales.PluralRuleOther, false)
	}

	fmt.Println(translator.T("welcome", "gebitang"))
	fmt.Println(translator.C("days", 1, 0, translator.FmtNumber(1, 0)))
	fmt.Println(translator.O("day-of-month", 1, 0, translator.FmtNumber(1, 0)))
	fmt.Println(translator.R("between", 1, 0, 2, 0, translator.FmtNumber(1, 0), translator.FmtNumber(2, 0)))
}
