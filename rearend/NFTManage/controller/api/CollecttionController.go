package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yqnft/NFTManage/config"
	"yqnft/NFTManage/models/model"
	result "yqnft/NFTManage/models/request_result"
	apiResult "yqnft/NFTManage/models/request_result/api"
	"yqnft/NFTManage/service"
)

type NFTController struct {
}

var nftCollectionService service.NftCollectionService

func (con NFTController) FindFairCollectionList(c *gin.Context) {
	//results := nftCollectionService.FindFairMarketList()
	ch := make(chan []apiResult.NFTMarketDTO)
	go nftCollectionService.FindFairMarketList(ch)
	results := <-ch
	close(ch)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": results, "message": "操作成功", "statusCode": 200})
}

func (con NFTController) ColletionDetail(c *gin.Context) {
	nftCollectionId := c.Query("nftCollectionId")
	ch := make(chan apiResult.CollectionDetailVO)
	//CollectionDetailVO := nftCollectionService.NftCollectionDetail(nftCollectionId)
	go nftCollectionService.NftCollectionDetail(ch, nftCollectionId)
	CollectionDetailVO := <-ch
	close(ch)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": CollectionDetailVO, "message": "操作成功", "statusCode": 200})
}

func (con NFTController) FindWorksInCollection(c *gin.Context) {
	nftCollectionId := c.Query("collectionId")
	ch := make(chan []apiResult.NftDetialWorksVO)
	go nftCollectionService.FindWorksInCollection(ch, nftCollectionId)
	nftDetialWorksVOS := <-ch
	close(ch)
	// TODO page开发
	page := result.Page{PageNo: 1, PageSize: 8, PageCount: 99}
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": nftDetialWorksVOS, "page": page, "message": "操作成功", "statusCode": 200})
}

func (con NFTController) FindNftMarketList(c *gin.Context) {
	var collectionName string
	if c.Query("searchText") != "" {
		collectionName = c.Query("searchText")
	}
	typeId := c.Query("typeId")
	ch := make(chan []apiResult.NFTMarketDTO)
	go nftCollectionService.FindNftMarketList(ch, collectionName, typeId)
	nftMarketDTOS := <-ch
	close(ch)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": nftMarketDTOS, "page": nil, "message": "操作成功", "statusCode": 200})
}

func (con NFTController) Update(c *gin.Context) {
	var nftCollectionVO apiResult.NftCollectionVO
	if err := c.ShouldBindJSON(&nftCollectionVO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ch := make(chan model.NftCollection)
	go nftCollectionService.UpdateCreateCollection(ch, nftCollectionVO)
	nftCollection := <-ch
	close(ch)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": nftCollection, "message": "操作成功", "statusCode": 200})
}

func (con NFTController) Create(c *gin.Context) {
	var nftCollectionVO apiResult.NftCollectionVO
	if err := c.ShouldBindJSON(&nftCollectionVO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ch := make(chan model.NftCollection)
	go nftCollectionService.SaveCreateCollection(ch, nftCollectionVO)
	nftCollection := <-ch
	close(ch)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": nftCollection, "message": "操作成功", "statusCode": 200})
}

func (con NFTController) LabelValuePair(c *gin.Context) {
	userId := config.SessionGet("userId").(string)
	ch := make(chan []apiResult.LabelValuePairVO)
	labelValuePairVOs := <-ch
	go nftCollectionService.GetLabelValuePair(ch, userId)
	close(ch)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": labelValuePairVOs, "page": nil, "message": "操作成功", "statusCode": 200})
}

func (con NFTController) FindCollectionsInUser(c *gin.Context) {
	ch := make(chan []apiResult.NFTMarketDTO)
	go nftCollectionService.FindCollectionInUser(ch)
	results := <-ch
	close(ch)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": results, "message": "操作成功", "statusCode": 200})
}
