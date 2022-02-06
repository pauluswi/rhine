package validator

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	ModeDefault = "default"
	ModeCompact = "compact"
	ModeVerbose = "verbose"
)

var (
	validateinst *validator.Validate
)

func GetGlobal() *validator.Validate {
	if validateinst == nil {
		validateinst = validator.New()
	}

	return validateinst
}

func SetGlobal(v *validator.Validate) {
	validateinst = v
}

// ***

type Error struct {
	origin error
	mode   string
	loc    string
}

func (e Error) Error() (out string) {
	// check if valid validation error
	verr, ok := e.origin.(validator.ValidationErrors)
	if !ok {
		return e.origin.Error()
	}

	// build message
	switch e.mode {
	case ModeCompact:
		fields := []string{}
		for _, q := range verr {
			fields = append(fields, q.Field())
		}
		out += fmt.Sprintf("invalid fields: %s", strings.Join(fields, ", "))

	case ModeVerbose:
		out += fmt.Sprintf("invalid fields at %s: ", e.loc)
		fallthrough

	default:
		errs := []string{}
		for _, q := range verr {
			msg := ""

			msg += q.StructNamespace()
			msg += " must be " + q.ActualTag()

			// Print condition parameters, e.g. oneof=red blue -> { red blue }
			if q.Param() != "" {
				msg += "{" + q.Param() + "}"
			}

			if q.Value() != nil && q.Value() != "" {
				msg += fmt.Sprintf(", actual value is %v", q.Value())
			}

			errs = append(errs, msg)
		}

		out += strings.Join(errs, "; ")
	}

	return
}

func (e Error) Unwrap() error {
	return e.origin
}

// ***

type Opts struct {
	Mode string
}

func Validate(s interface{}) (err error) {
	verr := GetGlobal().Struct(s)
	if verr == nil {
		return
	}
	return Error{origin: verr, mode: ModeDefault, loc: getCallerFuncName()}
}

func ValidateWithOpts(s interface{}, opt Opts) (err error) {
	verr := GetGlobal().Struct(s)
	if verr == nil {
		return
	}
	return Error{origin: verr, mode: opt.Mode, loc: getCallerFuncName()}
}

// ***

func getCallerFuncName() (fname string) {
	pc, _, _, ok := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		f, l := details.FileLine(pc)
		fname = fmt.Sprintf("%s:%d", f, l)
	}
	return
}
