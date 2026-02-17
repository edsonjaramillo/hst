package ports

// Environment resolves user/environment specific values.
type Environment interface {
	HistoryFilePath() (string, error)
}
