package service

import (
	"context"

	"github.com/ArthurDotSaito/gRPC-go/internal/database"
	"github.com/ArthurDotSaito/gRPC-go/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(cxt context.Context, aNewCategory *pb.CreateCategoryRequest) (*pb.CategoryRespose, error) {
	category, err := c.CategoryDB.Create(aNewCategory.Name, aNewCategory.Description)
	if err != nil {
		return nil, err
	}
	categoryResponse := &pb.Category{
		Id: category.ID,
		Name: category.Name,
		Description: category.Description,

	}

	return &pb.CategoryRespose{Category: categoryResponse}, nil
}