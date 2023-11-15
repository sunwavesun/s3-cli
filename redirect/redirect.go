package redirect

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func CreateRedirect(bucket, object, path string) error {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	output, err := uploadObject(bucket, object, path)
	if err != nil {
		log.Printf(err.Error())
		return err
	}

	log.Printf(string(output))

	return nil
}

func uploadObject(bucket, object, path string) ([]byte, error) {
	// Create temp file named object
	tempFile, err := os.Create(object)
	if err != nil {
		return nil, err
	}
	defer tempFile.Close()

	// Call S3-cli command and write to log
	bucketPath := fmt.Sprintf("s3://%s/%s", bucket, object)
	cmd := exec.Command("aws", "s3", "cp", object, bucketPath, "--acl", "public-read", "--website-redirect", path)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// Delete temp file
	err = os.Remove(object)
	if err != nil {
		return nil, err
	}

	return output, nil
}
