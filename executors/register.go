package executors

func (e *ExecutionEngine) RegisterTypes() error {
	for _, eType := range types {
		e.Executors = append(e.Executors, eType)
	}
	return nil
}
