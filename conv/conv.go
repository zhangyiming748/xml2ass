package conv

import (
	"fmt"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/xml2ass/constant"
	"log/slog"
	"os"
	"os/exec"
	"strings"
)

func GetXmls() {
	var allFiles []GetFileInfo.BasicInfo
	if isExist(constant.BILI) {
		files := GetFileInfo.GetAllFilesInfo(constant.BILI, "xml")
		allFiles = append(allFiles, files...)
	}
	if isExist(constant.HD) {
		files := GetFileInfo.GetAllFilesInfo(constant.HD, "xml")
		allFiles = append(allFiles, files...)
	}
	if isExist(constant.GLOBAL) {
		files := GetFileInfo.GetAllFilesInfo(constant.GLOBAL, "xml")
		allFiles = append(allFiles, files...)
	}

	for _, xml := range allFiles {
		slog.Info("获取到的文件信息", slog.Any("xml", xml))
		if fp, err := Conv(xml); err != nil {
			slog.Warn("当前字幕转换出错")
		} else {
			slog.Info("当前字幕转换成功", slog.String("输出路径", fp))
		}
	}
}
func Conv(xml GetFileInfo.BasicInfo) (string, error) {
	// danmaku2ass danmaku.xml -s 1280x720 -dm 15 -fs 45 -a 50 -o danmaku.ass
	ass := strings.Replace(xml.FullPath, ".xml", ".ass", 1)
	bat := strings.Join([]string{"danmaku2ass", xml.FullPath, "-s", "640x480", "-dm", "15", "-fs", "35", "-r", "-o", ass}, " ")
	cmd := exec.Command("bash", "-c", bat)
	//cmd := exec.Command("danmaku2ass", xml.FullPath, "-s", "1280x720", "-dm", "15", "-fs", "45", "-a", "50", "-r", "-o", ass)
	output, err := cmd.CombinedOutput()
	slog.Debug("生成命令", slog.String("命令原文", fmt.Sprint(cmd)))
	if err != nil {
		slog.Warn("当前弹幕文件转换错误", slog.Any("文件信息", xml), slog.Any("错误原文", err))
		return "", err
	} else {
		slog.Info("当前弹幕文件转换成功", slog.Any("命令输出", string(output)))
		return ass, nil
	}
}
func isExist(dir string) bool {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		fmt.Println("文件夹不存在")
		return false
	} else {
		fmt.Println("文件夹存在")
		return true
	}
}
