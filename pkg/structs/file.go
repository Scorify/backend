package structs

import (
	"path/filepath"

	"github.com/google/uuid"
)

type File struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (file *File) Path(fileType string, parentID uuid.UUID) string {
	return filepath.Join("/api/files/", fileType, parentID.String(), file.ID.String(), file.Name)
}
