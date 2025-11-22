package validator

import (
	"fmt"
	"regexp"
	"strings"
)


type Validator struct {
    rules []func() error
}


func New() *Validator {
    return &Validator{rules: []func() error{}}
}

type Field struct {
    validator *Validator
    value     interface{}
    name      string
}

func (v *Validator) Field(value interface{}, name string) *Field {
    return &Field{
        validator: v,
        value:     value,
        name:      name,
    }
}

func (f *Field) String(message string) *Field {
    f.validator.rules = append(f.validator.rules, func() error {
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


func (f *Field) Min(length int, message string) *Field {
    f.validator.rules = append(f.validator.rules, func() error {
		
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

func (f *Field) Max(length int, message string) *Field {
    f.validator.rules = append(f.validator.rules, func() error {
		
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

func (f *Field) MinLength(length int, message string) *Field {
    f.validator.rules = append(f.validator.rules, func() error {
		
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

func (f *Field) MaxLength(length int, message string) *Field {
    f.validator.rules = append(f.validator.rules, func() error {
		
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

func (f *Field) Number(message string) *Field {
    f.validator.rules = append(f.validator.rules, func() error {
        _, ok := f.value.(int)
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



func (f *Field) Phone(message string) *Field {
    f.validator.rules = append(f.validator.rules, func() error {
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


func (v *Validator) Validate(stopOnFirst bool) error {
	var allErrors []string
	for _, rule := range v.rules {
		if err := rule(); err != nil {
			if stopOnFirst {
				return err
			}
			allErrors = append(allErrors, err.Error())
		}
	}
	if len(allErrors) > 0 {
		return fmt.Errorf(strings.Join(allErrors, "; "))
	}
	return nil
}

