package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:arasy@ci@tcp(127.0.0.1:3308)/perkemi_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.CheckEmailAvailability)
	router.Run()

	// router.GET("/handler", handler)
	// router.Run(":8080")

	// userInput := user.RegisterUserInput{}
	// userInput.Name = "Tes simpan dari service"
	// userInput.Email = "contoh@gmail.com"
	// userInput.Occupation = "programmer"
	// userInput.Password = "oyoyoy"

	// userService.RegisterUser(userInput)

	// user := user.User{
	// 	Name:             "Test simpan",
	// 	Occupation:       "TEST",
	// 	Email:            "TEST",
	// 	Password_hash:    "TESTJUGA",
	// 	Avatar_file_name: "TESTJG",
	// 	Role:             "TESTJG",
	// }
	// userRepository.Save(user)
	// var users []user.User
	// length := len(users)
	// fmt.Println(length)

	// db.Find(&users)

	// length = len(users)
	// fmt.Println(length)

	// for _, user := range users {
	// 	fmt.Println(user.Name)
	// 	fmt.Println(user.Email)
	// 	fmt.Println(user.Created_at)
	// 	fmt.Println(user.Updated_at)

	// }
}

// func handler(c *gin.Context) {
// 	dsn := "root:arasy@ci@tcp(127.0.0.1:3308)/perkemi_db?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	var users []user.User
// 	db.Find(&users)

// 	c.JSON(http.StatusOK, users)

// 	//INPUT
// 	//HANDLER (MENANGKAP DARI INPUT USER, MAPPING INPUTdari user  KE STRUCT input)
// 	//SERVICE (MAPPING KE STRUCT User)
// 	//=REPOSITORY (SAVE STRUCT User ke DB)
// 	//DB
// }
