package structs

import (
	"path/filepath"

	"github.com/google/uuid"
)

type FileType string

const (
	FileTypeInject     FileType = "inject"
	FileTypeSubmission FileType = "submission"
)

type File struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (file *File) Path(fileType FileType, parentID uuid.UUID) string {
	return filepath.Join("/api/files/", string(fileType), parentID.String(), file.ID.String(), file.Name)
}
