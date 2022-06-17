package gilc

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mattn/go-colorable"
)

func YesNoDialog(prompt string, acceptEnter bool) (bool, error) {
	for {
		fmt.Printf("%s [y/n]: ", prompt)
		answer, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return false, err
		}
		switch strings.TrimSpace(strings.ToLower(answer)) {
		case "y", "yes":
			return true, nil
		case "n", "no":
			return false, nil
		default:
			fmt.Printf("'%s' is not a valid answer. Type 'y' or 'n'\n", answer)
		}
	}
}

func ReadLine(prompt string) (string, error) {
	fmt.Print(prompt)
	return bufio.NewReader(os.Stdin).ReadString('\n')
}

func WriteLine(msg string) {
	Write(msg)
	fmt.Print("\n")
}

func Write(msg string) {
	out := colorable.NewColorableStdout()
	indicator := false
	for _, c := range msg {
		// number after the "§"
		if indicator {
			indicator = false
			switch c {
			case '0', '8':
				fmt.Fprint(out, "\033[30m")
				break
			case '1', '9':
				fmt.Fprint(out, "\033[34")
				break
			case '2', 'a':
				fmt.Fprintf(out, "\033[32m")
				break
			case '3', 'b':
				fmt.Fprintf(out, "\033[36m")
				break
			case '4', 'c':
				fmt.Fprintf(out, "\033[31m")
				break
			case '5', 'd':
				fmt.Fprintf(out, "\033[35m")
				break
			case '6', 'e':
				fmt.Fprintf(out, "\033[33m")
				break
			case '7', 'f':
				fmt.Fprintf(out, "\033[37m")
				break
			case 'r':
				fmt.Fprint(out, "\033[0m")
				break
			default:
				break
			}
			continue
		}
		// § -> next is indicator
		if c == '§' {
			indicator = true
			continue
		}
		// just normal letter
		fmt.Print(string(c))
	}
	fmt.Fprint(out, "\033[0m")
}
