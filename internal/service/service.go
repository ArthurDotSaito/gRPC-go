package service

import (
	"context"
	"io"

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

func (c *CategoryService) CreateCategory(cxt context.Context, aNewCategory *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(aNewCategory.Name, aNewCategory.Description)
	if err != nil {
		return nil, err
	}
	categoryResponse := &pb.Category{
		Id: category.ID,
		Name: category.Name,
		Description: category.Description,

	}

	return categoryResponse, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Empty) (*pb.CategoryListResponse, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var categoryResponses []*pb.Category
	
	for _, category := range categories {
		categoryResponses = append(categoryResponses, &pb.Category{
			Id: category.ID,
			Name: category.Name,
			Description: category.Description,
		})
	}

	return &pb.CategoryListResponse{Categories: categoryResponses}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.CategoryGetRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.FindByCourseID(in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Category{
		Id: category.ID,
		Name: category.Name,
		Description: category.Description,
	}, nil
}

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error{
	categories := &pb.CategoryListResponse{}

	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(categories);
		}
		if err != nil {
			return err
		}

		categoryResult, err := c.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}
		categories.Categories = append(categories.Categories, &pb.Category{
			Id: categoryResult.ID,
			Name: categoryResult.Name,
			Description: categoryResult.Description,
		})
	}
}