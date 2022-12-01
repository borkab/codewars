package kata

import "strings"

/*
Complete the solution so that it returns true if the first argument(string) passed in ends with the 2nd argument (also a string).

Examples:

solution("abc", "bc") // returns true
solution("abc", "d") // returns false
*/

func solution(str, ending string) bool {
	// Your code here!
	return strings.HasSuffix(str, ending)
	//it gives the same result:
	//return len(str) >= len(ending) && str[len(str)-len(ending):] == ending

	//it gives also the same result:
	/*if len(str) < len(ending) == true{ return false }
	runes := []rune(str)
	substring := string(runes[(len(str)-len(ending)):len(str)])
	result := substring == ending
	return result
	*/

	//also another solution:
	/*endingLength, strLength := len(ending), len(str)
		index := strLength - endingLength
	if index < 0 {
			return false
			}
		strSl := strings.Split(str, "")
		end := strings.Join(strSl[index:], "")
		return ending == end
	*/

	//and an other one:
	/*
		if len(str) >= len(ending) {
			strLength := len(str)
			endingLength := len(ending)
			startIndex := strLength - endingLength
			comparedData := ""
			for i := startIndex; i < strLength; i++ {
				comparedData += string(str[i])
			}
			if comparedData == ending {
				return true
			}else{
				return false
			}
		}else {
			return false
		}
	*/
}
