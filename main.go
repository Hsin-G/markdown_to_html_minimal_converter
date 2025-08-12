package main

import (
    "io/ioutil"
    "log"
    "os"
    "strings"

    "github.com/yuin/goldmark"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Usage: go run main.go fichier.mk")
    }
    inPath := os.Args[1]

    input, err := ioutil.ReadFile(inPath)
    if err != nil {
        log.Fatal(err)
    }

    var buf strings.Builder
    if err := goldmark.Convert(input, &buf); err != nil {
        log.Fatal(err)
    }

    outPath := strings.TrimSuffix(inPath, ".mk") + ".html"
    err = ioutil.WriteFile(outPath, []byte(buf.String()), 0644)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Converti %s en %s\n", inPath, outPath)
}

