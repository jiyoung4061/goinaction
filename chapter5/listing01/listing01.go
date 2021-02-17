/* 고루틴을 생성하는 방법과 스케줄러의 동작을 설명하는 예제 */
package listing01

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 스케줄러가 사용할 하나의 논리 프로세서를 할당.
	runtime.GOMAXPROCS(1) // 스케줄러가 사용할 논리 프로세서의 개수를 스스로 조정

	// wg는 프로그램의 종료를 대기하기 위해 사용
	// 각각의 고루틴마다 하나씩 총 두개의 카운트를 추가
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("고루틴을 실행합니다.")

	// 익명 함수를 선언하고 고루틴을 생성
	go func() {
		// main함수에게 종료를 알리기위한 Done함수 호출을 예약
		defer wg.Done()

		// 알파벳을 세번 출력
		for count := 0; count < 3 ; count++ {
			for char := 'a' ; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// 익명 함수를 선언하고 고루틴 생성
	go func() {
		// main 함수에게 종료를 알리기위한 Done 함수 호출 예약
		defer wg.Done()

		// 알파벳 세번 출력
		for count:=0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++{
				fmt.Printf("%c ", char)
			}
		}
	}()

	// 고루틴이 종료될 때까지 대기
	fmt.Println("대기 중...")
	wg.Wait()

	fmt.Println("\n 프로그램을 종료합니다.")
}