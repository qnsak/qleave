package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh_tw"
)

var (
	uni   *ut.UniversalTranslator
	trans ut.Translator
)

func Init() {
	zh := zh_Hant_TW.New()
	uni = ut.New(zh)
	trans, _ = uni.GetTranslator("zh_Hant_TW")
	v := binding.Validator.Engine().(*validator.Validate)
	zh_translations.RegisterDefaultTranslations(v, trans)
}

// 驗證文本轉換
func Translate(err error) map[string][]string {
	var result = make(map[string][]string)

	errors := err.(validator.ValidationErrors)

	for _, err := range errors {
		result[err.Field()] = append(result[err.Field()], err.Translate(trans))
	}

	return result
}
