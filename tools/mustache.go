package tools

import "strings"

func GetMustacheTemplate(template string) string {
	mustache := strings.Replace(template, "{", "{{", -1)
	mustache = strings.Replace(mustache, "}", "}}", -1)

	return mustache
}
