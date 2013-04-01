package main

import (
	"path/filepath"
	"io/ioutil"
	"os"
	"os/exec"
)


type ArchiveExplorer interface {
	// Fetches all data for archive
	Fetch() error
	//Removes all fetched data
	Clean() error
	
	PathInArchive(path string) (string, error)

}


type LocalFS struct {
}


func (fs LocalFS) Fetch() error {
	return nil
}

func (fs LocalFS) Clean() error {
	return nil
}

func (fs LocalFS) PathInArchive(path string) (string,error) {
	var err error = nil
	var expanded string
	if expanded, err = filepath.EvalSymlinks(path); err != nil {
		return "",err
	}

	var absolute string
	if absolute,err = filepath.Abs(expanded); err != nil {
		return "",err
	}

	return absolute,nil
}


