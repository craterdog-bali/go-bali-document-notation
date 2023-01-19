/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package utilities

import (
	fmt "fmt"
	sts "strings"
	osx "os"
)

// TYPE DEFINITIONS

type (
)

// CONSTANT DEFINITIONS

const (
	EOL = "\n" // The POSIX end of line character.
)

// CONFIGURATOR IMPLEMENTATION

// This constructor creates a new configurator using the specified directory and
// filename.
func Configurator(directory, filename string) *configurator {
	var err error
	// Validate the configuration filename.
	if len(filename) == 0 {
		var message = "The configurator requires a filename."
		panic(message)
	}
	// Validate the configuration directory.
	if len(directory) == 0 {
		directory, err = osx.UserHomeDir()
		if err != nil {
			var message = fmt.Sprintf("Could not determine the user's home directory: %v.", err)
			panic(message)
		}
	}
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	// Create the configuration directory (if necessary).
	directory += ".bali/"
	err = osx.MkdirAll(directory, 0700)
	if err != nil {
		var message = fmt.Sprintf("Could not create the configuration directory: %v.", err)
		panic(message)
	}
	// Create the configurator.
	var file = directory + filename;
	return &configurator{file}
}

// This type defines the structure and methods associated with a configurator
// agent.
type configurator struct {
	file string
}

// PUBLIC INTERFACE

// This method loads the current configuration stored in the configuration file.
func (v *configurator) Load() string {
	var bytes, err = osx.ReadFile(v.file)
	if err != nil {
		var message = fmt.Sprintf("Could not read the configuration file: %v.", err)
		panic(message)
	}
	var configuration = string(bytes[:len(bytes)-1]) // Remove the POSIX EOL character.
	return configuration
}

// This method stores the current configuration in the configuration file.
func (v *configurator) Store(configuration string) {
	var bytes = []byte(configuration + EOL) // Append the POSIX EOL character.
	var err = osx.WriteFile(v.file, bytes, 0600)
	if err != nil {
		var message = fmt.Sprintf("Could not create the configuration file: %v.", err)
		panic(message)
	}
}

// This method deletes the configuration file.
func (v *configurator) Delete() {
	var err = osx.Remove(v.file)
	if err != nil {
		var message = fmt.Sprintf("Could not delete the configuration file: %v.", err)
		panic(message)
	}
}
