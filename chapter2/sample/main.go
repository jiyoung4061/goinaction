package main

import (
	"log"
	"os"

	_ "github.com/webgenie/go-in-action/chapter2/sample/matchers"
	"github.com/webgenie/go-in-action/chapter2/sample/search"
)

// init함수는 main함수보다 먼저 호출!
func init(){
	// 표준출력으로 로그 출력하도록 변경
	log.SetOutput(os.Stdout)
}

// 프로그램 진입점
func main() {
	// 지정된 검색어로 검색 수행
	search.Run("Sherlock Holmes")
}
