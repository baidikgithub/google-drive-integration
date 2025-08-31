package drive

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

// DriveService struct to hold the Drive API service
type DriveService struct {
	Service *drive.Service
}

// NewDriveService initializes a new Google Drive API service
func NewDriveService(credFile string) (*DriveService, error) {
	b, err := os.ReadFile(credFile)
	if err != nil {
		return nil, fmt.Errorf("error reading credentials file: %v", err)
	}

	config, err := google.JWTConfigFromJSON(b, drive.DriveFileScope)
	if err != nil {
		return nil, fmt.Errorf("error parsing credentials file: %v", err)
	}

	client := config.Client(context.Background())
	srv, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("error creating Drive service: %v", err)
	}

	return &DriveService{Service: srv}, nil
}
