package base

// 下载请求
type Request struct {
	// 下载链接
	URL string `json:"url"`
	// 附加信息
	Extra interface{} `json:"extra"`
}

// 资源信息
type Resource struct {
	Req *Request `json:"req"`
	// 资源总大小
	TotalSize int64 `json:"total_size"`
	// 是否支持断点下载
	Range bool `json:"range"`
	// 资源所包含的文件列表
	Files []*File `json:"files"`
}

type File struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Size int64  `json:"size"`
}

// 下载选项
type Options struct {
	// 保存文件名
	Name string `json:"name"`
	// 保存目录
	Path string `json:"path"`
	// 并发连接数
	Connections int `json:"connections"`
}