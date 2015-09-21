package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

func match(line string, exps [](*regexp.Regexp)) bool {
	for _, re := range exps {
		if re.MatchString(line) {
			return true
		}
	}
	return false
}

func usage() {
	fmt.Println("command | timestamping regexp [regexp] ...")
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		usage()
		return
	}

	t := time.Now()

	var exps [](*regexp.Regexp)
	for _, r := range args {
		compiled, err := regexp.Compile(r)
		if err != nil {
			log.Panic(err)
			return
		}
		exps = append(exps, compiled)
	}

	fscanner := bufio.NewScanner(os.Stdin)
	for fscanner.Scan() {
		line := fscanner.Text()
		if match(line, exps) {
			last := t
			t = time.Now()
			fmt.Println(time.Now().Unix(), "secs", fmt.Sprintf("%.2f", t.Sub(last).Seconds()), line)
		} else {
			fmt.Println(line)
		}
	}

}
