package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/firerplayer/stash-task/backend/configs"
	"github.com/firerplayer/stash-task/backend/internal/domain/gateway"
	"github.com/firerplayer/stash-task/backend/internal/infra/repository"
	"github.com/firerplayer/stash-task/backend/internal/infra/web"
	"github.com/firerplayer/stash-task/backend/internal/infra/web/handlers"
	ucTask "github.com/firerplayer/stash-task/backend/internal/usecase/task"
	ucUser "github.com/firerplayer/stash-task/backend/internal/usecase/user"
	//import pg driver
	pg "github.com/jackc/pgx/v5"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	connPg, err := pg.Connect(context.Background(), fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.DBDriver,
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	),
	)
	if err != nil {
		panic(errors.New("Error connecting to database --> " + err.Error()))
	}
	//conn, err := sql.Open(config.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true",
	//	config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName))
	//if err != nil {
	//	panic(err)
	//}
	//defer func(conn *sql.DB) {
	//	e := conn.Close()
	//	if e != nil {
	//		_ = fmt.Errorf(e.Error())
	//	}
	//}(conn)

	// Configuração do webserver e rotas
	webServer := web.NewWebServer(config.WebServerPort, "Stash Task", config.JwtSecret)
	repoUserPg := repository.NewUserRepositoryPg(connPg)
	repoTaskPg := repository.NewTaskRepositoryPg(connPg)

	setupUsersHandlers(*webServer, repoUserPg, config.JwtSecret)
	setupTaskHandlers(*webServer, repoTaskPg)
	fmt.Println("Server running on port: ", webServer.WebServerPort)

	if err = webServer.Start(); err != nil {
		panic(err)
	}
}

func setupUsersHandlers(webserver web.WebServer, repoUserPg gateway.UsersGateway, jwtSecretKey string) {
	// Users
	createUserUseCase := ucUser.NewCreateUserUseCase(repoUserPg)
	deleteUserByIdUseCase := ucUser.NewDeleteUserByIDUseCase(repoUserPg)
	findUserByIdUseCase := ucUser.NewGetUserByIdUseCase(repoUserPg)
	findUserByEmailUseCase := ucUser.NewFindUserByEmailUseCase(repoUserPg)
	updateUserByIdUseCase := ucUser.NewUpdateUserUseCase(repoUserPg)
	listAllUsersUseCase := ucUser.NewListAllUsersUseCase(repoUserPg)
	loginUserUseCase := ucUser.NewLoginUseCase(repoUserPg, jwtSecretKey)

	//usersHandlers
	userHandlers := handlers.NewUserHandlers(
		*createUserUseCase,
		*listAllUsersUseCase,
		*findUserByIdUseCase,
		*findUserByEmailUseCase,
		*updateUserByIdUseCase,
		*deleteUserByIdUseCase,
		*loginUserUseCase,
	)
	userHandlers.RegisterRoutes(webserver)
}

func setupTaskHandlers(webserver web.WebServer, repoTaskPg gateway.TasksGateway) {
	taskHandlers := handlers.NewTasksHandlers(
		*ucTask.NewCreateTaskUseCase(repoTaskPg),
		*ucTask.NewGetTaskByIDUseCase(repoTaskPg),
		*ucTask.NewListAllTasksUseCase(repoTaskPg),
		*ucTask.NewListTasksByUser(repoTaskPg),
		*ucTask.NewDeleteTaskUseCase(repoTaskPg),
		*ucTask.NewUpdateTaskUseCase(repoTaskPg),
		*ucTask.NewUncompleteTaskUseCase(repoTaskPg),
		*ucTask.NewCompletTaskUseCase(repoTaskPg),
	)
	taskHandlers.RegisterRoutes(webserver)
}
