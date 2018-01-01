package defination

// WebServiceReturn WebServiceReturn
type WebServiceReturn struct {
	Code int
	Msg  string
}

type RepoStatus string

var (
	Error            RepoStatus = "error"
	Stopped          RepoStatus = "stopped"
	Success          RepoStatus = "success"
	Waiting          RepoStatus = "waiting"
	CloneConecting   RepoStatus = "cloneConnecting"
	CloneReceiving   RepoStatus = "cloneReceiving"
	CloneResolving   RepoStatus = "cloneResolving"
	UpdateConnecting RepoStatus = "updateConnecting"
	UpdateReceiving  RepoStatus = "updateReceiving"
	UpdateResolving  RepoStatus = "updateResolving"
)
