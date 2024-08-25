/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package repository

type entityIndex map[uint]uint // id -> index in {Entity}List (NOT ENTITY ID!!!)

func (ei entityIndex) MaxKey() (max uint) {

	max = 0

	if len(ei) > 0 /* implies `ei != nil` */ {
		for k := range ei {
			if k > max {
				max = k
			}
		}
	}

	return max
}

// inlined
//
//go:nosplit
func makeEntityIndex(sz int) entityIndex {
	return make(entityIndex, sz)
}
