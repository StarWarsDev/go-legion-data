package legiondata

import (
	"encoding/json"
	"gopkg.in/src-d/go-git.v4"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	repoURL  = "https://github.com/zacharyp/sw-legion-data"
	repoPath = "/tmp/legion-data"
	jsonPath = repoPath + "/lib/legion-data.json"
)

// GetData fetches the latest out/legion-data.json file, parses it and returns a Data struct
func GetData() (Data, error) {
	var data Data

	// remove cloned repo if it is already there
	if fileExists(repoPath) {
		err := removeDir(repoPath)
		if err != nil {
			log.Println("error removing repo temp directory", repoPath, err)
			return data, err
		}
	}

	// clone the repo
	_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
	})
	if err != nil {
		return data, err
	}

	// yarn install
	err = yarn("install", repoPath)
	if err != nil {
		return data, err
	}

	// yarn json
	err = yarn("json", repoPath)
	if err != nil {
		return data, err
	}

	// read JSON into memory
	bytes, err := getFileContent(jsonPath)
	if err != nil {
		return data, err
	}
	_ = json.Unmarshal(bytes, &data)

	// remove cloned repo
	err = removeDir(repoPath)
	if err != nil {
		// Getting this error is not a big deal.
		// At this point we have the data, return it.
		// But we should still log the error.
		log.Println("error removing repo temp directory", repoPath, err)
	}

	return data, nil
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func removeDir(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func yarn(script, path string) error {
	cmd := exec.Command("yarn", script)
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}

func getFileContent(path string) ([]byte, error) {
	var bytes []byte
	file, err := os.Open(path)
	if err != nil {
		return bytes, err
	}
	defer file.Close()

	bytes, _ = ioutil.ReadAll(file)

	return bytes, nil
}
