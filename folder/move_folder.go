package folder

import (
	"fmt"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Find source and destination folders
	var sourceFolder *Folder
	var destFolder *Folder

	// First pass to find destination folder
	for _, folder := range f.folders {
		if folder.Name == dst {
			destFolder = &folder
			break
		}
	}

	// Find the specific source folder (first one that matches the name)
	for _, folder := range f.folders {
		if folder.Name == name {
			sourceFolder = &folder
			break
		}
	}

	// Error checking
	if sourceFolder == nil {
		return nil, fmt.Errorf("source folder does not exist")
	}
	if destFolder == nil {
		return nil, fmt.Errorf("destination folder does not exist")
	}
	if sourceFolder.Name == destFolder.Name {
		return nil, fmt.Errorf("cannot move a folder to itself")
	}
	if sourceFolder.OrgId != destFolder.OrgId {
		return nil, fmt.Errorf("cannot move a folder to a different organization")
	}

	// Check if trying to move to a child of itself
	if strings.HasPrefix(destFolder.Paths, sourceFolder.Paths+".") {
		return nil, fmt.Errorf("cannot move a folder to a child of itself")
	}

	// Create new folders array for the result
	newFolders := make([]Folder, len(f.folders))
	copy(newFolders, f.folders)

	// Get the old path prefix and new path prefix
	oldPrefix := sourceFolder.Paths
	newPrefix := destFolder.Paths + "." + sourceFolder.Name

	// Update paths for the source folder and all its children
	for i, folder := range newFolders {
		if folder.Name == sourceFolder.Name && folder.Paths == oldPrefix {
			// Update source folder
			newFolders[i].Paths = newPrefix
		} else if strings.HasPrefix(folder.Paths, oldPrefix+".") {
			// Update child folders
			newFolders[i].Paths = strings.Replace(folder.Paths, oldPrefix, newPrefix, 1)
		}
	}

	return newFolders, nil
}
