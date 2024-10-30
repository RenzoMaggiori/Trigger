package action

import (
	"context"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"trigger.com/trigger/pkg/errors"
)

func (m Model) About(remoteAddr string) (AboutModel, error) {
	host, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return AboutModel{}, err
	}

	actions, err := m.Get()
	if err != nil {
		return AboutModel{}, err
	}

	areaServices := map[string]ServiceModel{}
	for _, a := range actions {
		service := areaServices[a.Provider]
		if a.Type == "trigger" {
			service.Actions = append(service.Actions, AreaModel{
				Name:        a.Action,
				Description: a.Action,
			})
		} else {
			service.Reactions = append(service.Reactions, AreaModel{
				Name:        a.Action,
				Description: a.Action,
			})
		}
		areaServices[a.Provider] = service
	}

	services := make([]ServiceModel, 0)
	for k, v := range areaServices {
		services = append(services, ServiceModel{
			Name:      k,
			Actions:   v.Actions,
			Reactions: v.Reactions,
		})
	}

	about := AboutModel{
		Client: ClientModel{
			Host: host,
		},
		Server: ServerModel{
			CurrentTime: time.Now().Unix(),
			Services:    services,
		},
	}
	return about, nil
}

func getService(services []ServiceModel, name string) *ServiceModel {
	for _, s := range services {
		if s.Name != name {
			continue
		}
		return &s
	}
	return nil
}

func (m Model) Get() ([]ActionModel, error) {
	actions := make([]ActionModel, 0)
	ctx := context.TODO()
	filter := bson.M{}
	cursor, err := m.Collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &actions); err != nil {
		return nil, err
	}
	return actions, nil
}

func (m Model) GetById(id primitive.ObjectID) (*ActionModel, error) {
	var action ActionModel
	action.Input = make([]string, 0)
	action.Output = make([]string, 0)
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	err := m.Collection.FindOne(ctx, filter).Decode(&action)

	if err != nil {
		return nil, errors.ErrActionNotFound
	}
	return &action, nil
}

func (m Model) GetByProvider(provider string) ([]ActionModel, error) {
	actions := make([]ActionModel, 0)
	ctx := context.TODO()
	filter := bson.M{"provider": provider}
	cursor, err := m.Collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &actions); err != nil {
		return nil, err
	}
	return actions, nil
}

func (m Model) GetByActionName(actionName string) (*ActionModel, error) {
	var action ActionModel
	action.Input = make([]string, 0)
	action.Output = make([]string, 0)
	ctx := context.TODO()
	filter := bson.M{"action": actionName}
	err := m.Collection.FindOne(ctx, filter).Decode(&action)

	if err != nil {
		return nil, errors.ErrActionNotFound
	}
	return &action, nil
}

func (m Model) Add(add *AddActionModel) (*ActionModel, error) {
	ctx := context.TODO()

	if add.Type != "trigger" && add.Type != "reaction" {
		return nil, errors.ErrActionTypeNone
	}
	newAction := ActionModel{
		Id:       primitive.NewObjectID(),
		Input:    add.Input,
		Output:   add.Output,
		Provider: add.Provider,
		Type:     add.Type,
		Action:   add.Action,
	}
	_, err := m.Collection.InsertOne(ctx, newAction)

	if err != nil {
		return nil, err
	}
	return &newAction, nil
}
