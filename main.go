package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
)

func error(msg string) {
	os.Stderr.WriteString(fmt.Sprintf("%v\n", msg))
	os.Exit(1)
}

func input() string {
	reader := bufio.NewReader(os.Stdin)

	line, err := reader.ReadString('\n')

	if err != nil {
		error(err.Error())
	}

	return line[:len(line)-1]
}

func decode(text string, depth int) string {
	if depth == 0 {
		return text
	}

	d, err := url.QueryUnescape(text)

	if err != nil {
		error(err.Error())
	}

	return decode(d, depth-1)
}

func encode(text string, depth int) string {
	if depth == 0 {
		return text
	}

	return encode(url.PathEscape(text), depth-1)
}

func print(msg string) {
	os.Stdout.WriteString(fmt.Sprintf("%v\n", msg))
}

func main() {
	shouldEncode := flag.Bool("e", false, "encode text")
	depth := flag.Int("depth", 1, "how many times the text should be encoded/decoded")

	flag.Parse()

	text := input()

	if *shouldEncode {
		print(encode(text, *depth))
	} else {
		print(decode(text, *depth))
	}
}
