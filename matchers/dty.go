package matchers

import "LogAnalysis/analysis"

type dtyMatcher struct{}

func init() {
	var match dtyMatcher
	analysis.Register("dty", match)
}

func (m dtyMatcher) Analysis(logFile string) ([]*analysis.Result, error) {
	panic("implement me")
}
