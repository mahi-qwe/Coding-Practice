import (
    "strings"
)

func checkIfPangram(sentence string) bool {
    str := strings.ToLower(sentence)
    seen := [26]bool{}
    
    for _, ch := range str{
        if ch >= 'a' && ch <= 'z' {
			seen[ch -'a'] = true
		}
    }

    for _, value := range seen{
        if !value{
            return false
        }
    }

    return true
}