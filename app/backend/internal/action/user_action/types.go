package useraction

const UserActionCtxKey string = "UserActionCtxKey"

// type Service interface {
// 	Get() ([]SessionModel, error)
// 	GetById(primitive.ObjectID) (*SessionModel, error)
// 	GetByUserId(primitive.ObjectID) ([]SessionModel, error)
// 	Add(*AddSessionModel) (*SessionModel, error)
// 	UpdateById(primitive.ObjectID, *UpdateSessionModel) (*SessionModel, error)
// 	DeleteById(primitive.ObjectID) error
// 	GetByToken(string) (*SessionModel, error)
// }

// type Handler struct {
// 	Service
// }

// type Model struct {
// 	Collection *mongo.Collection
// }

// type UserActionModel struct {
// 	Id       primitive.ObjectID   `json:"id" bson:"_id"`
// 	ActionId primitive.ObjectID   `json:"actionId" bson:"actionId"`
// 	UserId   primitive.ObjectID   `json:"userId" bson:"userId"`
// 	Children []primitive.ObjectID `json:"children" bson:"children"`
// 	Fields   map[string]string    `json:"fields" bson:"fields"`
// }

// type AddSessionModel struct {
// 	UserId       primitive.ObjectID `json:"userId" bson:"userId"`
// 	ProviderName *string            `json:"providerName,omitempty" bson:"providerName,omitempty"`
// 	AccessToken  string             `json:"accessToken" bson:"accessToken"`
// 	RefreshToken *string            `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
// 	Expiry       time.Time          `json:"expiry" bson:"expiry"`
// 	IdToken      *string            `json:"idToken,omitempty" bson:"idToken,omitempty"`
// }

// type UpdateSessionModel struct {
// 	AccessToken  *string    `json:"accessToken,omitempty" bson:"accessToken,omitempty"`
// 	RefreshToken *string    `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
// 	Expiry       *time.Time `json:"expiry,omitempty" bson:"expiry,omitempty"`
// 	IdToken      *string    `json:"idToken,omitempty" bson:"idToken,omitempty"`
// }
