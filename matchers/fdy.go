package matchers

import "LogAnalysis/analysis"

type ftyMatcher struct{}

func init() {
	var match ftyMatcher
	analysis.Register("fdy", match)
}

func (m ftyMatcher) Analysis(logFile string) ([]*analysis.Result, error) {
	panic("implement me")
}
