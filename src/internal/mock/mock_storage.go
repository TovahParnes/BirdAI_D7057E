package mock

import (
	"birdai/src/internal/models"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockMongoInstance struct {
	Collections map[string]models.IMongoCollection
}

func (m MockMongoInstance) GetCollection(name string) models.IMongoCollection {
	return m.Collections[name]
}

func (m MockMongoInstance) AddCollection(name string) {
	m.Collections[name] = &mockCollection{[]models.HandlerObject{}}
}

func (m MockMongoInstance) DisconnectDB() {
	fmt.Println("Disconnected")
}

func NewMockMongoInstance() models.IMongoInstance {
	return MockMongoInstance{map[string]models.IMongoCollection{}}
}

type mockCollection struct {
	data []models.HandlerObject
}

func (m *mockCollection) FindOne(id string) (models.HandlerObject, error) {
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

func (m *mockCollection) FindAll() ([]models.HandlerObject, error) {
	return m.data, nil
}

func (m *mockCollection) UpdateOne(query bson.D) (models.HandlerObject, error) {
	doc, err := bson.Marshal(query)
	if err != nil {
		return nil, errors.New("wrong format")
	}
	id := query[0].Value
	for _, one := range m.data {
		if one.GetID() == id {
			switch one.(type) {
			case models.User:
				var test models.User
				err = bson.Unmarshal(doc, &test)
				one = test
				return test, err
			case models.Admin:
				var test models.Admin
				err = bson.Unmarshal(doc, &test)
				one = test
				return test, err
			case models.Bird:
				var test models.Bird
				err = bson.Unmarshal(doc, &test)
				one = test
				return test, err
			case models.Post:
				var test models.Post
				err = bson.Unmarshal(doc, &test)
				one = test
				return test, err
			case models.Sound:
				var test models.Sound
				err = bson.Unmarshal(doc, &test)
				one = test
				return test, err
			case models.Image:
				var test models.Image
				err = bson.Unmarshal(doc, &test)
				one = test
				return test, err

			}
		}
	}

	return nil, errors.New("could not find")
}

func (m *mockCollection) DeleteOne(id string) (models.HandlerObject, error) {
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

func (m *mockCollection) CreateOne(object models.HandlerObject) error {
	m.data = append(m.data, object)
	return nil
}
