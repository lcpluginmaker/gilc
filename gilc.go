package gilc

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/alexcoder04/arrowprint"
)

type IData struct {
	Username     string `json:"Username"`
	SavePath     string `json:"SavePath"`
	DownloadPath string `json:"DownloadPath"`
	Version      string `json:"Version"`
}

type IPlugin struct {
	Name        string
	Description string
	Data        IData
	Main        func(IData)
	Command     func(IData, []string)
	Shutdown    func(IData)
}

var Plugin IPlugin = IPlugin{}

func parseData() IData {
	if len(os.Args) <= 2 {
		fmt.Println("something went wrong, no arguments passed")
		os.Exit(1)
	}
	jsonBytes, err := base64.StdEncoding.DecodeString(os.Args[2])
	if err != nil {
		os.Exit(1)
	}
	res := IData{}
	err = json.Unmarshal(jsonBytes, &res)
	if err != nil {
		os.Exit(1)
	}
	return res
}

func Setup(description string, pmain func(IData), pcommand func(IData, []string), pshutdown func(IData)) {
	p, err := os.Executable()
	if err != nil {
		arrowprint.Err0("cannot get executable path")
		os.Exit(1)
	}
	_, Plugin.Name = path.Split(p)
	Plugin.Description = description
	Plugin.Data = parseData()
	Plugin.Main = pmain
	Plugin.Command = pcommand
	Plugin.Shutdown = pshutdown
}

func Run() {
	switch os.Args[1] {
	case "init":
		fmt.Printf(`{"Description": "%s"}`+"\n", Plugin.Description)
	case "main":
		Plugin.Main(Plugin.Data)
	case "shutdown":
		Plugin.Shutdown(Plugin.Data)
	default:
		Plugin.Command(Plugin.Data, os.Args[3:])
	}
}
