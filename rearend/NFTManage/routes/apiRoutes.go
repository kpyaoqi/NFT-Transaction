package routes

import (
	"github.com/gin-gonic/gin"
	"yqnft/NFTManage/controller/api"
)

func ApiRoutesInit(route *gin.Engine) {
	apiNftCollectionRoute := route.Group("/nft/api/nft/collection")
	apiNftCollectionRoute.POST("/findFairCollectionList", api.NFTController{}.FindFairCollectionList)
	apiNftCollectionRoute.POST("/colletionDetail", api.NFTController{}.ColletionDetail)
	apiNftCollectionRoute.POST("/findWorksInCollection", api.NFTController{}.FindWorksInCollection)
	apiNftCollectionRoute.POST("/findNftMarketList", api.NFTController{}.FindNftMarketList)
	apiNftCollectionRoute.POST("/update", api.NFTController{}.Update)
	apiNftCollectionRoute.POST("/create", api.NFTController{}.Create)
	apiNftCollectionRoute.POST("/labelValuePair", api.NFTController{}.LabelValuePair)
	apiNftCollectionRoute.POST("/findCollectionsInUser", api.NFTController{}.FindCollectionsInUser)

	apiNftWorksRoute := route.Group("/nft/api/nft/works")
	apiNftWorksRoute.POST("/buyWorks", api.WorkController{}.BuyWorks)
	apiNftWorksRoute.POST("/buyWorksInfo", api.WorkController{}.BuyWorksInfo)
	apiNftWorksRoute.POST("/logicDel", api.WorkController{}.LogicDel)
	apiNftWorksRoute.POST("/findWorks", api.WorkController{}.FindWorks)
	apiNftWorksRoute.POST("/updateWorks", api.WorkController{}.UpdateWorks)
	apiNftWorksRoute.POST("/createWorks", api.WorkController{}.CreateWorks)
	apiNftWorksRoute.POST("/hisTrans", api.WorkController{}.HisTrans)
	apiNftWorksRoute.POST("/findWorksInUser", api.WorkController{}.FindWorksInUser)
	apiNftWorksRoute.POST("/findFollowsInUser", api.WorkController{}.FindFollowsInUser)
	apiNftWorksRoute.POST("/buyNftWork", api.WorkController{}.BuyNftWork)

	apiNftOrderRoute := route.Group("/nft/api/nft/order")
	apiNftOrderRoute.POST("/generateOrderWorksIn", api.OrderController{}.GenerateOrderWorksIn)
	apiNftOrderRoute.POST("/checkOrderState", api.OrderController{}.CheckOrderState)

	apiNftSystemUserRoute := route.Group("/nft/api/system/user")
	apiNftSystemUserRoute.POST("/menu", api.SystemUserController{}.Menu)
	apiNftSystemUserRoute.POST("/save", api.SystemUserController{}.Save)
	apiNftSystemUserRoute.POST("/getInfo", api.SystemUserController{}.GetInfo)
	apiNftSystemUserRoute.POST("/info", api.SystemUserController{}.Info)
	apiNftSystemUserRoute.POST("/userTypeList", api.SystemUserController{}.UserTypeList)
	apiNftSystemUserRoute.POST("/fllowWorks", api.SystemUserController{}.FllowWorks)
	apiNftSystemUserRoute.POST("/hisTranInUser", api.SystemUserController{}.HisTranInUser)
	apiNftSystemUserRoute.POST("/followStatus", api.SystemUserController{}.FollowStatus)
	apiNftSystemUserRoute.POST("/list", api.SystemUserController{}.List)

	apiNftSystemRoleRoute := route.Group("/nft/api/system/role")
	apiNftSystemRoleRoute.POST("/list", api.SystemRoleController{}.List)
	apiNftSystemRoleRoute.POST("/pageList", api.SystemRoleController{}.PageList)
	apiNftSystemRoleRoute.POST("/getMenusByRoleId", api.SystemRoleController{}.GetMenusByRoleId)

	apiNftSystemOrgRoute := route.Group("/nft/api/system/org")
	apiNftSystemOrgRoute.POST("/treeselect", api.SystemOrgController{}.Treeselect)
	apiNftSystemOrgRoute.POST("/pageList", api.SystemOrgController{}.PageList)
	apiNftSystemOrgRoute.POST("/orgTypeList", api.SystemOrgController{}.OrgTypeList)

	apiNftSystemMenuRoute := route.Group("/nft/api/system/menu")
	apiNftSystemMenuRoute.POST("/all/list/json", api.SystemMenuController{}.AllListJson)
	apiNftSystemMenuRoute.POST("/list", api.SystemMenuController{}.List)

	apiNftSystemNotifyRoute := route.Group("/nft/api/system/notify")
	apiNftSystemNotifyRoute.POST("/message", api.SystemNotifyController{}.Message)

	apiNftTypeRoute := route.Group("/nft/api/nft/type")
	apiNftTypeRoute.POST("/list", api.TypeController{}.List)

	apiNftAttachmentRoute := route.Group("/nft/attachment")
	apiNftAttachmentRoute.POST("/upload", api.AttachmentController{}.Upload)
}
