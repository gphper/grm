package model

type OpenDbReq struct {
	Index string `form:"index" label:"index" json:"index" binding:"required"`
}

type GetKeysReq struct {
	Index string `form:"index" label:"index" json:"index" binding:"required"`
}

type KeyReq struct {
	Id string `form:"id" label:"id" json:"id" binding:"required"`
	Sk string `form:"sk" label:"sk" json:"sk" binding:"required"`
	Db int    `form:"db" label:"db" json:"db" binding:"required"`
}

type TtlKeyReq struct {
	Id  string `form:"id" label:"id" json:"id" binding:"required"`
	Sk  string `form:"sk" label:"sk" json:"sk" binding:"required"`
	Db  int    `form:"db" label:"db" json:"db" binding:"required"`
	Ttl string `form:"ttl" label:"ttl" json:"ttl" binding:"required"`
}

type ListKeyReq struct {
	Id    string `form:"id" label:"id" json:"id" binding:"required"`
	Sk    string `form:"sk" label:"sk" json:"sk" binding:"required"`
	Db    int    `form:"db" label:"db" json:"db" binding:"required"`
	Page  int    `form:"page" label:"page" json:"page" binding:"required"`
	Limit int    `form:"limit" label:"limit" json:"limit" binding:"required"`
}

type ListItemKeyReq struct {
	Id   string `form:"id" label:"id" json:"id" binding:"required"`
	Sk   string `form:"sk" label:"sk" json:"sk" binding:"required"`
	Db   int    `form:"db" label:"db" json:"db" binding:"required"`
	Item string `form:"item" label:"item" json:"item" binding:"required"`
}

type HashKeyReq struct {
	Id    string `form:"id" label:"id" json:"id" binding:"required"`
	Sk    string `form:"sk" label:"sk" json:"sk" binding:"required"`
	Db    int    `form:"db" label:"db" json:"db" binding:"required"`
	Page  int    `form:"page" label:"page" json:"page" binding:"required"`
	Limit int    `form:"limit" label:"limit" json:"limit" binding:"required"`
}

type HashItemKeyReq struct {
	Id   string `form:"id" label:"id" json:"id" binding:"required"`
	Sk   string `form:"sk" label:"sk" json:"sk" binding:"required"`
	Db   int    `form:"db" label:"db" json:"db" binding:"required"`
	Item string `form:"item" label:"item" json:"item" binding:"required"`
}

type AddHashItemKeyReq struct {
	Id    string `form:"id" label:"id" json:"id" binding:"required"`
	Sk    string `form:"sk" label:"sk" json:"sk" binding:"required"`
	Db    int    `form:"db" label:"db" json:"db" binding:"required"`
	Itemk string `form:"itemk" label:"itemk" json:"itemk" binding:"required"`
	Itemv string `form:"itemv" label:"itemv" json:"itemv" binding:"required"`
}

// set 类型
type SetKeyReq struct {
	Id    string `form:"id" label:"id" json:"id" binding:"required"`
	Sk    string `form:"sk" label:"sk" json:"sk" binding:"required"`
	Db    int    `form:"db" label:"db" json:"db" binding:"required"`
	Page  int    `form:"page" label:"page" json:"page" binding:"required"`
	Limit int    `form:"limit" label:"limit" json:"limit" binding:"required"`
}

type SetItemKeyReq struct {
	Id   string `form:"id" label:"id" json:"id" binding:"required"`
	Sk   string `form:"sk" label:"sk" json:"sk" binding:"required"`
	Db   int    `form:"db" label:"db" json:"db" binding:"required"`
	Item string `form:"item" label:"item" json:"item" binding:"required"`
}
