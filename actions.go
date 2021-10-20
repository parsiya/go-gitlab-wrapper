package wrapper

import (
	"encoding/base64"

	gl "github.com/xanzy/go-gitlab"
)

// CommitActionOptions is an alias of CommitActionOptions
type CommitActionOptions gl.CommitActionOptions

// CreateFileAction returns a CommitActionOptions to create a new file at path
// with content.
func CreateFileAction(path, content string) *CommitActionOptions {

	encoded := base64.StdEncoding.EncodeToString([]byte(content))
	return &CommitActionOptions{
		Action:   gl.FileAction(gl.FileCreate),
		FilePath: &path,
		Content:  &encoded,
		Encoding: base64Encoding,
	}
}

// DeleteFileAction returns a CommitActionOptions to delete the file at path.
func DeleteFileAction(path string) *CommitActionOptions {

	return &CommitActionOptions{
		Action:   gl.FileAction(gl.FileDelete),
		FilePath: &path,
	}
}

// MoveFileAction returns a CommitActionOptions to move a file from previousPath
// to path.
func MoveFileAction(path, previousPath string) *CommitActionOptions {

	return &CommitActionOptions{
		Action:       gl.FileAction(gl.FileMove),
		FilePath:     &path,
		PreviousPath: &previousPath,
	}
}

// UpdateFileAction returns a CommitActionOptions to update the file at path with
// content.
func UpdateFileAction(path, content string) *CommitActionOptions {

	encoded := base64.StdEncoding.EncodeToString([]byte(content))
	return &CommitActionOptions{
		Action:   gl.FileAction(gl.FileUpdate),
		FilePath: &path,
		Content:  &encoded,
		Encoding: base64Encoding,
	}
}

// ChmodFileAction returns a CommitActionOptions that changes the executable status of
// the file at path to execute_filemode.
func ChmodFileAction(path string, executeFilemode bool) *CommitActionOptions {

	return &CommitActionOptions{
		Action:          gl.FileAction(gl.FileChmod),
		FilePath:        &path,
		ExecuteFilemode: &executeFilemode,
	}
}

// NewDirectoryAction returns a CommitActionOptions to create a .gitkeep file
// in the new directory. This tells git to create the directory and store that
// file there.
//
// path: Just supply the path to the directory. .gitkeep will be added by the
// function.
func NewDirectoryAction(path string) *CommitActionOptions {

	path += "/.gitkeep"
	return CreateFileAction(path, "")
}
