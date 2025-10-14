package interfaces

import (
    "context"
    "coop/internal/models"
    "mime/multipart"
)

type PublicDocService interface {
    GetPublicDocByID(ctx context.Context, id int) (models.PublicDocWithDetails, error)
    GetAllPublicDocs(ctx context.Context, page, pageSize int) (models.PublicDocListResponse, error)
    CreatePublicDoc(ctx context.Context, req models.CreatePublicDocRequest, file *multipart.FileHeader) (models.PublicDocWithDetails, error)
    UpdatePublicDoc(ctx context.Context, id int, req models.UpdatePublicDocRequest) (models.PublicDocWithDetails, error)
    DeletePublicDoc(ctx context.Context, id int) error
    SearchPublicDocs(ctx context.Context, query string, page, pageSize int) (models.PublicDocListResponse, error)
    DownloadPublicDoc(ctx context.Context, id int) (string, error)
}