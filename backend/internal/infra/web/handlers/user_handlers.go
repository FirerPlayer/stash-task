package handlers

import (
	"github.com/firerplayer/stash-task/backend/internal/infra/web"
	"github.com/firerplayer/stash-task/backend/internal/usecase/dto"
	usecase "github.com/firerplayer/stash-task/backend/internal/usecase/user"
	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	CreateUserUseCase      usecase.CreateUserUseCase
	ListAllUsersUseCase    usecase.ListAllUsersUseCase
	FindUserByIdUseCase    usecase.GetUserByIDUseCase
	FindUserByEmailUseCase usecase.FindUserByEmailUseCase
	UpdateUserByIDUseCase  usecase.UpdateUserByIDUseCase
	DeleteUserByIDUseCase  usecase.DeleteUserByIDUseCase
	LoginUseCase           usecase.LoginUseCase
}

func NewUserHandlers(
	createUserUseCase usecase.CreateUserUseCase,
	listAllUsersUseCase usecase.ListAllUsersUseCase,
	findUserByIdUseCase usecase.GetUserByIDUseCase,
	findUserByEmailUseCase usecase.FindUserByEmailUseCase,
	updateUserByIDUseCase usecase.UpdateUserByIDUseCase,
	deleteUserByIDUseCase usecase.DeleteUserByIDUseCase,
	loginUseCase usecase.LoginUseCase,
) *UserHandlers {
	return &UserHandlers{
		CreateUserUseCase:      createUserUseCase,
		ListAllUsersUseCase:    listAllUsersUseCase,
		FindUserByIdUseCase:    findUserByIdUseCase,
		FindUserByEmailUseCase: findUserByEmailUseCase,
		UpdateUserByIDUseCase:  updateUserByIDUseCase,
		DeleteUserByIDUseCase:  deleteUserByIDUseCase,
		LoginUseCase:           loginUseCase,
	}
}

func (u *UserHandlers) RegisterRoutes(server web.WebServer) {
	server.Post("/users", u.CreateUser)
	server.Post("/users/login", u.Login)

	//server.Protected().Get("/users", u.GetAllLimitUsers)
	server.Protected().Get("/users/:id", u.GetByID)
	server.Protected().Patch("/users/:id", u.UpdateByID)
	server.Protected().Delete("/users/:id", u.DeleteByID)
}

func (u *UserHandlers) GetByID(c *fiber.Ctx) error {
	id := c.Params("id", "invalid")
	if id == "invalid" {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse("id is required"))
	}
	usrOut, err := u.FindUserByIdUseCase.Execute(c.Context(), id)
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
	id := c.Params("id", "invalid")
	if id == "invalid" {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse("id is required"))
	}
	err := u.DeleteUserByIDUseCase.Execute(c.Context(), dto.DeleteUserByIDInputDTO{
		ID: id,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResponse(err.Error()))
	}
	return c.SendStatus(fiber.StatusOK)
}

func (u *UserHandlers) UpdateByID(c *fiber.Ctx) error {
	id := c.Params("id", "invalid")
	if id == "invalid" {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse("id is required"))
	}
	var input dto.UpdateUserInputDTO
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse(err.Error()))
	}
	err = u.UpdateUserByIDUseCase.Execute(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResponse(err.Error()))
	}
	return c.SendStatus(fiber.StatusOK)

}

func (u *UserHandlers) Login(c *fiber.Ctx) error {
	var input dto.LoginInputDTO
	c.MultipartForm()
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
