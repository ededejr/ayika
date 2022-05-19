# ayika

the yoruba word for "environment". this package provides convenient methods for parsing environment variables into your go program.

## usage
### reading env files
a `Load` interface is exported by the package which supports reading .env files in a few different ways

#### reading from the nearest .env file
```go
ayika.Load.Default()
```

#### reading from multiple local .env files
```go
ayika.Load.Local("path1", "path2", ..."pathN")
```

#### using a reader to read from remote sources
```go
ayika.Load.WithReader(reader)
```

#### reading into a variable
```go
envMap := ayika.Load.IntoVariable()
```

### parsing values
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

## developing

### tests
unit tests are located in the `tests` folder and can be run with `cd tests && go test`.
