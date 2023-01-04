package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/validator/v10"

	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var uni *ut.UniversalTranslator

func Validate(i interface{}) error {
	en := en.New()
	uni = ut.New(en, en)

	trans, _ := GetTranslator("en")

	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}
		return name
	})

	en_translations.RegisterDefaultTranslations(v, trans)

	if err := v.Struct(i); err != nil {
		return err
	}

	return nil
}

func GetTranslator(lang string) (ut.Translator, error) {
	trans, _ := uni.GetTranslator(lang)
	return trans, nil
}
