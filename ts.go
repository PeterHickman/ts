package main

import (
	"bufio"
	"fmt"
	"github.com/lestrrat-go/strftime"
	"os"
	"time"
)

const format = "%FT%T.%f %z"

func main() {
	f, _ := strftime.New(format, strftime.WithMicroseconds('f'))

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		f.Format(os.Stdout, time.Now())
		fmt.Printf(" %s\n", scanner.Text())
	}
}
