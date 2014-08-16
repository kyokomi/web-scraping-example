package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {

	configFile, err := ReadConfigFile()
	if err != nil {
		log.Fatal("read config error!", err)
	}

	writeDir := strings.Join([]string{configFile.OutputDir, configFile.Keyword}, "/")
	fmt.Println("writeDir ", writeDir)
	if err := os.MkdirAll(writeDir, os.FileMode(0755)); err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	for _, page := range configFile.PageSettings {
		imageUrls := page.GetImagePaths(configFile.Keyword)
		for _, url := range imageUrls {
			wg.Add(1)
			go func(writeDir, url string) {
				writeImage(writeDir, url)
				wg.Done()
			}(writeDir, url)
		}
	}
	wg.Wait()
}

func writeImage(writeDir, url string) {

	idx := strings.LastIndex(url, "/")
	fileName := strings.Join([]string{writeDir, url[idx+1:]}, "/")

	_, err := ioutil.ReadFile(fileName)
	if err == nil {
		return
	}
	fmt.Println("fileName ", fileName)

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	io.Copy(file, res.Body)

	return
}
