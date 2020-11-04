package defangIPaddr

import "fmt"

func DefangIPaddr(address string) string {
	result := ``
	for _, char := range address {
		if string(char) == "." {
			result += fmt.Sprintf("[%v]", string(char))
		} else {
			result += string(char)
		}
	}
	return result
}
