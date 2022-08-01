package util

import (
	"cleaningservice/common/variables"
	"strconv"
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

// String to int64 array
func StringToIntArray(str string) []int64 {
	result := []int64{}
	if str == "" {
		return result
	}

	strList := strings.Split(str, variables.Separator)
	for _, val := range strList {
		if val == "" {
			continue
		}

		id, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			continue
		}

		result = append(result, id)
	}

	return result
}

// Int64 array to string
func IntArrayToString(arr []int64) string {
	result := []string{}
	for _, val := range arr {
		result = append(result, strconv.FormatInt(val, 10))
	}

	return strings.Join(result[:], variables.Separator)
}

// Disjoint elements from two int64 array
// arr1 is the array containing elements wanted to be delete
// arr2 is the origin array
func DisjointIntArray(arr1 []int64, arr2 []int64) []int64 {
	result := []int64{}
	checkMap := make(map[int64]int)

	for _, val := range arr1 {
		checkMap[val] = 1
	}
	for _, val := range arr2 {
		_, exist := checkMap[val]
		if !exist {
			result = append(result, val)
		}
	}

	return result
}

// Union of two int64 array
func UnionIntArray(arr1 []int64, arr2 []int64) []int64 {
	result := []int64{}
	checkMap := make(map[int64]int)

	for _, val := range arr1 {
		checkMap[val] = 1
	}
	for _, val := range arr2 {
		checkMap[val] = 1
	}

	for key := range checkMap {
		result = append(result, key)

	}

	return result
}
