package internal

import "os"


var(
	FolderName string = ".gogit"
)

func CreateGoGitFolder() {
	os.Mkdir(FolderName, os.ModePerm)
}