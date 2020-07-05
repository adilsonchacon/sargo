# sargo

Simple (but awesome) args in Go: command line parser

## Install

```
$> go get github.com/adilsonchacon/sargo
```

## Setup

To setup the args in your application you should use the method *SET*. You can setup as many args as you want, but each name and shortcut must be unique.

```
Set(name string, shortCut string, defaultValue interface{}, description string)
```

* *name* is the full name of the arg.
* *shortCut* is the simplest abbreviation for the arg name.
* *defaultValue* is the value that the arg will assume if the arg is not in command line.
* *description* full explanation of the arg.

```go
...

import ("github.com/adilsonchacon/sargo")

...

sargo.Set("host", "h", "localhost", "http server host. Default value is \"localhost\"")
sargo.Set("port", "p", 8081, "http server port. Default value is \"8081\"")
sargo.Set("rate", "r", 5.5, "site rate. Default value is \"5.5\"")
sargo.Set("activated", "a", true, "activated. Default value is \"true\"")

...

```

To understand how *sargo* works properly, check out the following complete examples:

## Main Example
```go
package main

import (
	"fmt"
	"github.com/adilsonchacon/sargo"
)

func main() {
	sargo.Set("host", "h", "localhost", "http server host. Default value is \"localhost\"")
	sargo.Set("port", "p", 8081, "http server port. Default value is \"8081\"")
	sargo.Set("rate", "r", 5.5, "site rate. Default value is \"5.5\"")
	sargo.Set("activated", "a", true, "activated. Default value is \"true\"")

	host, err := sargo.Get("host")
	// host, err := sargo.GetString("host")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("host:", host)
	}

	port, err := sargo.GetInt("port")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("port:", port)
	}

	rate, err := sargo.GetFloat("rate")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("rate:", rate)
	}

	activated, err := sargo.GetBool("activated")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("activated:", activated)
	}
}
```

#### Example 1

```
$> go run main.go

host: localhost
port: 8081
rate: 5.5
activated: true
```

#### Example 2

```
$> go run main.go --host=10.0.0.1 --port=9091 --rate=3.14 --activated=false

host: 10.0.0.1
port: 9091
rate: 3.14
activated: false
```


#### Example 3

```
$> go run main.go -h 10.0.0.1 -p 9091 -r 3.14 -a false

host: 10.0.0.1
port: 9091
rate: 3.14
activated: false
```

#### Example 4

```
$> go run main.go --host=127.0.0.1 -p 9091 -a 1

host: 127.0.0.1
port: 9091
rate: 5.5
activated: true
```

#### Example 5

```
$> go run main.go -r 938.291

host: localhost
port: 8081
rate: 938.291
activated: true
```

#### Example 6

```
$> go run main.go --activated=f

host: localhost
port: 8081
rate: 5.5
activated: false
```

### Print Help Example

```go
package main

import (
	"github.com/adilsonchacon/sargo"
	"os"
)

func main() {
	sargo.Set("host", "h", "localhost", "http server host. Default value is \"localhost\"")
	sargo.Set("port", "p", 8081, "http server port. Default value is \"8081\"")
	sargo.Set("rate", "r", 5.5, "site rate. Default value is \"5.5\"")
	sargo.Set("activated", "a", true, "is site activated? Default value is \"true\"")
  
	sargo.PrintHelp()
}
```

```
$> go run main.go

Usage:
  go run main.go [options]

Options:
  -h, [--host]       # http server host. Default value is "localhost"
  -p, [--port]       # http server port. Default value is "8081"
  -r, [--rate]       # site rate. Default value is "5.5"
  -a, [--activated]  # is site activated? Default value is "true"
```

### Print Help Custom Usage Example

```go
package main

import (
	"github.com/adilsonchacon/sargo"
	"os"
)

func main() {
	sargo.SetUsage("my_exec_file -c [PATH_TO_CONFIG_FILE]")
	sargo.Set("config", "c", "./conf.yml", "path for configuration file")
  
	sargo.PrintHelp()
}
```

```
$> go run main.go

Usage:
  my_exec_file -c [PATH_TO_CONFIG_FILE]

Options:
  -c, [--config]  # path for configuration file
```

### TIP!

Most of the time, you are expecting a *--help* or *-h* to print the help. In that case you can do the following:

```go

if len(os.Args) > 1 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
	sargo.PrintHelp()
}

```

```
$> go run main.go --help
```

or just:

```
$> go run main.go -h
```

### All Methods

| List |
| --- |
| Set(name string, shortCut string, defaultValue interface{}, description string) |
| Get(name string) (string, error) |
| GetString(name string) (string, error) |
| GetInt(name string) (int, error) |
| GetInt32(name string) (int32, error) |
| GetInt64(name string) (int64, error) |
| GetUint(name string) (uint, error) |
| GetUint32(name string) (uint32, error) |
| GetUint64(name string) (uint64, error) |
| GetFloat(name string) (float, error) |
| GetFloat32(name string) (float32, error) |
| GetFloat64(name string) (float64, error) |
| GetBool(name string) (bool, error) |
| PrintHelp() |
| PrintHelpAndExit() |
| SetUsage(tUsage string) |
| GetUsage() string |
