package handlers

import (
	"github.com/firerplayer/stash-task/backend/internal/infra/web"
	"github.com/firerplayer/stash-task/backend/internal/usecase/dto"
	usecase "github.com/firerplayer/stash-task/backend/internal/usecase/user"
	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	CreateUserUseCase     usecase.CreateUserUseCase
	ListAllUsersUseCase   usecase.ListAllUsersUseCase
	GetUserByIdUseCase    usecase.GetUserByIDUseCase
	GetUserByEmailUseCase usecase.GetUserByEmailUseCase
	UpdateUserByIDUseCase usecase.UpdateUserByIDUseCase
	DeleteUserByIDUseCase usecase.DeleteUserByIDUseCase
	LoginUseCase          usecase.LoginUseCase
}

func NewUserHandlers(
	createUserUseCase usecase.CreateUserUseCase,
	listAllUsersUseCase usecase.ListAllUsersUseCase,
	getUserByIdUseCase usecase.GetUserByIDUseCase,
	getUserByEmailUseCase usecase.GetUserByEmailUseCase,
	updateUserByIDUseCase usecase.UpdateUserByIDUseCase,
	deleteUserByIDUseCase usecase.DeleteUserByIDUseCase,
	loginUseCase usecase.LoginUseCase,
) *UserHandlers {
	return &UserHandlers{
		CreateUserUseCase:     createUserUseCase,
		ListAllUsersUseCase:   listAllUsersUseCase,
		GetUserByIdUseCase:    getUserByIdUseCase,
		GetUserByEmailUseCase: getUserByEmailUseCase,
		UpdateUserByIDUseCase: updateUserByIDUseCase,
		DeleteUserByIDUseCase: deleteUserByIDUseCase,
		LoginUseCase:          loginUseCase,
	}
}

func (u *UserHandlers) RegisterRoutes(server web.WebServer) {
	server.Post("/users", u.CreateUser)
	server.Post("/users/login", u.Login)

	//server.Protected().Get("/users/all", u.GetAllLimitUsers)
	server.Protected().Get("/users", u.GetByID)
	server.Protected().Patch("/users", u.UpdateByID)
	server.Protected().Delete("/users", u.DeleteByID)
}

func (u *UserHandlers) GetByID(c *fiber.Ctx) error {

	id := GetUserIDFromCtx(c)
	usrOut, err := u.GetUserByIdUseCase.Execute(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResponse(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(usrOut)
}

func (u *UserHandlers) GetAllLimitUsers(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 20)
	offset := c.QueryInt("offset", 0)
	usrOut, err := u.ListAllUsersUseCase.Execute(c.Context(), dto.ListAllUsersInputDTO{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResponse(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(usrOut)
}

func (u *UserHandlers) CreateUser(c *fiber.Ctx) error {
	var input dto.CreateUserInputDTO
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.SystemResponse{
			Message: err.Error(),
		})
	}
	usrOut, err := u.CreateUserUseCase.Execute(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResponse(err.Error()))
	}
	return c.Status(fiber.StatusCreated).JSON(usrOut)
}

func (u *UserHandlers) DeleteByID(c *fiber.Ctx) error {
	//id := c.Params("id", "invalid")
	//if id == "invalid" {
	//	return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse("id is required"))
	//}
	err := u.DeleteUserByIDUseCase.Execute(c.Context(), dto.DeleteUserByIDInputDTO{
		ID: GetUserIDFromCtx(c),
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResponse(err.Error()))
	}
	return c.SendStatus(fiber.StatusOK)
}

func (u *UserHandlers) UpdateByID(c *fiber.Ctx) error {
	userID := GetUserIDFromCtx(c)
	var input dto.UpdateUserInputDTO
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse(err.Error()))
	}
	input.UserID = userID
	err = u.UpdateUserByIDUseCase.Execute(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResponse(err.Error()))
	}
	return c.SendStatus(fiber.StatusOK)

}

func (u *UserHandlers) Login(c *fiber.Ctx) error {
	var input dto.LoginInputDTO
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse(err.Error()))
	}

	out, err := u.LoginUseCase.Execute(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResponse(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(out)

}
