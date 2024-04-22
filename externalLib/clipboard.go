package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"os/exec"
	"strings"
)

func main1() {
	text, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("Failed to read from clipboard:", err)
		return
	}
	fmt.Println("Clipboard contains:", text)
}
func main() {
	// AppleScript 命令来获取粘贴板中的文件路径
	appleScript := `tell application "Finder"
                        try
                            set the clipboardItems to the clipboard as text
                            return clipboardItems
                        on error
                            return ""
                        end try
                    end tell`
	// 运行 osascript 执行 AppleScript
	cmd := exec.Command("osascript", "-e", appleScript)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Failed to run AppleScript:", err)
		return
	}

	// 处理输出结果
	content := string(output)
	if content == "" {
		fmt.Println("No files in clipboard.")
	} else {
		// 如果粘贴板中包含文件路径，它们会被输出
		files := strings.Split(content, ", ")
		fmt.Println("Files in clipboard:")
		for _, file := range files {
			fmt.Println(file)
		}
	}
}
