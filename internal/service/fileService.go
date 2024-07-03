package service

import (
	"crypto/sha256"
	"fmt"
	"im-GIN/config"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type FileService struct{}

var (
	fileService *FileService
	once        sync.Once
)

func GetFileService() *FileService {
	once.Do(func() {
		fileService = &FileService{}
	})
	return fileService
}

func (*FileService) FileUpload(file *multipart.FileHeader) (string, error) {
	// 创建临时文件
	tempFile, err := file.Open()
	if err != nil {
		return "打开文件失败", err
	}
	defer tempFile.Close()
	// 计算SHA-256哈希
	hasher := sha256.New()
	if _, err := io.Copy(hasher, tempFile); err != nil {
		return "计算哈希失败", err
	}
	fileHash := hasher.Sum(nil)
	fileHashStr := fmt.Sprintf("%x", fileHash)
	prefix := fileHashStr[:8]
	tsStr := fmt.Sprintf("%d", time.Now().Unix())
	fileName := file.Filename
	fileType := filepath.Ext(fileName)
	newFileName := prefix + tsStr + fileType
	// 查表存不存在相同的哈希
	// 存在则更新引用次数
	// 不存在则新增记录并保存到本地

	// 重新打开文件句柄，确保在复制到目标路径时使用新鲜的句柄
	if _, err := tempFile.Seek(0, io.SeekStart); err != nil {
		return "文件上传出错", err
	}
	// 构建完整的保存路径
	targetPath := filepath.Join(config.Cfg.File.Dir, newFileName)
	// 创建目标目录如果不存在
	if err := os.MkdirAll(config.Cfg.File.Dir, 0755); err != nil {
		return "文件上传失败", err
	}
	// 创建文件
	dst, err := os.Create(targetPath)
	if err != nil {
		return "上传文件失败", err
	}
	defer dst.Close()
	// 将上传的文件内容写入新创建的文件
	if _, err := io.Copy(dst, tempFile); err != nil {
		return "上传文件时发生错误", err
	}
	return newFileName, nil
}

func (*FileService) FileDownload(filePath string) ([]byte, string, error) {
	// 打开文件
	file, err := os.Open(config.Cfg.File.Dir + filePath)
	if err != nil {
		return nil, "找不到文件了", err
	}
	defer file.Close()
	// 获取文件名
	filename := filepath.Base(filePath)
	// 读取文件内容
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return nil, "读取文件失败", err
	}
	return fileContent, filename, nil
}
