package kata

import "strings"

/*
Your task is to convert strings to how they would be written by Jaden Smith.
The strings are actual quotes from Jaden Smith, but they are not capitalized
in the same way he originally typed them.

Example:

Not Jaden-Cased: "How can mirrors be real if our eyes aren't real"
Jaden-Cased:     "How Can Mirrors Be Real If Our Eyes Aren't Real"
*/

func ToJadenCase(str string) string {
	// your code here...
	return strings.Title(str) //strings.Title() is depricated

	/* a better solution on your own:

	   words := strings.Split(str, " ")

	   for i, word := range words {
	     if len(word) == 0 {
	       continue
	     }

	     w := strings.ToUpper(string(word[0]))
	     words[i] = w
	     if len(word) == 1 {
	       continue
	     }

	     words[i] = w+word[1:]
	   }

	   return strings.Join(words, " ")
	*/

	// an even better solution:
	/*
		   for i, c := range str {
			if i == 0 || str[i-1] == ' ' {
				res += strings.ToUpper(string(c))
			}else {
				res += string(c)
			}
		}
		return
	*/

	//and an other one with strings.ToTitle()
	/*
			prev := ' '
	  result := make([]byte, len(str))
	  for i := 0; i < len(str); i++ {
	    c := rune(str[i])
	    if unicode.IsSpace(prev) {
	      c = unicode.ToTitle(c)
	    }
	    prev = c
	    result[i] = byte(c)
	  }
	  return string(result)
	*/
}
