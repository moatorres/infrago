<p style="text-align: center;">
  <h1>Loggo</h1>
</p>

<div style="text-align: left;">

![License](https://img.shields.io/npm/l/nx.svg?style=flat-square)

</div>

<hr/>

## Zero-dependency JSON-based logger written with native Go modules.

Loggo is an open-source JSON-based logger that uses native Go modules to parse messages and errors into JSON format.

## Usage

```go
package mypackage

import (
  loggo "github.com/moatorres/go/modules/logger"
)

var logger = loggo.New(loggo.LoggerOptions{
  Service: "my-service",
})

logger.Log("Hello 👋") // `AnsiColor` GREEN
// → {"service":"my-service","level":"Log","message":"Hello 👋","timestamp":"2023-10-12T10:15:18-03:00"}
```

### `logger.Info`

`AnsiColor` **BLUE**

```go
var logger = loggo.New(loggo.LoggerOptions{
  Service: "awesome-golang",
})

logger.Info("Informative ℹ️")
// → {"service":"awesome-golang","level":"INFO",":"Informative ℹ️","timestamp":"2023-10-12T10:15:18-03:00"}
```

### `logger.Warn`

`AnsiColor` **YELLOW**

```go
var logger = loggo.New(loggo.LoggerOptions{
  Service: "pedestrian-angel",
})

logger.Warn("Crossing 🚸")
// → {"service":"pedestrian-angel","level":"WARN","message":"Crossing 🚸","timestamp":"2023-10-12T10:20:47-03:00"}
```

### `logger.Error`

`AnsiColor` **LIGHT_RED**

```go
var logger = New(LoggerOptions{
	Service: "chaos-monkey",
})

type MonkeyError struct{}

func (m *MonkeyError) Error() string {
	return "Boom 🙉 💥"
}

func chaos() (string, error) {
	return "", &MonkeyError{}
}

func unleash() {
	s, err := chaos()
	if err != nil {
		logger.Error("%s", err)
		os.Exit(1)
	}
}

unleash()
// → {"service":"chaos-monkey","level":"ERROR","message":"Boom 🙉 💥","timestamp":"2023-10-12T10:20:47-03:00"}
```

### `logger.Debug`

`AnsiColor` **PURPLE**

```go
var logger = loggo.New(loggo.LoggerOptions{
  Service: "bug-nest",
})

var bug := "🐛🐞🪲"

logger.Debug("What's going on here? %s", bug)
// → {"service":"bug-nest","level":"DEBUG","message":"What's going on here? 🐛🐞🪲","timestamp":"2023-10-12T10:20:47-03:00"}
```

### `logger.Fatal`

`AnsiColor` **RED**

```go
var logger = New(LoggerOptions{
	Service: "dead-service",
})

logger.Fatal("Uhoh 💀")
// → {"service":"dead-service","level":"FATAL","message":"Uhoh 💀","timestamp":"2023-10-12T10:15:18-03:00"}
// exit status 1
```

## Feature Roadmap

- [ ] Disable `LogLevel` colors
- [ ] Customize `LogLevel` colors
- [ ] Set `LogLevel` via environment variable
- [ ] Edit `LogMessage` properties

<sub>⚡️ Powered by **OSS** — `< >` with ☕️ by [**Moa Torres**](https://github.com/moatorres)</sub>
