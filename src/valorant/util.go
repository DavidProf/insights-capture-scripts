package valorant

import "strings"

func getNameFromCode(data map[string]string, code string) string {
	code = strings.ToLower(code)
	if data[code] == "" {
		return code
	}
	return data[code]
}
