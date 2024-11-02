package folder

import (
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	// First find the parent folder to verify it exists and get its path
	var parentFolder *Folder
	for _, folder := range f.folders {
		if folder.OrgId == orgID && folder.Name == name {
			parentFolder = &folder
			break
		}
	}

	// Handle error cases
	if parentFolder == nil {
		return nil
	}

	childFolders := []Folder{}
	parentPath := parentFolder.Paths

	// Find all child folders by checking if their paths start with parent path
	for _, folder := range f.folders {
		// Skip if not in same org
		if folder.OrgId != orgID {
			continue
		}

		// A child folder will:
		// 1. Have a path that starts with parent path + "."
		// 2. Not be the parent folder itself
		if strings.HasPrefix(folder.Paths, parentPath+".") {
			childFolders = append(childFolders, folder)
		}
	}

	return childFolders
}
