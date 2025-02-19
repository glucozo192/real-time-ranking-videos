package utils

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/klauspost/compress/zip"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func DownloadFile(url, destination string) error {
	// Create a file to write the downloaded content
	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data from the URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error: %s", resp.Status)
	}

	// Copy the response body to the output file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func MkdirAll(PathStorage string) error {
	if _, err := os.Stat(PathStorage); os.IsNotExist(err) {
		return os.MkdirAll(PathStorage, 0777)
	}
	return nil
}

func CheckFileFromUrl(ctx context.Context, Url string) error {
	//download file
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodHead, Url, nil)
	if err != nil {
		return err
	}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 && response.StatusCode != 307 {
		return errors.New(fmt.Sprintf("link does not exist :%s", Url))
	}
	return nil
}

func DownloadFileFromUrl(ctx context.Context, Url string, pathStorage string, fileName string) error {
	Url = strings.TrimSpace(Url)
	//Mkdir
	if err := MkdirAll(pathStorage); err != nil {
		return errors.New("folder creation failed " + pathStorage)
	}
	if err := CheckFileFromUrl(ctx, Url); err != nil {
		return err
	}
	//download file
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, Url, nil)
	if err != nil {
		return err
	}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 && response.StatusCode != 307 {
		return errors.New(fmt.Sprintf("link does not exist :%s", Url))
	}
	defer response.Body.Close()
	body := response.Body

	// Create the file
	if fileName == "" {
		fileName = path.Base(Url)
	}
	dir := path.Dir(fileName)
	if dir != "" && dir != "." {
		errMkdir := MkdirAll(pathStorage + "/" + dir)
		if errMkdir != nil {
			return errors.New("folder creation failed " + pathStorage + "/" + dir)
		}
	}
	out, err := os.Create(pathStorage + "/" + fileName)
	defer out.Close()
	if err != nil {
		return err
	}

	// Write the body to file
	_, err = io.Copy(out, body)
	if err != nil {
		return err
	}

	return nil
}

func DownloadFileFromUrlV2(ctx context.Context, Url string, pathStorage string, fileName string) error {
	Url = strings.TrimSpace(Url)
	//Mkdir
	if err := MkdirAll(pathStorage); err != nil {
		return errors.New("folder creation failed " + pathStorage)
	}

	if strings.Contains(Url, "drive.google.com") {
		err := DownloadFileFromGGDrive(ctx, Url, pathStorage, fileName)
		if err != nil {
			return err
		}
		return nil
	}

	if err := CheckFileFromUrl(ctx, Url); err != nil {
		return err
	}
	//download file
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, Url, nil)
	if err != nil {
		return err
	}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 && response.StatusCode != 307 {
		return errors.New(fmt.Sprintf("link does not exist :%s", Url))
	}
	defer response.Body.Close()
	body := response.Body

	// Create the file
	if fileName == "" {
		fileName = path.Base(Url)
	}
	dir := path.Dir(fileName)
	if dir != "" && dir != "." {
		errMkdir := MkdirAll(pathStorage + "/" + dir)
		if errMkdir != nil {
			return errors.New("folder creation failed " + pathStorage + "/" + dir)
		}
	}
	out, err := os.Create(pathStorage + "/" + fileName)
	defer out.Close()
	if err != nil {
		return err
	}

	// Write the body to file
	_, err = io.Copy(out, body)
	if err != nil {
		return err
	}

	return nil
}

func ZipFolder2(source, target string) error {
	err := os.MkdirAll(path.Dir(target), fs.ModePerm)
	if err != nil {
		return err
	}
	// 1. Create a ZIP file and zip.Writer
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	// 2. Go through all the files of the source
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if source == path {
			return nil
		}
		if err != nil {
			return err
		}

		// 3. Create a local file header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// set compression
		header.Method = zip.Deflate

		// 4. Set relative path of a file as the header name
		header.Name, err = filepath.Rel(source, path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}

		// 5. Create writer for the file header and save content of the file
		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(headerWriter, f)

		if err != nil {
			return err
		}

		return nil
	})
}

func GetGGDriveDownloadUrl(url string) (string, error) {
	if strings.Contains(url, "drive.google.com") {
		var idFile string
		var startIndex int
		var endIndex int
		if strings.Contains(url, "&usp=drive_copy") {
			startIndex = strings.Index(url, "id=") + 3
			endIndex = strings.Index(url, "&usp=drive_copy")
		} else if strings.Contains(url, "file/d/") {
			startIndex = strings.Index(url, "file/d/") + 7
			endIndex = strings.Index(url, "/view")
		} else {
			return "", fmt.Errorf("invalid url: %s", url)
		}
		if startIndex == -1 || endIndex == -1 {
			return "", fmt.Errorf("invalid url: %s", url)
		}
		idFile = url[startIndex:endIndex]
		url = fmt.Sprintf("https://drive.google.com/u/0/uc?id=%s&export=download", idFile)
	}

	return url, nil
}

func DownloadFileFromGGDrive(ctx context.Context, Url string, pathStorage string, fileName string) error {
	serviceAccountFile := "google-drive-key.json"
	outputFilePath := pathStorage + "/" + fileName

	fileId, err := GetIDFileGGDrive(Url)

	b, err := os.ReadFile(serviceAccountFile)
	if err != nil {
		return err
	}

	config, err := google.JWTConfigFromJSON(b, drive.DriveReadonlyScope)
	if err != nil {
		return err
	}

	client := config.Client(ctx)

	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return err
	}

	res, err := srv.Files.Get(fileId).Download()
	if err != nil {
		return err
	}
	defer res.Body.Close()

	outFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	if _, err := io.Copy(outFile, res.Body); err != nil {
		return err
	}
	fmt.Printf("File downloaded successfully:%s\n", fileName)
	return nil
}

func GetIDFileGGDrive(url string) (string, error) {
	var idFile string
	var startIndex int
	var endIndex int
	if strings.Contains(url, "&usp=drive_copy") {
		startIndex = strings.Index(url, "id=") + 3
		endIndex = strings.Index(url, "&usp=drive_copy")
	} else if strings.Contains(url, "file/d/") {
		startIndex = strings.Index(url, "file/d/") + 7
		endIndex = strings.Index(url, "/view")
	} else {
		return "", errors.New(fmt.Sprintf("invalid url: %s", url))
	}
	if startIndex == -1 || endIndex == -1 {
		return "", errors.New(fmt.Sprintf("invalid url: %s", url))
	}
	idFile = url[startIndex:endIndex]
	return idFile, nil
}
