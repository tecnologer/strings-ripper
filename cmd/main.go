package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/tecnologer/string-shunker/cmd/cli"
)

func main() {
	app := cli.NewApp()

	setInputToArgs()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func setInputToArgs() {
	input, err := readStringFromPipe()
	if err != nil {
		log.Fatal(err)
	}

	aliasIndex := indexOfArg("-i")
	flagIndex := indexOfArg("--input")

	if aliasIndex != -1 || flagIndex != -1 {
		return
	}

	os.Args = append(os.Args, "-i", input)
}

func indexOfArg(arg string) int {
	for i, v := range os.Args {
		if v == arg {
			return i
		}
	}

	return -1
}

func readStringFromPipe() (string, error) {
	nBytes, nChunks := int64(0), int64(0)
	r := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, 4*1024)

	var input string

	for {

		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]

		if n == 0 {

			if err == nil {
				continue
			}

			if err == io.EOF {
				break
			}

			log.Fatal(err)
		}

		nChunks++
		nBytes += int64(len(buf))
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		input += string(buf)
	}

	return input, nil
}
