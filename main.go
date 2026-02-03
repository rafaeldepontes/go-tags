package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type User struct {
	Name string `Validate:"min=3,max=32"`

	// FIXME: I don't personally like email validation by some regex...
	//
	// Not only its a bad thing, but it can also break my whole system
	// with a big enough regex! So I'm not gonna use it here.
	Email string `Validate:"required"`
}

func Validate(value any) error {
	val := reflect.ValueOf(value)
	if val.Kind() != reflect.Struct {
		return errors.New("This is not a struct anymore buddy!!!")
	}

	var errs error

	for i := range val.NumField() {
		tag := val.Type().Field(i).Tag.Get("Validate")
		if tag == "" {
			continue
		}

		rules := strings.Split(tag, ",")
		for _, rule := range rules {
			if err := checkRules(&val, i, rule); err != nil {
				errs = errors.Join(errs, err)
			}
		}
	}

	return errs
}

func checkRules(value *reflect.Value, index int, rule string) error {
	field := value.Field(index)
	name := value.Type().Field(index).Name

	switch {
	case rule == "required":
		if field.IsZero() {
			return fmt.Errorf("%s is required", name)
		}

	case strings.HasPrefix(rule, "min="):
		min, err := strconv.Atoi(strings.TrimPrefix(rule, "min="))
		if err != nil {
			return fmt.Errorf("invalid min value for %s: %v", name, err)
		}

		if len(field.String()) < min {
			return fmt.Errorf("%s must have a length of at least %d", name, min)
		}

	case strings.HasPrefix(rule, "max="):
		max, err := strconv.Atoi(strings.TrimPrefix(rule, "max="))
		if err != nil {
			return fmt.Errorf("invalid max value for %s: %v", name, err)
		}

		if len(field.String()) > max {
			return fmt.Errorf("%s must have a maximum length of %d", name, max)
		}
	}
	return nil
}

func main() {
	// NOTHING HERE... IF YOU WANT THE TEST CASES
	// GO CHECK THE "main_test.go".
}
