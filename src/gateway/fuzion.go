// Copyright 2020 CBDT Technologies. All rights reserved.
// Use of this source code is governed by CBDT Technologies.
//
// Main entrypoint of the application
// Author Aravindh Chinnasamy Subburayar
package main

import (
	"fmt"

	"worker"
)

//
//Main entrypoint of the walkover application
//
func main() {
	// start application
	fmt.Println("Application Started !!!")

	worker.StartWorking(worker.FireStoreWorker{})

	// End application
	fmt.Println("Application Ended. !!!")
}
