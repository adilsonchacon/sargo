# sargo

Simple Args in Go: command line parser


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

### Command Line Example 1

```
$> go run main.go
```


### Output

```
host: localhost
port: 8081
rate: 5.5
activated: true
```

### Command Line Example 2

```
$> go run main.go --host=10.0.0.1 --port=9091 --rate=3.14 --activated=false
```

### Output

```
host: 10.0.0.1
port: 9091
rate: 3.14
activated: false
```


### Command Line Example 3

```
$> go run main.go -h 10.0.0.1 -p 9091 -r 3.14 -a false
```

### Output

```
host: 10.0.0.1
port: 9091
rate: 3.14
activated: false
```

### Command Line Example 4

```
$> go run main.go --host=127.0.0.1 -p 9091 -a 1
```

### Output

```
host: 127.0.0.1
port: 9091
rate: 5.5
activated: true
```

### Command Line Example 5

```
$> go run main.go -r 938.291
```

### Output

```
host: localhost
port: 8081
rate: 938.291
activated: true
```

### Command Line Example 6

```
$> go run main.go --activated=f
```

### Output

```
host: localhost
port: 8081
rate: 5.5
activated: false
```

## Print Help Example

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
	sargo.Set("activated", "a", true, "activated. Default value is \"true\"")
  
  sargo.PrintHelp()
}
```

### Output

```
Usage:
  go run main.go [options]

Options:
  --h, [--host]       # http server host. Default value is "localhost"
  --p, [--port]       # http server port. Default value is "8081"
  --r, [--rate]       # site rate. Default value is "5.5"
  --a, [--activated]  # activated. Default value is "true"
```