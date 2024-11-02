package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()

	org1ID := uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"))
	org2ID := uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a"))

	tests := []struct {
		name    string
		srcName string
		dstName string
		folders []folder.Folder
		want    []folder.Folder
		wantErr string
	}{
		{
			name:    "source folder does not exist",
			srcName: "nonexistent",
			dstName: "creative-scalphunter",
			folders: []folder.Folder{
				{
					Name:  "creative-scalphunter",
					OrgId: org1ID,
					Paths: "creative-scalphunter",
				},
			},
			want:    nil,
			wantErr: "source folder does not exist",
		},
		{
			name:    "destination folder does not exist",
			srcName: "creative-scalphunter",
			dstName: "nonexistent",
			folders: []folder.Folder{
				{
					Name:  "creative-scalphunter",
					OrgId: org1ID,
					Paths: "creative-scalphunter",
				},
			},
			want:    nil,
			wantErr: "destination folder does not exist",
		},
		{
			name:    "move to self",
			srcName: "creative-scalphunter",
			dstName: "creative-scalphunter",
			folders: []folder.Folder{
				{
					Name:  "creative-scalphunter",
					OrgId: org1ID,
					Paths: "creative-scalphunter",
				},
			},
			want:    nil,
			wantErr: "cannot move a folder to itself",
		},
		{
			name:    "move to different organization",
			srcName: "creative-scalphunter",
			dstName: "clear-arclight",
			folders: []folder.Folder{
				{
					Name:  "creative-scalphunter",
					OrgId: org1ID,
					Paths: "creative-scalphunter",
				},
				{
					Name:  "clear-arclight",
					OrgId: org2ID,
					Paths: "clear-arclight",
				},
			},
			want:    nil,
			wantErr: "cannot move a folder to a different organization",
		},
		{
			name:    "move to child of itself",
			srcName: "creative-scalphunter",
			dstName: "clear-arclight",
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
			want:    nil,
			wantErr: "cannot move a folder to a child of itself",
		},
		{
			name:    "successful simple move",
			srcName: "clear-arclight",
			dstName: "topical-micromax",
			folders: []folder.Folder{
				{
					Name:  "clear-arclight",
					OrgId: org1ID,
					Paths: "creative-scalphunter.clear-arclight",
				},
				{
					Name:  "topical-micromax",
					OrgId: org1ID,
					Paths: "topical-micromax",
				},
			},
			want: []folder.Folder{
				{
					Name:  "clear-arclight",
					OrgId: org1ID,
					Paths: "topical-micromax.clear-arclight",
				},
				{
					Name:  "topical-micromax",
					OrgId: org1ID,
					Paths: "topical-micromax",
				},
			},
			wantErr: "",
		},
		{
			name:    "move folder with children",
			srcName: "clear-arclight",
			dstName: "topical-micromax",
			folders: []folder.Folder{
				{
					Name:  "clear-arclight",
					OrgId: org1ID,
					Paths: "creative-scalphunter.clear-arclight",
				},
				{
					Name:  "bursting-lionheart",
					OrgId: org1ID,
					Paths: "creative-scalphunter.clear-arclight.bursting-lionheart",
				},
				{
					Name:  "topical-micromax",
					OrgId: org1ID,
					Paths: "topical-micromax",
				},
			},
			want: []folder.Folder{
				{
					Name:  "clear-arclight",
					OrgId: org1ID,
					Paths: "topical-micromax.clear-arclight",
				},
				{
					Name:  "bursting-lionheart",
					OrgId: org1ID,
					Paths: "topical-micromax.clear-arclight.bursting-lionheart",
				},
				{
					Name:  "topical-micromax",
					OrgId: org1ID,
					Paths: "topical-micromax",
				},
			},
			wantErr: "",
		},
		{
			name:    "move folder with deep nesting",
			srcName: "clear-arclight",
			dstName: "echo",
			folders: []folder.Folder{
				{
					Name:  "clear-arclight",
					OrgId: org1ID,
					Paths: "creative-scalphunter.clear-arclight",
				},
				{
					Name:  "bursting-lionheart",
					OrgId: org1ID,
					Paths: "creative-scalphunter.clear-arclight.bursting-lionheart",
				},
				{
					Name:  "topical-micromax",
					OrgId: org1ID,
					Paths: "creative-scalphunter.clear-arclight.bursting-lionheart.topical-micromax",
				},
				{
					Name:  "echo",
					OrgId: org1ID,
					Paths: "echo",
				},
			},
			want: []folder.Folder{
				{
					Name:  "clear-arclight",
					OrgId: org1ID,
					Paths: "echo.clear-arclight",
				},
				{
					Name:  "bursting-lionheart",
					OrgId: org1ID,
					Paths: "echo.clear-arclight.bursting-lionheart",
				},
				{
					Name:  "topical-micromax",
					OrgId: org1ID,
					Paths: "echo.clear-arclight.bursting-lionheart.topical-micromax",
				},
				{
					Name:  "echo",
					OrgId: org1ID,
					Paths: "echo",
				},
			},
			wantErr: "",
		},
		{
			name:    "same name folders in different paths",
			srcName: "clear-arclight",
			dstName: "echo",
			folders: []folder.Folder{
				{
					Name:  "clear-arclight",
					OrgId: org1ID,
					Paths: "creative-scalphunter.clear-arclight",
				},
				{
					Name:  "bursting-lionheart",
					OrgId: org1ID,
					Paths: "creative-scalphunter.clear-arclight.bursting-lionheart",
				},
				{
					Name:  "clear-arclight",
					OrgId: org1ID,
					Paths: "topical-micromax.clear-arclight",
				},
				{
					Name:  "echo",
					OrgId: org1ID,
					Paths: "echo",
				},
			},
			want: []folder.Folder{
				{
					Name:  "clear-arclight",
					OrgId: org1ID,
					Paths: "echo.clear-arclight",
				},
				{
					Name:  "bursting-lionheart",
					OrgId: org1ID,
					Paths: "echo.clear-arclight.bursting-lionheart",
				},
				{
					Name:  "clear-arclight",
					OrgId: org1ID,
					Paths: "topical-micromax.clear-arclight",
				},
				{
					Name:  "echo",
					OrgId: org1ID,
					Paths: "echo",
				},
			},
			wantErr: "",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := folder.NewDriver(tt.folders)
			got, err := f.MoveFolder(tt.srcName, tt.dstName)
			if tt.wantErr != "" {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
