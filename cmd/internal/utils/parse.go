package utils

func Parse(input string) []string {
	var args []string
	var current []rune

	inSingleQuotes := false
	inDoubleQuotes := false

	for _, c := range input {

		switch c {
		case '\'':
			if inDoubleQuotes {
				current = append(current, c)
			} else {
				inSingleQuotes = !inSingleQuotes
			}
		case '"':
			if inSingleQuotes {
				current = append(current, c)
			} else {
				inDoubleQuotes = !inDoubleQuotes
			}
		case ' ':
			if inSingleQuotes || inDoubleQuotes {
				current = append(current, c)
			} else if len(current) > 0 {
				args = append(args, string(current))
				current = []rune{}
			}
		default:
			current = append(current, c)
		}
	}
	if len(current) > 0 {
		args = append(args, string(current))
	}
	return args
}
