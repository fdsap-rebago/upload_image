package main

type Uploaded_Images struct {
	ImgID   int    `json:"img_id" gorm:"primaryKey"`
	ImgData []byte `json:"img_data"`
	ImgType string `json:"img_type"`
}

type ImageSrc struct {
	ImgSrc string `json:"img_src"`
}

type ImageID struct {
	ImgId int `json:"img_id" example:"1"`
}

type UserManagement struct {
	// RoleID   int    `json:"role_id"`
	// RoleName string `json:"role_name"`
	// RoleDesc string `json:"role_desc"`
	UserName    string `json:"user_name"`
	GivenName   string `json:"given_name"`
	MiddleName  string `json:"middle_name"`
	LastName    string `json:"last_name"`
	BranchNames string `json:"branch_names"`
	Roles       string `json:"roles"`
	CheckStatus string `json:"check_status"`
}
