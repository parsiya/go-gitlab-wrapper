# go-gitlab-wrapper - Work in Progress
My wrapper for some [go-gitlab][go-gitlab-link] operations.

[go-gitlab-link]: https://github.com/xanzy/go-gitlab

# Usage
Only some parts of the API are wrapped. However, the wrapper embeds the
go-gitlab client so we can call its method directly. Please see the docs at
[https://pkg.go.dev/github.com/xanzy/go-gitlab][go-gitlab-docs].

[go-gitlab-docs]: https://pkg.go.dev/github.com/xanzy/go-gitlab

## Create a Wrapper
This must be done first.

* project: This is an interface similar to `go-gitlab`. This can be either an
  int with the project ID or a string with the project's path like
  `parsiya/project`.
* token: This can be a project or personal token as a string. Make sure the
  access token has the `api` scope.
* baseURL: The base URL of your GitLab instance. E.g., `gitlab.com` or
  `gitlab.example.net`.
* email and name: Email address and name used in the commits and notes.

[personal-token]: https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html
[project-token]: https://docs.gitlab.com/ee/user/project/settings/project_access_tokens.html

```go
//            func Client(project interface{}, token,       baseURL,              email,             name string)
wr, err := wrapper.Client("parsiya/project"  , "yourtoken", "https://gitlab.com", "bot@example.net", "My GitLab Bot")
if err != nil {
    panic(err)
}
```

## Create a New Branch
This code creates a new branch named `new-branch` from `main`.

```go
br, err := wr.CreateBranch("new-branch", "main")
if err != nil {
    panic(err)
}
fmt.Println(br)
```

## Create a New File
Now, we can create a new file in the branch above.

```go
commitMsg := "Add new-new-file.txt from the API"
content := "new-new-file.txt content"

file, err := wr.NewFile("new-branch", commitMsg, "new-new-file.txt", content)
if err != nil {
    panic(err)
}
fmt.Println(file)
```

## Create a Commit
If we need add/modify/delete multiple files, we can create a commit with
multiple actions. We need one action per operation.

```go
// Add some files to BranchName
// Each file should have its own action.
file1Contents := "File 1 contents"
file2Contents := "File 2 contents"

file1Opt := wrapper.NewFileAction("dir1/file1.txt", file1Contents)
file2Opt := wrapper.NewFileAction("dir2/file2.txt", file2Contents)

// Create an empty directory - creates a directory with just the ".gitkeep" file.
dir3Opt := wrapper.NewDirectoryAction("dir3")
dir4Opt := wrapper.NewDirectoryAction("dir3/dir4")

options := [4]*CommitActionOptions{
    file1Opt, file2Opt, dir3Opt, dir4Opt,
}
```

Add all actions to a commit

```go
commitMessage := "add a bunch of files and directories"
commit, err := wr.Commit(options[:], "new-branch", commitMessage)
if err != nil {
	panic(err)
}
fmt.Println(commit)
```

## Create a Merge Request
Create a merge request from `new-branch` branch to `main`.

```go
mr, err := wr.CreateMerge("Merge request title", "Merge request description", "new-branch", "main", true)
if err != nil {
	panic(err)
}
fmt.Println(mr)
// mr.IID is used in the rest of the example to modify this merge request.
```

## List Merge Requests
We can list all merge requests that the token has access to (if it's not a
project token). To see all merge requests, pass empty strings to this method.

```go
// List all merge requests from new-branch to main.
mrList, err := wr.ListMergesByBranch("new-branch", "main")
if err != nil {
	panic(err)
}
fmt.Println(mrList)
```

We can also list all merge requests accessible to the token for a specific
project with a different method.

```go
// List all merge requests for a project.
prjMRList, err := wr.ListProjectMergeRequests("", "")
if err != nil {
	panic(err)
}
fmt.Println(prjMRList)
```

## See Merge Request Notes
Comments for merge requests are called notes. Using the merge request IID from
before we can see all of its comments.

```go
// List all notes for a merge request.
notes, err := wr.ListMergeRequestNotes(mr.IID)
if err != nil {
	panic(err)
}
// fmt.Println(notes)
noteID := notes[0].ID
```

## Add a Note to a Merge Request
The note text will be rendered in markdown in the web interface. We can add
notes to closed merge requests, too.

```go
// Create a new note for the merge request.
note, err := wr.CreateMergeRequestNote(mr.IID, "test merge request note")
if err != nil {
	panic(err)
}
fmt.Println(note)
```

## List Notes for a Merge Request
See all notes for the previous merge request including the new note.

```go
notes, err := wr.ListMergeRequestNotes(mr.IID,)
if err != nil {
	panic(err)
}
fmt.Println(notes)
```

## Update a Merge Request Note
Update the previous note.

```go
// Add a note to the merge request.
note, err := wr.UpdateMergeRequestNote(mr.IID, noteID, "test merge request note")
if err != nil {
	panic(err)
}
fmt.Println(note)
```

## Delete a Merge Request Note
Delete the previous note.

```go
// Delete the note from the merge reques.
if err := wr.DeleteMergeRequestNote(mr.IID, noteID); err != nil {
	panic(err)
}
```

# License
The original project is licensed under the Apache Version 2.0 license. This
wrapper has the same license. Please see [LICENSE](LICENSE) for details.
