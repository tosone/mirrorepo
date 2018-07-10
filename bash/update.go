package bash

// Update ..
func Update(dir string) (err error) {
	_, err = Run(dir, "git remote update --prune")
	return
}
