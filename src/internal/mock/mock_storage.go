package mock

import (
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
	"birdai/src/internal/utils"
	"errors"
	"fmt"
	"net/http"

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

func (m *mockCollection) FindOne(query bson.M) (models.Response) {

	doc, err := bson.Marshal(query)
	if err != nil {
		return utils.ErrorParams(err.Error())
	}
	//objectID, err := primitive.ObjectIDFromHex(id)
	for _, one := range m.data {
		switch one.(type) {
		case *models.User:
			var test *models.User
			err = bson.Unmarshal(doc, &test)
			if test.Id != "" && test.Id == one.(*models.User).Id || test.AuthId != "" && test.AuthId == one.(*models.User).AuthId {
				return utils.Response(one)
			}
		case *models.Admin:
			var test *models.Admin
			err = bson.Unmarshal(doc, &test)
			if test.Id != "" && test.Id == one.(*models.Admin).Id {
				return utils.Response(one)
			}
		case *models.Bird:
			var test *models.Bird
			err = bson.Unmarshal(doc, &test)
			if test.Id != "" && test.Id == one.(*models.Bird).Id {
				return utils.Response(one)
			}
		case *models.Post:
			var test *models.Post
			err = bson.Unmarshal(doc, &test)
			if test.Id != "" && test.Id == one.(*models.Post).Id {
				return utils.Response(one)
			}
		case *models.Media:
			var test *models.Media
			err = bson.Unmarshal(doc, &test)
			if test.Id != "" && test.Id == one.(*models.Media).Id {
				return utils.Response(one)
			}
		}
	}
	return utils.ErrorNotFoundInDatabase("User collection")
}

func (m *mockCollection) FindAll() (models.Response) {
	if len(m.data) == 0 {
		return utils.ErrorNotFoundInDatabase("User collection")
	}
	switch m.data[0].(type) {
	case *models.User:
		var list []*models.User
		for _, ob := range m.data {
			list = append(list, ob.(*models.User))
		}
		return utils.Response(list)
	case *models.Admin:
		var list []*models.Admin
		for _, ob := range m.data {
			list = append(list, ob.(*models.Admin))
		}
		return utils.Response(list)
	case *models.Bird:
		var list []*models.Bird
		for _, ob := range m.data {
			list = append(list, ob.(*models.Bird))
		}
		return utils.Response(list)
	case *models.Post:
		var list []*models.Post
		for _, ob := range m.data {
			list = append(list, ob.(*models.Post))
		}
		return utils.Response(list)
	case *models.Media:
		var list []*models.Media
		for _, ob := range m.data {
			list = append(list, ob.(*models.Media))
		}
		return utils.Response(list)
	default:
		return utils.ErrorToResponse(http.StatusBadRequest, "Wrong model", errors.New("Could not find objects").Error())
	}
}

func (m *mockCollection) UpdateOne(query bson.M) (models.Response) {
	doc, err := bson.Marshal(query)
	if err != nil {
		return utils.ErrorToResponse(http.StatusBadRequest, "Could not update object", err.Error())
	}
	id := query["_id"]
	for i, one := range m.data {
		if one.GetId() == id {
			switch one.(type) {
			case *models.User:
				var test *models.User
				err = bson.Unmarshal(doc, &test)
				m.data[i] = test
				return utils.Response(test)
			case *models.Admin:
				var test *models.Admin
				err = bson.Unmarshal(doc, &test)
				m.data[i] = test
				return utils.Response(test)
			case *models.Bird:
				var test *models.Bird
				err = bson.Unmarshal(doc, &test)
				m.data[i] = test
				return utils.Response(test)
			case *models.Post:
				var test *models.Post
				err = bson.Unmarshal(doc, &test)
				m.data[i] = test
				return utils.Response(test)
			case *models.Media:
				var test *models.Media
				err = bson.Unmarshal(doc, &test)
				m.data[i] = test
				return utils.Response(test)
			}
		}
	}

	return utils.ErrorNotFoundInDatabase("User collection")
}

func (m *mockCollection) DeleteOne(query bson.M) (models.Response) {
	//objectID, err := primitive.ObjectIDFromHex(id)
	for i, one := range m.data {
		if one.GetId() == query["_id"] {
			m.data = append(m.data[:i], m.data[i+1:]...)
			return utils.Response(one)
		}
	}
	return utils.ErrorNotFoundInDatabase("User collection")
}

func (m *mockCollection) CreateOne(object models.HandlerObject) (models.Response) {
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
		return utils.ErrorToResponse(http.StatusBadRequest, "Could not create object", "")
	}
	m.data = append(m.data, newObject)
	return utils.Response(newObject)
}
