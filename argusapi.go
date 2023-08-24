package argusapi

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type MonitoringRecordList struct {
	Data       []MonitoringRecord `json:"data"`
	Pagination Pagination         `json:"pagination"`
}
type Report struct {
	Cron           *string         `json:"cron,omitempty"`
	Period         *int64          `json:"period,omitempty"`
	EmailReceivers *EmailReceivers `json:"emailReceivers,omitempty"`
}
type Pair struct {
	Right string `json:"right"`
	Left  string `json:"left"`
}
type HttpBody struct {
	ContentType *ContentType `json:"contentType,omitempty"`
	Raw         *string      `json:"raw,omitempty"`
	FormPayload *[]Pair      `json:"formPayload,omitempty"`
	Form        BodyForm     `json:"form"`
}
type MonitorList struct {
	Data       []Monitor  `json:"data"`
	Pagination Pagination `json:"pagination"`
}
type EmailReceivers struct {
	To  []string `json:"to"`
	Cc  []string `json:"cc"`
	Bcc []string `json:"bcc"`
}
type Notification struct {
	Type           NotificationType `json:"type"`
	EmailReceivers *EmailReceivers  `json:"emailReceivers,omitempty"`
}
type Pagination struct {
	Total int64 `json:"total"`
	Index int64 `json:"index"`
	Limit int64 `json:"limit"`
}
type Monitor struct {
	Notifications        []Notification `json:"notifications"`
	Description          string         `json:"description"`
	Type                 MonitorType    `json:"type"`
	Interval             int64          `json:"interval"`
	Url                  string         `json:"url"`
	Headers              *[]Pair        `json:"headers,omitempty"`
	Method               *HttpMethod    `json:"method,omitempty"`
	Body                 *HttpBody      `json:"body,omitempty"`
	Timeout              int64          `json:"timeout"`
	Status               MonitorStatus  `json:"status"`
	Retries              int64          `json:"retries"`
	AcceptedStatusCodes  *[]string      `json:"acceptedStatusCodes,omitempty"`
	NotificationInterval int64          `json:"notificationInterval"`
	Id                   string         `json:"id"`
	Name                 string         `json:"name"`
}
type MonitoringRecord struct {
	CheckedAt    int64            `json:"checkedAt"`
	StatusCode   string           `json:"statusCode"`
	ResponseTime *int64           `json:"responseTime,omitempty"`
	Result       MonitoringResult `json:"result"`
	Id           string           `json:"id"`
	MonitorId    string           `json:"monitorId"`
}
type PutMonitorRequest struct {
	Name                 string         `json:"name"`
	Type                 MonitorType    `json:"type"`
	Method               *HttpMethod    `json:"method,omitempty"`
	Headers              *[]Pair        `json:"headers,omitempty"`
	Description          string         `json:"description"`
	Timeout              int64          `json:"timeout"`
	Status               MonitorStatus  `json:"status"`
	Url                  string         `json:"url"`
	AcceptedStatusCodes  *[]string      `json:"acceptedStatusCodes,omitempty"`
	NotificationInterval int64          `json:"notificationInterval"`
	Interval             int64          `json:"interval"`
	Retries              int64          `json:"retries"`
	Notifications        []Notification `json:"notifications"`
	Body                 *HttpBody      `json:"body,omitempty"`
}
type ContentType string

const TEXT ContentType = "TEXT"
const JSON ContentType = "JSON"
const XML ContentType = "XML"

type MonitoringResult string

const OK MonitoringResult = "OK"
const TIMEOUT MonitoringResult = "TIMEOUT"
const DOWN MonitoringResult = "DOWN"

type Ordering string

const ASCENDING Ordering = "ASCENDING"
const DESCENDING Ordering = "DESCENDING"

type MonitorStatus string

const ACTIVED MonitorStatus = "ACTIVED"
const DISACTIVED MonitorStatus = "DISACTIVED"

type BodyForm string

const RAW BodyForm = "RAW"
const X_WWW_FORM_URLENCODED BodyForm = "X_WWW_FORM_URLENCODED"

type MonitorType string

const HTTP MonitorType = "HTTP"

type HttpMethod string

const HEAD HttpMethod = "HEAD"
const GET HttpMethod = "GET"
const POST HttpMethod = "POST"
const PUT HttpMethod = "PUT"
const PATCH HttpMethod = "PATCH"

type NotificationType string

const EMAIL NotificationType = "EMAIL"

type MonitorApiInterface interface {
	ListMonitoringRecords(gin_context *gin.Context, id string, index int64, limit int64, startAt int64, endAt int64)
	CreateMonitor(gin_context *gin.Context, gin_body PutMonitorRequest)
	ListMonitors(gin_context *gin.Context, ordering Ordering, index int64, limit int64)
	GetMonitor(gin_context *gin.Context, id string)
	UpdateMonitor(gin_context *gin.Context, id string, gin_body PutMonitorRequest)
	DeleteMonitor(gin_context *gin.Context, id string)
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
func MonitorApiInterfaceMounter(gin_router *gin.Engine, gwg_api_label MonitorApiInterface) {
	gin_router.GET("/monitors/:id/records", ListMonitoringRecordsBuilder(gwg_api_label))
	gin_router.POST("/monitors", CreateMonitorBuilder(gwg_api_label))
	gin_router.GET("/monitors", ListMonitorsBuilder(gwg_api_label))
	gin_router.GET("/monitors/:id", GetMonitorBuilder(gwg_api_label))
	gin_router.PUT("/monitors/:id", UpdateMonitorBuilder(gwg_api_label))
	gin_router.DELETE("/monitors/:id", DeleteMonitorBuilder(gwg_api_label))
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
