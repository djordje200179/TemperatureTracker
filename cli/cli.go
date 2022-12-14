package cli

import (
	"TemperatureTracker/data/storage"
	"TemperatureTracker/data/storage/cache"
	"bufio"
	"fmt"
	"io"
	"strings"
)

type CLI struct {
	Storage storage.Storage
	Cache   storage.Cache

	finished bool

	*bufio.Reader
	*bufio.Writer
}

func New(storage storage.Storage, reader io.Reader, writer io.Writer) *CLI {
	cli := &CLI{
		Storage: storage,
		Cache:   cache.Instance(),

		Reader: bufio.NewReader(reader),
		Writer: bufio.NewWriter(writer),
	}

	return cli
}

func (cli *CLI) Handle() {
	for !cli.finished {
		line := cli.readLine()
		cli.handleLine(line)
	}
}

func (cli *CLI) readLine() string {
	cli.WriteString("> ")
	cli.Flush()

	text, _ := cli.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	return text
}

func (cli *CLI) handleLine(line string) {
	command := line

	switch command {
	case "exit":
		cli.finished = true
	case "get_latest":
		latestData := cli.Cache.GetLatestReadings()

		if len(latestData) != 0 {
			for _, data := range latestData {
				cli.WriteString(data.String())
			}
		} else {
			fmt.Println("No readings found")
		}
	}
}
