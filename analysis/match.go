package analysis

type Result struct {
}

type Matcher interface {
	Analysis(logFile string) ([]*Result, error)
}
