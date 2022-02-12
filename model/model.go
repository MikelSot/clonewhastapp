package model

import (
	"errors"
	"reflect"
)

var (
	ErrNilPointer = errors.New("el modelo recibido es nulo")
)

// ValidateStructNil returns an error if the model is nil
func ValidateStructNil(i interface{}) error {
	// omit struct type
	if reflect.ValueOf(i).Kind() == reflect.Struct {
		return nil
	}

	// Type: nil, Value: nil
	if i == nil {
		return ErrNilPointer
	}

	// Type: StructPointer, Value: nil
	// example: Type: *Cashbox, Value: nil
	if reflect.ValueOf(i).IsNil() {
		return ErrNilPointer
	}

	// Type: StructPointer, Value: ZeroValue
	// example: Type: *CashBox, Value: &CashBox{}
	return nil
}
