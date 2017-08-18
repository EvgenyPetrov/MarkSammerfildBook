package main

import (
    "flag"
    "log"
    "strings"
)

func main() {
    minSize, maxSize, suffixes, files := handleCommandLine()
    sink(filterSize(minSize, maxSize, filterSuffixes(suffixes, source(files))))
}

func handleCommandLine() ( minSize, maxSize int64, suffixes, files []string) {
    flag.Int64Var(&minSize, "min", -1,
        "minimum file size (-1 means no minimum)")
    flag.Int64Var(&maxSize, "max", -1,
        "maximum file size (-1 means no maximum)")
    var suffixesOpt *string = flag.String("suffixes", "",
        "comma-separated list of file suffixes")
    flag.Parse()
    if minSize > maxSize && maxSize != -1 {
        log.Fatalln("minimum size must be < maximum size")
    }
    suffixes = []string{}
    if *suffixesOpt != "" {
        suffixes = strings.Split(*suffixesOpt, ",")
    }
    files = flag.Args()
    return  minSize, maxSize, suffixes, files
}
