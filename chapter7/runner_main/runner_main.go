package runner_main

import(
	"log"
	"os"
	"time"

	"github.com/webgenie/go-in-action/chapter7/runner"
)

// 프로그램 실행 시간
const timeout = 3 * time.Second

func main() {
	log.Println("작업 시작...")
	
	r := runner.New(timeout) // 새로운 작업 실행기 생성

	r.Add(createTask(), createTask(), createTask()) // 수행 작업 등록

	// 작업 실행 및 결과 처리
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("지정된 시간 초과")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("인터럽트 발생")
			os.Exit(2)
		}
	}

	log.Println("프로그램 종료")
}

func createTask() func(int){
	return func(id int) {
		log.Printf("프로세스 - 작업 #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}