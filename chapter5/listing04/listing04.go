/* 단일 스레드 환경에서 고루틴이 스케줄러에 의해 분할 실행되는 것을 보여주는 예제 */
package listing04

import (
	"fmt"
	"runtime"
	"sync"
)

// wg는 프로그램의 종료를 대기하기 위해 사용
var wg sync.WaitGroup

func main() {
	// 스케줄러에 하나의 논리 프로세서만 할당
	runtime.GOMAXPROCS(1)
	

	// 고루틴마다 하나씩, 2개 카운터 추가
	wg.Add(2)

	// 두개의 고루틴 생성
	fmt.Println("고루틴 실행")
	go printPrime("A")
	go printPrime("B")

	// 고루틴 종료 대기
	fmt.Println("대기 중...")
	wg.Wait()

	fmt.Println("프로그램 종료")
}

// 소수 중 처음 5000개 출력
func printPrime(prefix string) {
	// 작업이 완료되면 Done 함수 호출 예약
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++{
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}

	fmt.Println("완료: ", prefix)
}