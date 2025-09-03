package model

import (
	"database/sql/driver"
	"errors"
)

type JSON []byte

func (j *JSON) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	if len(*j) == 0 || string(*j) == "null" {
		return nil, nil
	}

	return string(*j), nil
}

func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}

	s, ok := value.([]byte)
	if !ok {
		return errors.New("invalid scan source")
	}

	*j = append((*j)[0:0], s...)

	return nil
}

func (j *JSON) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}

	return *j, nil
}

func (j *JSON) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("null point exception")
	}

	*j = append((*j)[0:0], data...)

	return nil
}
