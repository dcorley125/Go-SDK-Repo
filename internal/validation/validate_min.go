package validation

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/dcorley125/Go-SDK-Repo/internal/utils"
)

func validateMin(field reflect.StructField, value reflect.Value) error {
	minValue, found := field.Tag.Lookup("min")
	if !found || minValue == "" {
		return nil
	}

	min, err := strconv.Atoi(minValue)
	if err != nil {
		return err
	}

	if value.IsNil() {
		return fmt.Errorf("field %s is required", field.Name)
	}

	val := utils.GetReflectValue(value)

	if val.CanInt() {
		if val.Int() < int64(min) {
			return fmt.Errorf("field %s is less than min value", field.Name)
		}
	} else if val.CanFloat() {
		if val.Float() < float64(min) {
			return fmt.Errorf("field %s is less than min value", field.Name)
		}
	} else {
		return fmt.Errorf("field %s is not a number", field.Name)
	}

	return nil
}
