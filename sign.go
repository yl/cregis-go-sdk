package cregis

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Sign(data map[string]any, apiKey string) string {
	m := make(map[string]any, len(data))
	for k, v := range data {
		var s string
		switch v.(type) {
		case float64:
			s = strconv.FormatFloat(v.(float64), 'f', -1, 64)
		case int, int64:
			s = strconv.FormatInt(v.(int64), 10)
		default:
			s = fmt.Sprintf("%s", v)
		}

		//not null
		if len(s) < 1 {
			continue
		}

		// not sign filed
		if k == "sign" {
			continue
		}

		m[k] = s
	}

	sortStr := paramsAsciiAsc(m)
	sortStr = fmt.Sprintf("%s%s", apiKey, sortStr)
	md5Str := md5.Sum([]byte(sortStr))
	sortStr = fmt.Sprintf("%x", md5Str)
	sortStr = strings.ToLower(sortStr)
	return sortStr
}

func Verify(data map[string]any, apiKey, signStr string) bool {
	return Sign(data, apiKey) == signStr
}

func paramsAsciiAsc(paramMap map[string]any) string {
	list := make([]string, 0, len(paramMap))
	for k, _ := range paramMap {
		list = append(list, k)
	}

	sort.Strings(list)

	r := strings.Builder{}
	for _, k := range list {
		r.WriteString(fmt.Sprintf("%v%v", k, paramMap[k]))
	}

	return r.String()
}
