package main

import (
    "fmt"
    "github.com/cbuschka/semver/internal/git"
    "github.com/cbuschka/semver/internal/version"
    "log"
)

func main() {
    tags, err := git.ListTags()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("The output is %v\n", tags)

    versions, _, err := version.FilterConvertAndSortAsc(tags)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("The output is %v\n", versions)
}
