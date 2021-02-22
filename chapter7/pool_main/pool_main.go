package pool_main

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/webgenie/go-in-action/chapter7/pool"
)

const (
	maxGoroutines = 25
	pooledResources = 2
)

type dbConnection struct {
	ID int32
}

func (dbConn *dbConnection) Close() error{
	log.Println("닫힘: 데이터베이스 연결", dbConn.ID)
	return nil
}

var idCounter int32
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter,1)
	log.Println("생성: 새 데이터베이스 연결", id)

	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()

	log.Println("프로그램 종료")
	p.Close()
}

func performQueries(query int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	defer p.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("질의 : QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}