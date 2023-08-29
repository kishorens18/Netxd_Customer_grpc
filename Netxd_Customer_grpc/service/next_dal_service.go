package service

import (
	"context"
	"netxd_dal/interfaces"
	"netxd_dal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx                context.Context
}

func InitCustomerService(collection *mongo.Collection, ctx context.Context) interfaces.ICustomer {
	return &CustomerService{collection, ctx}
}

func (p *CustomerService) CreateCustomer(customer *models.Customer) (*models.DBResponse, error) {
	// Insert the customer into the MongoDB collection
	res, err := p.CustomerCollection.InsertOne(p.ctx, customer)
	if err != nil {
		return nil, err
	}

	// Construct the response with customer ID (primitive.ObjectID) and CreatedAt
	response := &models.DBResponse{
		CustomerID: res.InsertedID.(primitive.ObjectID),
		CreatedAt:  customer.CreatedAt,
	}

	return response, nil
}
