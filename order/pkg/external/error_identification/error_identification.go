package error_identification

import (
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"strconv"
)

type (
	ErrorIdentification struct {
	}

	ErrorIdentificationInterface interface {
		createId() string
		addToContext(ctx *fiber.Ctx)
		ResignInMiddleware(c *fiber.Ctx) error // resign with middleware in frame
	}
)

const NameFields = "error_identification_debug"

func (e *ErrorIdentification) createId() string {
	return strconv.FormatUint(rand.Uint64(), 10)
}

func (e *ErrorIdentification) addToContext(ctx *fiber.Ctx) {
	ctx.Locals(NameFields, e.createId())
}

func (e *ErrorIdentification) ResignInMiddleware(ctx *fiber.Ctx) error {
	e.addToContext(ctx)
	return ctx.Next()
}

func GetErrIdFrCtx(ctx *fiber.Ctx) string {
	value := ctx.Locals(NameFields).(string)
	return value
}

func NewErrorIdentification() ErrorIdentificationInterface {
	return &ErrorIdentification{}
}
