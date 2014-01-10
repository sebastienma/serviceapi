package main

import (
        "github.com/subosito/iglo"
        "log"
        "flag"
        "os"
)

var fname = flag.String("out", "output.html", "Filename of the HTML output")
var fmd = flag.String("mdfile", "API.md", "Blueprint API filename")

func main() {
        flag.Parse()

        f, err := os.Open(*fmd)
        if err != nil {
                log.Fatal(err)
        }
        defer f.Close()

        w, err := os.Create(*fname)
        if err != nil {
                log.Fatal(err)
        }
        defer w.Close()

        err = iglo.MarkdownToHTML(w, f)
        if err != nil {
                log.Fatal(err)
        }
}
