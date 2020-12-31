package interpret

// Interpret Given the string command, return the Goal Parser's interpretation of command.
// https://leetcode.com/problems/goal-parser-interpretation/
func Interpret(command string) string {
	var result string
	for i, letter := range command {
		if string(letter) == "G" {
			result += "G"
		} else if string(letter) == "(" {
			if string(command[i+1]) == ")" {
				result = result + "o"
			} else if string(command[i+1]) == "a" {
				result = result + "al"
			}
		}
	}
	return result
}
