# go-validator

A lightweight, chainable validation library for Go — with support for strings, numbers, emails, phone numbers, and custom messages.

---

## Installation

```sh
go get github.com/Olumuyiwaray/go-validator
```

## Features

- Chainable, expressive validation
- String, number, email & phone validation
- Required field validation
- Min / Max value checks
- MinLength / MaxLength checks
- Supports custom error messages
- Validate all fields or stop on first error
- Zero dependencies

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/Olumuyiwaray/go-validator"
)

func main() {
    v := validator.New()

    v.Field("test@example.com", "Email").
        Required().
        String().
        Email()

    if err := v.Validate(false); err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Validation passed")
    }
}
```

## Usage Examples

#### Validate Required Email

```
v.Field("hello@example.com", "Email").
    Required().
    String().
    Email()
```

#### Validate Integer with Min/Max

```
v.Field(25, "Age").
    Required().
    Number().
    Min(18).
    Max(60)
```

#### Validate String Length

```
v.Field("Samuel", "Username").
    Required().
    MinLength(3).
    MaxLength(15)
```

#### Validate Phone Number

```
v.Field("+12345678901", "Phone").
    Required().
    Phone()
```

#### Custom Error Messages

```
v.Field("abc", "Email").
    Email("invalid email format provided")
```

### Validation Modes

#### Stop on First Error

```
err := v.Validate(true)
if err != nil {
    fmt.Println("Error:", err)
}
```

#### Collect All Errors

```
err := v.Validate(false)
if err != nil {
    fmt.Println("Errors:", err)
}
```

### Contributing

Pull requests are welcome.
For major changes, open an issue first to discuss your idea.

### Support

If you find this library useful, please star the repository.
