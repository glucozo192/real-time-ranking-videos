package config

import (
	"fmt"
	"os"
)

func (c conf) GetMediaCdn() string {
	return c.MediaCdn
}

func (c conf) GetMediaDisplay() string {
	return c.MediaDisplay
}

func (c conf) GetMediaS3() string {
	return c.MediaS3
}

func (c conf) GetPathUpload() string {
	path, _ := os.Getwd()
	return fmt.Sprintf("%s/src/storage/app", path)
	if c.App.Env == "dev" {
		return "/mnt/volume_sgp1_dev/data/platform-dev"
	} else if c.App.Env == "live" {
		return "/opt/data/platform"
	}

	return "/var/www/html/data_story"
}
