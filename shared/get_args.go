package shared

import "regexp"

// package switch expression {
// case condition:

// }
func GetArgs(a []string) map[string]interface{} {
	final := make(map[string]interface{})
	for i, v := range a {
		if i >= 1 {
			validParam := regexp.MustCompile(`^[-][-]?`)
			regexKey := regexp.MustCompile(`^[-][-]?`)
			if validParam.MatchString(v) {
				final[regexKey.ReplaceAllString(v, "")] = true
			} else if !validParam.MatchString(v) && validParam.MatchString(a[i-1]) {
				final[regexKey.ReplaceAllString(a[i-1], "")] = v
			}
		}
	}
	return final
}
