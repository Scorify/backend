package structs

import (
	"io"
	"os"
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

func (file *File) FilePath(fileType FileType, parentID uuid.UUID) string {
	return filepath.Join("./files/", string(fileType), parentID.String(), file.ID.String(), file.Name)
}

func (file *File) APIPath(fileType FileType, parentID uuid.UUID) string {
	return filepath.Join("/api/files/", string(fileType), parentID.String(), file.ID.String(), file.Name)

}

func (file *File) WriteFile(fileType FileType, parentID uuid.UUID, reader io.ReadSeeker) error {
	filePath := file.FilePath(fileType, parentID)

	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return err
	}

	fileHandle, err := os.Create(filePath)
	if err != nil {
		return err
	}

	_, err = io.Copy(fileHandle, reader)
	return err
}

func (file *File) DeleteFile(fileType FileType, parentID uuid.UUID) error {
	return os.Remove(file.FilePath(fileType, parentID))
}
