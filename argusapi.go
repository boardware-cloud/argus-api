package argusapi

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type MonitoringRecordList struct {
	Data       []MonitoringRecord `json:"data"`
	Pagination Pagination         `json:"pagination"`
}
type Pair struct {
	Left  string `json:"left"`
	Right string `json:"right"`
}
type MonitorProperties struct {
	Method               *HttpMethod    `json:"method,omitempty"`
	Notifications        []Notification `json:"notifications"`
	Name                 string         `json:"name"`
	Description          string         `json:"description"`
	Type                 MonitorType    `json:"type"`
	Url                  string         `json:"url"`
	Status               MonitorStatus  `json:"status"`
	AcceptedStatusCodes  *[]Pair        `json:"acceptedStatusCodes,omitempty"`
	NotificationInterval int64          `json:"notificationInterval"`
	Headers              []Pair         `json:"headers"`
	Interval             int64          `json:"interval"`
	Timeout              int64          `json:"timeout"`
	Retries              int64          `json:"retries"`
	Reports              []Report       `json:"reports"`
}
type PutMonitorRequest struct {
	Timeout              int64          `json:"timeout"`
	AcceptedStatusCodes  *[]Pair        `json:"acceptedStatusCodes,omitempty"`
	NotificationInterval int64          `json:"notificationInterval"`
	Body                 *HttpBody      `json:"body,omitempty"`
	Type                 MonitorType    `json:"type"`
	Interval             int64          `json:"interval"`
	Headers              *[]Pair        `json:"headers,omitempty"`
	Url                  string         `json:"url"`
	Retries              int64          `json:"retries"`
	Notifications        []Notification `json:"notifications"`
	Name                 string         `json:"name"`
	Method               *HttpMethod    `json:"method,omitempty"`
	Description          string         `json:"description"`
	Status               MonitorStatus  `json:"status"`
}
type HttpBody struct {
	Form        BodyForm     `json:"form"`
	ContentType *ContentType `json:"contentType,omitempty"`
	Raw         *string      `json:"raw,omitempty"`
	FormPayload *[]Pair      `json:"formPayload,omitempty"`
}
type Report struct {
	Cron           *string         `json:"cron,omitempty"`
	Period         *int64          `json:"period,omitempty"`
	EmailReceivers *EmailReceivers `json:"emailReceivers,omitempty"`
}
type Monitor struct {
	Type                 MonitorType    `json:"type"`
	Retries              int64          `json:"retries"`
	Id                   string         `json:"id"`
	Notifications        []Notification `json:"notifications"`
	AcceptedStatusCodes  *[]Pair        `json:"acceptedStatusCodes,omitempty"`
	Interval             int64          `json:"interval"`
	Url                  string         `json:"url"`
	Method               *HttpMethod    `json:"method,omitempty"`
	NotificationInterval int64          `json:"notificationInterval"`
	Name                 string         `json:"name"`
	Timeout              int64          `json:"timeout"`
	Status               MonitorStatus  `json:"status"`
	Headers              *[]Pair        `json:"headers,omitempty"`
	Body                 *HttpBody      `json:"body,omitempty"`
	Description          string         `json:"description"`
}
type MonitorList struct {
	Data       []Monitor  `json:"data"`
	Pagination Pagination `json:"pagination"`
}
type EmailReceivers struct {
	Cc  []string `json:"cc"`
	Bcc []string `json:"bcc"`
	To  []string `json:"to"`
}
type Pagination struct {
	Total int64 `json:"total"`
	Index int64 `json:"index"`
	Limit int64 `json:"limit"`
}
type MonitoringRecord struct {
	Result       MonitoringResult `json:"result"`
	Id           string           `json:"id"`
	MonitorId    string           `json:"monitorId"`
	CheckedAt    int64            `json:"checkedAt"`
	StatusCode   string           `json:"statusCode"`
	ResponseTime *int64           `json:"responseTime,omitempty"`
}
type Notification struct {
	Type           NotificationType `json:"type"`
	EmailReceivers *EmailReceivers  `json:"emailReceivers,omitempty"`
}
type BodyForm string

const RAW BodyForm = "RAW"
const X_WWW_FORM_URLENCODED BodyForm = "X_WWW_FORM_URLENCODED"

type MonitorStatus string

const ACTIVED MonitorStatus = "ACTIVED"
const DISACTIVED MonitorStatus = "DISACTIVED"

type MonitoringResult string

const OK MonitoringResult = "OK"
const TIMEOUT MonitoringResult = "TIMEOUT"
const DOWN MonitoringResult = "DOWN"

type HttpMethod string

const HEAD HttpMethod = "HEAD"
const GET HttpMethod = "GET"
const POST HttpMethod = "POST"
const PUT HttpMethod = "PUT"
const PATCH HttpMethod = "PATCH"

type NotificationType string

const EMAIL NotificationType = "EMAIL"

type ContentType string

const TEXT ContentType = "TEXT"
const JSON ContentType = "JSON"
const XML ContentType = "XML"

type MonitorType string

const HTTP MonitorType = "HTTP"

type Ordering string

const ASCENDING Ordering = "ASCENDING"
const DESCENDING Ordering = "DESCENDING"

type MonitorApiInterface interface {
	DeleteMonitor(gin_context *gin.Context, id string)
	GetMonitor(gin_context *gin.Context, id string)
	UpdateMonitor(gin_context *gin.Context, id string, gin_body PutMonitorRequest)
	ListMonitoringRecords(gin_context *gin.Context, id string, index int64, limit int64, startAt int64, endAt int64)
	CreateMonitor(gin_context *gin.Context, gin_body PutMonitorRequest)
	ListMonitors(gin_context *gin.Context, ordering Ordering, index int64, limit int64)
}

func DeleteMonitorBuilder(api MonitorApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		id := gin_context.Param("id")
		api.DeleteMonitor(gin_context, id)
	}
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
	gin_router.DELETE("/monitors/:id", DeleteMonitorBuilder(gwg_api_label))
	gin_router.GET("/monitors/:id", GetMonitorBuilder(gwg_api_label))
	gin_router.PUT("/monitors/:id", UpdateMonitorBuilder(gwg_api_label))
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
