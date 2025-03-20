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
	NotFile
	RepoFailed
	InvalidType
	SBOMFailed
	FileUnzipFailed
	FileNotFound
	UploadFailed
)
const (
	Success int = iota
)
