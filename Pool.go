package goroutinepool

import (
	"time"
)

// Task 定义任务接口
type Task interface {
	running()
}

// GoroutinePool 协程池对象
type GoroutinePool struct {
	Num      int
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
	for i := 1; i <= t.Num; i++ {
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

// 添加任务
func (t *GoroutinePool) addTask(task Task)  {
	t.taskChan <- task
}