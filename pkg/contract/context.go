package contract

type ClientIPCtxKey struct{}
type UserAgentCtxKey struct{}

var (
	ClientIPKey  = ClientIPCtxKey{}
	UserAgentKey = UserAgentCtxKey{}
)
