package common

import (
	"strconv"
	"strings"
)

/**
* 解析redis服务信息
 */
func ParseInfo(infoString string) (result map[string]map[string]interface{}) {
	result = map[string]map[string]interface{}{}
	lines := strings.Split(infoString, "\n")
	section := "default"
	for i := range lines {

		if strings.TrimSpace(lines[i]) == "" {
			continue
		}

		if strings.HasPrefix(lines[i], "#") {
			sections := strings.Fields(lines[i])
			if len(sections) > 1 {
				section = sections[1]
			}
			continue
		}

		items := strings.Split(lines[i], ":")
		if len(items) < 1 {
			continue
		}
		if strings.Contains(items[1], ",") {
			valueItems := strings.Split(items[1], ",")
			subMap := make(map[string]interface{}, len(valueItems))
			for i := range valueItems {
				subValueitems := strings.Split(valueItems[i], "=")
				if len(subValueitems) < 1 {
					continue
				}
				floatValue, err := strconv.ParseFloat(strings.TrimSpace(subValueitems[1]), 10)
				if err == nil {
					subMap[subValueitems[0]] = floatValue
				} else {
					subMap[subValueitems[0]] = strings.TrimSpace(subValueitems[1])
				}
			}
			if _, ok := result[section]; ok {
				result[section][items[0]] = subMap
			} else {
				result[section] = map[string]interface{}{
					items[0]: subMap,
				}
			}

		} else {
			floatValue, err := strconv.ParseFloat(strings.TrimSpace(items[1]), 10)
			var vlaue interface{} = strings.TrimSpace(items[1])
			if err == nil {
				vlaue = floatValue
			}
			if _, ok := result[section]; ok {
				result[section][items[0]] = vlaue
			} else {
				result[section] = map[string]interface{}{
					items[0]: vlaue,
				}
			}
		}
	}
	return
}
