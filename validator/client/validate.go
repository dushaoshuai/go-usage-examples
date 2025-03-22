package validate

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/samber/lo"
)

var (
	// use a single instance of Validate, it caches struct info.
	validate *validator.Validate

	trans    ut.Translator
	zhHansCN = "zh_Hans_CN"
)

func init() {
	// validation

	validate = validator.New(validator.WithRequiredStructEnabled())

	lo.Must0(validate.RegisterValidation(cnmobilephonenumber, CNMobilePhoneNumber))

	// translation

	zh := zh_Hans_CN.New()
	zht := ut.New(zh, zh)
	trans = lo.Must(zht.GetTranslator(zhHansCN))

	lo.Must0(registerTranslations(validate, trans))
	lo.Must0(zh_translations.RegisterDefaultTranslations(validate, trans))

	// custom struct field name parser
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get(zhHansCN)
	})
}

type Validator struct{}

func (v Validator) Validate(_ *http.Request, data any) error {
	err := validate.Struct(data)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok { // 应该不可能
			return err
		}

		var errs error
		for _, fe := range err.(validator.ValidationErrors) {
			errs = errors.Join(errs, fe)
		}
		return errs
	}

	return nil
}
