package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"yqnft/NFTManage/config"
	"yqnft/NFTManage/models"
	"yqnft/NFTManage/models/model"
	"yqnft/NFTManage/service"
	"yqnft/NFTManage/utils"
)

type AttachmentController struct {
}

var staticdir string = "./static"
var userRoleOrgService service.UserRoleOrgService

func (con AttachmentController) Upload(c *gin.Context) {
	userId := config.SessionGet("userId").(string)
	query := c.Query("type")
	file, _ := c.FormFile("file")

	folder := "/attachmentFile/"
	filePath := staticdir + folder
	os.MkdirAll(filePath, os.ModePerm)
	originalFilename := file.Filename // 请替换为实际的文件名
	// 获取文件后缀
	var suffix string
	dotIndex := strings.LastIndex(file.Filename, ".")
	if dotIndex != -1 {
		suffix = originalFilename[dotIndex:]
	}
	// 生成唯一的文件名
	now := time.Now().Format("20060102150405")
	uuid := uuid.New().String()
	fileUrl := now + uuid + suffix
	upLoadPath := staticdir + folder + fileUrl
	size := file.Size
	c.SaveUploadedFile(file, upLoadPath)
	attachment := model.TAttachment{}
	orgIdByUserId := userRoleOrgService.FindOrgIdByUserId(userId)
	if orgIdByUserId != "" {
		attachment.OrgID = orgIdByUserId
	}
	attachment.CreateUser = userId
	attachment.CreateTime = time.Now()
	attachment.ID = utils.GetTimeNO()
	attachment.FileName = originalFilename
	attachment.FileURL = folder + fileUrl
	attachment.FileSize = size / 1024
	attachment.AttachmentType, _ = strconv.Atoi(query)
	attachment.Suffix = suffix
	attachment.Sortno = 0
	attachment.Active = 1
	attachment.Md5Code, _ = utils.CalculateMD5(staticdir + attachment.FileURL)
	models.DB.Create(&attachment)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": attachment, "message": "操作成功", "statusCode": 200})
}
