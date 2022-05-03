package main

import "testing"

func TestOla(t *testing.T) {
	result := Ola("Ser!")
	esp := "Olá Ser!"

	if result != esp {
		t.Errorf("res '%s', esp '%s'", result, esp)
	}

}
