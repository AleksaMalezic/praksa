package iteration

const repeatCounter = 5

func Repeat(char_a string, num_a int) string {
	var repeated string
	if num_a == 0{
		for i := 0; i < repeatCounter; i++{
			repeated += char_a
		}
		return repeated
	}
	for i := 0; i < num_a; i++{
		repeated += char_a
	}
	return repeated
}