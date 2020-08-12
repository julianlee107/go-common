package main

import (
	"fmt"
	"strings"
)

func main() {
	m := make(map[string]interface{})
	m["tag"] = "_undef"
	m["traceId"] = "123123"
	fmt.Println(parseParams(m))
}

func parseParams(m map[string]interface{}) string {
	var tag = "_undef"
	if _tag, ok := m["tag"]; ok {
		if val, ok := _tag.(string); ok {
			tag = val
		}
	}
	for key, val := range m {
		if key == "tag" {
			continue
		}
		tag = tag + "||" + fmt.Sprintf("%v=%v", key, val)
	}
	tag = strings.Trim(fmt.Sprintf("%q", tag), "\"")
	return tag
}
