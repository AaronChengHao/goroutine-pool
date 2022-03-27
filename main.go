package main

import (
	"fmt"
	"net/http"
	"time"
)

// Task 定义任务接口
type Task interface {
	running()
}

// GoroutinePool 协程池对象
type GoroutinePool struct {
	num      int
	taskChan chan Task
}

// 启动协程池
func (t *GoroutinePool) start() {
	t.taskChan = make(chan Task)
	// - 根据预设的数值启动相应数量的协程
	t.createGoroutine()

	// - 为了不产生死锁，起一个协程保持间歇运行状态
	go func() {
		for {
			time.Sleep(time.Second * 3)
		}
	}()
}

// 创建协程
func (t *GoroutinePool) createGoroutine() {
	for i := 1; i <= t.num; i++ {
		go func(taskChan chan Task) {
			for {
				select {
				case task := <-taskChan:
					task.running()
				}
			}
		}(t.taskChan)
	}
}

// GetHtml 任务对象
type GetHtml struct {
}

func (t *GetHtml) running() {
	fmt.Println("开始执行获取html内容任务")
	c := http.Client{}
	res, _ := c.Get("https://www.baidu.com")
	content := make([]byte, 1024)
	res.Body.Read(content)
	fmt.Println(string(content))
}

func main() {
	pool := &GoroutinePool{num: 100}
	pool.start()
	// - 投递任务
	pool.taskChan <- &GetHtml{}
	select {}
}
