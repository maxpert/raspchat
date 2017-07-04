package rasutils

import (
    "os"
)

// CreatePathIfMissing creates a given path if it's missing
func CreatePathIfMissing(path string) error {
    if exists, err := PathExists(path); exists == false && err == nil {
        return os.MkdirAll(path, os.ModePerm)
    } else if err != nil {
        return err
    }

    return nil
}

// PathExists indicates if a path exists or not
func PathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err) {
        return false, nil
    }
    return true, err
}
