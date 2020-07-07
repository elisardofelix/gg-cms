package DataRepos

import (
	"context"
	"fmt"
	"gg-cms/DB"
	"gg-cms/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepo interface {
	Insert(user Models.User) (Models.User, error)
	Update(user Models.User) (Models.User, error)
	Delete(ID string) error
	Get(username string) (Models.User, error)
	GetAllUsers(setLimit int64, setSkip int64) (*mongo.Cursor, error)
	ExistsCredential(username string, password string) (bool, bool, error)
	ExistsAny(active bool) (bool, error)
}

type userRepo struct {
	client *mongo.Client
	dbName string
	colletionName string
}

func NewUserRepo () UserRepo {
	return newUserRepo(DB.Conf.DataBase)
}

func NewUserRepoTest () UserRepo {
	return newUserRepo(DB.Conf.TestDataBase)
}

func newUserRepo(database string) UserRepo {
	dbclient, err := DB.GetMongoDBClient()
	//Create index in collection for unique values in the field permaLink
	if err == nil {
		collection := dbclient.Database(database).Collection(collections["users"])
		createUniqueIndex(collection, "userName")
		createUniqueIndex(collection, "email")
	}

	return &userRepo {
		client : dbclient,
		dbName: database,
		colletionName: collections["users"],
	}
}

func createUniqueIndex(collection *mongo.Collection, field string){
	mod := mongo.IndexModel{
		Keys: bson.M{
			field: 1,
		},
		Options: options.Index().SetUnique(true),
	}
	collection.Indexes().CreateOne(context.TODO(), mod)

}

func (ur *userRepo) Insert(user Models.User) (Models.User, error) {
	collection := ur.client.Database(ur.dbName).Collection(ur.colletionName)
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println(err.Error())
		return Models.User{}, err
	}
	oid, _ := insertResult.InsertedID.(primitive.ObjectID)
	user.ID = oid.Hex()
	return user, err
}

func (ur *userRepo) Update(user Models.User) (Models.User, error) {
	collection := ur.client.Database(ur.dbName).Collection(ur.colletionName)
	primID,_ := primitive.ObjectIDFromHex(user.ID)
	filter := bson.D{{"_id", primID}}

	update := bson.D{
		{"$set", bson.D{
			{"userName", user.UserName},
			{"status", user.Status},
			{"email", user.Email},
			{"isAdmin", user.IsAdmin},
		}},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return Models.User{}, err
	} else {
		var updatedUser Models.User
		collection.FindOne(context.TODO(), filter).Decode(&updatedUser)
		return updatedUser, nil
	}
}

func (ur *userRepo) Delete(ID string) error {
	collection := ur.client.Database(ur.dbName).Collection(ur.colletionName)
	primID,_ := primitive.ObjectIDFromHex(ID)
	filter := bson.D{{"_id", primID}}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}


func (ur *userRepo) Get(username string) (Models.User, error) {
	var result Models.User
	collection := ur.client.Database(ur.dbName).Collection(ur.colletionName)
	filter := bson.D{{"userName", username}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	return result, err
}


func (ur *userRepo) GetAllUsers(setLimit int64, setSkip int64) (*mongo.Cursor, error) {
	collection := ur.client.Database(ur.dbName).Collection(ur.colletionName)

	// Pass these options to the Find method
	findOptions := options.Find()
	//Example for Sorting
	//findOptions.SetSort(map[string]int{"when": -1})
	findOptions.SetSkip(setSkip)
	findOptions.SetLimit(setLimit)

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	return cur, err
}


func (ur *userRepo) ExistsCredential(username string, password string) (bool, bool, error){
	var user Models.User

	collection := ur.client.Database(ur.dbName).Collection(ur.colletionName)
	filter := bson.D{
		{"userName", username},
		{"password", password},
		{"status", "Active"}}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	if user.ID != "" && err == nil {
		return true, user.IsAdmin, nil
	}

	return false, false, err
}

func (ur *userRepo) ExistsAny(active bool) (bool, error){
	collection := ur.client.Database(ur.dbName).Collection(ur.colletionName)
	filter := bson.D{}
	if active {
		filter = bson.D{{"status", "Active"}}
	}
	qty, err := collection.CountDocuments(context.TODO(), filter, nil)
	if qty > 0 && err == nil {
		return true, nil
	}

	return false, err
}
