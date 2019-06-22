/**
资源池
*/
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"io"
)

const (
	maxGoroutine = 5

	poolRes = 2
)

func main() {
	var wg sync.WaitGroup

	wg.Add(maxGoroutine)

	p, err := New(createConn, poolRes)

	if err != nil {
		fmt.Println(err)
		return
	}

	// 模拟查询数据库
	for query := 0; query < maxGoroutine; query++ {
		go func(q int){
			dbQuery(q, p)
			wg.Done()
		}(query)
	 }
	 
	 wg.Wait()
	 fmt.Println("start close")
	 p.Close()
}

func dbQuery(query int, pool *Pool) {
	conn, err := pool.Acquire()
	if err != nil {
		fmt.Println(err)
		return 
	}

	defer pool.Release(conn)

	// 模拟查询
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Println("第", query, "个查询，第",conn.(*dbConnect).ID,"个连接")
}

type dbConnect struct {
	ID int32
}

func (db *dbConnect) Close() error {
	fmt.Println("关闭连接", db.ID)
	return nil
}

var idCount int32
func createConn() (io.Closer, error) {
	id := atomic.AddInt32(&idCount, 1)
	return &dbConnect{id}, nil
}

var ErrPoolClosed = errors.New("资源池已经关闭")

type Pool struct {
	m sync.Mutex
	res chan io.Closer
	factory func() (io.Closer, error)
	closed bool
}

// 创建资源池
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size too small...")
	}

	return &Pool{
		factory: fn,
		res: make(chan io.Closer, size),
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.res:
		fmt.Println("Acquire share")
		if !ok {
			return nil, ErrPoolClosed
		}

		return r, nil
	default:
		fmt.Println("Acquire new")
		return p.factory()
	}
}

// 关闭资源池
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	close(p.res)

	for r := range p.res {
		r.Close()
	}
}

// 释放资源池
func (p *Pool)Release (r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.res <- r:
		fmt.Println("资源池释放")
	default:
		fmt.Println("资源池满了")
		r.Close()
	}
}