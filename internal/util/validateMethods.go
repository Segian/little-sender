package util

import (
	"strings"

	"github.com/Segian/little-sender/internal/config/constants"
)

func ValidateMethod(method string) bool {
	UpperMethod := strings.ToUpper(method)
	switch UpperMethod {
	case constants.MethodGet, constants.MethodPost, constants.MethodPut, constants.MethodDelete:
		return true
	default:
		return false
	}
}

func MethodToUpper(method *string) {
	if method != nil {
		*method = strings.ToUpper(*method)
	}
}
