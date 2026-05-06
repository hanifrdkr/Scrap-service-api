// Package validatorx
package validatorx

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"helicopter-hr/internal/app_rest/util"
	"helicopter-hr/pkg/ginx"
	"regexp"
	"time"
)

const (
	AlphaNumericDash     = `^[0-9a-zA-Z\-]+$`
	AlphaNumeric         = `^[0-9a-zA-Z]+$`
	Numeric              = `^[0-9]+`
	AlphaNumericSpace    = `^[0-9a-zA-Z\s]+$`
	Alpha                = `^[a-zA-Z]+$`
	AlphaSpace           = `^[a-zA-Z\s]+$`
	AlphaDashSpace       = `^[0-9a-zA-Z\-\s]+$`
	IndonesianPeopleName = `^[a-zA-Z\'’.,\s]+$`
	RtRw                 = `^\d{1,3}\/\d{1,3}$`
	SubDistrict          = `^[0-9a-zA-Z\-\s\(\)]+$`
	Address              = `^[A-Za-z0-9'\.\-\s\,/#_()\[\]]+$`
	Pob                  = `^[A-Za-z'\.\-\s\,/#_()\[\]]+$`
	// LayoutDateTimeFormat date time format
	LayoutDateTimeFormat     = `2006-01-02 15:04:05`
	LayoutDateFormatYYYYMMDD = `2006-01-02`
	Email                    = `^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`
	UUID                     = `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`
)

var (
	AlphabetNumericSpaceChar     = regexp.MustCompile(`^[\w\s\(\)\-\+\,\.\!\?\/\\]+$`)
	AlphabetNumericSpaceCharRule = validation.Match(AlphabetNumericSpaceChar).Error(`must be among or combination these characters (a-z, A-Z, 0-9, space, enter, tab, comma(,), dot(.), slash and back slash(\/), question mark(?), exclamation mark(!), underscore(_), plus and minus(-+))`)
)

// ValidateAlphaNumericDash for reference transaction id
func validateAlphaNumericDash(v string) bool {
	pattern := `^[0-9a-zA-Z\-]+$`

	rgx, err := regexp.Compile(pattern)

	if err != nil {
		return false
	}

	return rgx.MatchString(v)
}

// validateAlphaNumericSpaceCommaDot for reference transaction id
func validateAlphaNumericSpaceCommaDot(v string) bool {
	pattern := `^[0-9a-zA-Z\.,\s]+$`

	rgx, err := regexp.Compile(pattern)

	if err != nil {
		return false
	}

	return rgx.MatchString(v)
}

// validateAlphaNumericSpaceCommaDotDash for reference transaction id
func validateAlphaNumericSpaceCommaDotDash(v string) bool {
	pattern := `^[a-zA-Z0-9\s,.-]+$`

	rgx, err := regexp.Compile(pattern)

	if err != nil {
		return false
	}

	return rgx.MatchString(v)
}

// validateDecimal for reference transaction id
func validateDecimal(v string) bool {
	pattern := `^[0-9]\d*(\.\d+)?$`

	rgx, err := regexp.Compile(pattern)

	if err != nil {
		return false
	}

	return rgx.MatchString(v)
}

// validateNumeric for reference transaction id
func validateNumeric(v string) bool {
	pattern := `^[0-9]*$`

	rgx, err := regexp.Compile(pattern)

	if err != nil {
		return false
	}

	return rgx.MatchString(v)
}

// validateEmail for validate email.
func validateEmail(v string) bool {
	rgx, err := regexp.Compile(Email)

	if err != nil {
		return false
	}

	return rgx.MatchString(v)
}

// validateUUID for validate uuid.
func validateUUID(v string) bool {
	rgx, err := regexp.Compile(UUID)

	if err != nil {
		return false
	}

	return rgx.MatchString(v)
}

// validateLength for length of character.
func validateLength(min, max int) func(v string) bool {
	return func(v string) bool {
		if len(v) < min {
			return false
		}
		if len(v) > max {
			return false
		}
		return true
	}
}

// validateValue for value of number.
func validateValue(min, max int) func(v int) bool {
	return func(v int) bool {
		if v < min {
			return false
		}
		if v > max {
			return false
		}
		return true
	}
}

func Regex(pattern string) func(v string) bool {
	return func(v string) bool {
		if len(v) == 0 {
			return true
		}
		rgx, err := regexp.Compile(pattern)

		if err != nil {
			return false
		}

		return rgx.MatchString(v)
	}
}

func validDOB(v string) bool {
	var (
		f  = []string{`2006-01-02`, `02-01-2006`}
		tm time.Time
	)

	if len(v) == 0 {
		return true
	}

	for i := 0; i < len(f); i++ {
		t, err := time.Parse(f[i], v)

		if err != nil && i == 0 {
			continue
		}

		if err != nil {
			return false
		}

		tm = t
		break
	}

	if tm.After(time.Now().AddDate(-15, 0, 0)) {
		return false
	}

	return true
}

func validDateTime(v string) bool {
	if len(v) == 0 {
		return true
	}

	_, err := time.Parse(LayoutDateTimeFormat, v)
	if err != nil {
		return false
	}

	return true
}

func validDateFormatYYYYMMDD(v string) bool {
	if len(v) == 0 {
		return true
	}

	_, err := time.Parse(LayoutDateFormatYYYYMMDD, v)
	if err != nil {
		return false
	}

	return true
}

func validIn(in []string) func(string) bool {
	return func(v string) bool {
		return util.InArray(v, in)
	}
}

func ValidAlphaNumericDash() validation.StringRule {
	return validation.NewStringRuleWithError(
		validateAlphaNumericDash,
		validation.NewError("validation_is_alphanumeric", "must contain alpha, digits and dash only"))
}

func ValidAlphaNumericSpaceCommaDot() validation.StringRule {
	return validation.NewStringRuleWithError(
		validateAlphaNumericSpaceCommaDot,
		validation.NewError("validation_is_alphanumeric", "must contain alpha, digits, space, comma(,) and dot(.) only"))
}

func ValidAlphaNumericSpaceCommaDotDash() validation.StringRule {
	return validation.NewStringRuleWithError(
		validateAlphaNumericSpaceCommaDotDash,
		validation.NewError("validation_is_alphanumeric", "must contain alpha, digits, space, comma(,), dash(-) and dot(.) only"))
}

func ValidDecimal() validation.StringRule {
	return validation.NewStringRuleWithError(
		validateDecimal,
		validation.NewError("validation_is_decimal", "must contain digits decimal"))
}

func ValidNumeric() validation.StringRule {
	return validation.NewStringRuleWithError(
		validateNumeric,
		validation.NewError("validation_is_numeric", "must contain digits"))
}

func ValidDOB() validation.StringRule {
	return validation.NewStringRuleWithError(
		validDOB,
		validation.NewError("validation_is_dob", "must be valid date of bird YYYY-MM-DD or DD-MM-YYYY"))
}

func ValidRegex(pattern string) validation.StringRule {
	return validation.NewStringRuleWithError(
		Regex(pattern),
		validation.NewError("validation_regex", fmt.Sprintf("must be valid regex %s", util.SubstringAfter(pattern, "^"))))
}

func ValidLength(min, max int) validation.StringRule {
	return validation.NewStringRuleWithError(
		validateLength(min, max),
		validation.NewError("validation_regex", fmt.Sprintf("must be at least %d character and no more than %d character", min, max)))
}

func ValidIn(in []string) validation.StringRule {
	return validation.NewStringRuleWithError(
		validIn(in),
		validation.NewError("validation_is_in", fmt.Sprintf("must be one of : %s", util.StringJoin(in, ",", ""))))
}

func ValidDateTime() validation.StringRule {
	return validation.NewStringRuleWithError(
		validDateTime,
		validation.NewError("validation_is_date_time", fmt.Sprintf("must be valid date fromat: YYYY-MM-DD H:m:s")))
}

func ValidDateFormatYYYYMMDD() validation.StringRule {
	return validation.NewStringRuleWithError(
		validDateFormatYYYYMMDD,
		validation.NewError("validation_is_date", fmt.Sprintf("must be valid date fromat: YYYY-MM-DD")))
}

func ValidEmail() validation.StringRule {
	return validation.NewStringRuleWithError(
		validateEmail,
		validation.NewError("validation_is_email", fmt.Sprintf("invalid email format")))
}

func ValidUUID() validation.StringRule {
	return validation.NewStringRuleWithError(
		validateUUID,
		validation.NewError("validation_is_uuid", fmt.Sprintf("invalid uuid")))
}

func Required() validation.RequiredRule {
	return validation.Required
}

// Min returns a validation rule that checks if a value is greater or equal than the specified value.
// By calling Exclusive, the rule will check if the value is strictly greater than the specified value.
// Note that the value being checked and the threshold value must be of the same type.
// Only int, uint, float and time.Time types are supported.
// An empty value is considered valid. Please use the Required rule to make sure a value is not empty.
func Min(min interface{}) validation.ThresholdRule {
	return validation.Min(min)

}

// Max returns a validation rule that checks if a value is less or equal than the specified value.
// By calling Exclusive, the rule will check if the value is strictly less than the specified value.
// Note that the value being checked and the threshold value must be of the same type.
// Only int, uint, float and time.Time types are supported.
// An empty value is considered valid. Please use the Required rule to make sure a value is not empty.
func Max(max interface{}) validation.ThresholdRule {
	return validation.Max(max)
}

type ErrorViolation struct {
	Field string
	Error error
}
type ErrorViolations []ErrorViolation

func New() *ErrorViolations {
	return &ErrorViolations{}
}

func (e ErrorViolations) Set(field string, err error) *ErrorViolations {
	if err != nil {
		e = append(e, ErrorViolation{
			Field: field,
			Error: err,
		})
	}
	return &e
}

func (e *ErrorViolations) Apply() []ginx.ErrorField {
	var violations []ginx.ErrorField
	if e != nil {
		for _, value := range *e {
			violations = append(violations, ginx.ErrorField{
				Field: value.Field,
				Error: value.Error.Error(),
			})
		}

	}

	return violations
}
