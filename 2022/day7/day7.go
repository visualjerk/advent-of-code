package main

import (
	"regexp"
	"sort"
	"strings"

	"aoc.io/utils"
)

type File struct {
	name   string
	ext    string
	size   int
	parent *Directory
}

type Directory struct {
	name   string
	parent *Directory
	size   int
}

func (directory *Directory) isEqual(otherDir *Directory) bool {
	if directory.name != otherDir.name {
		return false
	}
	if directory.parent != otherDir.parent {
		return false
	}
	return true
}

type FileTree struct {
	files       []File
	directories []Directory
}

func (tree *FileTree) getExistingDirectory(directory *Directory) *Directory {
	for _, existingDir := range tree.directories {
		if existingDir.isEqual(directory) {
			return &existingDir
		}
	}
	return nil
}

func (tree *FileTree) getContainedFiles(directory Directory) []File {
	files := []File{}

	for _, file := range tree.files {
		if *file.parent == directory {
			files = append(files, file)
		}
	}

	for _, otherDir := range tree.directories {
		if otherDir.parent != nil && *otherDir.parent == directory {
			files = append(files, tree.getContainedFiles(otherDir)...)
		}
	}

	return files
}

func (tree *FileTree) calcSize(directory Directory) int {
	files := tree.getContainedFiles(directory)
	size := 0

	for _, file := range files {
		size += file.size
	}

	return size
}

func (tree *FileTree) calcUsedSpace() int {
	return tree.directories[0].size
}

func (tree *FileTree) precalcDirectorySizes() {
	for index, dir := range tree.directories {
		tree.directories[index].size = tree.calcSize(dir)
	}
}

func (tree *FileTree) getDirectoryWithMinSize(minSize int) Directory {
	sortedDirs := make([]Directory, len(tree.directories))
	copy(sortedDirs, tree.directories)

	sort.Slice(sortedDirs, func(i, j int) bool {
		return sortedDirs[i].size < sortedDirs[j].size
	})

	for index, dir := range sortedDirs {
		if minSize <= dir.size {
			return sortedDirs[index]
		}
	}
	return Directory{}
}

func (tree *FileTree) deleteDirectory(directory Directory) {
	println("deleted dir", directory.name, "with size", directory.size)
}

type Disk struct {
	fileTree       FileTree
	totalDiskSpace int
}

func (disk *Disk) freeUpSpace(amount int) {
	usedSpace := disk.fileTree.calcUsedSpace()
	freeSpace := disk.totalDiskSpace - usedSpace
	neededSpace := amount - freeSpace

	if neededSpace < 0 {
		return
	}

	directory := disk.fileTree.getDirectoryWithMinSize(neededSpace)
	disk.fileTree.deleteDirectory(directory)
}

type ParserContext struct {
	currentDirectory *Directory
	fileTree         FileTree
}

func (context *ParserContext) addDirectory(name string) *Directory {
	directory := Directory{
		name:   name,
		parent: context.currentDirectory,
	}

	existingDir := context.fileTree.getExistingDirectory(&directory)
	if existingDir != nil {
		return existingDir
	}

	context.fileTree.directories = append(context.fileTree.directories, directory)
	return &directory
}

func (context *ParserContext) addFile(name string, ext string, size int) {
	file := File{
		name:   name,
		parent: context.currentDirectory,
		ext:    ext,
		size:   size,
	}

	context.fileTree.files = append(context.fileTree.files, file)
}

type LineParserApplier func(context *ParserContext, params map[string]string)

type LineParser struct {
	pattern string
	apply   LineParserApplier
}

func (parser *LineParser) parse(context *ParserContext, rawLine string) {
	regEx := regexp.MustCompile(parser.pattern)
	matches := regEx.FindStringSubmatch(rawLine)

	if len(matches) > 0 {
		result := utils.GetRegexGroups(*regEx, matches)
		parser.apply(context, result)
	}
}

var cmdOpenDir = LineParser{
	`^\$ cd (?P<name>\S+)`,
	func(context *ParserContext, params map[string]string) {
		name := params["name"]

		if name == ".." {
			if context.currentDirectory == nil {
				return
			}
			context.currentDirectory = context.currentDirectory.parent
			return
		}

		directory := context.addDirectory(name)
		context.currentDirectory = directory
	},
}
var cmdList = LineParser{
	`^\$ ls`,
	func(context *ParserContext, params map[string]string) {
		// do nothing
	},
}
var outputDir = LineParser{
	`^dir (?P<name>\S+)`,
	func(context *ParserContext, params map[string]string) {
		context.addDirectory(params["name"])
	},
}
var outputFile = LineParser{
	`^(?P<size>\d+) (?P<name>\S+?)(.(?P<ext>\S+))?$`,
	func(context *ParserContext, params map[string]string) {
		context.addFile(params["name"], params["ext"], utils.SafeStringToInt(params["size"]))
	},
}

var parsers = []LineParser{
	cmdOpenDir,
	cmdList,
	outputDir,
	outputFile,
}

func parseLine(context *ParserContext, rawLine string) {
	for _, parser := range parsers {
		parser.parse(context, rawLine)
	}
}

func parseCommandLineOutput(output string) FileTree {
	parserContext := ParserContext{
		currentDirectory: nil,
		fileTree:         FileTree{},
	}
	rawLines := strings.Split(output, "\n")

	for _, rawLine := range rawLines {
		parseLine(&parserContext, rawLine)
	}

	return parserContext.fileTree
}

func main() {
	data := utils.LoadData()

	fileTree := parseCommandLineOutput(data)
	fileTree.precalcDirectorySizes()

	disk := Disk{fileTree: fileTree, totalDiskSpace: 70000000}
	disk.freeUpSpace(30000000)
}
