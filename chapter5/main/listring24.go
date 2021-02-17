/* 버퍼가 있는 채널을 이용해 미리 정해진 고루틴의 개수만큼 다중 작업을 수행하는 예제 */
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4 // 실행할 고루틴의 개수
	taskLoad = 10 // 처리할 작업의 개수
)

var wg sync.WaitGroup

func init() {
	// 랜덤 값 생성기 초기화
	rand.Seed(time.Now().Unix())
}

func main() {
	// 작업 부하를 관리하기 위한 버퍼가 있는 채널을 생성
	tasks := make(chan string, taskLoad)

	// 작업을 처리할 고루틴 실행
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// 실행할 작업 추가
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("작업: %d", post)
	}

	// 작업을 모두 처리하면 채널을 닫는다.
	close(tasks)

	wg.Wait()
}

// 버퍼가 있는 채널에서 수행할 작업을 가져가는 고루틴
func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		// 작업이 할당될 때까지 대기
		task, ok := <- tasks
		if !ok {
			// 채널이 닫힌 경우
			fmt.Printf("작업자 : %d : 종료합니다.\n", worker)
			return
		}

		// 작업을 시작하는 메시지 출력
		fmt.Printf("작업자 : %d : 작업시작 : %s \n", worker, task)

		// 작업을 처리(임의 시간 대기)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// 작업 완료 메세지
		fmt.Printf("작업자 : %d : 작업완료: %s \n", worker, task)
	}
}