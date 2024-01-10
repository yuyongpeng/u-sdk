package asynq

import (
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"golang.org/x/net/context"
	"log"
	"strings"
	"time"
)

// A list of task types.
const (
	TypeEmailDelivery = "email:deliver"
	TypeImageResize   = "image:resize"
)

type EmailDeliveryPayload struct {
	UserID     int
	TemplateID string
}

type ImageResizePayload struct {
	SourceURL string
}

// 创建一个异步任务,将数据放入payload中
func NewEmailDeliveryTask(userID int, tmplID string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailDeliveryPayload{UserID: userID, TemplateID: tmplID})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailDelivery, payload), nil
}

// 创建一个异步任务,将数据放入payload中
func NewImageResizeTask(src string) (*asynq.Task, error) {
	payload, err := json.Marshal(ImageResizePayload{SourceURL: src})
	if err != nil {
		return nil, err
	}
	// task options can be passed to NewTask, which can be overridden at enqueue time.
	return asynq.NewTask(TypeImageResize, payload, asynq.MaxRetry(5), asynq.Timeout(20*time.Minute)), nil
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// 处理异步任务
func HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var p EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Sending Email to User: user_id=%d, template_id=%s", p.UserID, p.TemplateID)
	// Email delivery code ...
	return nil
}

// ImageProcessor implements asynq.Handler interface.
type ImageProcessor struct {
	// ... fields for struct
}

func (processor *ImageProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p ImageResizePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Resizing image: src=%s", p.SourceURL)
	// Image resizing code ...
	return nil
}

func NewImageProcessor() *ImageProcessor {
	return &ImageProcessor{}
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// group 处理
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 将相同group的任务合并为一个任务
func aggregate(group string, tasks []*asynq.Task) *asynq.Task {
	log.Printf("Aggregating %d tasks from group %q", len(tasks), group)
	var b strings.Builder
	for _, t := range tasks {
		b.Write(t.Payload())
		b.WriteString("\n")
	}
	return asynq.NewTask("aggregated-task", []byte(b.String()))
}

// 聚合后的任务处理函数
func handleAggregatedTask(ctx context.Context, task *asynq.Task) error {
	log.Print("Handler received aggregated task")
	log.Printf("aggregated messags: %s", task.Payload())
	return nil
}
