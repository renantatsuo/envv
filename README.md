# envv

envv is a type-safe env var parser for Go.

It currently supports:

- Type safety for basic types (string, int, bool, float64)
- Default values
- Required field validation
- Panic on parse errors for early failure detection

## Installation

Although you can install it, I'd recommend you to just copy it into you project.

```bash
go get github.com/renantatsuo/envv
```

## Usage

### Basic Usage

```go
// String value
str := envv.Get("APP_NAME").String().Optional().Parse()

// Integer value
port := envv.Get("PORT").Int().Optional().Parse()

// Boolean value
debug := envv.Get("DEBUG").Bool().Optional().Parse()

// Float value
factor := envv.Get("SCALE_FACTOR").Float64().Optional().Parse()
```

### Default Values

Provide fallback values when environment variables are not set:

```go
port := envv.Get("PORT").Int().Default(8080).Parse()
debug := envv.Get("DEBUG").Bool().Default(false).Parse()
factor := envv.Get("SCALE_FACTOR").Float64().Default(1.0).Parse()
```

### Required Values

Mark environment variables as required - will panic if not set:

```go
apiKey := envv.Get("API_KEY").String().Required().Parse()
```

## Error Handling

The package panics fast-fail during application startup. It will panic:

- if a required variable is not set
- if a value cannot be parsed into the expected type
