package errcheckanalyzer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestErrCheckAnalizer(t *testing.T) {
	// функция analysistest.Run применяет тестируемый анализатор ErrCheckAnalyzer
	// к пакетам из папки testdata и проверяет ожидания
	analysistest.Run(t, analysistest.TestData(), ErrCheckAnalyzer, "./...")
}
