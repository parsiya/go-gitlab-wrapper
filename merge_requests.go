// Merge request wrappers.
package wrapper

import (
	gl "github.com/xanzy/go-gitlab"
)

// CreateMerge creates a merge request from srcBranch to destBranch.
func (w *Wrapper) CreateMerge(title, description, srcBranch, destBranch string, squash bool) (*gl.MergeRequest, error) {

	opts := &gl.CreateMergeRequestOptions{
		Title:        &title,
		Description:  &description,
		SourceBranch: &srcBranch,
		TargetBranch: &destBranch,
		Squash:       &squash,
	}

	mr, _, err := w.MergeRequests.CreateMergeRequest(w.project, opts)
	return mr, err
}

// ListMergeRequests lists all merge requests from srcBranch to destBranch
// across all projects. Pass an empty value to search everything. E.g.,
// ListMergeRequests("", "") displays all merge requests.
func (w *Wrapper) ListMergeRequests(srcBranch, destBranch string) ([]*gl.MergeRequest, error) {

	opts := &gl.ListMergeRequestsOptions{
		SourceBranch: &srcBranch,
		TargetBranch: &destBranch,
		Scope:        gl.String("all"),
		// by default it returns only merge requests created by the user, this
		// returns all merge requests the user has access to.
	}

	mergeRequests, _, err := w.MergeRequests.ListMergeRequests(opts)
	return mergeRequests, err
}

// ListProjectMergeRequests lists all merge requests for a project. Pass empty
// values to show everything.
func (w *Wrapper) ListProjectMergeRequests(srcBranch, destBranch string) ([]*gl.MergeRequest, error) {

	opts := &gl.ListProjectMergeRequestsOptions{
		SourceBranch: &srcBranch,
		TargetBranch: &destBranch,
		Scope:        gl.String("all"),
	}

	mergeRequests, _, err := w.MergeRequests.ListProjectMergeRequests(w.project, opts)
	return mergeRequests, err
}

// CreateMergeRequestNote adds a note/comment to mergeRequestID.
func (w *Wrapper) CreateMergeRequestNote(mergeRequestID int, noteBody string) (*gl.Note, error) {

	opts := &gl.CreateMergeRequestNoteOptions{
		Body: &noteBody,
	}

	note, _, err := w.Notes.CreateMergeRequestNote(w.project, mergeRequestID, opts)
	return note, err
}

// ListMergeRequestNotes lists all notes for mergeRequestID.
func (w *Wrapper) ListMergeRequestNotes(mergeRequestID int) ([]*gl.Note, error) {

	opts := &gl.ListMergeRequestNotesOptions{} // opts supports sort and orderby and we do not care about these

	notes, _, err := w.Notes.ListMergeRequestNotes(w.project, mergeRequestID, opts)
	return notes, err
}

// UpdateMergeRequestNote replaces mergeRequestID's noteID with noteBody.
func (w *Wrapper) UpdateMergeRequestNote(mergeRequestID, noteID int, noteBody string) (*gl.Note, error) {

	opts := &gl.UpdateMergeRequestNoteOptions{
		Body: &noteBody,
	}

	note, _, err := w.Notes.UpdateMergeRequestNote(w.project, mergeRequestID, noteID, opts)
	return note, err
}

// DeleteMergeRequestNote deletes mergeRequestID's noteID.
func (w *Wrapper) DeleteMergeRequestNote(mergeRequestID, noteID int) error {

	_, err := w.Notes.DeleteMergeRequestNote(w.project, mergeRequestID, noteID)
	return err
}
