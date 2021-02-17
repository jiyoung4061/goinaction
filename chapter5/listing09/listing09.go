/* 경쟁 상태를 재현한 예제 */
package listing09

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// 모든 고루틴이 값을 증가하려고 시도하는 변수
	counter int

	// 프로그램이 종료될 때까지 대기할 WaitGroup
	wg sync.WaitGroup
)

func main() {
	// 고루틴당 하나씩, 2개의 카운트를 추가.
	wg.Add(2)

	// 고루틴 생성
	go incCounter(1)
	go incCounter(2)

	// 고루틴 실행이 종료될때까지 대기
	wg.Wait()
	fmt.Println("최종결과 : " , counter)
}

// 패키지 수준에 정의된 counter 변수의 값을 증가시키는 함수
func incCounter(id int) {
	// 함수 실행이 종료되면 main 함수에 알리기 위해 Done 함수 호출을 예약
	defer wg.Done()

	for count := 0; count <2; count++ {
		// counter 변수 값을 읽는다.
		value := counter

		// 스레드를 양보해 큐로 돌아가도록 한다.
		runtime.Gosched()

		// 현재 카운터 값 증가
		value++

		// 원래 변수에 증가된 값 저장
		counter = value
	}

}