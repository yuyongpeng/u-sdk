package asynq

import (
	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
	"time"
)

type AsynqTestSUite struct {
	suite.Suite
	Redis string
}

func (suite *AsynqTestSUite) SetupSuite() {
	suite.Redis = "127.0.0.1:6379"
}

// 发布数据到redis中
func (suite *AsynqTestSUite) TestProduce() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: suite.Redis})
	defer client.Close()

	// 创建任务
	task, err := NewEmailDeliveryTask(42, "some:template:id")
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}
	// 发送要处理的任务到redis
	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	// asynq.ProcessIn(10*time.Second) : 在 10 秒后执行任务
	// asynq.ProcessAt(time.Now().Add(10 * time.Second)) : 在 10 秒后执行任务
	// asynq.MaxRetry(5) 最多重试5次， 最大默认就25次。
	// asynq.TaskID("mytaskid") 自动生成唯一任务ID
	// asynq.Unique(time.Hour) 为任务创建唯一性锁
	// asynq.Timeout(time.Minute) 任务超时时间
	// asynq.Deadline(time.Now().Add(2 * time.Hour)) 任务截止时间
	info, err = client.Enqueue(task, asynq.ProcessIn(24*time.Hour))
	if err != nil {
		log.Fatalf("could not schedule task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	task, err = NewImageResizeTask("https://example.com/myassets/image.jpg")
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}
	// 指定重试次数和超时时间
	info, err = client.Enqueue(task, asynq.MaxRetry(10), asynq.Timeout(3*time.Minute))
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}

// 消费队列中的数据
func (suite *AsynqTestSUite) TestConsume() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: suite.Redis},
		asynq.Config{
			// 任务的最大并发数
			Concurrency: 10,
			// 可以选择指定具有不同优先级的多个队列。
			Queues: map[string]int{
				"critical": 6, // 高优先级
				"default":  3,
				"low":      1, // 低优先级
			},
			StrictPriority:  false,           // 有先处理优先级高的队列，没数据了才处理优先级低的队列
			ShutdownTimeout: 8 * time.Second, // 关闭进程后，让工作携程继续工作完成任务的时间。默认8秒

			HealthCheckInterval:      15 * time.Second, // 指定健康检查的时间间隔
			DelayedTaskCheckInterval: 5 * time.Second,  // 指定对“计划”和“重试”任务运行的检查之间的间隔，以及如果已准备好处理这些任务，则将其转发到“挂起”状态。
			// 指定任务的等待时间。如果在此时间段内收到任务，则服务器将等待另一个相同长度的时间段. 最多等待 GroupMaxDelay 时间。
			//如果未设置或为零，则宽限期设置为 1 分钟。GroupGracePeriod 的最短持续时间为 1 秒。如果指定的值小于一秒，则对 NewServer 的调用将崩溃
			GroupGracePeriod: 15 * time.Second, //
			// GroupMaxDelay 指定服务器在将任务聚合到组中之前等待传入任务的最长时间。
			//如果未设置或为零，则不使用延迟限制
			GroupMaxDelay: 15 * time.Second,
			// GroupMaxSize 指定可以聚合到组中单个任务中的最大任务数。如果达到 GroupMaxSize，服务器会立即将任务聚合为一个任务。
			//如果未设置或为零，则不使用大小限制
			GroupMaxSize: 10,
			// GroupAggregator 指定用于将组中的多个任务聚合为一个任务的聚合函数。
			//如果未设置或为零，则将在服务器上禁用组聚合功能
			//GroupAggregator:
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeEmailDelivery, HandleEmailDeliveryTask)
	mux.Handle(TypeImageResize, NewImageProcessor())

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// group 处理
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 测试 任务的分组功能
func (suite *AsynqTestSUite) TestGroupProduce() {
	c := asynq.NewClient(asynq.RedisClientOpt{Addr: suite.Redis, DB: 1})
	defer c.Close()
	message := "ONE"
	task := asynq.NewTask("aggregation-tutorial", []byte(message))
	info, err := c.Enqueue(task, asynq.Queue("tutorial"), asynq.Group("example-group"))
	if err != nil {
		log.Fatalf("Failed to enqueue task: %v", err)
	}
	log.Printf("Successfully enqueued task: %s", info.ID)
}

// 测试 任务的分组消费
func (suite *AsynqTestSUite) TestGroupConsume() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: suite.Redis, DB: 1},
		asynq.Config{
			Queues:          map[string]int{"tutorial": 10},       // 队列
			GroupAggregator: asynq.GroupAggregatorFunc(aggregate), // 聚合的处理函数
			// 指定任务的等待时间。如果在此时间段内收到任务，则服务器将等待另一个相同长度的时间段. 最多等待 GroupMaxDelay 时间。
			//如果未设置或为零，则宽限期设置为 1 分钟。GroupGracePeriod 的最短持续时间为 1 秒。如果指定的值小于一秒，则对 NewServer 的调用将崩溃
			GroupGracePeriod: 10 * time.Second,
			GroupMaxDelay:    30 * time.Second, // 组 等待数据聚合的时间
			GroupMaxSize:     3,                // 组 的大小
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc("aggregated-task", handleAggregatedTask)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

// 运行suite中的所有测试方法
func TestSuite(t *testing.T) {
	suite.Run(t, new(AsynqTestSUite))
}
