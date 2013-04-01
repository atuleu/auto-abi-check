package main

import (
	"path/filepath"
	"io/ioutil"
	"os"
	"os/exec"
)



type DebExplorer struct {
	Packages []string
	tmpDir,curDir string
	architecture string
}

func NewDebExplorer(packages []string) (*DebExplorer,error) {
	var err error
	var tmpDir,curDir,architecture string

	if tmpDir, err = ioutil.TempDir("","auto-abi-check_debian"); err != nil {
		return nil,err;
	}

	if curDir,err = os.Getwd(); err != nil {
		return nil,err
	}


	if architecture , err  = exec.Command(dpkg,"--print-architecture").Output(); err != nil {
		return nil,err
	}

	return &DebExplorer{packages,tmpDir,curDir,architecture},nil
	
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
		
		cmd := exec.Command("apt-get","download","-q",pack)
		out, err := cmd.Output(); 
		if err != nil {
			return err
		}
		
		reg, err := regex.Compile("Get:[0-9]+ " + pack + " ([\S]+)")
		if err != nil {
			return err
		}

		m, err := reg.FindStringSubmatch(out)
		if err != nil {
			return err
		}
		
		name = pack+ "_" + m[1] + "_" + de.architecture + ".deb"

		absPath = filepath.join(de.tmpDir,name)
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
