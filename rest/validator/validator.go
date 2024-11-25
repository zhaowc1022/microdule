package validator

import (
	"errors"
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"log"
	"strings"
)

type (
	Validator interface {
		GetValidator() *validator.Validate
		RegisterValidator(tag, msg string, fn any) error
		FetchValidatorError(q any) error
	}
	ValidatorData struct {
		Validator *validator.Validate
		Trans     ut.Translator
	}
)

// NewValidator 创建gin验证器
func NewValidator(locale string) *ValidatorData {
	validators := &ValidatorData{}

	validators.Validator = validator.New()

	zhT := zh.New() // 中文翻译器
	enT := en.New() // 英文翻译器
	uni := ut.New(enT, zhT, enT)

	// locale 通常取决于 rest 请求头的 'Accept-Language'
	var ok bool
	// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
	validators.Trans, ok = uni.GetTranslator(locale)
	if !ok {
		log.Println(fmt.Errorf("uni.GetTranslator(%s) failed", locale))
		return validators
	}

	// 注册翻译器
	switch locale {
	case "en":
		_ = enTranslations.RegisterDefaultTranslations(validators.Validator, validators.Trans)
	case "zh":
		_ = zhTranslations.RegisterDefaultTranslations(validators.Validator, validators.Trans)
	default:
		_ = enTranslations.RegisterDefaultTranslations(validators.Validator, validators.Trans)
	}

	return validators
}

func (g *ValidatorData) GetValidator() *validator.Validate {
	return g.Validator
}

// FetchValidatorError 获取gin验证器错误
func (g *ValidatorData) FetchValidatorError(q any) error {
	var errs validator.ValidationErrors

	err := g.Validator.Struct(q)
	ok := errors.As(err, &errs)
	if ok {
		result := errs.Translate(g.Trans)
		errs := make([]string, 0)
		for _, s2 := range result {
			errs = append(errs, s2)
		}
		return errors.New(strings.Join(errs, " | "))
	}
	return err
}

func (g *ValidatorData) RegisterValidator(tag, msg string, fn any) error {
	err := g.Validator.RegisterValidation(tag, fn.(func(validator.FieldLevel) bool))
	if err != nil {
		return err
	}

	return g.Validator.RegisterTranslation(tag, g.Trans, registerTranslator(tag, msg), translate)
}

// registerTranslator 为自定义字段添加翻译功能
func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// translate 自定义字段的翻译方法
func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return fe.Field() + msg
}
