package httpx

import "github.com/gofiber/fiber/v2"

// Pagination represents pagination metadata
type Pagination struct {
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
	TotalItems int  `json:"total_items"`
	TotalPages int  `json:"total_pages"`
	HasNext    bool `json:"has_next"`
	HasPrev    bool `json:"has_prev"`
}

// APIResponse is a generic structure for JSON API responses
type APIResponse[T any] struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message,omitempty"`
	Data       T           `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// -----------------------------------------------------------------------------
// 🔹 Convenience Constructors
// -----------------------------------------------------------------------------

// Success creates a successful response with data.
func Success[T any](data T) APIResponse[T] {
	return APIResponse[T]{
		Success: true,
		Data:    data,
	}
}

// SuccessWithMessage creates a successful response with data and a message.
func SuccessWithMessage[T any](data T, message string) APIResponse[T] {
	return APIResponse[T]{
		Success: true,
		Message: message,
		Data:    data,
	}
}

// SuccessWithPagination creates a paginated successful response.
func SuccessWithPagination[T any](data T, pagination *Pagination) APIResponse[T] {
	return APIResponse[T]{
		Success:    true,
		Data:       data,
		Pagination: pagination,
	}
}

// Error creates an error response with a message.
func Error(message string) APIResponse[any] {
	return APIResponse[any]{
		Success: false,
		Error:   message,
	}
}

// -----------------------------------------------------------------------------
// 🔹 Fiber JSON Helpers
// -----------------------------------------------------------------------------

func JSON[T any](c *fiber.Ctx, status int, data T) error {
	return c.Status(status).JSON(Success(data))
}

func JSONWithMessage[T any](c *fiber.Ctx, status int, data T, message string) error {
	return c.Status(status).JSON(SuccessWithMessage(data, message))
}

func JSONPaginated[T any](c *fiber.Ctx, status int, data T, pagination *Pagination) error {
	return c.Status(status).JSON(SuccessWithPagination(data, pagination))
}

// -----------------------------------------------------------------------------
// 🔹 Error Response Shortcuts
// -----------------------------------------------------------------------------

func BadRequest(c *fiber.Ctx, msg string) error {
	return c.Status(fiber.StatusBadRequest).JSON(Error(msg))
}

func Unauthorized(c *fiber.Ctx, msg string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(Error(msg))
}

func Forbidden(c *fiber.Ctx, msg string) error {
	return c.Status(fiber.StatusForbidden).JSON(Error(msg))
}

func NotFound(c *fiber.Ctx, msg string) error {
	return c.Status(fiber.StatusNotFound).JSON(Error(msg))
}

func Conflict(c *fiber.Ctx, msg string) error {
	return c.Status(fiber.StatusConflict).JSON(Error(msg))
}

func InternalServerError(c *fiber.Ctx, msg string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(Error(msg))
}

func UnprocessableEntity(c *fiber.Ctx, msg string) error {
	return c.Status(fiber.StatusUnprocessableEntity).JSON(Error(msg))
}

func NoContent(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}
