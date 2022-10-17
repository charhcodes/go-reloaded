for l := 0; l < len(emptyString); l++ {
	if emptyString[i] == "a" {
		if emptyString[i+1] == " " {
			if emptyString[i+2] == "a" || emptyString[i+2] == "e" || emptyString[i+2] == "i" || emptyString[i+2] == "o" || emptyString[i+2] == "u" ||  emptyString[i+2] == "h" {
				anstr := "an"
				astr := "a"
				return strings.Replace(emptyString, astr, anstr, 1)
			}
		}
	}
}