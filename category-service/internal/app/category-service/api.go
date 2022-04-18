package category_service

import (
	"github.com/ozon-edu-go-2021/week-3-workshop/category-service/internal/service/category"
	desc "github.com/ozon-edu-go-2021/week-3-workshop/category-service/pkg/category-service"
)

type Implementation struct {
	desc.UnimplementedCategoryServiceServer

	categoryService category.Service
}

func NewCategoryService(categoryService category.Service) desc.CategoryServiceServer {
	return &Implementation{
		categoryService: categoryService,
	}
}
