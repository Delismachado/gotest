package iteracao

func Repit(c string) string {
	var rept string
	for i := 0; i < 5; i++ {
		rept = rept + c
	}

	return rept
}
