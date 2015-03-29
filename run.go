package main

import (
	"github.com/haridas/sdiscover/store"
)

const (
	masterMode = "master"
	agentMode  = "agent"
)

// Struct which abstract all the required objects to do the application
// logics.
type SDiscover struct {
	Store store.Store
	Conf  SConf
}

// All Application speicfic Errors.
type SError struct {
	Eid     string
	Message string
}

// Initiate the process.
func (sd *SDiscover) Run() error {
	return nil
}

// Run as Master mode
func (sd *SDiscover) RunAsMaster() error {
	return nil
}

// Run as Agent mode
func (sd *SDiscover) RunAsAgent() error {
	return nil
}

func (sd *SDiscover) checkWhichModeToRun() (string, error) {
	if sd.Conf.ConsulMode == masterMode {
		//
	} else if sd.Conf.ConsulMode == agentMode {
		//
	} else {
		// UnKnown Working mode. panic.
	}
	return "", nil
}

func (sd *SDiscover) checkNumberOfMasters() int {
	return 0
}
