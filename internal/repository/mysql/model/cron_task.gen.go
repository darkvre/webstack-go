// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCronTask = "cron_task"

// CronTask mapped from table <cron_task>
type CronTask struct {
	ID                  int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`
	Name                string    `gorm:"column:name;not null;comment:任务名称" json:"name"`
	Spec                string    `gorm:"column:spec;not null;comment:crontab 表达式" json:"spec"`
	Command             string    `gorm:"column:command;not null;comment:执行命令" json:"command"`
	Protocol            int64     `gorm:"column:protocol;not null;default:1;comment:执行方式 1:shell 2:http" json:"protocol"`
	HTTPMethod          int64     `gorm:"column:http_method;not null;default:1;comment:http 请求方式 1:get 2:post" json:"http_method"`
	Timeout             int64     `gorm:"column:timeout;not null;default:60;comment:超时时间(单位:秒)" json:"timeout"`
	RetryTimes          int64     `gorm:"column:retry_times;not null;default:3;comment:重试次数" json:"retry_times"`
	RetryInterval       int64     `gorm:"column:retry_interval;not null;default:60;comment:重试间隔(单位:秒)" json:"retry_interval"`
	NotifyStatus        int64     `gorm:"column:notify_status;not null;comment:执行结束是否通知 1:不通知 2:失败通知 3:结束通知 4:结果关键字匹配通知" json:"notify_status"`
	NotifyType          int64     `gorm:"column:notify_type;not null;default:1;comment:通知类型 1:邮件 2:webhook" json:"notify_type"`
	NotifyReceiverEmail string    `gorm:"column:notify_receiver_email;not null;comment:通知者邮箱地址(多个用,分割)" json:"notify_receiver_email"`
	NotifyKeyword       string    `gorm:"column:notify_keyword;not null;comment:通知匹配关键字(多个用,分割)" json:"notify_keyword"`
	Remark              string    `gorm:"column:remark;not null;comment:备注" json:"remark"`
	IsUsed              int64     `gorm:"column:is_used;not null;default:1;comment:是否启用 1:是  -1:否" json:"is_used"`
	CreatedAt           time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`
	CreatedUser         string    `gorm:"column:created_user;not null;comment:创建人" json:"created_user"`
	UpdatedAt           time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`
	UpdatedUser         string    `gorm:"column:updated_user;not null;comment:更新人" json:"updated_user"`
}

// TableName CronTask's table name
func (*CronTask) TableName() string {
	return TableNameCronTask
}