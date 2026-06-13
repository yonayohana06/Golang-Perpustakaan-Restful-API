package main

import (
	// "os"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http"

	userService "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/application/usecase/user"
	userController "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/user"
	userRepository "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/repository/mysql/user"

	bukuService "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/application/usecase/buku"
	bukuController "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/buku"
	bukuRepository "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/repository/mysql/buku"

	authService "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/application/usecase/auth"
	authController "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/auth"
	authRepository "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/repository/mysql/auth"

	peminjamanService "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/application/usecase/peminjaman"
	peminjamanController "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/peminjaman"
	peminjamanRepository "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/repository/mysql/peminjaman"

	dendaService "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/application/usecase/denda"
	dendaController "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/denda"
	dendaRepository "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/repository/mysql/denda"

	anggotaUseCase "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/application/usecase/anggota"
	anggotaHandler "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/anggota"
	anggotaRepository "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/repository/mysql/anggota"

	conf "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/config"
	"github.com/gofiber/fiber/v2"

	// "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/database"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/middleware"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/utility"
)

func main() {
	db := conf.MySQLConn()
	// database.Migrate()
	// database.Seeder(db)

	// Init JWT Config
	jwtConfig := conf.GetJwtConfig()
	jwtUtility := utility.NewJWTUtility(jwtConfig)

	// User module
	userRepo := userRepository.NewUserRepository(db)
	userServices := userService.NewUserService(userRepo)
	userCon := userController.NewUserController(userServices)

	// Buku module
	bukuRepo := bukuRepository.NewBukuRepository(db)
	bukuServices := bukuService.NewBukuService(bukuRepo)
	bukuCon := bukuController.NewBukuController(bukuServices)

	// Auth module
	authRepo := authRepository.NewAuthRepository(db)
	authService := authService.NewAuthService(authRepo, jwtUtility)
	authController := authController.NewAuthController(authService)

	// Peminjaman module
	peminjamanRepo := peminjamanRepository.NewPeminjamanRepository(db)
	peminjamanServices := peminjamanService.NewPeminjamanService(peminjamanRepo)
	peminjamanCon := peminjamanController.NewPeminjamanController(peminjamanServices)

	// Denda module
	dendaRepo := dendaRepository.NewDendaRepository(db)
	dendaServices := dendaService.NewDendaService(dendaRepo)
	dendaCon := dendaController.NewDendaController(dendaServices)

	// 1. Inisiasi (Dependency Injection) - Anggota
	anggotaRepo := anggotaRepository.NewAnggotaRepository(db) // db adalah instance *gorm.DB Anda
	anggotaUC := anggotaUseCase.NewAnggotaUseCase(anggotaRepo)
	anggotaHan := anggotaHandler.NewAnggotaHandler(anggotaUC)

	// Middleware
	jwtAuthMiddleware := middleware.NewJWTAuthMiddleware(jwtUtility, jwtConfig)

	appConfig := conf.ServerTimeOut()
	app := fiber.New(appConfig)

	http.RegisterPath(
		app,
		userCon,
		bukuCon,
		authController,
		peminjamanCon,
		dendaCon,
		anggotaHan,
		jwtAuthMiddleware,
	)

	// port := os.Getenv("PORT")

	// if port == "" {
	// 	port = "8001"
	// }

	// app.Listen(":" + port)

	app.Listen(":8001")

}
