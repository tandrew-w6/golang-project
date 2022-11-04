package new_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	newpb "gateway-service/pb"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	new_response "gateway-service/domain/new/response"

	"google.golang.org/protobuf/encoding/protojson"
)

type handler struct {
	ctx        context.Context
	log        *logrus.Logger
	newClient  newpb.NewServiceClient
	httpClient *http.Client
}

type Handler interface {
	GetListData(*gin.Context)
}

func NewHandler(ctx context.Context, log *logrus.Logger, conn newpb.NewServiceClient, httpClient *http.Client) Handler {
	return &handler{
		ctx:        ctx,
		log:        log,
		newClient:  conn,
		httpClient: httpClient,
	}
}

func (h *handler) GetListData(c *gin.Context) {
	//REQUEST
	var (
		lowPrice  *float64
		highPrice *float64
	)

	if c.Query("lowPrice") != "" {
		lowPriceVal, _ := strconv.ParseFloat(c.Query("lowPrice"), 64)
		lowPrice = &lowPriceVal
	}

	if c.Query("highPrice") != "" {
		highPriceVal, _ := strconv.ParseFloat(c.Query("highPrice"), 64)
		highPrice = &highPriceVal
	}

	in := newpb.GetListDataRequest{
		Search:     c.Query("search"),
		Restaurant: c.Query("restaurant"),
		LowPrice:   lowPrice,
		HighPrice:  highPrice,
	}

	res, err := h.newClient.GetListData(h.ctx, &in)
	if err != nil {
		return
	}

	//RETURN
	var out new_response.GetListDataResponse
	// Convert JSON to Proto with CamelCase
	byt, err := protojson.Marshal(res)
	if err != nil {
		return
	}

	err = json.Unmarshal(byt, &out)
	if err != nil {
		fmt.Println("Error unmarshal Proto to JSON", err)
		return
	}

	response := new_response.Response{
		StatusCode: http.StatusOK,
		Message:    "success",
		Status:     "OK",
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Data:       out,
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
	c.JSON(http.StatusOK, response)
}
