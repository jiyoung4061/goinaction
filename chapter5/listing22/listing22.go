/* 버퍼가 없는 채널을 이용해 계주 경기를 묘사하는 예제 */
package listing22

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main(){
	baton := make(chan int)

	// 마지막 주자를 위해 카운터 하나 생성
	wg.Add(1)

	// 첫번째 주자 경기 준비
	go Runner(baton)

	// 경기 시작
	baton <- 1

	wg.Wait()
}

// 계주의 각 주자를 표현하는 Runner 함수
func Runner(baton chan int) {
	var newRunner int

	// 바톤을 전달받을 때까지 기다린다.
	runner := <- baton

	// 트랙을 달린다
	fmt.Printf("%d번째 주자가 바통을 받아 달리기 시작했습니다.\n", runner)

	// 새로운 주자가 교체 지점에서 대기한다.
	if runner != 4 {
		newRunner = runner +1
		fmt.Printf("%d 번째 주자가 대기합니다.\n", newRunner)
		go Runner(baton)
	}

	// 트랙을 달린다.
	time.Sleep(100 * time.Millisecond)

	// 경기 끝났는지 검사
	if runner == 4 {
		fmt.Printf("%d 번째 주자가 도착했습니다. 경기가 끝났습니다. \n", runner)
		wg.Done()
		return
	}

	// 다음 주자에게 바통을 넘긴다.
	fmt.Printf("%d번재 주자가 %d 번째 주자에게 바통을 넘겼습니다.\n", runner, newRunner)

	baton <- newRunner
}