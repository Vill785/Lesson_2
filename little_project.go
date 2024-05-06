package main

import "fmt"

const (
	exit        = "exit"
	auth        = "auth"
	reg         = "reg"
	add_product = "add_product"
	order       = "order"
)

func main() {
	var command string
	userList := []string{"user1_password1", "user2_password2"}
	productList := make([]string, 0, 10)
	for command != exit {
		fmt.Println("Введите команду") // Сделать красивый вывод, вывести список команд на этом шаге
		fmt.Scan(&command)
		switch command {
		case exit:
			break
		case reg:
			fmt.Println("Введите логин и пароль в таком виде login_password")
			fmt.Scan(&command) // Сделать так, чтобы выводил сообщение если пользователь уже существует
			userList = append(userList, command)
			message := fmt.Sprintf("Пользователь %s успешно добавлен", command)
			fmt.Println(message)
		case auth:
			fmt.Println("Введите логин и пароль в таком виде login_password")
			fmt.Scan(&command) // Сделать так, чтобы выводил сообщение если пользователь уже существует
			for _, v := range userList {
				if v == command {
					fmt.Println("Добро пожаловать в магазин")
				} else {
					fmt.Println("Вы не зарегистрированы")
				}
			}
		case add_product:
			fmt.Println("Добавьте продукт в product_list")
			fmt.Scan(&command)
			productList = append(productList, command)
			messageProduct := fmt.Sprintf("Продукт %s успешно добавлен в лист", command)
			fmt.Println(messageProduct)
		case order:
			fmt.Println(productList)
			productList = nil
			//fmt.Println(productList)
		}
	}
}

// добавить команду order, которая выводит сообщение, что вы купили, и очищает корзину
