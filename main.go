package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	var targetPath string
	// 获取命令行参数
	args := os.Args
	// 判断是否有命令行参数传入
	if len(args) > 1 {
		// 遍历命令行参数，并打印出来
		for _, arg := range args[1:] {
			targetPath = arg
		}
	} else {
		fmt.Println("没有传入任何参数")
		return
	}
	//targetPath = "input/Apollo_11_moonwalk_montage_720p.mp4"
	fileName := filepath.Base(targetPath)
	fileExt := filepath.Ext(fileName)
	audioPath := "output/" + strings.Split(fileName, ".")[0] + ".wav"
	fmt.Println("文件名称:", fileName)
	fmt.Println("文件后缀:", fileExt)
	fmt.Println("audioPath:", audioPath)
	cmd := exec.Command("ffmpeg", []string{
		"-y",
		"-i", targetPath,
		"-vn",
		"-acodec", "pcm_s16le",
		"-ar", "16000",
		"-ac", "2",
		audioPath,
	}...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	cmd = exec.Command("./whisper/main", []string{
		"-f", audioPath,
		"-osrt",
		"-m",
		"./whisper/models/ggml-base.bin",
	}...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
