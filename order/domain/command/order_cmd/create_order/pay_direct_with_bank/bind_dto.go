package pay_direct_with_bank

import (
	"github.com/gofiber/fiber/v2"
)

func BindDataDtoRequestCreateOrder(c *fiber.Ctx) (ParamRequestDto, error) {
	var o ParamRequestDto
	if errBP := c.BodyParser(&o); errBP != nil {
		return ParamRequestDto{}, errBP
	}

	return o, nil
}
