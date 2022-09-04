package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/fatih/color"
)

type data struct {
	Photo        string
	Message      string
	Name         string
	Delay        time.Duration
	RandomAge    bool
	MaxRandomAge int
	Age          int
}

func main() {

	config, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(" [-] Заполните файл ./config.json")
	}
	var res data
	json.Unmarshal(config, &res)

	yellow := color.New(color.FgYellow)
	yellow.Println(" [!] Для запуска нажмите enter")
	fmt.Scanln()

	spam(res.Photo, res.Message, res.Name, res.Delay, res.Age, res.MaxRandomAge, res.RandomAge)
}
