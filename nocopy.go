package crema

// noCopy is used with go vet's -copylocks to prevent accidental copies.
type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}
