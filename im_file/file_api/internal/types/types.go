// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type FileRequest struct {
}

type FileResponse struct {
	Src []string `json:"src"`
}

type ImagePreviewRequest struct {
	ImageType string `path:"imageType"`
	ImageName string `path:"imageName"`
}

type ImageRequest struct {
}

type ImageResponse struct {
	Url string `json:"url"`
}
