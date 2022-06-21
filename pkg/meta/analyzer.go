package meta

// Analyzer is a single text analyzer or token generator.
type Analyzer int16

//go:generate enumer -type=Analyzer -json -trimprefix=Analyzer -transform=lower

// TODO: Support remaining analyzer and filter.

const (
	// AnalyzerUnset indicates that no analyzer is set.
	AnalyzerUnset Analyzer = iota
	// AnalyzerStandard is the standard analyzer.
	AnalyzerStandard
	// AnalyzerSimple is the simple analyzer.
	AnalyzerSimple
	AnalyzerKeyword
	AnalyzerWeb
	AnalyzerRegexp
	AnalyzerStop
	AnalyzerWhitespace
)
