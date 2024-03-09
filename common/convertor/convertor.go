package convertor

import (
	"log"
	"strconv"
)

func String2int(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result
}

func String2bool(str string) bool {
	var result bool
	boolValue, err := strconv.ParseBool(str)
	if err == nil {
		result = boolValue
	} else {
		log.Fatal(err.Error())
	}
	return result
}
