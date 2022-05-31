package util

import (
	"cleaningservice/common/variables"
	"strings"
)

// Combine two string arraies
func CombineStringArray(a1 []string, a2 []string) ([]string, string) {
	res := append(a1, a2...)
	return res, strings.Join(res[:], variables.Separator)
}

//Remove union elements on array one
func RemoveUnionStringArray(target []string, reference []string) ([]string, string) {
	if len(reference) >= len(target) {
		return []string{}, ""
	}

	// Get different elements from target referring to reference
	resultList := []string{}
	checkMap := make(map[string]int)

	for _, val := range reference {
		checkMap[val] = 1
	}
	for _, val := range target {
		_, exist := checkMap[val]
		if !exist {
			resultList = append(resultList, val)
		}
	}

	return resultList, strings.Join(resultList[:], variables.Separator)
}
