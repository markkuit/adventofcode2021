package main

import (
        "bufio"
        "fmt"
        "log"
        "os"
        "strconv"
)

var (
        last      int
        increases = -1
)

func main() {
        f, err := os.Open("input")
        if err != nil {
                log.Fatal(err)
        }
        defer f.Close()

        s := bufio.NewScanner(f)
        for s.Scan() {
                value, err := strconv.Atoi(s.Text())
                if err != nil {
                        log.Fatal(err)
                }

                if value > last {
                        increases++
                }
                last = value
        }

        fmt.Println(increases)
}
