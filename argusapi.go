package argusapi

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Notification struct {
	EmailReceivers *EmailReceivers  `json:"emailReceivers,omitempty"`
	Type           NotificationType `json:"type" binding:"required"`
}
type Pagination struct {
	Index int64 `json:"index" binding:"required"`
	Limit int64 `json:"limit" binding:"required"`
	Total int64 `json:"total" binding:"required"`
}
type Monitor struct {
	Notifications        []Notification `json:"notifications" binding:"required"`
	Description          string         `json:"description" binding:"required"`
	Type                 MonitorType    `json:"type" binding:"required"`
	Interval             int64          `json:"interval" binding:"required"`
	Headers              *[]Pair        `json:"headers,omitempty"`
	Retries              int64          `json:"retries" binding:"required"`
	Method               *HttpMethod    `json:"method,omitempty"`
	AcceptedStatusCodes  *[]Pair        `json:"acceptedStatusCodes,omitempty"`
	NotificationInterval int64          `json:"notificationInterval" binding:"required"`
	Name                 string         `json:"name" binding:"required"`
	Timeout              int64          `json:"timeout" binding:"required"`
	Status               MonitorStatus  `json:"status" binding:"required"`
	Url                  string         `json:"url" binding:"required"`
}
type MonitoringRecordList struct {
	Pagination Pagination         `json:"pagination" binding:"required"`
	Data       []MonitoringRecord `json:"data" binding:"required"`
}
type EmailReceivers struct {
	To  []string `json:"to" binding:"required"`
	Cc  []string `json:"cc" binding:"required"`
	Bcc []string `json:"bcc" binding:"required"`
}
type Pair struct {
	Right string `json:"right" binding:"required"`
	Left  string `json:"left" binding:"required"`
}
type MonitorProperties struct {
	Name                 string         `json:"name" binding:"required"`
	Method               *HttpMethod    `json:"method,omitempty"`
	Headers              []Pair         `json:"headers" binding:"required"`
	AcceptedStatusCodes  *[]Pair        `json:"acceptedStatusCodes,omitempty"`
	Notifications        []Notification `json:"notifications" binding:"required"`
	Reports              []Report       `json:"reports" binding:"required"`
	Timeout              int64          `json:"timeout" binding:"required"`
	Status               MonitorStatus  `json:"status" binding:"required"`
	Description          string         `json:"description" binding:"required"`
	Type                 MonitorType    `json:"type" binding:"required"`
	Interval             int64          `json:"interval" binding:"required"`
	Retries              int64          `json:"retries" binding:"required"`
	Url                  string         `json:"url" binding:"required"`
	NotificationInterval int64          `json:"notificationInterval" binding:"required"`
}
type MonitorList struct {
	Data       []Monitor  `json:"data" binding:"required"`
	Pagination Pagination `json:"pagination" binding:"required"`
}
type PutMonitorRequest struct {
	Type                 MonitorType    `json:"type" binding:"required"`
	Interval             int64          `json:"interval" binding:"required"`
	Url                  string         `json:"url" binding:"required"`
	Retries              int64          `json:"retries" binding:"required"`
	Method               *HttpMethod    `json:"method,omitempty"`
	AcceptedStatusCodes  *[]Pair        `json:"acceptedStatusCodes,omitempty"`
	Notifications        []Notification `json:"notifications" binding:"required"`
	Name                 string         `json:"name" binding:"required"`
	Description          string         `json:"description" binding:"required"`
	Timeout              int64          `json:"timeout" binding:"required"`
	Status               MonitorStatus  `json:"status" binding:"required"`
	Headers              *[]Pair        `json:"headers,omitempty"`
	NotificationInterval int64          `json:"notificationInterval" binding:"required"`
}
type MonitoringRecord struct {
	StatusCode   string           `json:"statusCode" binding:"required"`
	ResponseTime *int64           `json:"responseTime,omitempty"`
	Result       MonitoringResult `json:"result" binding:"required"`
	Id           string           `json:"id" binding:"required"`
	MonitorId    string           `json:"monitorId" binding:"required"`
	CheckedAt    int64            `json:"checkedAt" binding:"required"`
}
type Report struct {
	Cron           *string         `json:"cron,omitempty"`
	Period         *int64          `json:"period,omitempty"`
	EmailReceivers *EmailReceivers `json:"emailReceivers,omitempty"`
}
type NotificationType string

const EMAIL NotificationType = "EMAIL"

type MonitorStatus string

const ACTIVED MonitorStatus = "ACTIVED"
const DISACTIVED MonitorStatus = "DISACTIVED"

type HttpMethod string

const HEAD HttpMethod = "HEAD"
const GET HttpMethod = "GET"
const POST HttpMethod = "POST"
const PUT HttpMethod = "PUT"
const PATCH HttpMethod = "PATCH"

type MonitoringResult string

const OK MonitoringResult = "OK"
const TIMEOUT MonitoringResult = "TIMEOUT"
const DOWN MonitoringResult = "DOWN"

type MonitorType string

const HTTP MonitorType = "HTTP"

type Ordering string

const ASCENDING Ordering = "ASCENDING"
const DESCENDING Ordering = "DESCENDING"

type MonitorApiInterface interface {
	GetMonitor(gin_context *gin.Context, id string)
	UpdateMonitor(gin_context *gin.Context, id string, gin_body PutMonitorRequest)
	DeleteMonitor(gin_context *gin.Context, id string)
	ListMonitoringRecords(gin_context *gin.Context, id string, index int64, limit int64, startAt int64, endAt int64)
	CreateMonitor(gin_context *gin.Context, gin_body PutMonitorRequest)
	ListMonitors(gin_context *gin.Context, ordering Ordering, index int64, limit int64)
}

func GetMonitorBuilder(api MonitorApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		api.GetMonitor(gin_context, id)
	}
}
func UpdateMonitorBuilder(api MonitorApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		var putMonitorRequest PutMonitorRequest
		if err := gin_context.ShouldBindJSON(&putMonitorRequest); err != nil {
			gin_context.JSON(400, gin.H{})
			return
		}
		api.UpdateMonitor(gin_context, id, putMonitorRequest)
	}
}
func DeleteMonitorBuilder(api MonitorApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		api.DeleteMonitor(gin_context, id)
	}
}
func ListMonitoringRecordsBuilder(api MonitorApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		index := gin_context.Query("index")
		limit := gin_context.Query("limit")
		startAt := gin_context.Query("startAt")
		endAt := gin_context.Query("endAt")
		api.ListMonitoringRecords(gin_context, id, stringToInt64(index), stringToInt64(limit), stringToInt64(startAt), stringToInt64(endAt))
	}
}
func CreateMonitorBuilder(api MonitorApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		var putMonitorRequest PutMonitorRequest
		if err := gin_context.ShouldBindJSON(&putMonitorRequest); err != nil {
			gin_context.JSON(400, gin.H{})
			return
		}
		api.CreateMonitor(gin_context, putMonitorRequest)
	}
}
func ListMonitorsBuilder(api MonitorApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		ordering := gin_context.Query("ordering")
		index := gin_context.Query("index")
		limit := gin_context.Query("limit")
		api.ListMonitors(gin_context, Ordering(ordering), stringToInt64(index), stringToInt64(limit))
	}
}
func MonitorApiInterfaceMounter(gin_router *gin.Engine, gwg_api_label MonitorApiInterface) {
	gin_router.GET("/monitors/:id", GetMonitorBuilder(gwg_api_label))
	gin_router.PUT("/monitors/:id", UpdateMonitorBuilder(gwg_api_label))
	gin_router.DELETE("/monitors/:id", DeleteMonitorBuilder(gwg_api_label))
	gin_router.GET("/monitors/:id/records", ListMonitoringRecordsBuilder(gwg_api_label))
	gin_router.POST("/monitors", CreateMonitorBuilder(gwg_api_label))
	gin_router.GET("/monitors", ListMonitorsBuilder(gwg_api_label))
}
func stringToInt32(s string) int32 {
	if value, err := strconv.ParseInt(s, 10, 32); err == nil {
		return int32(value)
	}
	return 0
}
func stringToInt64(s string) int64 {
	if value, err := strconv.ParseInt(s, 10, 64); err == nil {
		return value
	}
	return 0
}
func stringToFloat32(s string) float32 {
	if value, err := strconv.ParseFloat(s, 32); err == nil {
		return float32(value)
	}
	return 0
}
func stringToFloat64(s string) float64 {
	if value, err := strconv.ParseFloat(s, 64); err == nil {
		return value
	}
	return 0
}
