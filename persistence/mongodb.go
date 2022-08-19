package persistence

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/kasuma0/gobozito/conf"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mongotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go.mongodb.org/mongo-driver/mongo"
)

var (
	mongocli *mongo.Client
)

//MongoDb ... struct for mongoDB conection
type MongoDb struct {
	Collection string
	DB         string
}

func init() {
	ops := options.Client().SetMonitor(mongotrace.NewMonitor())
	ops.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	ops.SetMaxPoolSize(conf.DiscordConfiguration.Credentials.MongoDB.MaxPoolSize)
	ops.SetMinPoolSize(conf.DiscordConfiguration.Credentials.MongoDB.MinPoolSize)
	ops.SetMaxConnIdleTime(conf.DiscordConfiguration.Credentials.MongoDB.MaxConnIdleTime * time.Minute)
	ops.ApplyURI(conf.DiscordConfiguration.Credentials.MongoDB.URL)
	client, err := mongo.Connect(context.Background(), ops)
	if err != nil {
		panic(err)
	}
	mongocli = client
	_ = mongocli.Ping(context.TODO(), nil)
	log.Info("MongoDB connection success")
}

//GetManyDocument ... Get all documents without any filter from a collection
func (mon *MongoDb) GetManyDocument(ctx context.Context) (*mongo.Cursor, error) {
	db := mongocli.Database(mon.DB).Collection(mon.Collection)
	resultDb, err := db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	return resultDb, nil
}

/*GetManyDocumentFiltered ...
Método de búsqueda con datos filtrados.
se recomienda usar Cursor.All() como en el ejemplo:
	var result Type{}
	defer Cursor.Close()
	err = Cursor.All(context.TODO(), &result)
	if err != nil {
		return nil, err
	}*/
func (mon *MongoDb) GetManyDocumentFiltered(ctx context.Context, filter bson.D, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	db := mongocli.Database(mon.DB).Collection(mon.Collection)
	resultDb, err := db.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}

	return resultDb, nil
}

/*GetDocumentFiltered ...
Método de búsqueda con datos filtrados.
se recomienda usar Decode() como en el ejemplo:
	var result Type{}
	err = result.Decode(&result) (result respuesta de metodo GetDocumentFiltered)
	if err != nil {
		return nil, err
	}*/
func (mon *MongoDb) GetDocumentFiltered(ctx context.Context, filter bson.D, opts ...*options.FindOneOptions) *mongo.SingleResult {
	db := mongocli.Database(mon.DB).Collection(mon.Collection)
	resultDb := db.FindOne(ctx, filter, opts...)
	return resultDb
}

//InsertDocuments ... Insert many documents in a mongo collection.
func (mon *MongoDb) InsertDocuments(ctx context.Context, doc []interface{}) (interface{}, error) {
	db := mongocli.Database(mon.DB).Collection(mon.Collection)
	resultDb, err := db.InsertMany(ctx, doc)
	if err != nil {
		return nil, err
	}
	return resultDb, nil
}

//InsertDocument ... Insert one document in a mongo collection.
func (mon *MongoDb) InsertDocument(ctx context.Context, doc interface{}) (*mongo.InsertOneResult, error) {
	db := mongocli.Database(mon.DB).Collection(mon.Collection)
	resultDb, err := db.InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}
	return resultDb, nil
}

//UpdateDocument ... Update one document in a mongo collection
func (mon *MongoDb) UpdateDocument(ctx context.Context, filter bson.D, update bson.D) (*mongo.UpdateResult, error) {
	db := mongocli.Database(mon.DB).Collection(mon.Collection)
	result, err := db.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//UpdateManyDocuments ... Update many documents in a mongo collection
func (mon *MongoDb) UpdateManyDocuments(ctx context.Context, filter bson.D, update bson.D) (*mongo.UpdateResult, error) {
	db := mongocli.Database(mon.DB).Collection(mon.Collection)
	result, err := db.UpdateMany(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//DeleteDocument ... Delete one document in a mongo collection
func (mon *MongoDb) DeleteDocument(ctx context.Context, filter bson.D, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	db := mongocli.Database(mon.DB).Collection(mon.Collection)
	result, err := db.DeleteOne(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//GetCollectionsName ... get an array of collections name
func (mon *MongoDb) GetCollectionsName(ctx context.Context, filter bson.D, opts ...*options.ListCollectionsOptions) ([]string, error) {
	collections, err := mongocli.Database(mon.DB).ListCollectionNames(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}

	return collections, nil
}
