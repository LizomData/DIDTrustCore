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
	NotUser
	NotPrivileged
)
const (
	Success int = iota
)
