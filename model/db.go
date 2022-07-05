package model

type OpenDbReq struct {
	Index string `form:"index" label:"index" json:"index" binding:"required"`
}

type GetKeysReq struct {
	Index string `form:"index" label:"index" json:"index" binding:"required"`
}

type ShowKeyReq struct {
	Id string `form:"id" label:"id" json:"id" binding:"required"`
	Sk string `form:"sk" label:"sk" json:"sk" binding:"required"`
	Db int    `form:"db" label:"db" json:"db" binding:"required"`
}
