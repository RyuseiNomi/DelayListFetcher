package JsonWorker

import (
	"fmt"
	"log"
	"os"
)

var (
	tempDir    = "/tmp/json/"
	bucketName = "delay-list"
	key        = "delay-list.json"
)

/**
 * Create empty JSON file on temp file and write fetched delay-list
 */
func CreateJSON(delayList []byte) error {
	if delayList == nil {
		return fmt.Errorf("nil bytes was given")
	}

	if err := os.MkdirAll(tempDir, 0777); err != nil {
		return err
	}

	file, err := os.Create(tempDir + key)
	if err != nil {
		return err
	}

	_, err = file.Write(delayList)
	if err != nil {
		return err
	}

	if isExist := isExistTempFile(tempDir); isExist != true {
		return fmt.Errorf("Temp file does not exist")
	}

	log.Printf("Success to create JSON File!")
	return nil
}

/**
 * Verify if can create temp directory
 */
func isExistTempFile(tempFile string) bool {
	_, err := os.Stat(tempFile)
	return !os.IsNotExist(err)
}
