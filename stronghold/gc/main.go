package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calcGC(s string) float32 {
	GC := 0
	for i := range s {
		if s[i] == 'G' || s[i] == 'C' {
			GC += 1
		}
	}

	return 100 * float32(GC) / float32(len(s))
}

func main() {
	file, err := os.Open("gc.in")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var maxGC float32
	var header, maxGCheader string
	builder := strings.Builder{}

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		prevheader := header
		line := scanner.Text()

		if line[0] == '>' {
			header = line[1:]
			if builder.Len() > 0 {
				GCcontent := calcGC(builder.String())
				builder.Reset()

				if GCcontent > maxGC {
					maxGCheader = prevheader
					maxGC = GCcontent
				}
			}
		} else {
			builder.Write([]byte(line))
		}
	}

	GCcontent := calcGC(builder.String())
	builder.Reset()

	if GCcontent > maxGC {
		maxGCheader = header
		maxGC = GCcontent
	}

	fmt.Printf("%s\n", maxGCheader)
	fmt.Printf("%f\n", maxGC)
}
