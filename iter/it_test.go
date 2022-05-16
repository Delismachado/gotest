package iteracao

import "testing"

func TestRepet(t *testing.T) {
	rep := Repit("a")
	exp := "aaaaa"

	if rep != exp {
		t.Errorf("esperado '%s' mas obteve '%s'", exp, rep)
	}
}

func BenchmarRepit(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Repit("a")
	}

}
