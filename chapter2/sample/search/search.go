package search

import (
	"log"
	"sync"
)

// 검색을 처리할 검색기의 매핑 정보를 저장할 맵(map)
var matchers = make(map[string]Matcher)

// 검색로직을 수행할 Run 함수
func Run(searchTerm string) {
	// 검색할 피드의 목록을 조회
	feeds, err := RetrieveFeeds()
	/* 
		변수 선언 연산자 ( := )
		변수의 선언과 초기화를 동시에 수행
	*/
	if err != nil {
		log.Fatal(err) // log.Fatal 함수 : 오류값을 전달받아 종료전 오류 내용을 터미널창에 출력
	}

	// 버퍼가 없는 채널을 생성하여 화면에 표시할 검색 결과를 전달받는다.
	results := make(chan *Result)

	/* 처리 완료후 프로그램이 곧바로 종료되는 것을 방지 */
	// 1. 모든 피드를 처리할 때까지 기다릴 대기 그룹(Wait group)을 설정
	var waitGroup sync.WaitGroup // sync 패키지의 WaitGroup : 특정 고루틴이 작업을 완료했는지 추적가능
	// 고루틴의 실행이 종료될 때마다 전체 개수를 하나씩 줄여나감.

	// 2. 개별피드를 처리하는동안 대기해야할 고루틴의 개수 설정
	waitGroup.Add(len(feeds))

	// 각기 다른 종류의 피드를 처리할 고루틴을 실행
	for _, feed := range feeds {
		/*
		for range : 요소 하나당 2개 값 return (요소의 index, 요소의 복사본)
		_ (빈 식별자) : index값이 대입될 변수를 대체 => 여러개 값을 return할 경우 빈 식별자를 사용해 특정 리턴값 무시 가능
		*/

		// 검색을 위해 검색기 초기화
		// Matcher타입 값이 맵에 존재하는지 확인
		matcher, exists := matchers[feed.Type]
		/*
			matchers의 return 값은 2개 : 검색된 키 값, 키 존재여부(boolean)
										==> 키 존재시 해당 키 값의 복사본 return 
		*/
		if !exists { 
			matcher = matcher["default"]
		}

		// 검색을 실행하기위해 고루틴 실행
		go func(matcher Matcher, feed *Feed){
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// 모든 작업이 완료되었는지를 모니터링할 고루틴 실행
	go func(){
		// 모든 작업이 처리될때까지 기다린다.
		waitGroup.Wait()

		// Display 함수에게 프로그램을 종료할수 있음을 알리기위해 채널을 닫는다.
		close(results)
	}()

	// 검색 결과를 화면에 표시하고 마지막 결과를 표시한 뒤 리턴
	Display(results)
}

// 프로그램에서 사용할 검색기를 등록할 함수를 정의한다.
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "검색기가 이미 등록되었습니다.")
	}

	log.Println("등록완료:", feedType, " 검색기")
	matchers[feedType] = matcher
}