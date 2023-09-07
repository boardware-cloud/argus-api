package argusapi

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Monitor struct {
	Status               MonitorStatus  `json:"status"`
	Url                  string         `json:"url"`
	Body                 *HttpBody      `json:"body,omitempty"`
	AcceptedStatusCodes  *[]string      `json:"acceptedStatusCodes,omitempty"`
	NotificationInterval int64          `json:"notificationInterval"`
	Id                   string         `json:"id"`
	Description          string         `json:"description"`
	Type                 MonitorType    `json:"type"`
	Name                 string         `json:"name"`
	Retries              int64          `json:"retries"`
	Headers              *[]Pair        `json:"headers,omitempty"`
	Notifications        []Notification `json:"notifications"`
	Interval             int64          `json:"interval"`
	Timeout              int64          `json:"timeout"`
	Method               *HttpMethod    `json:"method,omitempty"`
}
type HttpBody struct {
	FormPayload *[]Pair      `json:"formPayload,omitempty"`
	Form        BodyForm     `json:"form"`
	ContentType *ContentType `json:"contentType,omitempty"`
	Raw         *string      `json:"raw,omitempty"`
}
type Report struct {
	Cron           *string         `json:"cron,omitempty"`
	Period         *int64          `json:"period,omitempty"`
	EmailReceivers *EmailReceivers `json:"emailReceivers,omitempty"`
}
type Pagination struct {
	Index int64 `json:"index"`
	Limit int64 `json:"limit"`
	Total int64 `json:"total"`
}
type CreateReservedRequest struct {
	AccountId *string `json:"accountId,omitempty"`
	StartAt   *int64  `json:"startAt,omitempty"`
	ExpiredAt *int64  `json:"expiredAt,omitempty"`
}
type ReservedList struct {
	Data       *[]Reserved `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}
type HttpRequest struct {
	Body    *HttpBody   `json:"body,omitempty"`
	Url     string      `json:"url"`
	Headers *[]Pair     `json:"headers,omitempty"`
	Method  *HttpMethod `json:"method,omitempty"`
}
type MonitoringRecord struct {
	Id           string           `json:"id"`
	MonitorId    string           `json:"monitorId"`
	CheckedAt    int64            `json:"checkedAt"`
	StatusCode   string           `json:"statusCode"`
	ResponseTime *int64           `json:"responseTime,omitempty"`
	Result       MonitoringResult `json:"result"`
}
type MonitoringRecordList struct {
	Data       []MonitoringRecord `json:"data"`
	Pagination Pagination         `json:"pagination"`
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
type Pair struct {
	Left  string `json:"left"`
	Right string `json:"right"`
}
type Reserved struct {
	ExpiredAt int64  `json:"expiredAt"`
	Id        string `json:"id"`
	StartAt   int64  `json:"startAt"`
}
type MonitorList struct {
	Data       []Monitor  `json:"data"`
	Pagination Pagination `json:"pagination"`
}
type PutMonitorRequest struct {
	Retries              int64          `json:"retries"`
	Notifications        []Notification `json:"notifications"`
	Type                 MonitorType    `json:"type"`
	Url                  string         `json:"url"`
	Method               *HttpMethod    `json:"method,omitempty"`
	NotificationInterval int64          `json:"notificationInterval"`
	Name                 string         `json:"name"`
	Interval             int64          `json:"interval"`
	Timeout              int64          `json:"timeout"`
	Body                 *HttpBody      `json:"body,omitempty"`
	Description          string         `json:"description"`
	Status               MonitorStatus  `json:"status"`
	Headers              *[]Pair        `json:"headers,omitempty"`
	AcceptedStatusCodes  *[]string      `json:"acceptedStatusCodes,omitempty"`
}
type ContentType string

const TEXT ContentType = "TEXT"
const JSON ContentType = "JSON"
const XML ContentType = "XML"

type MonitorStatus string

const ACTIVED MonitorStatus = "ACTIVED"
const DISACTIVED MonitorStatus = "DISACTIVED"

type MonitoringResult string

const OK MonitoringResult = "OK"
const TIMEOUT MonitoringResult = "TIMEOUT"
const DOWN MonitoringResult = "DOWN"

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

type BodyForm string

const RAW BodyForm = "RAW"
const X_WWW_FORM_URLENCODED BodyForm = "X_WWW_FORM_URLENCODED"

type Ordering string

const ASCENDING Ordering = "ASCENDING"
const DESCENDING Ordering = "DESCENDING"

type MonitorApiInterface interface {
	CreateMonitor(gin_context *gin.Context, gin_body PutMonitorRequest)
	ListMonitors(gin_context *gin.Context, ordering Ordering, index int64, limit int64)
	GetMonitor(gin_context *gin.Context, id string)
	UpdateMonitor(gin_context *gin.Context, id string, gin_body PutMonitorRequest)
	DeleteMonitor(gin_context *gin.Context, id string)
	ListMonitoringRecords(gin_context *gin.Context, id string, index int64, limit int64, startAt int64, endAt int64)
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
func MonitorApiInterfaceMounter(gin_router *gin.Engine, gwg_api_label MonitorApiInterface) {
	gin_router.POST("/monitors", CreateMonitorBuilder(gwg_api_label))
	gin_router.GET("/monitors", ListMonitorsBuilder(gwg_api_label))
	gin_router.GET("/monitors/:id", GetMonitorBuilder(gwg_api_label))
	gin_router.PUT("/monitors/:id", UpdateMonitorBuilder(gwg_api_label))
	gin_router.DELETE("/monitors/:id", DeleteMonitorBuilder(gwg_api_label))
	gin_router.GET("/monitors/:id/records", ListMonitoringRecordsBuilder(gwg_api_label))
}

type ReservedApiInterface interface {
	ListReserved(gin_context *gin.Context)
	CreateReservedMonitor(gin_context *gin.Context, gin_body CreateReservedRequest)
}

func ListReservedBuilder(api ReservedApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		api.ListReserved(gin_context)
	}
}
func CreateReservedMonitorBuilder(api ReservedApiInterface) func(c *gin.Context) {
	return func(gin_context *gin.Context) {
		var createReservedRequest CreateReservedRequest
		if err := gin_context.ShouldBindJSON(&createReservedRequest); err != nil {
			gin_context.JSON(400, gin.H{})
			return
		}
		api.CreateReservedMonitor(gin_context, createReservedRequest)
	}
}
func ReservedApiInterfaceMounter(gin_router *gin.Engine, gwg_api_label ReservedApiInterface) {
	gin_router.GET("/reserved", ListReservedBuilder(gwg_api_label))
	gin_router.POST("/reserved", CreateReservedMonitorBuilder(gwg_api_label))
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
