package pfs

import (
	"fmt"
	"regexp"

	"github.com/pachyderm/pachyderm/src/client/pfs"
	"github.com/pachyderm/pachyderm/src/client/pkg/grpcutil"
)

// ErrFileNotFound represents a file-not-found error.
type ErrFileNotFound struct {
	File *pfs.File
}

// ErrRepoNotFound represents a repo-not-found error.
type ErrRepoNotFound struct {
	Repo *pfs.Repo
}

// ErrRepoExists represents a repo-exists error.
type ErrRepoExists struct {
	Repo *pfs.Repo
}

// ErrCommitNotFound represents a commit-not-found error.
type ErrCommitNotFound struct {
	Commit *pfs.Commit
}

// ErrNoHead represents an error encountered because a branch has no head (e.g.
// inspectCommit(master) when 'master' has no commits)
type ErrNoHead struct {
	Branch *pfs.Branch
}

// ErrCommitExists represents an error where the commit already exists.
type ErrCommitExists struct {
	Commit *pfs.Commit
}

// ErrCommitFinished represents an error where the commit has been finished
// (e.g from PutFile or DeleteFile)
type ErrCommitFinished struct {
	Commit *pfs.Commit
}

// ErrCommitDeleted represents an error where the commit has been deleted (e.g.
// from InspectCommit)
type ErrCommitDeleted struct {
	Commit *pfs.Commit
}

// ErrParentCommitNotFound represents a parent-commit-not-found error.
type ErrParentCommitNotFound struct {
	Commit *pfs.Commit
}

func (e ErrFileNotFound) Error() string {
	return fmt.Sprintf("file %v not found in repo %v at commit %v", e.File.Path, e.File.Commit.Repo.Name, e.File.Commit.ID)
}

func (e ErrRepoNotFound) Error() string {
	return fmt.Sprintf("repo %v not found", e.Repo.Name)
}

func (e ErrRepoExists) Error() string {
	return fmt.Sprintf("repo %v already exists", e.Repo.Name)
}

func (e ErrCommitNotFound) Error() string {
	return fmt.Sprintf("commit %v not found in repo %v", e.Commit.ID, e.Commit.Repo.Name)
}

func (e ErrNoHead) Error() string {
	return fmt.Sprintf("the branch \"%s\" is nil", e.Branch.Name)
}

func (e ErrCommitExists) Error() string {
	return fmt.Sprintf("commit %v already exists in repo %v", e.Commit.ID, e.Commit.Repo.Name)
}

func (e ErrCommitFinished) Error() string {
	return fmt.Sprintf("commit %v in repo %v has already finished", e.Commit.ID, e.Commit.Repo.Name)
}

func (e ErrCommitDeleted) Error() string {
	return fmt.Sprintf("commit %v/%v was deleted", e.Commit.Repo.Name, e.Commit.ID)
}

func (e ErrParentCommitNotFound) Error() string {
	return fmt.Sprintf("parent commit %v not found in repo %v", e.Commit.ID, e.Commit.Repo.Name)
}

// ByteRangeSize returns byteRange.Upper - byteRange.Lower.
func ByteRangeSize(byteRange *pfs.ByteRange) uint64 {
	return byteRange.Upper - byteRange.Lower
}

var commitNotFoundRe = regexp.MustCompile("commit [^ ]+ not found in repo [^ ]+")

// IsCommitNotFoundErr returns true if 'err' has an error message that matches
// ErrCommitNotFound
func IsCommitNotFoundErr(err error) bool {
	if err == nil {
		return false
	}
	return commitNotFoundRe.MatchString(grpcutil.ScrubGRPC(err).Error())
}

var commitDeletedRe = regexp.MustCompile("commit [^ ]+/[^ ]+ was deleted")

// IsCommitDeletedErr returns true if 'err' has an error message that matches
// ErrCommitDeleted
func IsCommitDeletedErr(err error) bool {
	if err == nil {
		return false
	}
	return commitDeletedRe.MatchString(grpcutil.ScrubGRPC(err).Error())
}
