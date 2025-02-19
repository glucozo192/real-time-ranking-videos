package utils

import "strconv"

func ConvertStringToInt(value string) (int, error) {
	if s, err := strconv.ParseFloat(value, 64); err == nil {
		valueNumber := int(s)
		return valueNumber, nil
	} else {
		return strconv.Atoi(value)
	}
}
