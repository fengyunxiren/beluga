package utils

import "os"

func CheckAndCreatePath(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(path, os.ModePerm)
			return err
		} else {
			return err
		}
	}
	return nil
}
