package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	MAXSIZEDIR  = 100000
	CAPACITY    = 70000000
	NEEDEDSPACE = 30000000
)

type Directory struct {
	name      string
	parent    *Directory
	children  map[string]*Directory
	files     map[string]int
	totalSize int
}

func main() {
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	dirTree := parseInput(string(input))

	fmt.Printf(
		"Part 1 - Total size of directories smaller than 100000: %d\n",
		SumDirsSmallerThan(dirTree, MAXSIZEDIR),
	)

	minimalDirectorySize := NEEDEDSPACE - (CAPACITY - dirTree.totalSize)
	deletableDirectorySize := FindDirectoryToDelete(dirTree, minimalDirectorySize)
	fmt.Printf(
		"Part 2 - Size of the directory to delete: %d\n",
		deletableDirectorySize,
	)

}

func (d *Directory) ChangeDirectory(targetDir string) (cd *Directory) {
	if targetDir == ".." {
		d = d.parent
	} else {
		if _, ok := d.children[targetDir]; !ok {
			d.children[targetDir] = &Directory{
				name:     targetDir,
				parent:   d,
				children: map[string]*Directory{},
				files:    map[string]int{},
			}
		}

		d = d.children[targetDir]
	}
	return d
}

func (d *Directory) ListContents(rawList string) {
	if strings.HasPrefix(rawList, "dir") {
		subDir := strings.TrimPrefix(rawList, "dir ")
		if _, ok := d.children[subDir]; !ok {
			d.children[subDir] = &Directory{
				name:     subDir,
				parent:   d,
				children: map[string]*Directory{},
				files:    map[string]int{},
			}
		}
	} else {
		fileListing := strings.Split(rawList, " ")
		fileName := fileListing[1]
		fileSize, err := strconv.Atoi(fileListing[0])
		if err != nil {
			panic(err)
		}
		d.files[fileName] = fileSize
	}
}

func (d *Directory) SetDirectoryTotalSize() int {
	totalSize := 0
	for _, child := range d.children {
		totalSize += child.SetDirectoryTotalSize()
	}
	for _, fileSize := range d.files {
		totalSize += fileSize
	}
	d.totalSize = totalSize

	return d.totalSize
}

func SumDirsSmallerThan(dir *Directory, size int) int {
	sum := 0
	if dir.totalSize <= size {
		sum += dir.totalSize
	}

	for _, child := range dir.children {
		sum += SumDirsSmallerThan(child, size)
	}

	return sum
}

func FindDirectoryToDelete(dir *Directory, minSize int) int64 {
	var smallest int64 = math.MaxInt64
	directorySizes := []int64{
		smallest,
	}
	if dir.totalSize >= minSize {
		directorySizes = append(directorySizes, int64(dir.totalSize))
		smallest = min(smallest, int64(dir.totalSize))
	}

	for _, child := range dir.children {
		smallest = min(smallest, FindDirectoryToDelete(child, minSize))
	}

	return smallest
}

func parseInput(input string) *Directory {
	root := &Directory{
		name:      "/",
		children:  map[string]*Directory{},
		files:     map[string]int{},
		totalSize: 0,
	}
	dirTree := root

	commands := strings.Split(input, "\n")
	for _, command := range commands {

		if strings.HasPrefix(command, "$ cd") {
			dirTree = dirTree.ChangeDirectory(strings.TrimPrefix(command, "$ cd "))
		}

		if !strings.HasPrefix(command, "$") && command != "" {
			dirTree.ListContents(command)
		}
	}

	dirTree = root

	dirTree.SetDirectoryTotalSize()

	return dirTree
}
