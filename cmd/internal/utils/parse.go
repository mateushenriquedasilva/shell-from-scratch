package utils

func Parse(input string) []string {
	var args []string
	var current string

	inQuotes := false

	for i := 0; i < len(input); i++ {
		c := input[i]

		switch c {
		case '\'':
			inQuotes = !inQuotes
		case ' ':
			if inQuotes {
				current += string(c)
			} else if current != "" {
				args = append(args, current)
				current = ""
			}
		default:
			current += string(c)
		}
	}

	if current != "" {
		args = append(args, current)
	}

	return args
}
