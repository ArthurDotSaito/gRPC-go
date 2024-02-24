package service

import{
	"github.com/ta04/course-service/internal/database"
	"github.com/ta04/course-service/internal/pb"

}

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryServuice(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB
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