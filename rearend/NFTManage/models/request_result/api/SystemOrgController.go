package apiResult

import "yqnft/NFTManage/models/model"

type OrgTreeVO struct {
	Title    string      `json:"title"` // 唯一标识
	Key      string      `json:"key"`
	Children []OrgTreeVO `json:"children"`
}

func (orgTreeVO OrgTreeVO) OrgTreeConvertOrgTreeVO(orgList []model.TOrg) []OrgTreeVO {
	orgTreeVOS := []OrgTreeVO{}
	for _, org := range orgList {
		treeVO := OrgTreeVO{}
		treeVO.Title = org.Name
		treeVO.Key = org.ID
		if org.Children != nil {
			orgTreeVOS2 := orgTreeVO.OrgTreeConvertOrgTreeVO(org.Children)
			treeVO.Children = orgTreeVOS2
		}
		orgTreeVOS = append(orgTreeVOS, treeVO)
	}
	return orgTreeVOS
}
