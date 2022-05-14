package main

//	Throughout program execution, we often want to create data that isn't needed after the program exits.
//	Temporary files and directories are useful for this purpose since they don't pollute the file system over time

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//	The easiest way to create a temporary file is by calling os.CreateTemp. it creates a file and opens it for reading and writing.
	//	We provide "" as the first argument, so os.CreateTemp will create the file in the default location of our OS
	f, err := os.CreateTemp("", "sample")
	check(err)

	//	Display the name of the temporary file. On Unix-based OSes the directory will likely be /tmp. The file name starts
	//	with the prefix given as the second argument to os.CreateTemp and the rest is chosen automatically to ensure that concurrent calls will always create different filenames

	//	Cleanup the file after we're done. The OS is likely to clean up temporary files by itself after some time, but it's good practice to do this explicitly.
	defer os.Remove(f.Name())

	//	We can write some data to the file
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	//	If we intend to write many temporary files, we may prefer to create a temporary directory. os.MkDirTemp's arguments
	//	are the same as CreateTemp's, but it returns a directory name rather than an open file
	dirname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dirname)

	defer os.RemoveAll(dirname)

	//	Now we can synthesize temporary file names by prefixing them with our temporary directory
	filename := filepath.Join(dirname, "file1")
	err = os.WriteFile(filename, []byte{1, 2, 3, 4}, 0666)
	check(err)
}
