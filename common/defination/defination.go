package defination

// WebServiceReturn WebServiceReturn
type WebServiceReturn struct {
	Code int
	Msg  string
}

// RepoStatus ..
type RepoStatus string

var (
	// Error ..
	Error RepoStatus = "error"
	// Success ..
	Success RepoStatus = "success"
	// Waiting ..
	Waiting RepoStatus = "waiting"
	// CloneConecting ..
	CloneConecting RepoStatus = "cloneConnecting"
	// CloneReceiving ..
	CloneReceiving RepoStatus = "cloneReceiving"
	// CloneResolving ..
	CloneResolving RepoStatus = "cloneResolving"
	// UpdateConnecting ..
	UpdateConnecting RepoStatus = "updateConnecting"
	// UpdateReceiving ..
	UpdateReceiving RepoStatus = "updateReceiving"
	// UpdateResolving ..
	UpdateResolving RepoStatus = "updateResolving"
)
