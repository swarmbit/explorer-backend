package storage

import (
    "context"
    "errors"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

    "github.com/spacemeshos/explorer-backend/model"
)

func (s *Storage) InitLayersStorage(ctx context.Context) error {
    _, err := s.db.Collection("layers").Indexes().CreateOne(ctx, mongo.IndexModel{Keys: bson.D{{"number", 1}}, Options: options.Index().SetName("numberIndex").SetUnique(true)});
    return err
}

func (s *Storage) GetLayer(parent context.Context, query *bson.D) (*model.Layer, error) {
    ctx, cancel := context.WithTimeout(parent, 5*time.Second)
    defer cancel()
    cursor, err := s.db.Collection("layers").Find(ctx, query)
    if err != nil {
        return nil, err
    }
    if !cursor.Next(ctx) {
        return nil, errors.New("Empty result")
    }
    doc := cursor.Current
    account := &model.Layer{
        Number: uint32(doc.Lookup("number").Int32()),
        Status: int(doc.Lookup("status").Int32()),
    }
    return account, nil
}

func (s *Storage) GetLayers(parent context.Context, query *bson.D, opts ...*options.FindOptions) ([]*model.Layer, error) {
    ctx, cancel := context.WithTimeout(parent, 5*time.Second)
    defer cancel()
    cursor, err := s.db.Collection("layers").Find(ctx, query, opts...)
    if err != nil {
        return nil, err
    }
    var docs interface{} = []bson.D{}
    err = cursor.All(ctx, &docs)
    if err != nil {
        return nil, err
    }
    if len(docs.([]bson.D)) == 0 {
        return nil, nil
    }
    layers := make([]*model.Layer, len(docs.([]bson.D)), len(docs.([]bson.D)))
    for i, doc := range docs.([]bson.D) {
        layers[i] = &model.Layer{
            Number: uint32(doc[0].Value.(int32)),
            Status: int(doc[1].Value.(int32)),
        }
    }
    return layers, nil
}

func (s *Storage) SaveLayer(parent context.Context, in *model.Layer) error {
    ctx, cancel := context.WithTimeout(parent, 5*time.Second)
    defer cancel()
    _, err := s.db.Collection("layers").InsertOne(ctx, bson.D{
        {"number", in.Number},
        {"status", in.Status},
    })
    return err
}
