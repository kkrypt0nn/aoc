package aoc

import "testing"

func TestReverseString(t *testing.T) {
	name := "Krypton"
	reversed := ReverseString(name)
	expected := "notpyrK"
	if reversed == expected {
		t.Log("Successfully reversed the string")
		return
	}
	t.Errorf("Failed reversing the string (got: %s, expected: %s)", reversed, expected)
}

func TestAtoi(t *testing.T) {
	stringNumber := "1337"
	converted := Atoi(stringNumber)
	expected := 1337
	if converted == expected {
		t.Log("Successfully converted a string to an integer")
		return
	}
	t.Errorf("Failed converting a string to an integer (got: %d, expected: %d)", converted, expected)
}

func TestContainsCharacters(t *testing.T) {
	sentence := "the quick brown fox jumps over the lazy dog"
	characters := "xyz"
	contains := ContainsCharacters(sentence, characters)
	if contains {
		t.Log("Successfully checked if the characters of the substring is contained in the sentence")
		return
	}
	t.Errorf("Failed checking if the characters of the substring is contained in the sentence (got: %t, expected: true)", contains)
}

func TestContainsExactly(t *testing.T) {
	sentence := "thequickbrownfxjmpsvlazydg"
	characters := "abcdefghijklmnopqrstuvwxyz"
	contains := ContainsExactly(sentence, characters)
	if contains {
		t.Log("Successfully checked if the characters are in the sentence, and have the same length")
		return
	}
	t.Errorf("Failed checking if the characters are in the sentence, and have the same length (got: %t, expected: true)", contains)
}
