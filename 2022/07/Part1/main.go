package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	helpers "github.com/maxpowerfreak/advent-of-code"
)

func main() {
	body, err := helpers.GetInputResponseBody(2022, 7)
	if err != nil {
		panic(err)
	}
	defer body.Close()

	pageScanner := bufio.NewScanner(body)
	pageScanner.Split(bufio.ScanLines)

	currentDirectory := &fileSystemEntry{
		name:                 "/",
		parent:               nil,
		size:                 0,
		isDir:                true,
		subFileSystemEntries: make(map[string]*fileSystemEntry),
	}

	// skip first line, that's cd /
	pageScanner.Scan()

	// used to keep track of directories to calculate size of each later
	directories := []*fileSystemEntry{}
	// build directory tree
	for pageScanner.Scan() {
		line := strings.Fields(pageScanner.Text())

		// $ cd {anything}
		// other entries all have 2 values ($ ls, dir dirName, fileSize size)
		if len(line) > 2 {
			// $ cd ..
			if line[2] == ".." {
				currentDirectory = currentDirectory.parent
			} else {
				// else cd to one of the subdirectories
				currentDirectory = currentDirectory.subFileSystemEntries[line[2]]
			}
		} else if line[0] == "dir" {
			// create a new subdirectory
			currentDirectory.subFileSystemEntries[line[1]] = &fileSystemEntry{
				name:                 line[1],
				parent:               currentDirectory,
				size:                 0,
				isDir:                true,
				subFileSystemEntries: make(map[string]*fileSystemEntry),
			}

			// It's a directory, we append it to our slice
			directories = append(directories, currentDirectory.subFileSystemEntries[line[1]])
		} else if len(line[0]) > 1 {
			// we're effectively skipping `$ ls`
			// last case is `fileSize fileName`
			fileSize, _ := strconv.Atoi(line[0])
			currentDirectory.subFileSystemEntries[line[1]] = &fileSystemEntry{
				name:                 line[1],
				parent:               currentDirectory,
				size:                 fileSize,
				isDir:                false,
				subFileSystemEntries: nil,
			}
		}
	}

	var totalSmallSize int
	for _, directory := range directories {
		size := sumSizes(*directory)

		if size < 100000 {
			totalSmallSize += size
		}
	}

	fmt.Println(totalSmallSize)
}

func sumSizes(entry fileSystemEntry) int {
	if !entry.isDir {
		// it's a file, give me the size
		return entry.size
	}

	var size int
	// R E C U R S E
	for _, directory := range entry.subFileSystemEntries {
		size += sumSizes(*directory)
	}

	return size
}

type fileSystemEntry struct {
	name                 string
	parent               *fileSystemEntry
	size                 int
	isDir                bool
	subFileSystemEntries map[string]*fileSystemEntry
}
