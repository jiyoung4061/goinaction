/* 두개의 고루틴을 이용해 테니스 경기를 모방하는 예제 */
package listing20

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init(){
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// 버퍼가 없는 채널 생성
	court := make(chan int)

	// 고루틴당 하나씩 두개 카운터 추가
	wg.Add(2)

	// 선수 두명 입장
	go player("나달", court)
	go player("이형택", court)

	// 경기 시작
	court <- 1
	wg.Wait()
}

// 테니스 선수의 행동을 모방하는 player함수
func player(name string, court chan int){
	defer wg.Done()

	for { // 무한루프
		// 공이 되돌아올 때까지 기다린다.
		ball, ok := <-court
		if !ok {
			//  채널이 닫혔으면 승리한 것으로 간주
			fmt.Printf("%s 선수가 승리\n", name)
			return
		}

		// 랜덤 값을 이용해 공을 받아치지 못했는지 확인
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("%s 선수가 공을 받아치지 못했습니다.\n", name)

			// 채널을 닫아 현재 선수가 패배했음을 알림
			close(court)
			return
		}

		// 선수가 공을 받아친 횟수를 출력하고 그 값을 증가시킨다.
		fmt.Printf("%s 선수가 %d 번째 공을 받아쳤습니다\n", name, ball)

		ball++

		// 공을 상대 선수에게 보낸다.
		court <- ball
	}
}