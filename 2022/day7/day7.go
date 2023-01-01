package main

import (
	"aoc.io/utils"
	"fmt"
	"regexp"
	"strings"
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
}

type FileTree struct {
	files       []File
	directories []Directory
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

type ParserContext struct {
	currentDirectory *Directory
	fileTree         FileTree
}

func (context *ParserContext) addDirectory(name string) {
	directory := Directory{
		name:   name,
		parent: context.currentDirectory,
	}

	context.currentDirectory = &directory
	context.fileTree.directories = append(context.fileTree.directories, directory)
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
			context.currentDirectory = context.currentDirectory.parent
			return
		}

		context.addDirectory(name)
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

const TRESHOLD = 100000

func main() {
	data := utils.LoadData()
	fileTree := parseCommandLineOutput(data)

	sum := 0

	for _, dir := range fileTree.directories {
		size := fileTree.calcSize(dir)
		if size <= TRESHOLD {
			sum += size
		}
	}

	fmt.Println(sum)
}
