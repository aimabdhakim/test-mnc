package main

import (
	"fmt"
	"strings"
)

func question1() {
	var size int
	fmt.Scanln(&size)

	var arr = make([]string, size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%s", &arr[i])
	}

	indexResult := findPair(arr)

	var result string
	if len(indexResult) == 0 {
		result = "false"
	} else {
		// to get the output as intended (int int int ... )
		result = strings.Trim(strings.Replace(fmt.Sprint(indexResult), " ", " ", -1), "[]")
	}

	fmt.Println(result)
}

func findPair(arr []string) []int {
	var arrIndexPair []int
	var firstNonUniqueString string

	var uniqueString []string
	for i := 0; i < len(arr); i++ {
		if !contains(uniqueString, arr[i]) {
			uniqueString = append(uniqueString, arr[i])
		} else if firstNonUniqueString == "" {
			firstNonUniqueString = arr[i]
		}
	}

	for i := 0; i < len(arr); i++ {
		if strings.EqualFold(firstNonUniqueString, arr[i]) {
			arrIndexPair = append(arrIndexPair, i+1)
		}
	}

	return arrIndexPair
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.EqualFold(a, e) {
			return true
		}
	}
	return false
}
