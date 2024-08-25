/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// string consist of only valid utf-8 chars (any valid char is allowed)
//
//go:nosplit
func validateUtf8(fl validator.FieldLevel) bool {
	return utf8.ValidString(fl.Field().String())
}

// string consist of only valid AND printable chars
func validateUtf8String(fl validator.FieldLevel) bool {

	s := fl.Field().String()

	for i := 0; i < len(s); {

		r, w := utf8.DecodeRuneInString(s[i:])

		if (r == utf8.RuneError) || !unicode.IsPrint(r) {
			return false
		}

		i += w
	}

	return true
}

// NOTE \t \r \n are included
func validateUtf8Text(fl validator.FieldLevel) bool {

	s := fl.Field().String()

	for i := uint(0); /* BCE hint */ i < uint(len(s)); {

		r, w := utf8.DecodeRuneInString(s[i:])

		if (r == utf8.RuneError) || ((r != '\n') && (r != '\r') && (r != '\t') && !unicode.IsPrint(r)) {
			return false
		}

		i += uint(w)
	}

	return true
}

func validateListGenericUints(fl validator.FieldLevel, sep string) bool {

	list := fl.Field().String()

	// fast-check
	if len(list) == 0 {
		return false
	}

	for {
		var (
			num   string
			found bool
		)

		num, list, found = strings.Cut(list, sep)

		// empty num
		if num == "" {
			return false
		}

		if _, err := strconv.ParseUint(num, 10, 0); err != nil {
			// fail, not number
			return false
		}

		// its last part
		if !found {
			// success
			return true
		}
	}

	return false
}

// `^[0-9]+(?:;[0-9]+)*$`
func validateListSimpleNums(fl validator.FieldLevel) bool {
	return validateListGenericUints(fl, ";")
}

// `^[0-9]+(?:,[0-9]+)*$`
func validateListSimpleIds(fl validator.FieldLevel) bool {
	return validateListGenericUints(fl, ",")
}

func validateListGenericTokens(fl validator.FieldLevel, sep string, rxToken *regexp.Regexp) bool {

	list := fl.Field().String()

	// fast-check
	if len(list) == 0 {
		return false
	}

	for {
		var (
			token string
			found bool
		)

		token, list, found = strings.Cut(list, sep)

		// empty token
		if token == "" {
			return false
		}

		if !rxToken.MatchString(token) {
			return false
		}

		// its last part
		if !found {
			// success
			return true
		}
	}

	return false

}

var rxVldTokenWeak = regexp.MustCompile(`^(?:[a-z](?:[a-z0-9_\-]*[a-z0-9])?\.)*[a-z](?:[a-z0-9_\-]*[a-z0-9])?$`)

func validateListSimpleTokens(fl validator.FieldLevel) bool {
	return validateListGenericTokens(fl, ";", rxVldTokenWeak)
}

type validators map[string]validator.Func

func registerValidators(vs validators) (err error) {

	e := binding.Validator.Engine()

	ve, ok := e.(*validator.Validate)

	if !ok {
		return fmt.Errorf("wrong binding.Validator Engine type: want validator.Validate, got %T", e)
	}

	for tag, fn := range vs {

		if err = ve.RegisterValidation(tag, fn); err != nil {
			return fmt.Errorf("validator %s registration error: %w", tag, err)
		}
	}

	return nil
}

/**
 *	TODO
	- utf8, utf8-string, utf8-text - PR в go-validator репу
	- предложить универсальный интерфейс валидации для собственных типов значений

	type Validator interface {
		Valid() bool
	}
*/

func InitValidators() error {

	vs := validators{
		"utf8":               validateUtf8,
		"utf8-string":        validateUtf8String,
		"utf8-text":          validateUtf8Text,
		"list-simple-nums":   validateListSimpleNums,
		"list-simple-ids":    validateListSimpleIds,
		"list-simple-tokens": validateListSimpleTokens,
	}

	return registerValidators(vs)
}
