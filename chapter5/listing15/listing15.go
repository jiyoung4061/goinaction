/* atomic 패키지의 함수들을 이용해 숫자타입의 값을 안전하게 읽고 쓰는 예제 */
package listing15

import (
	"fmt"	
	"sync"
	"sync/atomic"
	"time"
)

var (
	// 실행중인 고루틴들의 종료 신호가 될 플래그
	shutdown int64

	// 프로그램이 종료될 때까지 대기할 WaitGroup
	wg sync.WaitGroup
)

func main() {
	// 고루틴당 하나씩 카운터 추가(총 2개)
	wg.Add(2)

	// 고루틴 두개 생성
	go doWork("A")
	go doWork("B")

	// 고루틴이 실행될 시간을 할애
	time.Sleep(1 * time.Second)

	// 종료 신호 플래그 설정
	fmt.Println("프로그램 종료!")
	atomic.StoreInt64(&shutdown, 1)

	// 고루틴 실행 종료까지 대기
	wg.Wait()
}

// 필요한 작업을 실행하다가 종료 플래그를 검사해
// 일찍 종료되는 함수를 흉내내는 함수
func doWork(name string) {
	// 함수 실행이 종료되면 main함수에 알리기 위해 Done 함수 호출 예약
	defer wg.Done()

	for {
		fmt.Printf("작업진행중: %s\n", name)
		time.Sleep(250 * time.Millisecond)

		// 종료 플래그를 확인하고 실행을 종료
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("작업을 종료합니다: %s\n", name)
			break
		}
	}
}
