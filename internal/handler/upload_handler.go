package handler

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/satudata/backend/internal/handler/dto"
	"github.com/satudata/backend/internal/service"
)

// UploadHandler handles file upload HTTP requests.
type UploadHandler struct {
	storageService service.StorageServiceInterface
}

// NewUploadHandler membuat instance baru UploadHandler.
func NewUploadHandler(storageService service.StorageServiceInterface) *UploadHandler {
	return &UploadHandler{
		storageService: storageService,
	}
}

// UploadDatasetFile godoc
// @Summary Upload file dataset
// @Description Upload file dataset (CSV, JSON, XLSX, Parquet)
// @Tags Protected
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File dataset"
// @Success 200 {object} dto.APIResponse{data=dto.UploadResponse}
// @Failure 400 {object} dto.APIResponse
// @Security BearerAuth
// @Router /protected/upload/dataset [post]
func (h *UploadHandler) UploadDatasetFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		dto.BadRequestResponse(c, "File is required")
		return
	}
	defer file.Close()

	// Validasi tipe file
	ext := strings.ToLower(filepath.Ext(header.Filename))
	allowedExts := map[string]string{
		".csv":     "text/csv",
		".json":    "application/json",
		".xlsx":    "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		".xls":     "application/vnd.ms-excel",
		".parquet": "application/octet-stream",
		".pdf":     "application/pdf",
	}
	contentType, ok := allowedExts[ext]
	if !ok {
		dto.BadRequestResponse(c, fmt.Sprintf("File type %s is not allowed. Allowed: CSV, JSON, XLSX, Parquet, PDF", ext))
		return
	}

	// Baca file
	data, err := io.ReadAll(file)
	if err != nil {
		dto.InternalErrorResponse(c, "Failed to read file")
		return
	}

	// Validasi ukuran (max 50MB)
	if header.Size > 50*1024*1024 {
		dto.BadRequestResponse(c, "File too large. Maximum size is 50MB")
		return
	}

	url, err := h.storageService.UploadDataset(c.Request.Context(), service.UploadFileRequest{
		FileName:    header.Filename,
		FileSize:    header.Size,
		ContentType: contentType,
		Data:        data,
	})
	if err != nil {
		dto.InternalErrorResponse(c, err.Error())
		return
	}

	dto.SuccessResponse(c, dto.UploadResponse{
		URL:        url,
		FileName:   header.Filename,
		FileSize:   header.Size,
		FileFormat: strings.TrimPrefix(ext, "."),
	})
}

// UploadArticleImageFile godoc
// @Summary Upload gambar artikel
// @Description Upload gambar untuk artikel/infografis
// @Tags Protected
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Gambar artikel"
// @Success 200 {object} dto.APIResponse{data=dto.UploadResponse}
// @Failure 400 {object} dto.APIResponse
// @Security BearerAuth
// @Router /protected/upload/article-image [post]
func (h *UploadHandler) UploadArticleImageFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		dto.BadRequestResponse(c, "File is required")
		return
	}
	defer file.Close()

	// Validasi tipe gambar
	ext := strings.ToLower(filepath.Ext(header.Filename))
	allowedExts := map[string]string{
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".png":  "image/png",
		".webp": "image/webp",
		".svg":  "image/svg+xml",
	}
	contentType, ok := allowedExts[ext]
	if !ok {
		dto.BadRequestResponse(c, fmt.Sprintf("File type %s is not allowed. Allowed: JPG, PNG, WebP, SVG", ext))
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		dto.InternalErrorResponse(c, "Failed to read file")
		return
	}

	// Validasi ukuran (max 10MB)
	if header.Size > 10*1024*1024 {
		dto.BadRequestResponse(c, "File too large. Maximum size is 10MB")
		return
	}

	url, err := h.storageService.UploadArticleImage(c.Request.Context(), service.UploadFileRequest{
		FileName:    header.Filename,
		FileSize:    header.Size,
		ContentType: contentType,
		Data:        data,
	})
	if err != nil {
		dto.InternalErrorResponse(c, err.Error())
		return
	}

	dto.SuccessResponse(c, dto.UploadResponse{
		URL:        url,
		FileName:   header.Filename,
		FileSize:   header.Size,
		FileFormat: strings.TrimPrefix(ext, "."),
	})
}

// UploadStandardDocFile godoc
// @Summary Upload dokumen standar
// @Description Upload dokumen standar data (PDF)
// @Tags Protected
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Dokumen PDF"
// @Success 200 {object} dto.APIResponse{data=dto.UploadResponse}
// @Failure 400 {object} dto.APIResponse
// @Security BearerAuth
// @Router /protected/upload/standard-doc [post]
func (h *UploadHandler) UploadStandardDocFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		dto.BadRequestResponse(c, "File is required")
		return
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(header.Filename))
	if ext != ".pdf" {
		dto.BadRequestResponse(c, "Only PDF files are allowed")
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		dto.InternalErrorResponse(c, "Failed to read file")
		return
	}

	// Validasi ukuran (max 50MB)
	if header.Size > 50*1024*1024 {
		dto.BadRequestResponse(c, "File too large. Maximum size is 50MB")
		return
	}

	url, err := h.storageService.UploadStandardDoc(c.Request.Context(), service.UploadFileRequest{
		FileName:    header.Filename,
		FileSize:    header.Size,
		ContentType: "application/pdf",
		Data:        data,
	})
	if err != nil {
		dto.InternalErrorResponse(c, err.Error())
		return
	}

	dto.SuccessResponse(c, dto.UploadResponse{
		URL:        url,
		FileName:   header.Filename,
		FileSize:   header.Size,
		FileFormat: "pdf",
	})
}
