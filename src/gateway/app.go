// Copyright 2020 CBDT Technologies. All rights reserved.
// Use of this source code is governed by CBDT Technologies.
//
// Main entrypoint of the application
// Author Aravindh Chinnasamy Subburayar
package main

import (
	"config"
	"firestoreconnector"
	"fmt"
	"os"
	"utils"
	"worker"
)

//
//Main entrypoint of the walkover application
//
func main() {
	// start application
	fmt.Println("Application Started...")

	argsWithoutProgramName := os.Args[1:]

	{
		utils.ValidateArguments(argsWithoutProgramName)
		utils.PopulateArguments(argsWithoutProgramName)
		config.InitializeConfigFile()
		if config.FRAMEWORK_STREAM == "FB" {
			hitFirebase()
		} else if config.FRAMEWORK_STREAM == "FS" {
			hitFireStore()
		}
	}

	// End application
	fmt.Println("Application Ended.")
}

func hitFirebase() {

	// inputURL := "https://webcat-284707.firebaseio.com/"

	// firebaseconnector.ConnectFireBase(inputURL)
	// firebaseconnector.IngestDataToFireBase("", "", "")

	worker.StartWorking(worker.FirebaseWorker{})
}

func hitFireStore() {

	firestoreconnector.FuzionMatch()
}
