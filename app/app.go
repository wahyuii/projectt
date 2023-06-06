package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"projectt/app/context"
	"projectt/app/database/repository"
	usersRepository "projectt/app/database/repository/users"
	usersResolver "projectt/app/graphql/resolver/users"
	"projectt/app/graphql/schema"
	"projectt/app/utils"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

type App struct{}

// Initialize function
func (a *App) Initialize() {
	fmt.Println("testing database...")

	tapDB, _ := utils.ConnectDB()

	if tapDB != nil {
		fmt.Println("testing database success...")
		transact := repository.NewTransact(tapDB)

		// cek
		usersRepository := usersRepository.NewUsersRepository(tapDB)
		usersResolver := usersResolver.NewUsersResolver(tapDB, transact, usersRepository)

		graphqlSchema := schema.InitGraphQLSchema()
		gqlHandler := context.ApplyContextMiddleware(
			handler.New(&handler.Config{
				Schema:     &graphqlSchema,
				Pretty:     true,
				Playground: true,
			}),
			usersResolver,
		)

		appHandler := context.ApplyAuthenticationMiddleware(gqlHandler)
		server := gin.Default()

		server.GET("/", appHandler)
		server.POST("/", appHandler)
		server.POST("/query", appHandler)
		server.POST("/mutation", appHandler)
		server.POST("/login", appHandler)
		server.OPTIONS("/", appHandler)
		server.OPTIONS("/query", appHandler)
		server.OPTIONS("/mutation", appHandler)
		server.OPTIONS("/login", appHandler)
		log.Fatal(server.Run(":" + os.Getenv("ACTIVE_PORT")))
	}
}

// Run Function
func (a *App) Run() {
	port := ":" + os.Getenv("ACTIVE_PORT")
	fmt.Println("GQL SERVICE RUN AT", port)
	fmt.Println(http.ListenAndServe(port, nil))
}
