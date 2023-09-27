package mock

import (
	"birdai/src/internal/storage"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockMongoInstance struct {
	Collections map[string]storage.IMongoCollection
}

func (m MockMongoInstance) GetCollection(name string) storage.IMongoCollection {
	return m.Collections[name]
}

func (m MockMongoInstance) AddCollection(name string) {
	m.Collections[name] = &mockCollection{[]storage.HandlerObject{}}
}

func (m MockMongoInstance) DisconnectDB() {
	fmt.Println("Disconnected")
}

func NewMockMongoInstance() storage.IMongoInstance {
	return MockMongoInstance{map[string]storage.IMongoCollection{}}
}

type mockCollection struct {
	data []storage.HandlerObject
}

func (m *mockCollection) FindOne(id string) (storage.HandlerObject, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	for _, one := range m.data {
		if one.GetID() == objectID {
			return one, nil
		}
	}
	return nil, errors.New("could not find")
}

func (m *mockCollection) FindAll() ([]storage.HandlerObject, error) {
	return m.data, nil
}

func (m *mockCollection) UpdateOne(query bson.D) (storage.HandlerObject, error) {
	doc, err := bson.Marshal(query)
	if err != nil {
		return nil, errors.New("wrong format")
	}
	id := query[0].Value
	for _, one := range m.data {
		if one.GetID() == id {
			switch one.(type) {
			case storage.User:
				var test storage.User
				err = bson.Unmarshal(doc, &test)
				one = test
				return test, err
			case storage.Admin:
				var test storage.Admin
				err = bson.Unmarshal(doc, &test)
				one = test
				return test, err
			case storage.Bird:
				var test storage.Bird
				err = bson.Unmarshal(doc, &test)
				one = test
				return test, err
			case storage.Post:
				var test storage.Post
				err = bson.Unmarshal(doc, &test)
				one = test
				return test, err
			case storage.Sound:
				var test storage.Sound
				err = bson.Unmarshal(doc, &test)
				one = test
				return test, err
			case storage.Image:
				var test storage.Image
				err = bson.Unmarshal(doc, &test)
				one = test
				return test, err

			}
		}
	}

	return nil, errors.New("could not find")
}

func (m *mockCollection) DeleteOne(id string) (storage.HandlerObject, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	for i, one := range m.data {
		if one.GetID() == objectID {
			m.data = append(m.data[:i], m.data[i+1:]...)
			return one, nil
		}
	}
	return nil, errors.New("could not find")
}

func (m *mockCollection) CreateOne(object storage.HandlerObject) error {
	m.data = append(m.data, object)
	return nil
}
