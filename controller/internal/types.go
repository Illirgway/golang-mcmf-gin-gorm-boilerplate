/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package internal

import (
	"bytes"
	"strconv"
)

type StringsPluckList []string

var (
	emptyString = []byte{'"', '"'}
)

// MarshalJSON implements json.Marshaler
func (spl StringsPluckList) MarshalJSON() ([]byte, error) {

	if len(spl) == 0 {
		return emptyString, nil
	}

	const (
		quote = '"'
		sep   = ','
	)

	n := 2 // len(`""`)

	for _, s := range spl {
		n += len(s) + 1 // ','
	}

	// да, я знаю, что n на 1 байт больше, чем надо

	var b bytes.Buffer

	b.Grow(n)

	b.WriteByte(quote) // open

	for i, s := range spl {

		if i > 0 {
			b.WriteByte(sep)
		}

		// TODO check if need quote
		if s != "" {
			// TODO use github.com/PurpleSec/Escape
			s = strconv.Quote(s)
			// remove the quotes themselves
			s = s[1 : len(s)-1]
		}

		b.WriteString(s)
	}

	b.WriteByte(quote) // close

	return b.Bytes(), nil
}
