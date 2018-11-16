// DEN
// Copyright (C) 2018 Andreas T Jonsson

package page437

import (
	"fmt"
	"testing"
	"unicode"
)

func TestPrintCodePage(t *testing.T) {
	t.SkipNow() //Skip this test for now.

	for i, c := range codePage {
		if unicode.IsPrint(c) {
			fmt.Printf("0x%X: %c\n", i, c)
		} else {
			fmt.Printf("0x%X: Not printable\n", i)
		}
	}
}

func TestASCII(t *testing.T) {
	for i := 32; i < 127; i++ {
		c := codePage[i]
		if int(c) != i {
			t.Logf("ASCII character missmatch: %d: %c", i, c)
			t.Fail()
		}
	}
}

func TestDuplicates(t *testing.T) {
	for i, a := range codePage {
		for j, b := range codePage {
			if a == b && i != j {
				t.Logf("Character is duplicated: %d == %d, 0x%X: %c", i, j, a, b)
				t.Fail()
			}
		}
	}
}
