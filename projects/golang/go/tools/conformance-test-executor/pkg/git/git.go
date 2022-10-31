package git

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
)

func GetHeadAndBaseHashes(repoUrl string, branch string) (headSha string, baseSha string, err error) {
	// Clones the given repository in memory, creating the remote, the local
	// branches and fetching the objects, exactly as:
	fmt.Printf("git cloning %s into memory...", repoUrl)

	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: repoUrl,
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
	})

	if err != nil{
		return "", "", fmt.Errorf("cloning: %bv", err)
	}

	// retrieve the branch
	ref, err := r.Head()
	if err != nil{
		return "", "", fmt.Errorf("retrieving HEAD: %v", err)
	}

	// ... retrieves the commit history
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil{
		return "", "", fmt.Errorf("cloning: %bv", err)
	}

	// iterate over the first two commits, getting what will be the head and the base for the branch
	head, err := cIter.Next()
	if err != nil {
		return "", "", fmt.Errorf("retrieving SHA of head commit: %v", err)
	}
	headSha = head.Hash.String()

	base, err := cIter.Next()
	if err != nil {
		return "", "", fmt.Errorf("retrieving SHA of base commit: %v", err)
	}
	baseSha = base.Hash.String()

	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Println(c)
		return nil
	})
	if err != nil{
		return "", "", fmt.Errorf("iterating over commits: %bv", err)
	}
	return headSha, baseSha, nil
}
