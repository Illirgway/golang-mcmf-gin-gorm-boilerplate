/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package repository

type entityIndex[ID ~uint] map[ID]uint // id -> index in {Entity}List (NOT ENTITY ID!!!)

func (ei entityIndex[ID]) MaxKey() (max ID) {

	max = ID(0)

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
func makeEntityIndex[ID ~uint](sz uint) entityIndex[ID] {
	return make(entityIndex[ID], sz)
}
