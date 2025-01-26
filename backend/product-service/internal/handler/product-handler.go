package handler

import (
	"fmt"
	"product-service/internal/models"
	"product-service/internal/service"

	"product-service/internal/middlewares"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productService *service.ProductService
}

func SetupProductRoute(app *fiber.App, productService *service.ProductService) {
	handler := ProductHandler{productService: productService}

	product := app.Group("/product")
	product.Use(middlewares.Authentication)
	product.Use(middlewares.IsSeller)
	product.Post("/", handler.CreateProduct)
	product.Get("/:page", handler.GetProduct)
	product.Delete("/:product_id", handler.DeleteProduct)
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	// form, _ := c.MultipartForm()

	//
	// mainImage, _ := c.FormFile("main_image")
	// fmt.Println(mainImage)

	// subImages := form.File["sub_image"]
	// for _, subImage := range subImages {
	// 	fmt.Println(subImage.Filename)
	// }
	type NameProd struct {
		NameProd string `json:"name_prod"`
	}
	var nameProd NameProd
	err := c.BodyParser(&nameProd.NameProd)
	if err != nil {
		return c.SendStatus(fiber.ErrBadGateway.Code)
	}
	fmt.Printf("name prod = %s", nameProd.NameProd)
	product.UserID = int(c.Locals("user").(models.User).UserId)
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := h.productService.CreateProduct(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "create product success",
	})
}

func (h *ProductHandler) GetProduct(c *fiber.Ctx) error {
	var productFilter models.ProductFilter
	page, err := strconv.Atoi(c.Params("page"))
	userId := c.Locals("user").(models.User).UserId
	if err != nil {
		return err
	}
	productFilter.ProductID = c.Query("product_id")
	productFilter.ProductName = c.Query("product_name")
	productFilter.MainCategory, err = strconv.Atoi(c.Query("main_cate"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	productFilter.SubCategory, err = strconv.Atoi(c.Query("sub_cate"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	results, err := h.productService.GetProduct(productFilter, page, userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": results,
	})
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := h.productService.DeleteProductById(product.ProductID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "delete product success",
	})
}
