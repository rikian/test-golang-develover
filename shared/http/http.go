package http

type HttpRequest interface {
	CreateDirectoryImage(id string) (*HttpResponse, error)
	DeleteImage(userId, productId string) (*HttpResponse, error)
}
