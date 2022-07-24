
# gilc - GoILeoConsole

[Why use gilc?](./why.html)

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
    ans, err := gilc.ReadLine("Your favourite fruit")
    if err != nil {
        return
    }
    gilc.WriteLine("Your favourite fruit is §4" + ans + "§r.");
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
func ReadLine(prompt string) (string, error) { }
```

Input prompt.

```go
func Write(msg string) { }
```

Writes text parsing it according to teh LeoConsole standard (`§n`).

```go
func WriteLine(msg string) { }
```

Same as `Write(msg)`, but adds a newline at the end.

## More recommended libraries

Other libraries that might be useful for plugin development:

 - [alexcoder04/arrowprint](https://github.com/alexcoder04/arrowprint) - uniform message formatting in LeoConsole (same as in C# ILeoConsole and apkg)
 - [alexcoder04/friendly](https://github.com/alexcoder04/friendly) - more nice functions to use

