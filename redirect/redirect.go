package redirect

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func RunAction(action, bucket, object, path string) error {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	var output []byte
	switch action {
	case "create":
		output, err = createRedirect(bucket, object, path)
	case "remove":
		output, err = removeRedirect(bucket, object)
	}

	if err != nil {
		log.Printf(err.Error())
		return err
	}

	log.Printf(string(output))

	return nil
}

func createRedirect(bucket, object, path string) ([]byte, error) {
	// Create temp file named object
	tempFile, err := os.Create(object)
	if err != nil {
		return nil, err
	}
	defer tempFile.Close()

	// Call S3-cli command and write to log
	bucketPath := getBucketPath(bucket, object)
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

func removeRedirect(bucket, object string) ([]byte, error) {
	bucketPath := getBucketPath(bucket, object)
	cmd := exec.Command("aws", "s3", "rm", bucketPath)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return output, nil
}
func getBucketPath(bucket, object string) string {
	return fmt.Sprintf("s3://%s/%s", bucket, object)
}
