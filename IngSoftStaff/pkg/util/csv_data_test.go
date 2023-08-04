package util

import (
    "testing"
)

func TestReadData(t *testing.T) {
    // Test with a valid file
    err := ReadData("../db/profesIngSoft.csv")
    if err != nil {
        t.Errorf("Expected nil error, but got %v", err)
    }

    // Test with a non-existent file
    err = ReadData("non_existent_file.csv")
    if err == nil {
        t.Errorf("Expected non-nil error, but got nil")
    }
}