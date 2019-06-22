/**
runner
*/
package main

import (
	"fmt"
	"time"
	"os"
	"os/signal"
	"errors"
)

var ErrTO = errors.New("执行者执行超时")
var ErrTI = errors.New("执行者被打断")

// 定义执行者
type Runner struct {
	tasks []func(int) // 等待执行的任务
	complete chan error // 任务通知
	timeout <-chan time.Time // 定义超时时间
	interrupt chan os.Signal // 监听信号
}

//执行所有任务
func (r *Runner) Start() error {
	//接收信号
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <- r.complete:
		return err
	case <-r.timeout:
		return ErrTO
	}
}

// 执行任务
func(r *Runner) run() error{
	for id, task := range r.tasks {
		if r.isInterrupt() {
			return ErrTI
		}

		task(id)
	}

	return nil
}

// 检查终端信号
func(r *Runner) isInterrupt() bool{
	select {
	case <- r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

// 初始化任务
func New(tm time.Duration) *Runner {
	return &Runner{
		complete:make(chan error),
		timeout:time.After(tm),
		interrupt:make(chan os.Signal, 1),
	}
}

// 添加任务
func(r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func main() {
	fmt.Println("start...")

	timeout := 5* time.Second
	r := New(timeout)

	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err!= nil {
		switch err {
		case ErrTO:
			fmt.Println(err)
			os.Exit(1)
		case ErrTI:
			fmt.Println(err)
			os.Exit(2)
		}
	}

	fmt.Println("...completed")
}

func createTask() func(int){
	return func(id int) {
		fmt.Println("执行任务id", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}