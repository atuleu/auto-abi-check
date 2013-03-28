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


type DebExplorer struct {
	Packages []string
	tmpDir,curDir string
}

func NewDebExplorer(packages []string) (*DebExplorer,error) {
	var err error
	var tmpDir,curDir string

	if tmpDir, err = ioutil.TempDir("","auto-abi-check_debian"); err != nil {
		return nil,err;
	}

	if curDir,err = os.Getwd(); err != nil {
		return nil,err
	}

	return &DebExplorer{packages,tmpDir,curDir},nil
	
}


func (de *DebExplorer) fetchPackage(pack string) error {
	//make sure that we always end up in curDir
	defer os.Chdir(de.curDir)	
	
	var err error
	var absPath string

	//check if this a local file
	if filepath.Ext(pack) != ".deb" {
		if err = os.Chdir(de.tmpDir); err != nil {
			return err
		}

		cmd := exec.Command("apt-get","download",pack)
		if err = cmd.Run(); err != nil {
			return err
		}
		
		//todo find approriate name, and set it in absPath
	} else {
		if absPath, err = filepath.Abs(pack); err != nil {
			return err
		}
	}

	cmd := exec.Command("dpkg","-x",absPath,de.tmpDir)
	if err = cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (de *DebExplorer) Fetch() error {

	for _,p := range de.Packages {
		if err := de.fetchPackage(p); err != nil {
			return err
		}
	}

	return nil
}


func (de *DebExplorer) Clean() error {
	if err := os.RemoveAll(de.tmpDir); err != nil {
		return err
	}
	return nil
}


func (de *DebExplorer) PathInArchive(path string) (string,error){
	return filepath.Join(de.tmpDir,path),nil
}
