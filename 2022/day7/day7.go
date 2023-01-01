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
		regEx := regexp.MustCompile(parser.pattern)
		matches := regEx.FindStringSubmatch(rawLine)

		if len(matches) > 0 {
			result := utils.GetRegexGroups(*regEx, matches)
			parser.apply(context, result)
		}
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

	for _, dir := range fileTree.directories {
		fmt.Println("dir", dir.name)
		if dir.parent != nil {
			fmt.Println("parent", dir.parent.name)
		}
	}

	for _, file := range fileTree.files {
		fmt.Println("file", file.name, "parent", file.parent.name, "size", file.size, "ext", file.ext)
	}
}
