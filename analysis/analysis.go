package analysis

import "log"

var matchers = make(map[string]Matcher)

func Run(file string) {

}

// 分析日志
func Analysis(match Matcher, logFile string, results chan<- *Result) {
	analysisResult, err := match.Analysis(logFile)
	if err != nil {
		log.Println(err)
		return
	}

	for _, result := range analysisResult {
		results <- result
	}
}

// 注册日志分析的匹配器
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
