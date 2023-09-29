package mock

import (
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockMongoInstance struct {
	Collections map[string]repositories.IMongoCollection
}

func (m MockMongoInstance) GetCollection(name string) repositories.IMongoCollection {
	return m.Collections[name]
}

func (m MockMongoInstance) AddCollection(name string) {
	m.Collections[name] = &mockCollection{[]models.HandlerObject{}}
}

func (m MockMongoInstance) DisconnectDB() {
	fmt.Println("Disconnected")
}

func NewMockMongoInstance() repositories.IMongoInstance {
	return MockMongoInstance{map[string]repositories.IMongoCollection{}}
}

type mockCollection struct {
	data []models.HandlerObject
}

func (m *mockCollection) FindOne(query bson.M) (models.HandlerObject, error) {
	for _, one := range m.data {
		if one.GetId() == query["_id"] {
			return one, nil
		}
	}
	return nil, errors.New("could not find")
}

func (m *mockCollection) FindAll() (interface{}, error) {
	return m.data, nil
}

func (m *mockCollection) UpdateOne(query bson.M) (models.HandlerObject, error) {
	doc, err := bson.Marshal(query)
	if err != nil {
		return nil, errors.New("wrong format")
	}
	id := query["_id"]
	for i, one := range m.data {
		if one.GetId() == id {
			switch one.(type) {
			case *models.User:
				var test *models.User
				err = bson.Unmarshal(doc, &test)
				m.data[i] = test
				return test, err
			case *models.Admin:
				var test *models.Admin
				err = bson.Unmarshal(doc, &test)
				m.data[i] = test
				return test, err
			case *models.Bird:
				var test *models.Bird
				err = bson.Unmarshal(doc, &test)
				m.data[i] = test
				return test, err
			case *models.Post:
				var test *models.Post
				err = bson.Unmarshal(doc, &test)
				m.data[i] = test
				return test, err
			case *models.Media:
				var test *models.Media
				err = bson.Unmarshal(doc, &test)
				m.data[i] = test
				return test, err
			}
		}
	}

	return nil, errors.New("could not find")
}

func (m *mockCollection) DeleteOne(query bson.M) (models.HandlerObject, error) {
	//objectID, err := primitive.ObjectIDFromHex(id)
	for i, one := range m.data {
		if one.GetId() == query["_id"] {
			m.data = append(m.data[:i], m.data[i+1:]...)
			return one, nil
		}
	}
	return nil, errors.New("could not find")
}

func (m *mockCollection) CreateOne(object models.HandlerObject) (string, error) {
	var newObject models.HandlerObject
	switch object.(type) {
	case *models.User:
		newObject = &models.User{
			Id:       primitive.NewObjectID().Hex(),
			Username: object.(*models.User).Username,
			AuthId:   object.(*models.User).AuthId,
			Active:   object.(*models.User).Active,
		}
		newObject.SetCreatedAt()
	case *models.Admin:
		newObject = &models.Admin{
			Id:     primitive.NewObjectID().Hex(),
			UserId: object.(*models.Admin).UserId,
			Access: object.(*models.Admin).Access,
		}

	case *models.Bird:
		newObject = &models.Bird{
			Id:          primitive.NewObjectID().Hex(),
			Name:        object.(*models.Bird).Name,
			Description: object.(*models.Bird).Description,
			ImageId:     object.(*models.Bird).ImageId,
			SoundId:     object.(*models.Bird).SoundId,
		}

	case *models.Post:
		newObject = &models.Post{
			Id:       primitive.NewObjectID().Hex(),
			UserId:   object.(*models.Post).UserId,
			BirdId:   object.(*models.Post).BirdId,
			Location: object.(*models.Post).Location,
			ImageId:  object.(*models.Post).ImageId,
			SoundId:  object.(*models.Post).SoundId,
		}
		newObject.SetCreatedAt()

	case *models.Media:
		newObject = &models.Media{
			Id:       primitive.NewObjectID().Hex(),
			Data:     object.(*models.Media).Data,
			FileType: object.(*models.Media).FileType,
		}

	default:
		return "", nil
	}
	m.data = append(m.data, newObject)
	return newObject.GetId(), nil
}
