package main

import (
	"fmt"
	"try/config"
	"try/internal/controller"
	"try/internal/models"
)

func main() {
	setup := config.ImportSetting();
	connection , err := config.ConnectDB(setup);

	if err != nil{
		fmt.Println("s Conenncting to database", err.Error())
		return
	};
	connection.AutoMigrate(&models.User{}, &models.Todo{});

	userModel := models.NewUserModel(connection);
	userController := controller.NewUserController(userModel);

	todoModel := models.NewTodoModel(connection);
	todoController := controller.NewTodoController(todoModel);

	var inputMenu int;

	if inputMenu != 9 {
		fmt.Println("1. Login");
		fmt.Println("2. Register");
		fmt.Println("9. Keluar");
		fmt.Println("Input Menu : ");
		fmt.Scanln(&inputMenu);

		if inputMenu == 1 {
			var inputTodo int
			var isLogin = true;
			data, err := userController.Login();
			
			if err != nil {
				panic(err.Error())
			}

			if isLogin {
				fmt.Println("welcome to todo", data.Name);
				fmt.Println("Pilih Todo");
				fmt.Println("1. Add Todo")
				fmt.Println("2. Update Todo")
				fmt.Println("3. Delete Todo")
				fmt.Println("4. View Todo")
				fmt.Println("9. Exit Todo")

				if inputTodo == 9 {
					isLogin = false;
				}else if inputTodo == 1{
					todoController.AddTodo(data.ID);
				}else if inputTodo == 2 {
					todoController.UpdateTodo(data.ID);
				}else if inputTodo == 3 {
					todoController.DeleteTodo(data.ID)
				}else if inputTodo == 4 {
					todoController.GetTodo();
				}
			}

		}else if inputMenu == 2 {
			userController.Register();
		}

	}


}