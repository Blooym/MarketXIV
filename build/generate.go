package main

import (
	"os"
)

var build_dir = "./.build_data/"

// Tasks to run as a pre-build hook
func main() {
	cleanup()
	createBuildDir(build_dir)

	generateAPTRepoFile()
	generateDNFRepoFile()
}

func createBuildDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}
}

func createBuildFile(fileName string, data string) {
	file, err := os.Create(build_dir + fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		panic(err)
	}

	file.Sync()
}

func generateDNFRepoFile() {
	fileData := `[BitsOfAByte]            
name=BitsOfAByte Packages         
baseurl=https://packages.bitsofabyte.dev/yum/
enabled=1
gpgcheck=0`
	createBuildFile("bitsofabyte.repo", fileData)
}

func generateAPTRepoFile() {
	fileData := `deb [trusted=yes] https://packages.bitsofabyte.dev/apt/ /`
	createBuildFile("bitsofabyte.list", fileData)
}

func cleanup() {
	os.RemoveAll(build_dir)
}
