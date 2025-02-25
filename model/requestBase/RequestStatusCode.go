package requestBase

const (
	LoginStatusInvalid int = iota + 1001
	NotLoggedIn
	InvalidTokens
	LoginFailed
	ParameterError
	TokenGenerationFailed
	RegisterFailed
	RegisterAlready
)
const (
	Success int = iota
)
