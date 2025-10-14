package interfaces

import (
    "context"
    "coop/internal/models"
)

type PublicDocRepository interface {
    GetPublicDocByID(ctx context.Context, id int) (models.PublicDocWithDetails, error)
    GetAllPublicDocs(ctx context.Context, limit, offset int) ([]models.PublicDocWithDetails, int, error)
    CreatePublicDoc(ctx context.Context, doc models.PublicDoc) (int, error)
    UpdatePublicDoc(ctx context.Context, id int, doc models.PublicDoc) error
    DeletePublicDoc(ctx context.Context, id int) error
    SearchPublicDocs(ctx context.Context, query string, limit, offset int) ([]models.PublicDocWithDetails, int, error)
}