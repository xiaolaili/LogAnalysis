package analysis

type defaultMatcher struct{}

func init() {
	var match defaultMatcher
	Register("default", match)
}

func (m defaultMatcher) Analysis(logFile string) ([]*Result, error) {
	return nil, nil
}
