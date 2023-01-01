package main

import (
	"aoc.io/utils"
	"fmt"
	"regexp"
	"strings"
)

type File struct {
	name string
	ext  string
	size int
}

type Directory struct {
	name   string
	parent *Directory
	files  []File
}

type FileTree struct {
	directories []Directory
}

type ParserContext struct {
	currentDirectory *Directory
	fileTree         FileTree
}

type LineParserApplier func(context ParserContext, params map[string]string)

type LineParser struct {
	pattern string
	apply   LineParserApplier
}

var CMD_OPEN_DIR = LineParser{
	`^\$ cd (?P<name>\S+)`,
	func(context ParserContext, params map[string]string) {},
}
var CMD_LIST = LineParser{
	`^\$ ls`,
	func(context ParserContext, params map[string]string) {},
}
var OUTPUT_DIR = LineParser{
	`^dir (?P<name>\S+)`,
	func(context ParserContext, params map[string]string) {},
}
var OUTPUT_FILE = LineParser{
	`^(?P<size>\d+) (?P<name>\S+?)(.(?P<ext>\S+))?$`,
	func(context ParserContext, params map[string]string) {},
}

var parsers = map[string]LineParser{
	"CMD_OPEN_DIR": CMD_OPEN_DIR,
	"CMD_LIST":     CMD_LIST,
	"OUTPUT_DIR":   OUTPUT_DIR,
	"OUTPUT_FILE":  OUTPUT_FILE,
}

func parseLine(context *ParserContext, rawLine string) {
	for _, parser := range parsers {
		regEx := regexp.MustCompile(parser.pattern)
		matches := regEx.FindStringSubmatch(rawLine)

		if len(matches) > 0 {
			result := utils.GetRegexGroups(*regEx, matches)
			parser.apply(*context, result)
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

	fmt.Println(fileTree)
}
