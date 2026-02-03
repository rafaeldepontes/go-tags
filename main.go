package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type User struct {
	Name string `validate:"min=3,max=32"`

	// FIXME: I don't personally like email validation by some regex...
	//
	// Not only its a bad thing, but it can also break my whole system
	// with a big enough regex! So I'm not gonna use it here.
	Email string `validate:"required"`
}

func validate(value any) error {
	val := reflect.ValueOf(value)
	if val.Kind() != reflect.Struct {
		return errors.New("This is not a struct anymore buddy!!!")
	}

	for i := range val.NumField() {
		field := val.Field(i)
		tag := val.Type().Field(i).Tag.Get("validate")
		if tag == "" {
			continue
		}

		rules := strings.Split(tag, ",")
		for _, rule := range rules {
			name := val.Type().Field(i).Name

			switch {
			case rule == "required":
				if field.IsZero() {
					return fmt.Errorf("%s is required", name)
				}
			case strings.HasPrefix(rule, "min="):
				min, _ := strconv.Atoi(strings.TrimPrefix(rule, "min="))

				if field.IsZero() || len(field.String()) < min {
					return fmt.Errorf("%s must have a length of at least %d", name, min)
				}
			case strings.HasPrefix(rule, "max="):
				max, _ := strconv.Atoi(strings.TrimPrefix(rule, "max="))

				if field.IsZero() || len(field.String()) > max {
					return fmt.Errorf("%s must have a maximum length of %d", name, max)
				}
			}
		}
	}
	return nil
}

func main() {
	shortName := User{
		Name:  "u",
		Email: "test@example.com",
	}

	longName := User{
		Name:  "I have at least 32 characters, but I'm not sure.",
		Email: "test@example.com",
	}

	badEmail := User{
		Name:  "rafael",
		Email: "",
	}

	validUser := User{
		Name:  "rafael",
		Email: "test@example.com",
	}

	if err := validate(shortName); err != nil {
		println(err.Error())
	}
	if err := validate(longName); err != nil {
		println(err.Error())
	}
	if err := validate(badEmail); err != nil {
		println(err.Error())
	}
	if err := validate(validUser); err != nil {
		println(err.Error())
	}
	println("No errors for valid User!")
}
