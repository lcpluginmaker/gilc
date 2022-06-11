
# gilc - GoILeoConsole

## Example plugin using gilc

```go
package main

import (
	"fmt"

	"github.com/alexcoder04/gilc"
)

func pmain(data gilc.IData) {
	fmt.Println("PluginMain runs")
	fmt.Println("LeoConsole Version: ", data.Version)
}

func pcommand(data gilc.IData, args []string) {
	fmt.Println("My plugin was called")
	fmt.Println("Args: ", args)
	fmt.Println("SavePath: ", data.SavePath)
}

func pshutdown(data gilc.IData) {
	fmt.Println("PluginShutdown runs")
}

func main() {
	gilc.Setup("test plugin", pmain, pcommand, pshutdown)
	gilc.Run()
}

```

## API Documentation

### IData type

```go
type IData struct {
	Username           string
	SavePath           string
	DownloadPath       string
	Version            string
}
```

### IO functions

```go
func YesNoDialog(prompt string, acceptEnter bool) (bool, error) { }
```

Yes-No dialog with custom prompt.

```go
func ReadLine(prompt string) (string, error) {
```

Input prompt.

```go
func Write(msg string) {
```

Writes text parsing it according to teh LeoConsole standard (`§n`).

```go
func WriteLine(msg string) {
```

Same as `Write(msg)`, but adds a newline at the end.

