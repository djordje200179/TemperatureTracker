package cli

import (
	"TemperatureTracker/data/storage"
	"TemperatureTracker/data/storage/cache"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CLI struct {
	storage storage.Storage
	cache   storage.Cache

	reader *bufio.Reader
}

func Start(storage storage.Storage) *CLI {
	cli := &CLI{
		storage: storage,
		cache:   cache.Instance(),

		reader: bufio.NewReader(os.Stdin),
	}

	go cli.Handle()

	return cli
}

func (cli *CLI) Handle() {
	for {
		fmt.Print("> ")

		text, _ := cli.reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")

		switch text {
		case "exit":
			return
		case "get_latest":
			latestData := cli.cache.GetLatestReadings()

			if len(latestData) != 0 {
				for _, data := range latestData {
					fmt.Println(data)
				}
			} else {
				fmt.Println("No readings found")
			}
		}
	}
}
