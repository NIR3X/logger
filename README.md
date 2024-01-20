# Logger Package - A Simple Logging Utility for Go Applications

The `logger` package provides a simple logging utility for Go applications, allowing you to easily log messages with timestamps, file information, and custom prefixes.

## Installation

To use this package in your Go project, you can install it using the following command:

```bash
go get -u github.com/NIR3X/logger
```

## Usage

Import the package in your Go code:

```go
import "github.com/NIR3X/logger"
```

## Fprintln

The `Fprintln` function writes formatted log messages to the specified `io.Writer` with a timestamp, file information, and additional content. Example:

```go
logger.Fprintln(os.Stdout, "This is a log message.")
```

## Eprintln

The `Eprintln` function writes formatted error messages to the standard error stream with a timestamp, file information, and the "error:" prefix. Example:

```go
logger.Eprintln("This is an error message.")
```

## Println

The `Println` function writes formatted log messages to the standard output with a timestamp and file information. Example:

```go
logger.Println("This is another log message.")
```

Feel free to explore and adapt the functions to suit your logging needs.

## License
[![GNU AGPLv3 Image](https://www.gnu.org/graphics/agplv3-155x51.png)](https://www.gnu.org/licenses/agpl-3.0.html)  

This program is Free Software: You can use, study share and improve it at your
will. Specifically you can redistribute and/or modify it under the terms of the
[GNU Affero General Public License](https://www.gnu.org/licenses/agpl-3.0.html) as
published by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
