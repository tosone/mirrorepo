package errCode

type ErrNo struct {
	Code     int
	Describe string
}

var (
	Normal               = ErrNo{Code: 10000, Describe: ""}
	AddressNull          = ErrNo{Code: 10001, Describe: "address is null"}
	DatabaseErr          = ErrNo{Code: 10002, Describe: "database error"}
	CloneCannotBeStopped = ErrNo{Code: 10003, Describe: "repo clone cannot be stopped"}
	RepoIDNotValid       = ErrNo{Code: 10004, Describe: "repo id is not valid"}
)
