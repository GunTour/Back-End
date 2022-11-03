package helper

import (
	"unicode"
)

func Password(pass string) string {
	var (
		upp, low, num bool
		tot           uint8
	)

	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		// case unicode.IsPunct(char) || unicode.IsSymbol(char):
		//     sym = true
		//     tot++
		default:
			return "tidak ada password"
		}
	}

	if !upp {
		return "password must contain uppercase"
	} else if !low {
		return "password must contain lowercase"
	} else if !num {
		return "password must contain numeric"
		// } else if !sym {
		//     return "password must contain symbol"
	} else if tot < 8 {
		return "password must have minumum 8 character"
	}

	return "Valid"
}
