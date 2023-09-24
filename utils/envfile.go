package utils

import (
    "fmt"
    "os"
    "strings"
)

func WriteToEnvFile(paramName string, paramValue string, path string) error {
    if err := os.MkdirAll(path, os.ModePerm); err != nil {
        return fmt.Errorf("error creating directory: %v", err)
    }

    content := strings.ToUpper(paramName) + "=" + paramValue
    f, err := os.OpenFile(path+"/.env", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return fmt.Errorf("error creating .env file: %v", err)
    }
    defer f.Close()

    fmt.Println("Writing to .env file", path)
    n, err := f.WriteString(content + "\n")
    if err != nil {
        return fmt.Errorf("error writing to .env file: %v", err)
    }
    fmt.Printf("Wrote %d bytes to .env file\n", n)

    err = f.Sync()
    if err != nil {
        return fmt.Errorf("error syncing .env file: %v", err)
    }

    return nil
}
