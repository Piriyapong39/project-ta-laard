package handler

import (
	"fmt"
	"product-service/internal/models"
	"product-service/internal/service"
	"product-service/internal/utils"
	"strings"

	"product-service/internal/middlewares"

	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	product.Delete("/", handler.DeleteProduct)
	product.Patch("/inactive", handler.InactivateProductById)
	product.Patch("/active", handler.ActivateProduct)
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	product.UserID = int(c.Locals("user").(models.User).UserId)
	uuid := uuid.New()
	product.ProductID = uuid.String()
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println(err)
	}

	// ----------------------------------------------- Bug start here -----------------------------------------------------
	// upload main picture
	mainImage, err := c.FormFile("mainImage")
	if err != nil {
		fmt.Println(err)
	}
	mainImageSplit := strings.Split(mainImage.Filename, ".")
	newMainImageName := "main." + mainImageSplit[len(mainImageSplit)-1]
	mainImage.Filename = newMainImageName
	picPath, err := utils.UploadPicture(mainImage, c, product.ProductID)
	if err != nil {
		fmt.Printf("error: %s", err)
	}

	product.ProductImage = append(product.ProductImage, picPath)
	// upload sub image
	subImages := form.File["subImage"]
	for i, subImage := range subImages {
		subImageSplit := strings.Split(subImage.Filename, ".")
		extName := subImageSplit[len(subImageSplit)-1]
		extNameToLowerCase := strings.ToLower(extName)
		subImage.Filename = strconv.Itoa(i+1) + "." + extNameToLowerCase

		picPath, err := utils.UploadPicture(subImage, c, product.ProductID)
		if err != nil {
			fmt.Printf("error: %s", err)
		}
		product.ProductImage = append(product.ProductImage, picPath)
	}

	// ----------------------------------------------- Bug end here -----------------------------------------------------

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
	page, _ := strconv.Atoi(c.Params("page"))
	userId := c.Locals("user").(models.User).UserId
	productFilter.ProductID = c.Query("productId")
	productFilter.ProductName = c.Query("productName")
	productFilter.MainCategory, _ = strconv.Atoi(c.Query("mainCategory"))
	productFilter.SubCategory, _ = strconv.Atoi(c.Query("subCategory"))

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
	userId := c.Locals("user").(models.User).UserId
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := h.productService.DeleteProductById(product.ProductID, userId); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "delete product success",
	})
}

func (h *ProductHandler) InactivateProductById(c *fiber.Ctx) error {
	var product models.Product
	userId := c.Locals("user").(models.User).UserId
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.productService.InactivateProductById(product.ProductID, userId); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "inactive product success",
	})
}

func (h *ProductHandler) ActivateProduct(c *fiber.Ctx) error {

	var product models.Product
	userId := c.Locals("user").(models.User).UserId
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := h.productService.ActivateProduct(product.ProductID, userId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "active product success",
	})
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	userId := c.Locals("user").(models.User).UserId
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := h.productService.UpdateProduct(product, userId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return nil
}
