// Package wrapper provides a wrapper for the go-gitlab package for common
// functions.
package wrapper

import (
	"encoding/base64"

	gl "github.com/xanzy/go-gitlab"
)

var base64Encoding = gl.String("base64")

// A Wrapper is a go-gitlab and communicates with the GitLab endpoint.
type Wrapper struct {
	*gl.Client
	project interface{} // project ID or path (e.g., "phakimian/myproject")
	email   string      // bot's email, email and name appear in the commits
	name    string      // bot's name
}

// Client creates a new Wrapper client.
func Client(project interface{}, token, baseURL, email, name string) (*Wrapper, error) {

	glab, err := gl.NewClient(token, gl.WithBaseURL(baseURL))
	// wrap := Wrapper{glab, project, email, name}
	return &Wrapper{glab, project, email, name}, err
}

// ListRepo list all files in the repository.
// path is the path inside the repository to list. branch is the target branch.
func (w *Wrapper) ListRepo(path, branch string) ([]*gl.TreeNode, error) {

	opts := &gl.ListTreeOptions{
		Path: &path,
		Ref:  &branch,
	}

	nodes, _, err := w.Repositories.ListTree(w.project, opts)
	return nodes, err
}

// NewFile creates a single new file in the target repository and branch.
func (w *Wrapper) NewFile(branch, commitMsg, fileName string, content []byte) (*gl.FileInfo, error) {

	// https://github.com/xanzy/go-gitlab/blob/master/examples/repository_files.go

	// Convert the file content to base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	opts := &gl.CreateFileOptions{
		AuthorEmail:   &w.email,
		AuthorName:    &w.name,
		Branch:        &branch,
		Content:       &encoded,       // send the file as base64
		Encoding:      base64Encoding, // "base64"
		CommitMessage: &commitMsg,
	}

	file, _, err := w.RepositoryFiles.CreateFile(w.project, fileName, opts)
	return file, err
}

// Commit creates a new commit with the included array of CommitActionOptions.
func (w *Wrapper) Commit(actions []*gl.CommitActionOptions, branch, commitMessage string) (*gl.Commit, error) {

	opts := &gl.CreateCommitOptions{
		Actions:       actions,
		Branch:        &branch,
		CommitMessage: &commitMessage,
		AuthorEmail:   &w.email,
		AuthorName:    &w.name,
	}

	commit, _, err := w.Commits.CreateCommit(w.project, opts)
	return commit, err
}
