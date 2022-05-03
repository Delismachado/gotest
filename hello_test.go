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
		result := Hello("anyone!")
		exp := "Hi anyone!"
		verifyMsg(t, result, exp)
	})

	t.Run("Say, Hi any, when has empty string", func(t *testing.T) {
		result := Hello("")
		exp := "Hi Any"
		verifyMsg(t, result, exp)
	})

}
