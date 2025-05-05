package repositories

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/smartnotes/user-service/internal/models"
	"github.com/smartnotes/user-service/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBRepository struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoDBRepository() (*MongoDBRepository, error) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	db := client.Database("smartnotes")
	return &MongoDBRepository{
		client: client,
		db:     db,
	}, nil
}

// User methods
func (r *MongoDBRepository) CreateUser(user *models.User) error {
	collection := r.db.Collection("users")

	// Check if user already exists
	var existingUser models.User
	err := collection.FindOne(context.Background(), bson.M{
		"$or": []bson.M{
			{"email": user.Email},
			{"username": user.Username},
		},
	}).Decode(&existingUser)

	if err == nil {
		return errors.New("user already exists")
	}
	if err != mongo.ErrNoDocuments {
		return err
	}

	user.CreatedAt = time.Now()
	user.Role = "user"

	_, err = collection.InsertOne(context.Background(), user)
	return err
}

func (r *MongoDBRepository) FindUserByEmail(email string) (*models.User, error) {
	collection := r.db.Collection("users")
	var user models.User

	err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *MongoDBRepository) FindUserByID(id primitive.ObjectID) (*models.User, error) {
	collection := r.db.Collection("users")
	var user models.User

	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Book methods
func (r *MongoDBRepository) Create(ctx context.Context, book *domain.Book) error {
	collection := r.db.Collection("books")
	book.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	book.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	_, err := collection.InsertOne(ctx, book)
	return err
}

func (r *MongoDBRepository) Update(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID, update *domain.Book) error {
	collection := r.db.Collection("books")
	update.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	_, err := collection.UpdateOne(ctx, bson.M{
		"_id":     id,
		"user_id": userID,
	}, bson.M{"$set": update})
	return err
}

func (r *MongoDBRepository) Delete(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) error {
	collection := r.db.Collection("books")
	_, err := collection.DeleteOne(ctx, bson.M{
		"_id":     id,
		"user_id": userID,
	})
	return err
}

func (r *MongoDBRepository) GetByID(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) (*domain.Book, error) {
	collection := r.db.Collection("books")
	var book domain.Book
	err := collection.FindOne(ctx, bson.M{
		"_id":     id,
		"user_id": userID,
	}).Decode(&book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *MongoDBRepository) List(ctx context.Context, userID primitive.ObjectID, search, sortBy, order string, limit, offset int) ([]*domain.Book, error) {
	collection := r.db.Collection("books")

	// Build filter
	filter := bson.M{"user_id": userID}
	if search != "" {
		filter["$or"] = []bson.M{
			{"title": bson.M{"$regex": search, "$options": "i"}},
			{"description": bson.M{"$regex": search, "$options": "i"}},
		}
	}

	// Build sort
	sort := 1
	if order == "desc" {
		sort = -1
	}
	opts := options.Find().
		SetSort(bson.D{{Key: sortBy, Value: sort}}).
		SetLimit(int64(limit)).
		SetSkip(int64(offset))

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var books []*domain.Book
	if err = cursor.All(ctx, &books); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *MongoDBRepository) Count(ctx context.Context, userID primitive.ObjectID, search string) (int64, error) {
	collection := r.db.Collection("books")

	filter := bson.M{"user_id": userID}
	if search != "" {
		filter["$or"] = []bson.M{
			{"title": bson.M{"$regex": search, "$options": "i"}},
			{"description": bson.M{"$regex": search, "$options": "i"}},
		}
	}

	return collection.CountDocuments(ctx, filter)
}
