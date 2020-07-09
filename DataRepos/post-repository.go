package DataRepos

import (
	"fmt"
	"context"
	"gg-cms/DB"
	"gg-cms/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

type PostRepo interface {
	Insert(post Models.Post) (Models.Post, error)
	Update(post Models.Post) (Models.Post, error)
	Delete(ID string) error
	Get(permaLink string) (Models.Post, error)
	GetAllActive(setLimit int64, setSkip int64, areActive bool) ([]Models.Post, error)
}

type postRepo struct {
	client *mongo.Client
	dbName string
	colletionName string
}

func NewPostRepo() PostRepo {
	dbclient, err := DB.GetMongoDBClient()
	//Create index in collection for unique values in the field permaLink
	if err == nil {
		collection := dbclient.Database(DB.Conf.DataBase).Collection(collections["posts"])
		mod := mongo.IndexModel{
			Keys: bson.M{
				"permaLink": 1,
			},
			Options: options.Index().SetUnique(true),
		}
		collection.Indexes().CreateOne(context.TODO(), mod)
	}

	return &postRepo {
		client : dbclient,
		dbName: DB.Conf.DataBase,
		colletionName: collections["posts"],
	}
}

func NewPostRepoTest() PostRepo {
	dbclient, err := DB.GetMongoDBClient()
	//Create index in collection for unique values in the field permaLink
	if err == nil {
		collection := dbclient.Database(DB.Conf.TestDataBase).Collection(collections["posts"])
		mod := mongo.IndexModel{
			Keys: bson.M{
				"permaLink": 1,
			},
			Options: options.Index().SetUnique(true),
		}
		collection.Indexes().CreateOne(context.TODO(), mod)
	}

	return &postRepo {
		client : dbclient,
		dbName: DB.Conf.TestDataBase,
		colletionName: collections["posts"],
	}
}

func (pr *postRepo) Insert(post Models.Post) (Models.Post, error) {
	collection := pr.client.Database(pr.dbName).Collection(pr.colletionName)
	insertResult, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		fmt.Println(err.Error())
		return Models.Post{}, err
	}
	oid, _ := insertResult.InsertedID.(primitive.ObjectID)
	post.ID = oid.Hex()
	return post, err
}

func (pr *postRepo) Update(post Models.Post) (Models.Post, error) {
	collection := pr.client.Database(pr.dbName).Collection(pr.colletionName)
	primID,_ := primitive.ObjectIDFromHex(post.ID)
	filter := bson.D{{"_id", primID}}

	update := bson.D{
		{"$set", bson.D{
			{"title", post.Title},
			{"content", post.Content},
			{"permaLink", post.PermaLink},
		}},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return Models.Post{}, err
	} else {
		var updatedPost Models.Post
		collection.FindOne(context.TODO(), filter).Decode(&updatedPost)
		return updatedPost, nil
	}
}

func (pr *postRepo) Delete(ID string) error {
	collection := pr.client.Database(pr.dbName).Collection(pr.colletionName)
	primID,_ := primitive.ObjectIDFromHex(ID)
	filter := bson.D{{"_id", primID}}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}


func (pr *postRepo) Get(permaLink string) (Models.Post, error) {
	var result Models.Post
	collection := pr.client.Database(pr.dbName).Collection(pr.colletionName)
	filter := bson.D{{"permaLink", permaLink}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	return result, err
}


func (pr *postRepo) GetAllActive(setLimit int64, setSkip int64, areActive bool) ([]Models.Post, error) {
	var results = make([]Models.Post,0)
	collection := pr.client.Database(pr.dbName).Collection(pr.colletionName)

	// Pass these options to the Find method
	findOptions := options.Find()
	//Example for Sorting
	//findOptions.SetSort(map[string]int{"when": -1})
	findOptions.SetSkip(setSkip)
	findOptions.SetLimit(setLimit)
	filter := bson.D{{}}

	if areActive {
		filter = bson.D{{"status", "Active"}}
	}

	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return results, err
	}
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Models.Post
		err := cur.Decode(&elem)
		if err != nil {
			return results, err
		}

        elem.Content = getFirstElementString(elem.Content, "p")
		results = append(results, elem)
	}

	return results, nil
}

func getFirstElementString(str string, elem string) string {
	start := fmt.Sprintf("<%s>", elem)
	end := fmt.Sprintf("</%s>", elem)

	s := strings.Index(str, strings.ToLower(start))
	e := strings.Index(str, strings.ToLower(end))

	if s >= 0 && e > 0 {
		return str[s : e+len(end)]
	}

	s = strings.Index(str, strings.ToUpper(start))
	e = strings.Index(str, strings.ToUpper(end))

	if s >= 0 && e > 0 {
		return str[s : e+len(end)]
	}
	strLen := len(str)

	if strLen > 1000 {
		strLen = 1000
	}


	return str[:strLen]
}