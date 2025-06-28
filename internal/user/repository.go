package user

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	db "github.com/SitaGomes/coins-exchange/internal/services/firebase"
)

type UserRepositoryInterface interface {
	GetAllUsers(ctx context.Context) ([]*User, error)
	GetUser(ctx context.Context) *User
	AddUser(ctx context.Context, user *User) error
}

type UserRepository struct {
	client *firestore.Client
}

func NewUserRepository() *UserRepository {
	firebase := db.GetInstance()
	return &UserRepository{
		client: firebase.GetFirestoreClient(),
	}
}

func (r *UserRepository) collection() *firestore.CollectionRef {
	return r.client.Collection("users")
}

func (r *UserRepository) toModel(snap *firestore.DocumentSnapshot) (*User, error) {
	if !snap.Exists() {
		return nil, errors.New("user document doesn't exist")
	}

	user := User{}
	if err := snap.DataTo(&user); err != nil {
		return nil, errors.New("failed to convert user document")
	}

	if user.ID == "" {
		user.ID = snap.Ref.ID
	}

	return &user, nil
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]*User, error) {
	snapshots, err := r.collection().Documents(ctx).GetAll()
	if err != nil {
		return nil, errors.New("failed to het all users")
	}

	users := make([]*User, 0, len(snapshots))
	for _, snap := range snapshots {

		user, err := r.toModel(snap)
		if err != nil {
			continue
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) AddUser(ctx context.Context, user *User) error {

	if user == nil {
		return errors.New("User doesn't exist")
	}

	if user.Email == "" || user.Name == "" || user.password == "" {
		return errors.New("User must have an email, name and password")
	}

	var docRef *firestore.DocumentRef
	if user.ID == "" {
		docRef = r.collection().NewDoc()
		user.ID = docRef.ID
	} else {
		docRef = r.collection().Doc(user.ID)
	}

	if _, err := docRef.Set(ctx, user); err != nil {
		return err
	}

	return nil
}
