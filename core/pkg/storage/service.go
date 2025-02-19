package storage_service

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io/fs"
	"io/ioutil"
	"os"
)

func Put(ctx context.Context, diskStorage string, path string, fileName string, data []byte) error {
	if diskStorage == "local" {
		filePath := fmt.Sprintf("%s/%s", path, fileName)

		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll(path, fs.ModePerm)
			if err != nil {
				return err
			}
		}

		file, err := os.Create(filePath)
		defer file.Close()
		if err != nil {
			zap.S().Error(err)
			return err
		}
		err = ioutil.WriteFile(filePath, data, 0644)
		if err != nil {
			zap.S().Error(err)
			return err
		}
	}
	return nil
}
