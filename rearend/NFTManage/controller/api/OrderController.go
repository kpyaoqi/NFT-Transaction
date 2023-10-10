package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	apiResult "yqnft/NFTManage/models/request_result/api"
	"yqnft/NFTManage/service"
	"yqnft/NFTManage/utils"

	"net/http"
)

type OrderController struct {
}

var nftOrderService service.NftOrderService
var inShelf apiResult.InShelfVO

func (con OrderController) GenerateOrderWorksIn(c *gin.Context) {
	//var inShelfVO apiResult.InShelfVO
	var params map[string]interface{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	inShelf = utils.ToInShelfVO(params)

	orderService := service.NftOrderService{}
	result := orderService.SaveGenerateOrderWorksIn(inShelf)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": result, "message": "操作成功", "statusCode": 200})
}

func (con OrderController) CheckOrderState(c *gin.Context) {
	orderId := c.Query("orderId")
	marshal, _ := json.Marshal(inShelf)
	state := nftOrderService.UpdateFindOrderState(orderId, marshal)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": state, "message": "操作成功", "statusCode": 200})
}
