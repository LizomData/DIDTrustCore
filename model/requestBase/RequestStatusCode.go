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
	IllegalCharacter
	IncorrectFormat
)
const (
	Success int = iota
)
