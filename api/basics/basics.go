package basics

import (
	"encoding/json"
	"errors"
)

/**
These are some example solutions to the problems in the exercises.
 */

/**
 Exercise 1
*/

//func getResult(name string) (string, error) {
//	resultString := fmt.Sprintf("Input string '%v' has '%v' characters", name, len(name))
//
//	return resultString, nil
//}

/**
Exercise 2
 */

//func getResult(name string) (string, error) {
//	if name == "world" {
//		return "hello", nil
//	} else if name == "hello" {
//		return "world", nil
//	}
//
//	return "", errors.New("Some error")
//}

//func getResult(name string) (string, error) {
//	switch name {
//	case "world":
//		return "hello", nil
//	case "hello":
//		return "world", nil
//	default:
//		return "", errors.New("some error")
//	}
//}

/**
Exercise 4
 */

//func getResult(name string) (string, error) {
//	return reverseString(name)
//}

func reverseString(input string) (string, error) {
	if len(input) <= 1 {
		return "", errors.New("input should be longer than 1 character")
	}

	var result string

	for _, character := range input {
		result = string(character) + result
	}

	return result, nil
}

/**
Exercise 5
 */

//func getResult(name string) (string, error) {
//	stringResult, reverseError := reverseString(name)
//
//	resultMap := map[string]string{}
//
//	if reverseError == nil {
//		resultMap["result"] = stringResult
//	} else {
//		resultMap["error"] = reverseError.Error()
//	}
//
//	jsonResult, _ := json.Marshal(resultMap)
//
//	return string(jsonResult), nil
//}

type Response struct {
	Result string `json:"result,omitempty"`
	Error string `json:"error,omitempty"`
}

func getResult(name string) (string, error) {
	stringResult, reverseError := reverseString(name)

	resultMap := Response{}

	if reverseError == nil {
		resultMap.Result = stringResult
	} else {
		resultMap.Error = reverseError.Error()
	}

	jsonResult, _ := json.Marshal(resultMap)

	return string(jsonResult), nil
}
