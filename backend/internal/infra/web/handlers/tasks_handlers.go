package handlers

import (
	"github.com/firerplayer/stash-task/backend/internal/infra/web"
	"github.com/firerplayer/stash-task/backend/internal/usecase/dto"
	tu "github.com/firerplayer/stash-task/backend/internal/usecase/task"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type TasksHandlers struct {
	CreateTaskUseCase      tu.CreateTaskUseCase
	GetTaskByIDUseCase     tu.GetTaskByIDUseCase
	ListAllTasksUseCase    tu.ListAllTasksUseCase
	ListTasksByUserUseCase tu.ListTasksByUser
	DeleteTaskUseCase      tu.DeleteTaskUseCase
	UpdateTaskUseCase      tu.UpdateTaskUseCase
	UncompleteTaskUseCase  tu.UncompleteTaskUseCase
	CompleteTaskUseCase    tu.CompletTaskUseCase
}

func NewTasksHandlers(
	createTaskUseCase tu.CreateTaskUseCase,
	getTaskByIDUseCase tu.GetTaskByIDUseCase,
	listAllTasksUseCase tu.ListAllTasksUseCase,
	listTasksByUserUseCase tu.ListTasksByUser,
	deleteTaskUseCase tu.DeleteTaskUseCase,
	updateTaskUseCase tu.UpdateTaskUseCase,
	uncompleteTaskUseCase tu.UncompleteTaskUseCase,
	completeTaskUseCase tu.CompletTaskUseCase,
) *TasksHandlers {
	return &TasksHandlers{
		CreateTaskUseCase:      createTaskUseCase,
		GetTaskByIDUseCase:     getTaskByIDUseCase,
		ListAllTasksUseCase:    listAllTasksUseCase,
		ListTasksByUserUseCase: listTasksByUserUseCase,
		DeleteTaskUseCase:      deleteTaskUseCase,
		UpdateTaskUseCase:      updateTaskUseCase,
		UncompleteTaskUseCase:  uncompleteTaskUseCase,
		CompleteTaskUseCase:    completeTaskUseCase,
	}
}

func (t *TasksHandlers) RegisterRoutes(server web.WebServer) {
	server.Protected().Post("/tasks", t.CreateTask)
	server.Protected().Get("/tasks/view/:id", t.GetTaskByID)
	server.Protected().Put("/tasks/:id", t.UpdateTask)
	server.Protected().Delete("/tasks/:id", t.DeleteTask)
	server.Protected().Get("/tasks", t.ListAllTasks)
	server.Protected().Get("/tasks/user", t.ListTasksByUser)
	server.Protected().Patch("/tasks/complete/:id", t.CompleteTask)
	server.Protected().Patch("/tasks/uncomplete/:id", t.UncompleteTask)
}

func (t *TasksHandlers) CreateTask(c *fiber.Ctx) error {
	var input dto.CreateTaskInputDTO
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse(err.Error()))
	}
	input.UserID = GetUserIDFromCtx(c)
	usrOut, err := t.CreateTaskUseCase.Execute(c.Context(), input)
	return c.Status(fiber.StatusCreated).JSON(usrOut)
}

func (t *TasksHandlers) GetTaskByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse("id is required"))
	}
	usrOut, err := t.GetTaskByIDUseCase.Execute(c.Context(), dto.GetTaskByIDInputDTO{
		ID: id,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResponse(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(usrOut)
}

func (t *TasksHandlers) ListAllTasks(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 20)
	offset := c.QueryInt("offset", 0)
	out, err := t.ListAllTasksUseCase.Execute(c.Context(), dto.ListAllTasksInputDTO{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResponse(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(out)
}

func (t *TasksHandlers) ListTasksByUser(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 20)
	offset := c.QueryInt("offset", 0)
	tsk, err := t.ListTasksByUserUseCase.Execute(c.Context(), dto.ListAllTasksByUserIDInputDTO{
		UserID: GetUserIDFromCtx(c),
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResponse(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(tsk)

}

func (t *TasksHandlers) DeleteTask(c *fiber.Ctx) error {

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse("id is required"))
	}
	err := t.DeleteTaskUseCase.Execute(c.Context(), dto.DeleteTaskInputDTO{ID: id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResponse(err.Error()))
	}
	return c.SendStatus(fiber.StatusOK)
}

func (t *TasksHandlers) UpdateTask(c *fiber.Ctx) error {

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse("id is required"))
	}
	var input dto.UpdateTaskInputDTO
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse(err.Error()))
	}
	input.ID = id
	err = t.UpdateTaskUseCase.Execute(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResponse(err.Error()))
	}
	return c.SendStatus(fiber.StatusOK)
}

func (t *TasksHandlers) CompleteTask(c *fiber.Ctx) error {

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse("id is required"))
	}
	err := t.CompleteTaskUseCase.Execute(c.Context(), dto.CompleteTaskInputDTO{ID: id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResponse(err.Error()))
	}
	return c.SendStatus(fiber.StatusOK)
}

func (t *TasksHandlers) UncompleteTask(c *fiber.Ctx) error {

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResponse("id is required"))
	}
	err := t.UncompleteTaskUseCase.Execute(c.Context(), dto.UncompleteTaskInputDTO{ID: id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResponse(err.Error()))
	}
	return c.SendStatus(fiber.StatusOK)
}

func GetUserIDFromCtx(c *fiber.Ctx) string {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)

}
