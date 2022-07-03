package model

type OpenDbReq struct {
	Index string `form:"index" label:"index" json:"index" binding:"required"`
}
