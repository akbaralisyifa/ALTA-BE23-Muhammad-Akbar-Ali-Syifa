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
	// connection.AutoMigrate(&models.User{}, &models.Todo{});

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
				fmt.Print("input Todo")
				fmt.Scanln(&inputTodo)

				if inputTodo == 9 {
					isLogin = false;
				}else if inputTodo == 1{
					_, err := todoController.AddTodo(data.ID);
					if err != nil {
						fmt.Println("Failed to add todo", err);
						return
					}
					fmt.Println("Success Add todo")

				}else if inputTodo == 2 {
					_, err := todoController.UpdateTodo(data.ID);
					if err != nil {
						fmt.Println("Failed to update todo", err.Error())
						return
					}
					fmt.Println("Success update todo.")

				}else if inputTodo == 3 {
					_, err := todoController.DeleteTodo(data.ID);
					if err != nil {
						fmt.Println("Failed To Delete Todo");
					}
				}else if inputTodo == 4 {
					data, err := todoController.FindTodos(data.ID);
					if err != nil {
						fmt.Println("Failed to find todos");
						return
					}
					fmt.Println("Success to find todos")
					fmt.Println(data)
				}
			}

		}else if inputMenu == 2 {
			userController.Register();
		}

	}


}