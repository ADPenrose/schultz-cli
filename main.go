package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// We start by creating a new virtual environment.
	cmd := exec.Command("python", "-m", "venv", ".venv")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// We activate the virtual environment and install Django.
	cmd = exec.Command("bash", "-c", "source .venv/bin/activate && pip install Django")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// We create a new Django project.
	cmd = exec.Command("bash", "-c", "source .venv/bin/activate && django-admin startproject mysite .")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Django project created successfully.")

	// We create a new Django app.
	cmd = exec.Command("bash", "-c", "source .venv/bin/activate && python manage.py startapp polls")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Django app created successfully.")

	// Now, we need to connect the app to the project.
	filePath := "mysite/settings.py"
	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Convert content to string
	fileContent := string(fileContents)

	// Locate where to append the object (e.g., INSTALLED_APPS)
	// Assuming INSTALLED_APPS is a list defined in settings.py like:
	// INSTALLED_APPS = [
	//     ...
	// ]
	// You would search for the line containing INSTALLED_APPS and modify it accordingly.
	searchTerm := "INSTALLED_APPS = ["
	newLine := "\n    'polls.apps.PollsConfig',"

	// Find the position to insert the new line
	insertIndex := strings.Index(fileContent, searchTerm)
	if insertIndex == -1 {
		log.Fatalf("Could not find %s in %s", searchTerm, filePath)
	}

	// Insert the new line just after the found index
	insertIndex += len(searchTerm)
	modifiedContent := fileContent[:insertIndex] + newLine + fileContent[insertIndex:]

	// Write the modified content back to the file
	err = os.WriteFile(filePath, []byte(modifiedContent), 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

	log.Println("File modified successfully.")
}
