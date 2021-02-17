/* 고루틴을 생성하는 방법과 논리 프로세서가 두개인 경우 스케줄러의 동작 설명 */
package listing07

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 스케줄러에 두개의 논리 프로세서 할당
	runtime.GOMAXPROCS(2)
	
	// wg는 프로그램의 종료를 대기하기 위해 사용
	// 고루틴마다 하나씩, 2개 카운터 추가
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("고루틴 실행")
	
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


	// 고루틴 종료 대기
	fmt.Println("대기 중...")
	wg.Wait()

	fmt.Println("프로그램 종료")
}
