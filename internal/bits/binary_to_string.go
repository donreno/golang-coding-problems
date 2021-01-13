package bits

import "strings"

// Max32Bits represents all 32 bits marked

// BinaryToString converts binary to string
func BinaryToString(in float64) string {
	if in >= 1 || in <= 0 {
		return "ERROR"
	}

	builder := new(strings.Builder)
	builder.WriteString("0.")

	for in > 0 {
		if builder.Len() > 32 {
			return "ERROR"
		}

		in = in * 2

		if in >= 1 {
			builder.WriteRune('1')
			in--
		} else {
			builder.WriteRune('0')
		}
	}

	return builder.String()
}
