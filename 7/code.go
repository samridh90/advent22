package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samridh90/advent22/shared"
)

type Dir struct {
	dirs  map[string]Dir
	files map[string]int
	size  int
}

func newDir() Dir {
	return Dir{dirs: map[string]Dir{}, files: map[string]int{}}
}

func isCmd(fields []string) bool {
	return len(fields) > 1 && fields[0] == "$"
}

func isLs(fields []string) bool {
	return isCmd(fields) && fields[1] == "ls"
}

func isChangeDir(fields []string) bool {
	return isCmd(fields) && fields[1] == "cd"
}

func isDir(fields []string) bool {
	return fields[0] == "dir"
}

func getDirName(fields []string) string {
	return fields[1]
}

func getFileSizeAndName(fields []string) (int, string) {
	fileSize, _ := strconv.ParseInt(fields[0], 0, 64)
	fileName := fields[1]
	return int(fileSize), fileName
}

func dirToChangeTo(fields []string) string {
	return fields[2]
}

func makeTree(input []string) Dir {
	root := newDir()
	root.dirs["/"] = newDir()
	root.dirs["/"].dirs[".."] = root

	currentDir := root
	for i := 0; i < len(input); {
		fields := strings.Fields(input[i])
		if isChangeDir(fields) {
			dirName := dirToChangeTo(fields)
			currentDir = currentDir.dirs[dirName]
			i++
		} else if isLs(fields) {
			i++
		} else {
			for !isCmd(fields) && i < len(input) {
				if isDir(fields) {
					dirName := getDirName(fields)
					currentDir.dirs[dirName] = newDir()
					currentDir.dirs[dirName].dirs[".."] = currentDir
				} else {
					fileSize, fileName := getFileSizeAndName(fields)
					currentDir.files[fileName] = fileSize
				}
				i++
				if i >= len(input) {
					break
				}
				fields = strings.Fields(input[i])
			}
		}
	}

	return root
}

func computeSize(root Dir) Dir {
	totalFileSize := 0
	for _, size := range root.files {
		totalFileSize += size
	}
	totalDirSize := 0
	for dirName, dir := range root.dirs {
		if dirName != ".." {
			sizedDir := computeSize(dir)
			root.dirs[dirName] = sizedDir
			totalDirSize += sizedDir.size
		}
	}
	root.size = totalFileSize + totalDirSize
	return root
}

func findDirs(root Dir, f func(int) bool) []int {
	resultSet := []int{}
	queue := []Dir{root}

	for len(queue) != 0 {
		dir := queue[0]
		queue = queue[1:]
		if f(dir.size) {
			resultSet = append(resultSet, dir.size)
		}
		for name, childDir := range dir.dirs {
			if name != ".." {
				queue = append(queue, childDir)
			}
		}
	}

	return resultSet
}

func testBench(input []string) {
	root := computeSize(makeTree(input))
	resultSet := findDirs(root, func(dirSize int) bool { return dirSize <= 100000 })
	totalSize := 0
	for _, size := range resultSet {
		totalSize += size
	}
	fmt.Println(totalSize)
}

func testBench2(input []string) {
	root := computeSize(makeTree(input))
	freeSize := 70000000 - root.size
	requiredSize := 30000000 - freeSize
	resultSet := findDirs(root, func(dirSize int) bool { return dirSize >= requiredSize })
	min := resultSet[0]
	for _, v := range resultSet {
		if v < min {
			min = v
		}
	}
	fmt.Println(min)
}

func main() {
	test, _ := shared.ReadFile("./7/test.txt")
	input, _ := shared.ReadFile("./7/input.txt")
	fmt.Println("Part1")
	testBench(test)
	testBench(input)
	fmt.Println("Part2")
	testBench2(test)
	testBench2(input)
}
