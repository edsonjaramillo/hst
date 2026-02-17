package ports

import "context"

// Selector chooses one or many commands from a candidate list.
type Selector interface {
	SelectMany(ctx context.Context, candidates []string, preselectAll bool) ([]string, error)
}
