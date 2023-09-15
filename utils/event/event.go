package event

import (
	"project/repository/db/model"
	"time"
)

type Observer interface {
	Update(event Event)
}

type Event struct {
	User model.User
	Time time.Time
	// 其他相关字段
}

type EmailNotifier struct {
	// Email通知器的具体实现
}

func (en *EmailNotifier) Update(event Event) {
	// 处理邮件通知的逻辑
}

type SMSNotifier struct {
	// 短信通知器的具体实现
}

func (sn *SMSNotifier) Update(event Event) {
	// 处理短信通知的逻辑
}
