package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// eventMap is a helper type for storing event-function mappings
type eventList [][2]string

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(dir)
	files, err := getProtoFiles("89t/websocket/proto/bigProtoV4") // TODO: 放 .proto文件的目录。会读取该目录下所有.proto文件
	if err != nil {
		fmt.Println("Error getting proto files:", err)
		return
	}
	var sendMap, pushMap eventList

	// The updated regex patterns handle multi-line comments before messages
	// regexSend := regexp.MustCompile(`(?m)//@send\s+([^\n]+)\n(?:.*\n)*?message\s+(\w+)`)
	// regexPush := regexp.MustCompile(`(?m)//@push\s+([^\n]+)\n(?:.*\n)*?message\s+(\w+)`)
	regex := regexp.MustCompile(`(//\s*@(send|push)\s+(.*?)\n)+message\s+(\w+)`)

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}

		sendMap, pushMap = parseEvents(regex, string(content), sendMap, pushMap)
	}
}

func parseEvents(regex *regexp.Regexp, content string, send, push eventList) (eventList, eventList) {
	matches := regex.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		rows := strings.Split(match[0], "\n")

		msgTypeTmp := strings.Split(rows[len(rows)-1], " ")
		messageType := msgTypeTmp[1]

		eventNameTmp := rows[:len(rows)-1]
		for _, tmpStr := range eventNameTmp {
			subRegex := regexp.MustCompile(`//\s*@(send|push)\s+(.+)`)
			subMatches := subRegex.FindAllStringSubmatch(tmpStr, -1)
			action := subMatches[0][1]
			events := strings.Fields(subMatches[0][2])
			for _, event := range events {
				switch action {
				case "send":
					send = append(send, [2]string{event, messageType})
				case "push":
					push = append(push, [2]string{event, messageType})
				}
			}
		}
	}

	return send, push
}

func getProtoFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".proto" {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
