package pkg

import (
	"trendyol/internal/dto"
	"trendyol/internal/handler"
)

const (
	CommandAddItem          = "addItem"
	CommandAddVasItemToItem = "addVasItemToItem"
	CommandRemoveItem       = "removeItem"
	CommandResetCart        = "resetCart"
	CommandDisplayCart      = "displayCart"
	CommandUnknown          = "unknown"
)

type CommandExecutor struct {
	handler handler.CartHandler
}

func NewCommandExecutor(h handler.CartHandler) *CommandExecutor {
	return &CommandExecutor{
		handler: h,
	}
}

func (c *CommandExecutor) Run(requests []dto.Request) []dto.Response {
	var responses []dto.Response

	for _, request := range requests {
		var response dto.Response
		switch request.Command {
		case CommandAddItem:
			response = c.handler.AddItemToCart(request.Payload)
		case CommandAddVasItemToItem:
			response = c.handler.AddVasItemToItem(request.Payload)
		case CommandRemoveItem:
			response = c.handler.RemoveItem(request.Payload)
		case CommandResetCart:
			response = c.handler.Reset()
		case CommandDisplayCart:
			response = c.handler.Display()
		default:
			response = dto.Response{Result: false, Message: CommandUnknown}
		}
		responses = append(responses, response)
	}
	return responses
}
