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

func (m *mockCollection) FindOne(query bson.M) models.Response {

	doc, err := bson.Marshal(query)
	if err != nil {
		return utils.ErrorParams(err.Error())
	}
	//objectID, err := primitive.ObjectIDFromHex(id)
	for _, one := range m.data {
		switch one.(type) {
		case *models.UserDB:
			var test *models.UserDB
			err = bson.Unmarshal(doc, &test)
			if test.Id != "" && test.Id == one.(*models.UserDB).Id || test.AuthId != "" && test.AuthId == one.(*models.UserDB).AuthId {
				return utils.Response(one)
			}
		case *models.AdminDB:
			var test *models.AdminDB
			err = bson.Unmarshal(doc, &test)
			if test.Id != "" && test.Id == one.(*models.AdminDB).Id {
				return utils.Response(one)
			}
		case *models.BirdDB:
			var test *models.BirdDB
			err = bson.Unmarshal(doc, &test)
			if test.Id != "" && test.Id == one.(*models.BirdDB).Id {
				return utils.Response(one)
			}
		case *models.PostDB:
			var test *models.PostDB
			err = bson.Unmarshal(doc, &test)
			if test.Id != "" && test.Id == one.(*models.PostDB).Id {
				return utils.Response(one)
			}
		case *models.MediaDB:
			var test *models.MediaDB
			err = bson.Unmarshal(doc, &test)
			if test.Id != "" && test.Id == one.(*models.MediaDB).Id {
				return utils.Response(one)
			}
		}
	}
	return utils.ErrorNotFoundInDatabase("User collection")
}

func (m *mockCollection) FindAll() models.Response {
	if len(m.data) == 0 {
		return utils.ErrorNotFoundInDatabase("")
	}
	switch m.data[0].(type) {
	case *models.UserDB:
		var list []*models.UserDB
		for _, ob := range m.data {
			list = append(list, ob.(*models.UserDB))
		}
		return utils.Response(list)
	case *models.AdminDB:
		var list []*models.AdminDB
		for _, ob := range m.data {
			list = append(list, ob.(*models.AdminDB))
		}
		return utils.Response(list)
	case *models.BirdDB:
		var list []*models.BirdDB
		for _, ob := range m.data {
			list = append(list, ob.(*models.BirdDB))
		}
		return utils.Response(list)
	case *models.PostDB:
		var list []*models.PostDB
		for _, ob := range m.data {
			list = append(list, ob.(*models.PostDB))
		}
		return utils.Response(list)
	case *models.MediaDB:
		var list []*models.MediaDB
		for _, ob := range m.data {
			list = append(list, ob.(*models.MediaDB))
		}
		return utils.Response(list)
	default:
		return utils.ErrorToResponse(http.StatusBadRequest, "Wrong model", errors.New("Could not find objects").Error())
	}
}

func (m *mockCollection) UpdateOne(query bson.M) models.Response {
	doc, err := bson.Marshal(query)
	if err != nil {
		return utils.ErrorToResponse(http.StatusBadRequest, "Could not update object", err.Error())
	}
	id := query["_id"]
	for i, one := range m.data {
		if one.GetId() == id {
			switch one.(type) {
			case *models.UserDB:
				var test *models.UserDB
				err = bson.Unmarshal(doc, &test)
				m.data[i] = test
				return utils.Response(test)
			case *models.AdminDB:
				var test *models.AdminDB
				err = bson.Unmarshal(doc, &test)
				m.data[i] = test
				return utils.Response(test)
			case *models.BirdDB:
				var test *models.BirdDB
				err = bson.Unmarshal(doc, &test)
				m.data[i] = test
				return utils.Response(test)
			case *models.PostDB:
				var test *models.PostDB
				err = bson.Unmarshal(doc, &test)
				m.data[i] = test
				return utils.Response(test)
			case *models.MediaDB:
				var test *models.MediaDB
				err = bson.Unmarshal(doc, &test)
				m.data[i] = test
				return utils.Response(test)
			}
		}
	}

	return utils.ErrorNotFoundInDatabase("User collection")
}

func (m *mockCollection) DeleteOne(query bson.M) models.Response {
	//objectID, err := primitive.ObjectIDFromHex(id)
	for i, one := range m.data {
		if one.GetId() == query["_id"] {
			m.data = append(m.data[:i], m.data[i+1:]...)
			return utils.Response(one)
		}
	}
	return utils.ErrorNotFoundInDatabase("User collection")
}

func (m *mockCollection) CreateOne(object models.HandlerObject) models.Response {
	var newObject models.HandlerObject
	switch object.(type) {
	case *models.UserDB:
		newObject = &models.UserDB{
			Id:       primitive.NewObjectID().Hex(),
			Username: object.(*models.UserDB).Username,
			AuthId:   object.(*models.UserDB).AuthId,
			Active:   object.(*models.UserDB).Active,
		}
		newObject.SetCreatedAt()
	case *models.AdminDB:
		newObject = &models.AdminDB{
			Id:     primitive.NewObjectID().Hex(),
			UserId: object.(*models.AdminDB).UserId,
			Access: object.(*models.AdminDB).Access,
		}

	case *models.BirdDB:
		newObject = &models.BirdDB{
			Id:          primitive.NewObjectID().Hex(),
			Name:        object.(*models.BirdDB).Name,
			Description: object.(*models.BirdDB).Description,
			ImageId:     object.(*models.BirdDB).ImageId,
			SoundId:     object.(*models.BirdDB).SoundId,
		}

	case *models.PostDB:
		newObject = &models.PostDB{
			Id:       primitive.NewObjectID().Hex(),
			UserId:   object.(*models.PostDB).UserId,
			BirdId:   object.(*models.PostDB).BirdId,
			Location: object.(*models.PostDB).Location,
			ImageId:  object.(*models.PostDB).ImageId,
			SoundId:  object.(*models.PostDB).SoundId,
		}
		newObject.SetCreatedAt()

	case *models.MediaDB:
		newObject = &models.MediaDB{
			Id:       primitive.NewObjectID().Hex(),
			Data:     object.(*models.MediaDB).Data,
			FileType: object.(*models.MediaDB).FileType,
		}

	default:
		return utils.ErrorToResponse(http.StatusBadRequest, "Could not create object", "")
	}
	m.data = append(m.data, newObject)
	return utils.Response(newObject)
}
