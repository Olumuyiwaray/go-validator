// Package validator provides a simple and chainable validation library
// for Go. It allows you to validate strings, numbers, emails, phone numbers,
// and more with optional custom error messages.
//
// Example usage:
//
//   v := validator.New()
//   v.Field("test@example.com", "Email").
//       Required().
//       String().
//       Email().
//
//   if err := v.Validate(false); err != nil {
//       fmt.Println(err)
//   }

package validator

import (
	"fmt"
	"regexp"
)

// Validator holds all the validation rules for multiple fields.
// Call Validate() to check all rules.
type Validator struct {
    rules []func() error
}

// New creates and returns a new Validator instance.
//
// Example:
//
//    v := validator.New()
//
func New() *Validator {
    return &Validator{rules: []func() error{}}
}

// Field represents a single value being validated.
// It stores the value, its display name, and a reference to the parent validator.
// You can chain any number of validation methods on Field.
//
// Example:
//
//    v.Field(username, "Username").
//        Required().
//        String().
//        MinLength(3)etc.
type Field struct {
    validator *Validator
    value     interface{}
    name      string
}

// Field registers a new field to validate.
// `value` is the actual value being validated.
// `name` is the field name used in error messages.
//
// Example:
//
//    v.Field("john@example.com", "Email").Email()
func (v *Validator) Field(value interface{}, name string) *Field {
    return &Field{
        validator: v,
        value:     value,
        name:      name,
    }
}

// String ensures the field value is a string.
// Optionally accepts a custom error message.
//
// Example:
//    f.String()
//    f.String("Username must be text")
func (f *Field) String(messages ...string) *Field {
    f.validator.rules = append(f.validator.rules, func() error {
        message := ""
    if len(messages) > 0 {
        message = messages[0]
    }
        _, ok := f.value.(string)
        if !ok {
            if message != "" {
                return fmt.Errorf("%s", message);
            }
            return fmt.Errorf("%s must be a string", f.name);
        }
        return nil;
    })
    return f;
}

// Required ensures the field value is not empty.
// Works for string, int, float64, bool, and nil.
//
// Example:
//    f.Required()
func (f *Field) Required() *Field {
    f.validator.rules = append(f.validator.rules, func() error {
        switch v := f.value.(type) {
        case string:
            if len(v) == 0 {
                return fmt.Errorf("%s is required", f.name)
            }
        case int:
            if v == 0 {
                return fmt.Errorf("%s is required", f.name)
            }
        case float64:
            if v == 0.0 {
                return fmt.Errorf("%s is required", f.name)
            }
        case bool:
            // usually boolean always has a value, skip if not needed
        default:
            if f.value == nil {
                return fmt.Errorf("%s is required", f.name)
            }
        }
        return nil
    })
    return f
}


// Email validates that the field value is a valid email address.
// Accepts an optional custom error message.
//
// Example:
//    f.Email()
//    f.Email("Invalid email format")
func (f *Field) Email(messages ...string) *Field {
    f.validator.rules = append(f.validator.rules, func() error {
 	message := ""
    if len(messages) > 0 {
        message = messages[0]
    }
        str, ok := f.value.(string)
        if !ok {
			if message != "" {
                return fmt.Errorf("%s", message);
            }
            return fmt.Errorf("%s must be a valid email", f.name)
        }

        re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
        if !re.MatchString(str) {
			if message != "" {
                return fmt.Errorf("%s", message);
            }
            return fmt.Errorf("%s must be a valid email", f.name)
        }
        return nil
    })
    return f
}


// Min checks that an integer value is greater than or equal to `length`.
// Accepts an optional custom error message.
//
// Example:
//    f.Min(10)
//    f.Min(10, "Value must be at least 10")
func (f *Field) Min(length int, messages ...string) *Field {
    f.validator.rules = append(f.validator.rules, func() error {
        message := ""
    if len(messages) > 0 {
        message = messages[0]
    }
		
      value, ok := f.value.(int);

	  if (!ok) {
			if message != "" {
                return fmt.Errorf("%s", message);
            }
		return fmt.Errorf("%s must be an integer ", f.name)
	  }

	  if (value < length) {
		if message != "" {
            return fmt.Errorf("%s", message);
        }
		return fmt.Errorf("%d cannot be less than %d", value, length)
	  }
        return nil
    })
    return f
}

// Max checks that an integer value does not exceed `length`.
// Accepts an optional custom error message.
//
// Example:
//    f.Max(100)
//    f.Max(100, "Too large")
func (f *Field) Max(length int, messages ...string) *Field {
    f.validator.rules = append(f.validator.rules, func() error {
        message := ""
    if len(messages) > 0 {
        message = messages[0]
    }
		
      value, ok := f.value.(int);

	  if (!ok) {
		if message != "" {
            return fmt.Errorf("%s", message);
        }
		return fmt.Errorf("%s must be an integer ", f.name)
	  }

	  if (value > length) {
		if message != "" {
            return fmt.Errorf("%s", message);
        }
		return fmt.Errorf("%d cannot be greater than %d", value, length)
	  }
        return nil
    })
    return f
}


// MinLength validates that a string's length is at least `length`.
// Accepts an optional custom error message.
//
// Example:
//    f.MinLength(3)
//    f.MinLength(3, "Too short")
func (f *Field) MinLength(length int, messages ...string) *Field {
    f.validator.rules = append(f.validator.rules, func() error {

    message := ""
    if len(messages) > 0 {
        message = messages[0]
    }
		
      value, ok := f.value.(string);

	  if (!ok) {
		if message != "" {
        	return fmt.Errorf("%s", message);
        }
		return fmt.Errorf("%s must be a string ", f.name)
	  }

	  if (len(value) < length) {
		if message != "" {
            return fmt.Errorf("%s", message);
        }
		return fmt.Errorf("%s cannot be less than %d characters ", value, length)
	  }
        return nil
    })
    return f
}

// MaxLength validates that a string's length is not greater than `length`.
// Accepts an optional custom error message.
//
// Example:
//    f.MaxLength(20)
//    f.MaxLength(20, "Too long")
func (f *Field) MaxLength(length int, messages ...string) *Field {
    f.validator.rules = append(f.validator.rules, func() error {
		message := ""
    if len(messages) > 0 {
        message = messages[0]
    }
      value, ok := f.value.(string);

	  if (!ok) {
		if message != "" {
            return fmt.Errorf("%s", message);
        }
		return fmt.Errorf("%s must be a string ", f.name)
	  }

	  if (len(value) > length) {
		if message != "" {
            return fmt.Errorf("%s", message);
        }
		return fmt.Errorf("%s cannot be more than %d characters", value, length)
	  }
        return nil
    })
    return f
}

func (f *Field) Number(messages ...string) *Field {
    f.validator.rules = append(f.validator.rules, func() error {
        message := ""
    if len(messages) > 0 {
        message = messages[0]
    }
        _, ok := f.value.(int)
        if !ok {
            if message != "" {
                return fmt.Errorf("%s", message);
            }
            return fmt.Errorf("%s must be a number", f.name);
        }
        return nil;
    })
    return f;
}


// Phone validates that the field value is a phone number.
// Supports optional "+" prefix and 10–15 digits.
// Accepts an optional custom error message.
//
// Example:
//    f.Phone()
//    f.Phone("Invalid phone format")
func (f *Field) Phone(messages ...string) *Field {
    f.validator.rules = append(f.validator.rules, func() error {
        message := ""
    if len(messages) > 0 {
        message = messages[0]
    }
        str, ok := f.value.(string)
        if !ok {
			if message != "" {
                return fmt.Errorf("%s", message);
            }
            return fmt.Errorf("%s must be a valid email", f.name)
        }

        re := regexp.MustCompile(`^\+?[0-9]{10,15}$`)
        if !re.MatchString(str) {
			if message != "" {
                return fmt.Errorf("%s", message);
            }
            return fmt.Errorf("%s must be a valid phone number", f.name)
        }
        return nil
    })
    return f
}


// Validate runs all validation rules.
// If stopOnFirst is true, it stops at the first error.
func (v *Validator) Validate(stopOnFirst bool) []error {
	var allErrors []error

	for _, rule := range v.rules {
		if err := rule(); err != nil {
			if stopOnFirst {
				return []error{err}
			}
			allErrors = append(allErrors, err)
		}
	}

	if len(allErrors) == 0 {
		return nil 
	}
	return allErrors
}