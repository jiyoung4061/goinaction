/* 접근 동기화가 필요한 코드에 뮤텍스를 이용한 경쟁상태 해결 예제 */
package listing16

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg sync.WaitGroup
	// 코드의 임계 지역을 설정할 때 사용할 뮤텍스
	mutex sync.Mutex
)

func main(){
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Printf("최종결과: %d\n", counter)
}

// 패키지 수준에 정의된 counter변수의 값을 뮤텍스를 이용해 
// 안전하게 증가시키는 함수
func incCounter(id int){
	defer wg.Done()

	for count:= 0; count< 2; count++ {
		// 이 임계 지역에는 한번에 하나의 고루틴만 접근 가능하다.
		mutex.Lock()
		{ // 임계지역으로 중괄호가 필수는 아님!
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
		// 대기 중인 다른 고루틴이 접근할 수 있도록 잠금을 해제
	}
}