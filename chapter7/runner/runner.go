package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner 타입 : 주어진 시간동안 작업 수행 후, 인터럽트에 의해 실행종료
type Runner struct {
	interrupt chan os.Signal 	// 인터럽트 신호 수신을 위한 채널
	complete chan error 		// 처리 종료를 알리는 채널(에러	발생O error / 발생X nil)
	timeout <-chan time.Time	// 지정된 시간 초과를 알리는 채널 = 작업처리를 위해 사용된 시간을 관리함
	tasks []func(int)			// 작업목록을 인덱스순서로 저장한 슬라이스
}
/*
// 운영체제 신호를 표현하는 타입 => 운영체제마다 다름
type Signal interface {
	String() string
	Signal() 
}
*/

var ErrTimeout = errors.New("시간 초과")
var ErrInterrupt = errors.New("인터럽트 발생")

// 실행할 Runner 타입 값 리턴하는 함수
func New(d time.Duration) *Runner {
	return &Runner {
		interrupt:	make(chan os.Signal, 1), // 버퍼 크기가 1인 채널
		complete:	make(chan error),
		timeout:	time.After(d),
		/*
			time.After()의 return 값 = time.Time 타입의 채널
			런타임은 지정 시간이 지나면 이채널에 time.Time 값을 보냄
		*/
		// task는 제로값이 nil 슬라이스이므로 초기화X
	}
}

// Runner 타입에 작업 추가하는 메서드 (작업 : int형 id를 매개변수로 받는 함수)
func (r *Runner) Add(tasks ...func(int)){ // 가변 길이 매개변수 ( 개수 상관없이 모두 받아들임)
	r.tasks = append(r.tasks, tasks...)
}

// 저장된 모든 작업 실행 및 채널 이벤트 관찰
func (r *Runner) Start() error{
	// 인터럽트 신호 수신
	signal.Notify(r.interrupt, os.Interrupt)

	// 각 고루틴마다 각 작업 실행
	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete: // 작업 완료 신호 수신
		return err

	case <-r.timeout:		// 작업 시간 초과 수신
		return ErrTimeout
		
	}
}

// 개별 작업 실행 메서드
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() { // 인터럽트 신호 수신 여부 확인
			return ErrInterrupt
		}

		// 작업 실행
		task(id)
	}

	return nil
}

// 인터럽트 신호 수신 여부 확인 메서드
func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:	// 인터럽트 이벤트 발생시
		signal.Stop(r.interrupt) //이후 발생되는 인터럽트 신호를 수신하지 않도록
		return true
	default:
		return false
	}
}