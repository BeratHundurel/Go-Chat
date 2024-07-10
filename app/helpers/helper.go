package helpers

import "strconv"

func ReturnIdAsStringFromUint(id uint) string {
	return strconv.FormatUint(uint64(id), 10)
}

func ReturnIdAsIntFromUint(id uint) int {
	return int(id)
}

func ReturnIdAsJsonFromUint(id uint) string {
	return `{"id":` + strconv.FormatUint(uint64(id), 10) + `}`
}

func ReturnIdAsIntFromString(id string) (int, error) {
	ID, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return ID, nil
}
