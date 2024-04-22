package validation

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/dcorley125/Go-SDK-Repo/internal/utils"
)

func validatePattern(field reflect.StructField, value reflect.Value) error {
	pattern, found := field.Tag.Lookup("pattern")
	if !found {
		return nil
	}

	compiledRegex, err := regexp.Compile(pattern)
	if err != nil {
		return fmt.Errorf("regex failed to compile")
	}

	kind := utils.GetReflectKind(value.Type())

	if kind != reflect.String {
		return fmt.Errorf("field %s with value %v cannot match pattern %s because it is not a string", field.Name, value, pattern)
	}

	if !compiledRegex.MatchString(value.String()) {
		return fmt.Errorf("field %s with value %v does not match pattern %s", field.Name, value, pattern)
	}

	return nil
}
