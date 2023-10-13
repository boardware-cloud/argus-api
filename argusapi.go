package argusapi

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type MonitoringRecord struct {
	CheckedAt    int64            `json:"checkedAt"`
	ResponseTime int64            `json:"responseTime"`
	Result       MonitoringResult `json:"result"`
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
type PingMonitor struct {
	Url      *string `json:"url,omitempty"`
	Interval *int64  `json:"interval,omitempty"`
	Timeout  *int64  `json:"timeout,omitempty"`
	Retries  *int64  `json:"retries,omitempty"`
}
type HttpRequest struct {
	Method  *HttpMethod `json:"method,omitempty"`
	Url     string      `json:"url"`
	Headers *[]Pair     `json:"headers,omitempty"`
}
type MonitorList struct {
	Pagination Pagination `json:"pagination"`
	Data       []Monitor  `json:"data"`
}
type NotificationGroup struct {
	Interval      *int64          `json:"interval,omitempty"`
	Notifications *[]Notification `json:"notifications,omitempty"`
}
type Monitor struct {
	Type              MonitorType        `json:"type"`
	HttpMonitor       *HttpMonitor       `json:"httpMonitor,omitempty"`
	PingMonitor       *PingMonitor       `json:"pingMonitor,omitempty"`
	NotificationGroup *NotificationGroup `json:"notificationGroup,omitempty"`
	Id                string             `json:"id"`
	Name              string             `json:"name"`
	Description       string             `json:"description"`
	Status            MonitorStatus      `json:"status"`
}
type HttpMonitor struct {
	AcceptedStatusCodes *[]string   `json:"acceptedStatusCodes,omitempty"`
	Method              *HttpMethod `json:"method,omitempty"`
	Url                 *string     `json:"url,omitempty"`
	Interval            *int64      `json:"interval,omitempty"`
	Timeout             *int64      `json:"timeout,omitempty"`
	Retries             *int64      `json:"retries,omitempty"`
	Headers             *[]Pair     `json:"headers,omitempty"`
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
type Pair struct {
	Right string `json:"right"`
	Left  string `json:"left"`
}
type Reserved struct {
	Id        string `json:"id"`
	StartAt   int64  `json:"startAt"`
	ExpiredAt int64  `json:"expiredAt"`
}
type Report struct {
	Period         *int64          `json:"period,omitempty"`
	EmailReceivers *EmailReceivers `json:"emailReceivers,omitempty"`
	Cron           *string         `json:"cron,omitempty"`
}
type Notification struct {
	EmailReceivers *EmailReceivers  `json:"emailReceivers,omitempty"`
	Interval       *int64           `json:"interval,omitempty"`
	Type           NotificationType `json:"type"`
}
type PutMonitorRequest struct {
	HttpMonitor       *HttpMonitor       `json:"httpMonitor,omitempty"`
	PingMonitor       *PingMonitor       `json:"pingMonitor,omitempty"`
	NotificationGroup *NotificationGroup `json:"notificationGroup,omitempty"`
	Name              *string            `json:"name,omitempty"`
	Description       *string            `json:"description,omitempty"`
	Type              *MonitorType       `json:"type,omitempty"`
	Status            *MonitorStatus     `json:"status,omitempty"`
}
type ReservedList struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Data       *[]Reserved `json:"data,omitempty"`
}
type MonitorStatus string

const ACTIVED MonitorStatus = "ACTIVED"
const DISACTIVED MonitorStatus = "DISACTIVED"

type BodyForm string

const RAW BodyForm = "RAW"
const X_WWW_FORM_URLENCODED BodyForm = "X_WWW_FORM_URLENCODED"

type NotificationType string

const EMAIL NotificationType = "EMAIL"

type MonitorType string

const HTTP MonitorType = "HTTP"
const PING MonitorType = "PING"

type ContentType string

const TEXT ContentType = "TEXT"
const JSON ContentType = "JSON"
const XML ContentType = "XML"

type Ordering string

const ASCENDING Ordering = "ASCENDING"
const DESCENDING Ordering = "DESCENDING"

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

type MonitorApiInterface interface {
	CreateMonitor(gin_context *gin.Context, gin_body PutMonitorRequest)
	ListMonitors(gin_context *gin.Context, ordering Ordering, index int64, limit int64)
	DeleteMonitor(gin_context *gin.Context, id string)
	GetMonitor(gin_context *gin.Context, id string)
	UpdateMonitor(gin_context *gin.Context, id string, gin_body PutMonitorRequest)
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
func MonitorApiInterfaceMounter(gin_router *gin.Engine, gwg_api_label MonitorApiInterface) {
	gin_router.POST("/monitors", CreateMonitorBuilder(gwg_api_label))
	gin_router.GET("/monitors", ListMonitorsBuilder(gwg_api_label))
	gin_router.DELETE("/monitors/:id", DeleteMonitorBuilder(gwg_api_label))
	gin_router.GET("/monitors/:id", GetMonitorBuilder(gwg_api_label))
	gin_router.PUT("/monitors/:id", UpdateMonitorBuilder(gwg_api_label))
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
