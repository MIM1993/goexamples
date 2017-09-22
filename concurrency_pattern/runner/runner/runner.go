// 代码来自《GO语言实战》
// runner包用于展示如何使用通道来监视程序的执行时间
// runner包管理处理任务的运行和生命周期
// 这个程序可能作为cron作业执行

// 在设计上，可支持以下终止点：
// ~程序可以在分配的时间内完成工作，正常终止
// ~程序没有及时完成工作，自杀
// ~接收到操作系统发送的中断时间，程序立刻试图清理状态并停止工作
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner在给定的超时时间内执行一组任务，并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	// interrupt通道报告从操作系统发送的信号
	interrupt chan os.Signal

	// complete通道报告处理任务已经完成
	complete chan error

	// timeout报告处理任务已经超时
	// 单向通道，只支持读
	timeout <-chan time.Time

	// tasks持有一组以索引顺序依次执行的函数
	tasks []func(int)
}

// ErrTimeout会在任务执行超时时返回
// func New(text string) error
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt会在接收到操作系统的事件时返回
var ErrInterrupt = errors.New("received interrupt")

// New返回一个新的准备使用的Runner
// type Duration int64
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		// func After(d Duration) <-chan Time
		timeout: time.After(d),
	}
}

// Add将一个任务附加到Runner上，这个任务是一个接收int类型的ID作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start执行所有任务，并监视通道事件
func (r *Runner) Start() error {
	// 我们希望接收所有中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	// 用不同的goroutine执行不同的任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	// 当任务处理完成时发出的信号
	case err := <-r.complete:
		return err
	// 当任务处理程序运行超时时发出的信号
	case <-r.timeout:
		return ErrTimeout
	}
}

// run执行每个已注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		// 检测操作系统的中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		// 执行已注册的任务
		task(id)
	}

	return nil
}

// gotInerrupt验证是否接收到了中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	// 当中断时间被触发时发出的信号
	case <-r.interrupt:
		// 停止接收后续的任何信号
		signal.Stop(r.interrupt)
		return true
	// 继续正常运行
	default:
		return false
	}
}
