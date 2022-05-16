package iteracao

const qttdRept = 5

func Repit(c string) string {
	var rept string

	for i := 0; i < qttdRept; i++ {
		rept += c
	}

	return rept
}
