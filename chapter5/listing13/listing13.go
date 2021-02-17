/* atomic 패키지의 함수들을 이용하여 숫자 타입에 안전하게 접근하는 예제 */
package listng13
import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	// 공유 자원으로 활용될 변수
	counter int64

	// 프로그램이 종료될 때까지 대기할 WaitGroup
	wg sync.WaitGroup
)

func main() {
	// 고루틴당 하나씩 두개 카운터 추가
	wg.Add(2)

	// 두개 고루틴 생성
	go incCounter(1)
	go incCounter(2)

	// 고루틴 실행이 종료될 때까지 대기
	wg.Wait()

	// 최종 결과 출력
	fmt.Println("최종 결과 : ", counter)
}

// 패키지 수준에 정의된 counter 변수의 값을 증가시키는 함수
func incCounter(id int) {
	// 함수 실행이 종료되면 main함수에 알리기 위해 Done 함수 호출을 예약한다.
	defer wg.Done()

	for count:= 0; count <2; count++ {
		// counter 변수에 안전하게 1을 더한다.
		atomic.AddInt64(&counter, 1)

		// 스레드를 양보하고 실행 큐로 되돌아간다.
		runtime.Gosched()
	}
}