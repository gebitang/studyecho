package main

import (
	"flag"
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	//"github.com/go-playground/validator/v10/translations/zh_tw"
	zh_translate "github.com/go-playground/validator/v10/translations/zh"
)

type User struct {
	Name  string `validate:"required"`
	Age   uint   `validate:"gte=1,lte=130"`
	Email string `validate:"required,email"`
}

var (
	name  string
	age   uint
	email string
)

//call before main
func init() {
	flag.StringVar(&name, "name", "", "输入名字")
	flag.UintVar(&age, "age", 0, "输入年龄")
	flag.StringVar(&email, "email", "", "输入邮箱")
}

// https://mp.weixin.qq.com/s/CMBxIR0y6LxiIpDt6JYuDQ
func main() {
	flag.Parse()

	user := &User{
		Name:  name,
		Age:   age,
		Email: email,
	}

	validate := validator.New()

	e := en.New()
	uniTrans := ut.New(e, e, zh.New(), zh_Hant_TW.New())
	translator, _ := uniTrans.GetTranslator("zh")
	_ = zh_translate.RegisterDefaultTranslations(validate, translator)

	err := validate.Struct(user)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, err := range errs {
			fmt.Println(err.Translate(translator))
		}
	}
}
