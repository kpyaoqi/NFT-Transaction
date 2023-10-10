package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	apiResult "yqnft/NFTManage/models/request_result/api"
	"yqnft/NFTManage/service"
)

type WorkController struct {
}

var nftWorksService service.NftWorksService
var nftUserAssetsService service.NftUserAssetsService

func (con WorkController) BuyWorks(c *gin.Context) {
	worksId := c.Query("worksId")
	worksService := service.NftWorksService{}
	details := worksService.NftworksBuyDetails(worksId)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": details, "message": "操作成功", "statusCode": 200})
}

func (con WorkController) BuyWorksInfo(c *gin.Context) {
	worksId := c.Query("worksId")
	info := nftWorksService.NftworksInfo(worksId)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": info, "message": "操作成功", "statusCode": 200})
}

func (con WorkController) LogicDel(c *gin.Context) {
	worksId := c.Query("worksId")
	ok := nftWorksService.UpdateLogicDel(worksId)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": ok, "message": "操作成功", "statusCode": 200})
}

func (con WorkController) FindWorks(c *gin.Context) {
	worksId := c.Query("worksId")
	nftWorkVO := nftWorksService.FindWorksById(worksId)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": nftWorkVO, "message": "操作成功", "statusCode": 200})
}

func (con WorkController) UpdateWorks(c *gin.Context) {
	nftWorkVO := apiResult.NftWorkVO{}
	if err := c.ShouldBindJSON(&nftWorkVO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	nftWorkVO.Num, _ = strconv.ParseInt(nftWorkVO.Nums, 10, 64)

	worksService := service.NftWorksService{}
	nftWork := worksService.UpdateWorks(nftWorkVO)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": nftWork, "message": "操作成功", "statusCode": 200})
}

func (con WorkController) CreateWorks(c *gin.Context) {
	nftWorkVO := apiResult.NftWorkVO{}
	if err := c.ShouldBindJSON(&nftWorkVO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	nftWorkVO.Num, _ = strconv.ParseInt(nftWorkVO.Nums, 10, 64)
	worksService := service.NftWorksService{}
	nftWork := worksService.SaveCreateNftWorks(nftWorkVO)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": nftWork, "message": "操作成功", "statusCode": 200})
}

func (con WorkController) HisTrans(c *gin.Context) {
	worksId := c.Query("worksId")
	result := nftWorksService.HisTransForWorks(worksId)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": result, "message": "操作成功", "statusCode": 200})
}

func (con WorkController) FindWorksInUser(c *gin.Context) {
	worksService := service.NftWorksService{}
	//worksService := service.NftWorksService{UserId: "u_10001"}
	nftWork := worksService.FindWorksByUserId()
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": nftWork, "message": "操作成功", "statusCode": 200})
}

func (con WorkController) FindFollowsInUser(c *gin.Context) {

	worksService := service.NftWorksService{}
	//worksService := service.NftWorksService{UserId: "u_10001"}
	nftDetialWorksVOS := worksService.FindFollowsInUser()
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": nftDetialWorksVOS, "message": "操作成功", "statusCode": 200})
}

func (con WorkController) BuyNftWork(c *gin.Context) {
	worksId := c.Query("worksId")
	worksService := service.NftWorksService{}
	result := worksService.BuyNftWork(worksId)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": result, "message": "操作成功", "statusCode": 200})
}
