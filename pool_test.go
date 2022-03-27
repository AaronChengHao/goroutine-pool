package goroutinepool

import (
	"fmt"
	"net/http"
	"testing"
)

// GetHtml 任务对象
type GetHtml struct {
}

func (t *GetHtml) Running() {
	fmt.Println("开始执行获取html内容任务")
	c := http.Client{}
	res, _ := c.Get("https://www.baidu.com")
	content := make([]byte, 1024)
	res.Body.Read(content)
	fmt.Println(string(content))
}

func TestPool(t *testing.T) {
	pool := New(100)
	pool.Start()
	pool.AddTask(&GetHtml{})
	// 让主协程保持等待，不退出
	select {}
}
