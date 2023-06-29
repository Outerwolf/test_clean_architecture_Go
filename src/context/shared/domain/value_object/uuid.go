package value_object

import "github.com/google/uuid"

func ValidateUuid(value string) (string, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return "", err
	}

	return v.String(), nil
}
