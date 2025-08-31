package drive

import (
	"fmt"
	"io"
	"strings"

	"google.golang.org/api/drive/v3"
)

// CreateFile uploads a new file to Google Drive
func (d *DriveService) CreateFile(name, mimeType, content string) (string, error) {
	fileMetadata := &drive.File{Name: name, MimeType: mimeType}
	file, err := d.Service.Files.Create(fileMetadata).Media(strings.NewReader(content)).Do()
	if err != nil {
		return "", fmt.Errorf("error creating file: %v", err)
	}

	fmt.Println("✅ File created successfully! File ID:", file.Id)
	return file.Id, nil
}

// GetFileContent retrieves file content from Google Drive
func (d *DriveService) GetFileContent(fileID string) (string, error) {
	resp, err := d.Service.Files.Get(fileID).Download()
	if err != nil {
		return "", fmt.Errorf("error getting file: %v", err)
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading file content: %v", err)
	}

	return string(content), nil
}

// UpdateFile modifies the content of an existing file
func (d *DriveService) UpdateFile(fileID, newContent string) error {
	_, err := d.Service.Files.Update(fileID, nil).Media(strings.NewReader(newContent)).Do()
	if err != nil {
		return fmt.Errorf("error updating file: %v", err)
	}

	fmt.Println("✅ File updated successfully!")
	return nil
}
