package folder_test

import (
	"sort"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()

	// Create test UUIDs
	org1ID := uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
	org2ID := uuid.Must(uuid.NewV4())
	org3ID := uuid.Must(uuid.NewV4())

	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name:    "empty folder list",
			orgID:   org1ID,
			folders: []folder.Folder{},
			want:    []folder.Folder{},
		},
		{
			name:  "single folder matching orgID",
			orgID: org1ID,
			folders: []folder.Folder{
				{
					Name:  "root",
					OrgId: org1ID,
					Paths: "root",
				},
			},
			want: []folder.Folder{
				{
					Name:  "root",
					OrgId: org1ID,
					Paths: "root",
				},
			},
		},
		{
			name:  "single folder non-matching orgID",
			orgID: org1ID,
			folders: []folder.Folder{
				{
					Name:  "root",
					OrgId: org2ID,
					Paths: "root",
				},
			},
			want: []folder.Folder{},
		},
		{
			name:  "multiple folders with mixed orgIDs",
			orgID: org1ID,
			folders: []folder.Folder{
				{
					Name:  "folder1",
					OrgId: org1ID,
					Paths: "folder1",
				},
				{
					Name:  "folder2",
					OrgId: org2ID,
					Paths: "folder2",
				},
				{
					Name:  "folder3",
					OrgId: org1ID,
					Paths: "folder3",
				},
			},
			want: []folder.Folder{
				{
					Name:  "folder1",
					OrgId: org1ID,
					Paths: "folder1",
				},
				{
					Name:  "folder3",
					OrgId: org1ID,
					Paths: "folder3",
				},
			},
		},
		{
			name:  "multiple folders none matching orgID",
			orgID: org3ID,
			folders: []folder.Folder{
				{
					Name:  "folder1",
					OrgId: org1ID,
					Paths: "folder1",
				},
				{
					Name:  "folder2",
					OrgId: org2ID,
					Paths: "folder2",
				},
			},
			want: []folder.Folder{},
		},
		{
			name:  "complex nested folders with mixed orgIDs",
			orgID: org1ID,
			folders: []folder.Folder{
				{
					Name:  "root1",
					OrgId: org1ID,
					Paths: "root1",
				},
				{
					Name:  "child1",
					OrgId: org1ID,
					Paths: "root1.child1",
				},
				{
					Name:  "root2",
					OrgId: org2ID,
					Paths: "root2",
				},
				{
					Name:  "child2",
					OrgId: org2ID,
					Paths: "root2.child2",
				},
				{
					Name:  "child3",
					OrgId: org1ID,
					Paths: "root1.child3",
				},
			},
			want: []folder.Folder{
				{
					Name:  "root1",
					OrgId: org1ID,
					Paths: "root1",
				},
				{
					Name:  "child1",
					OrgId: org1ID,
					Paths: "root1.child1",
				},
				{
					Name:  "child3",
					OrgId: org1ID,
					Paths: "root1.child3",
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable for parallel testing
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := folder.NewDriver(tt.folders)
			got := f.GetFoldersByOrgID(tt.orgID)

			// Sort both slices to ensure consistent comparison
			sortFolders(got)
			sortFolders(tt.want)

			assert.Equal(t, tt.want, got,
				"GetFoldersByOrgID(%v) = %v, want %v",
				tt.orgID, got, tt.want)
		})
	}
}

// Helper function to sort folders by path for consistent comparison
func sortFolders(folders []folder.Folder) {
	sort.Slice(folders, func(i, j int) bool {
		return folders[i].Paths < folders[j].Paths
	})
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()

	org1ID := uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"))
	org2ID := uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a"))

	tests := []struct {
		name       string
		orgID      uuid.UUID
		folderName string
		folders    []folder.Folder
		want       []folder.Folder
	}{
		{
			name:       "folder does not exist",
			orgID:      org1ID,
			folderName: "nonexistent",
			folders:    []folder.Folder{},
			want:       nil,
		},
		{
			name:       "folder exists but has no children",
			orgID:      org1ID,
			folderName: "creative-scalphunter",
			folders: []folder.Folder{
				{
					Name:  "creative-scalphunter",
					OrgId: org1ID,
					Paths: "creative-scalphunter",
				},
			},
			want: []folder.Folder{},
		},
		{
			name:       "folder with direct children",
			orgID:      org1ID,
			folderName: "creative-scalphunter",
			folders: []folder.Folder{
				{
					Name:  "creative-scalphunter",
					OrgId: org1ID,
					Paths: "creative-scalphunter",
				},
				{
					Name:  "clear-arclight",
					OrgId: org1ID,
					Paths: "creative-scalphunter.clear-arclight",
				},
				{
					Name:  "topical-micromax",
					OrgId: org1ID,
					Paths: "creative-scalphunter.clear-arclight.topical-micromax",
				},
			},
			want: []folder.Folder{
				{
					Name:  "clear-arclight",
					OrgId: org1ID,
					Paths: "creative-scalphunter.clear-arclight",
				},
				{
					Name:  "topical-micromax",
					OrgId: org1ID,
					Paths: "creative-scalphunter.clear-arclight.topical-micromax",
				},
			},
		},
		{
			name:       "folder in different org",
			orgID:      org2ID,
			folderName: "creative-scalphunter",
			folders: []folder.Folder{
				{
					Name:  "creative-scalphunter",
					OrgId: org1ID,
					Paths: "creative-scalphunter",
				},
				{
					Name:  "clear-arclight",
					OrgId: org1ID,
					Paths: "creative-scalphunter.clear-arclight",
				},
			},
			want: nil,
		},
		{
			name:       "multiple nested children",
			orgID:      org1ID,
			folderName: "creative-scalphunter",
			folders: []folder.Folder{
				{
					Name:  "creative-scalphunter",
					OrgId: org1ID,
					Paths: "creative-scalphunter",
				},
				{
					Name:  "clear-arclight",
					OrgId: org1ID,
					Paths: "creative-scalphunter.clear-arclight",
				},
				{
					Name:  "topical-micromax",
					OrgId: org1ID,
					Paths: "creative-scalphunter.topical-micromax",
				},
				{
					Name:  "bursting-lionheart",
					OrgId: org1ID,
					Paths: "creative-scalphunter.clear-arclight.bursting-lionheart",
				},
			},
			want: []folder.Folder{
				{
					Name:  "clear-arclight",
					OrgId: org1ID,
					Paths: "creative-scalphunter.clear-arclight",
				},
				{
					Name:  "topical-micromax",
					OrgId: org1ID,
					Paths: "creative-scalphunter.topical-micromax",
				},
				{
					Name:  "bursting-lionheart",
					OrgId: org1ID,
					Paths: "creative-scalphunter.clear-arclight.bursting-lionheart",
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := folder.NewDriver(tt.folders)
			got := f.GetAllChildFolders(tt.orgID, tt.folderName)
			assert.Equal(t, tt.want, got)
		})
	}
}
