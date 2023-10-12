package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// マルチパートフォームが利用できるメモリの制限を設定する(デフォルトは 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload/single", func(c *gin.Context) {
		// 単一のファイル
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// 特定のディレクトリにファイルをアップロードする
		filename := filepath.Base(file.Filename)
		c.SaveUploadedFile(file, filename)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.POST("/upload/multi", func(c *gin.Context) {
		// マルチパートフォーム
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			// 特定のディレクトリにファイルをアップロードする
			filename := filepath.Base(file.Filename)
			c.SaveUploadedFile(file, filename)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	router.Run(":8080")
}
