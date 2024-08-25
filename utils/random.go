/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package utils

import (
	"errors"
	"math/rand"
	"strings"
	"sync"
)

// SEE https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
// SEE https://golangdocs.com/generate-random-string-in-golang
// SEE https://gosamples.dev/random-string/

type RandomSet string

const (
	RandomSetDigits     RandomSet = "0123456789"
	RandomSetLowerAlpha           = "abcdefghijklmnopqrstuvwxyz"
	RandomSetUpperAlpha           = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	RandomSetAlpha                = RandomSetLowerAlpha + RandomSetUpperAlpha
	RandomSetAlphaNum             = RandomSetDigits + RandomSetAlpha
	RandomSetHex                  = "0123456789abcdefABCDEF"
)

type RandomGenerator struct {
	lock sync.Mutex
	rs   rand.Source
	set  RandomSet
}

var (
	ErrRandomGeneratorNoSetsProvided = errors.New("random char sets should be provided")
)

func NewRandomGenerator(seed int64, sets ...RandomSet) (g *RandomGenerator, err error) {

	if len(sets) == 0 {
		return nil, ErrRandomGeneratorNoSetsProvided
	}

	if seed == 0 {
		seed = rand.Int63()
	}

	// оптимизация на базовый случай - задан 1 набор символов сразу
	superset := sets[0]

	for i := /* !!! */ 1; /* !!! */ i < len(sets); i++ {
		superset += sets[i]
	}

	return &RandomGenerator{
		rs:  rand.NewSource(seed),
		set: superset,
	}, nil
}

func (g *RandomGenerator) RandomString(n int) string {

	// fast-path
	if n <= 0 {
		return ""
	}

	g.lock.Lock()
	defer g.lock.Unlock()

	var b strings.Builder
	b.Grow(n)

	rs := g.rs

	sz := int64(len(g.set))

	for i := 0; i < n; i++ {
		b.WriteByte(g.set[rs.Int63()%sz])
	}

	return b.String()
}
