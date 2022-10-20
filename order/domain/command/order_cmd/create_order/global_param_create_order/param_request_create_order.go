package global_param_create_order

import "github.com/gofiber/fiber/v2"

type (
	ParamCreateOrder struct {
		Ctx ContextRequest
	}
	ContextRequest *fiber.Ctx
)
