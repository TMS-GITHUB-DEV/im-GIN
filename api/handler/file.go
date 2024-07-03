package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"im-GIN/internal/errors"
	"im-GIN/internal/service"
	"net/http"
	"path/filepath"
	"strings"
)

var MimeTypes = map[string]string{
	".html":    "text/html",
	".css":     "text/css",
	".js":      "application/javascript",
	".json":    "application/json",
	".xml":     "application/xml",
	".pdf":     "application/pdf",
	".txt":     "text/plain",
	".jpg":     "image/jpeg",
	".jpeg":    "image/jpeg",
	".png":     "image/png",
	".gif":     "image/gif",
	".svg":     "image/svg+xml",
	".ico":     "image/vnd.microsoft.icon",
	".mp3":     "audio/mpeg",
	".wav":     "audio/wav",
	".mp4":     "video/mp4",
	".avi":     "video/x-msvideo",
	".zip":     "application/zip",
	".tar":     "application/x-tar",
	".gz":      "application/gzip",
	".tar.gz":  "application/x-gzip",
	".tar.bz2": "application/x-bzip2",
	".doc":     "application/msword",
	".docx":    "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	".xls":     "application/vnd.ms-excel",
	".xlsx":    "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	".ppt":     "application/vnd.ms-powerpoint",
	".pptx":    "application/vnd.openxmlformats-officedocument.presentationml.presentation",
	// 添加更多...
}

func FileUpload(c *gin.Context) {
	// 获取文件
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "%+v", err)
		c.Abort()
		return
	}

	fileService := service.GetFileService()

	msg, err := fileService.FileUpload(file)
	fmt.Println("%v", err)
	if err != nil {
		c.Error(errors.NewServerError(msg, err))
		c.Abort()
		return
	}

	//c.Set(resp.RES, resp.Success("上传成功"))
	fmt.Println(msg)
	c.JSON(200, gin.H{
		"errno": 0, // 注意：值是数字，不能是字符串
		"data": gin.H{
			"url": "http://192.168.97.80:10101/file/download/" + msg, // 图片 src ，必须 fixme 替换前缀
		},
	})
}

func FileDownload(c *gin.Context) {
	url := c.Param("url")
	fmt.Println(url)
	fileService := service.GetFileService()
	file, fileName, err := fileService.FileDownload(url)
	if err != nil {
		c.Error(errors.SimpleError(""))
		return
	}
	//c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
	//c.File(url)

	fileType := strings.ToLower(filepath.Ext(fileName))
	c.Data(http.StatusOK, MimeTypes[fileType], file)
}
