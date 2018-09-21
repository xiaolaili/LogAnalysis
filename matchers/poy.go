package matchers

import "LogAnalysis/analysis"

type poyMatcher struct{}

func init() {
	var match poyMatcher
	analysis.Register("poy", match)
}

func (m poyMatcher) Analysis(logFile string) ([]*analysis.Result, error) {
	panic("implement me")
}
