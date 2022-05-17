## ayika

the yoruba word for "environment". this package provides convenient methods for parsing environment variables into your go program.

### usage

#### parsing a value as it appears in the env file
```go
ayika.Value("<your env key>")
```

#### parsing a numeric value
all numeric values are parsed as `int64`.

```go
ayika.Num("<your env key>")
```

#### parsing a boolean value
```go
ayika.Bool("<your env key>")
```

#### parsing a float value
all float values are parsed as `float64`

```go
ayika.Float("<your env key>")
```

#### parsing a json value
```go
type KeyValue struct {
  Key string `json:"key"`
}

ayika.JSON[KeyValue]("<your env key>")
```

### developing

#### tests
unit tests are located in the `tests` folder and can be run with `cd tests && go test`.