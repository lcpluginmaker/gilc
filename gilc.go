package gilc

import (
	"encoding/base64"
	"encoding/json"
	"os"
)

type IData struct {
	Username           string `json:"Username"`
	SavePath           string `json:"SavePath"`
	DownloadPath       string `json:"DownloadPath"`
	CurrentWorkingPath string `json:"CurrentWorkingPath"`
	Version            string `json:"Version"`
}

func ParseData() (IData, error) {
	jsonBytes, err := base64.StdEncoding.DecodeString(os.Args[1])
	if err != nil {
		return IData{}, err
	}
	res := IData{}
	err = json.Unmarshal(jsonBytes, &res)
	return res, err
}
