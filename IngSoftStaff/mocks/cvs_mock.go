package mocks

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "io"
    "os"

    "golang.org/x/text/encoding/charmap"
)

type CSVReader interface {
    Read() ([]string, error)
    ReadAll() ([][]string, error)
}

