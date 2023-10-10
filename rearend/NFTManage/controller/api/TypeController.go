package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yqnft/NFTManage/service"
)

type TypeController struct {
}

var nftTypeService service.NftTypeService

func (con TypeController) List(c *gin.Context) {
	labelValuePairVO2s := nftTypeService.GetLabelValuePair()
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": labelValuePairVO2s, "message": "操作成功", "statusCode": 200})
}
