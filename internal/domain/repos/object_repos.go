package repos

import "context"

type AvatarObjectRepository interface {
	GetUploadURL(context context.Context, objectName string) (string, error)
	GetDownloadURL(context context.Context, objectName string) (string, error)
}
