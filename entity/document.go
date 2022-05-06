package entity

type Document struct {
	DeepLinkCode string `json:"deepLinkCode,omitempty"`
	FileType     string `json:"fileType,omitempty"`
	Id           string `json:"id,omitempty"`
}
