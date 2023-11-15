package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sunwavesun/s3-cli/redirect"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Action (select between 'create' and 'remove'): ")
	action, _ := reader.ReadString('\n')
	action = strings.TrimSpace(action)

	if action != "create" && action != "remove" {
		fmt.Println("Unsupported action :(")
		return
	}

	fmt.Print("Bucket Name: ")
	bucket, _ := reader.ReadString('\n')
	bucket = strings.TrimSpace(bucket)

	fmt.Print("Object Name: ")
	object, _ := reader.ReadString('\n')
	object = strings.TrimSpace(object)

	switch action {
	case "create":
		fmt.Print("Redirect Path: ")
		path, _ := reader.ReadString('\n')
		path = strings.TrimSpace(path)

		// New line for readability
		fmt.Println()

		err := redirect.RunAction("create", bucket, object, path)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		fmt.Printf("Redirect from s3://%s/%s to %s created!\n", bucket, object, path)

	case "remove":
		// New line for readability
		fmt.Println()

		err := redirect.RunAction("remove", bucket, object, "")
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		fmt.Printf("Redirect s3://%s/%s removed!\n", bucket, object)
	}

}
