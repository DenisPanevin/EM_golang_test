package server

import (
	"EM-Api-testTask/internal/jobs"
	jobsDelivery "EM-Api-testTask/internal/jobs/delivery"
	jobsRepository "EM-Api-testTask/internal/jobs/repository/psql"
	jobsUseCase "EM-Api-testTask/internal/jobs/usecase"
	"EM-Api-testTask/internal/tasks"
	tasksDelivery "EM-Api-testTask/internal/tasks/delivery"
	tasksRepoisitory "EM-Api-testTask/internal/tasks/repository/psql"
	tasksUseCase "EM-Api-testTask/internal/tasks/usecase"
	"EM-Api-testTask/internal/users"
	usersDelivery "EM-Api-testTask/internal/users/delivery"
	usersRepository "EM-Api-testTask/internal/users/repository/psql"
	usersUsecase "EM-Api-testTask/internal/users/usecase"
	"EM-Api-testTask/pkg"
	"github.com/asaskevich/govalidator"
	"github.com/kpango/glg"
	"os"
	"os/signal"
	"strconv"

	"EM-Api-testTask/config"

	"context"
	"github.com/gorilla/mux"
	"net/http"

	"time"
)

type App struct {
	Config *config.Config
	//	Db         *pgxpool.Pool
	httpServer *http.Server
	UserUc     users.UseCase
	TasksUc    tasks.UseCase
	JobsUc     jobs.UseCase
}

func NewApp(config_path string) (*App, error) {
	cfg, err := config.NewConfig(config_path)
	if err != nil {
		glg.Fatal(err)
		return nil, err
	}

	Db := pkg.InitDb((*cfg).Psql_connection)

	//userRepo := authmongo.NewUserRepository(db, viper.GetString("mongo.user_collection"))
	//bookmarkRepo := bmmongo.NewBookmarkRepository(db, viper.GetString("mongo.bookmark_collection"))
	ur := usersRepository.NewUsersRepo(Db)
	tr := tasksRepoisitory.NewTasksRepo(Db)
	jr := jobsRepository.NewJobsRepo(Db)

	return &App{
		Config: cfg,
		//	Db:     Db,
		UserUc:  usersUsecase.NewUserUseCase(ur),
		TasksUc: tasksUseCase.NewTasksUseCase(tr),
		JobsUc:  jobsUseCase.NewJobsUseCase(jr),

		/*bookmarkUC: bmusecase.NewBookmarkUseCase(bookmarkRepo),
		authUC: authusecase.NewAuthUseCase(
			userRepo,
			viper.GetString("auth.hash_salt"),
			[]byte(viper.GetString("auth.signing_key")),
			viper.GetDuration("auth.token_ttl"),
		),*/
	}, nil
}

func (a *App) Run() error {

	govalidator.SetFieldsRequiredByDefault(true)
	router := mux.NewRouter()
	//validators.ValidatorInit()
	usersDelivery.RegisterUserRoutes(router, a.UserUc)
	tasksDelivery.RegisterTaskRoutes(router, a.TasksUc)
	jobsDelivery.RegisterJobsRoutes(router, a.JobsUc)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + strconv.Itoa(a.Config.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			glg.Fatal("Failed to listen and serve: %+v", err)
		}
	}()
	glg.Info("Server running on port ", a.Config.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
