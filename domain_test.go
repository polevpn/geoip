package geoip

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"testing"
)

func TestGetDomianFile(t *testing.T) {
	f, err := os.Open("./direct.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer f.Close()

	br := bufio.NewReader(f)

	list := make([]string, 0)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		out := strings.Trim(string(line), "\n")
		out = strings.Trim(out, "\r")
		list = append(list, out)
	}

	sort.Strings(list)
	for i := 0; i < len(list); i++ {
		fmt.Printf("\"%s\",\n", list[i])
	}

}
