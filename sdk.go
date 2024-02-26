package main

import (
	"fmt"
	"os"
	"time"

	ld "github.com/launchdarkly/go-server-sdk/v7"
	"github.com/launchdarkly/go-server-sdk/v7/ldcomponents"
)

func main() {
	// Replace "YOUR_SDK_KEY" with the actual SDK key you want to verify
	sdkKey := "YOUR_SDK_KEY"

	// Define SDK configuration
	sdkConfig := ld.Config{
		HTTP: ldcomponents.HTTPConfiguration().ConnectTimeout(5 * time.Second),
		// Attempt to disable logging or set to a default minimal setup
		Logging: ldcomponents.NoLogging(), // This disables logging output
	}

	// Attempt to initialize the LaunchDarkly client with the SDK key
	client, err := ld.MakeCustomClient(sdkKey, sdkConfig, 10*time.Second)

	if err != nil {
		if err == ld.ErrInitializationFailed {
			fmt.Println("SDK Key verification failed: Initialization with SDK key failed")
		} else {
			fmt.Printf("SDK Key verification failed: %v\n", err)
		}
		os.Exit(1)
	} else {
		fmt.Println("SDK Key is valid.")
	}

	// Close the client when it is no longer needed
	defer client.Close()
}
