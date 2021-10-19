// branch wrappers.
package wrapper

import gl "github.com/xanzy/go-gitlab"

// CreateBranch creates a new branch in the target repository.
// branch: name of the new branch.
// refBranch: name of the parent branch. Cannot be empty.
func (w *Wrapper) CreateBranch(branch string, refBranch string) (*gl.Branch, error) {

	opts := &gl.CreateBranchOptions{
		Branch: &branch,
		Ref:    &refBranch,
	}

	br, _, err := w.Branches.CreateBranch(w.project, opts)
	return br, err
}

// ClearBranch deletes everything in branch with a single commit.
func (w *Wrapper) ClearBranch(branch, commitMessage string) (*gl.Commit, error) {

	// Get all nodes in the branch.
	nodes, err := w.ListRepo("", branch)
	if err != nil {
		return nil, err
	}

	commitActions := make([]*gl.CommitActionOptions, len(nodes))

	for i, node := range nodes {
		// We need one action per node.
		// Probably less-buggy if we used append instead, using append means we
		// cannot create the original array with any values because these new
		// ones will be appended and not replace the original empty ones.
		commitActions[i] = DeleteFileAction(node.Path)
	}

	return w.Commit(commitActions, branch, commitMessage)
}
