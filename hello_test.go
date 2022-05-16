package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	verifyMsg := func(t *testing.T, result, exp string) {
		t.Helper()
		if result != exp {
			t.Errorf("res '%s', esp '%s'", result, exp)
		}

	}

	t.Run("Say hi to people", func(t *testing.T) {
		result := Hello("Oi anyone!", "")
		exp := "Oi anyone!"
		verifyMsg(t, result, exp)
	})

	t.Run("Say, Hi any, when has empty string", func(t *testing.T) {
		result := Hello("", "")
		exp := "Any"
		verifyMsg(t, result, exp)
	})

	t.Run("En espanol", func(t *testing.T) {
		result := Hello("Fiona", "espanol")
		exp := "Hola, Fiona"
		verifyMsg(t, result, exp)
	})

	t.Run("In french", func(t *testing.T) {
		result := Hello("Fiona", "french")
		exp := "Bonjour, Fiona"
		verifyMsg(t, result, exp)
	})

	t.Run("In Eng", func(t *testing.T) {
		result := Hello("Fiona", "eng")
		exp := "Hi, Fiona"
		verifyMsg(t, result, exp)
	})

}
