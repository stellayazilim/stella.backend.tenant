package helpers

import (
	"fmt"
	"strconv"
)

func ConvertUint(str string) (uint, error) {
	n, err := strconv.ParseUint(str, 10, 32)

	fmt.Println(n, err)
	if err != nil {
		return 0, fmt.Errorf("%'is not a valid id", str)
	}
	return uint(n), nil
}
