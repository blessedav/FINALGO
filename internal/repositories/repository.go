package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/smartnotes/user-service/internal/models"
	"github.com/smartnotes/user-service/pkg/domain"
)

type Repository interface {
	// User methods
	CreateUser(user *models.User) error
	FindUserByEmail(email string) (*models.User, error)
	FindUserByID(id primitive.ObjectID) (*models.User, error)

	// Book methods
	Create(ctx context.Context, book *domain.Book) error
	Update(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID, update *domain.Book) error
	Delete(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) error
	GetByID(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) (*domain.Book, error)
	List(ctx context.Context, userID primitive.ObjectID, search, sortBy, order string, limit, offset int) ([]*domain.Book, error)
	Count(ctx context.Context, userID primitive.ObjectID, search string) (int64, error)
}
