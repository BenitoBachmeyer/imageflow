package fileupload

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type fileUploader struct {
}

func (f *fileUploader) fileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("error occurred while extracting file from form request: %w", err)})
		return
	}

	dst := filepath.Join("./files/", filepath.Base(file.Filename))
	if err := c.SaveUploadedFile(file, dst); err != nil {
		log.Printf(fmt.Errorf("error occurred while saving uploaded file to destination: %w, %s", err, dst).Error())
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded successfully", file.Filename))
}
