package main

import (
	"fmt"
	"github.com/gzltommy/xxl-job-executor-test/job"
	xxl "github.com/xxl-job/xxl-job-executor-go"
	"log"
)

func main() {
	StartXXLJobClient()
}

func StartXXLJobClient() {
	// 创建一个执行器
	exec := xxl.NewExecutor(
		xxl.RegistryKey("testJob"),                                  // 执行器名称
		xxl.ServerAddr("http://192.168.101.133:8080/xxl-job-admin"), // 调度中心地址
		xxl.AccessToken("default_token"),                            // 请求令牌(默认为空)
		xxl.ExecutorPort("9090"),                                    // 默认 9999（非必填）
		xxl.SetLogger(&logger{}),                                    // 自定义日志
		//xxl.ExecutorIp("127.0.0.1"),      // 可自动获取
	)

	// 初始化
	exec.Init()

	//设置日志查看 handler
	exec.LogHandler(func(req *xxl.LogReq) *xxl.LogRes {
		return &xxl.LogRes{
			Code: 200,
			Msg:  "",
			Content: xxl.LogResContent{
				FromLineNum: req.FromLineNum,
				ToLineNum:   2,
				LogContent:  "这个是自定义日志 handler",
				IsEnd:       true,
			}}
	})

	//注册任务 handler
	exec.RegTask("TestJob1", job.TestJob1)
	exec.RegTask("TestJob2", job.TestJob2)

	// 运行
	err := exec.Run()
	log.Fatal(err)
}

// xxl.Logger 接口实现
type logger struct{}

func (l *logger) Info(format string, a ...interface{}) {
	fmt.Println(fmt.Sprintf("自定义日志 - "+format, a...))
}

func (l *logger) Error(format string, a ...interface{}) {
	log.Println(fmt.Sprintf("自定义日志 - "+format, a...))
}
