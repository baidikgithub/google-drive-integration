package main

import (
	"fmt"
	"log"

	"MyGolangApp/drive"
	"MyGolangApp/input"
	"MyGolangApp/utils"
)

func main() {
	credFile := "credentials.json" // Change if needed
	driveService, err := drive.NewDriveService(credFile)
	if err != nil {
		log.Fatalf("❌ Failed to initialize Drive service: %v", err)
	}

	// Step 1: Get user input for file creation
	fileName := input.GetUserInput("📄 Enter CSV file name (e.g., data.csv): ")
	fileContent := input.GetUserInput("📝 Enter CSV content (e.g., Name,Age,City\\nJohn,25,NY): ")

	// Step 2: Create a file in Google Drive
	fileID, err := driveService.CreateFile(fileName, "text/csv", fileContent)
	if err != nil {
		log.Fatalf("❌ Error creating file: %v", err)
	}

	// Step 3: Fetch and display file contents
	content, err := driveService.GetFileContent(fileID)
	if err != nil {
		log.Fatalf("❌ Error retrieving file contents: %v", err)
	}
	fmt.Println("📂 Current File Contents:\n", content)

	// Step 4: Ask user if they want to update the file
	updateChoice := input.GetUserInput("🔄 Do you want to update the file? (yes/no): ")
	if utils.ConfirmAction(updateChoice) {
		newContent := input.GetUserInput("✏️ Enter new CSV content: ")
		err := driveService.UpdateFile(fileID, newContent)
		if err != nil {
			log.Fatalf("❌ Error updating file: %v", err)
		}
	}
}
