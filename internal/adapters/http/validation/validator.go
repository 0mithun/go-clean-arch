package validation

import (
	"reflect"
	"strings"

	en "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

type MessagesProvider interface {
	Messages() map[string]string
}

type ErrorResponse struct {
	Message string              `json:"message"`
	Errors  map[string][]string `json:"errors"`
}

var (
	validate = validator.New()
	trans    ut.Translator
)

func init() {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)

	translator, found := uni.GetTranslator("en")
	if !found {
		panic("validator translator not found")
	}

	if err := enTranslations.RegisterDefaultTranslations(validate, translator); err != nil {
		panic(err)
	}

	trans = translator
}

func ValidateStruct(payload any) (map[string][]string, error) {
	err := validate.Struct(payload)
	if err == nil {
		return nil, nil
	}

	validationErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil, err
	}

	return formatValidationErrors(validationErrs, messagesFor(payload)), nil
}

func NewErrorResponse(errors map[string][]string) ErrorResponse {
	return ErrorResponse{
		Message: "validation failed",
		Errors:  errors,
	}
}

func formatValidationErrors(errs validator.ValidationErrors, customMessages map[string]string) map[string][]string {
	result := make(map[string][]string, len(errs))

	for _, fieldErr := range errs {
		field := fieldErr.Field()
		if field == "" {
			field = strings.ToLower(fieldErr.StructField())
		}

		result[field] = append(result[field], resolveValidationMessage(field, fieldErr, customMessages))
	}

	return result
}

func messagesFor(payload any) map[string]string {
	if payload == nil {
		return nil
	}

	if provider, ok := payload.(MessagesProvider); ok {
		return provider.Messages()
	}

	return nil
}

func resolveValidationMessage(field string, fieldErr validator.FieldError, customMessages map[string]string) string {
	if len(customMessages) == 0 {
		return fieldErr.Translate(trans)
	}

	fieldRuleKey := field + "." + fieldErr.Tag()
	if msg, ok := customMessages[fieldRuleKey]; ok {
		return msg
	}

	anyFieldRuleKey := "*." + fieldErr.Tag()
	if msg, ok := customMessages[anyFieldRuleKey]; ok {
		return msg
	}

	return fieldErr.Translate(trans)
}
