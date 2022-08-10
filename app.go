package main

import (
	"fmt"
	"golang-gorm-fundamental/config"
	"golang-gorm-fundamental/repository"
	"golang-gorm-fundamental/usecase"
	"golang-gorm-fundamental/utils"
)

func main() {
	// dbHost := "localhost"
	// dbPort := "5432"
	// dbUser := "postgres"
	// dbPass := "12345678"
	// dbName := "db_enigma_shop_v2"

	cfg := config.NewConfig()
	db := cfg.DBConn()
	defer cfg.DBClose()

	repo := repository.NewCustomerRepository(db)
	productRepo := repository.NewProductRepository(db)

	// Insert customer
	// id := utils.GenerateId()
	// password, _ := utils.HashPassword("12345678")
	// customer01 := entity.Customer{
	// 	Id:   id,
	// 	Name: "Dolby Atoms",
	// 	Address: []entity.Address{
	// 		{
	// 			StreetName: "Jln. Kapuas",
	// 			City:       "Cilacap",
	// 			PostalCode: "53224",
	// 		}, {
	// 			StreetName: "Jln. Ketintang",
	// 			City:       "Surabaya",
	// 			PostalCode: "60118",
	// 		},
	// 	},
	// 	Balance: 200000,
	// 	UserCredential: entity.UserCredential{
	// 		UserName: "dolby",
	// 		Password: password,
	// 	},
	// 	Email: "dolby.atoms@gmail.com",
	// 	Phone: "08912311111",
	// }
	// err := repo.Create(&customer01)
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// // update with preload
	// customer02:=entity.Customer{
	// 	Id: "",
	// }
	// customer02, err := repo.FindFirstWithPreloaded(map[string]interface{}{
	// 	"id": "3921d98e-552f-4d71-8289-8f594a5bcc5a",
	// }, "Address")
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// log.Println(customer02.ToString())

	// cstUse := usecase.NewAuthenticationUseCase(repo)

	// // Authentication
	// user := "apep"
	// pass := "12345678"
	// customer, err := cstUse.Login(user, pass)
	// utils.IsError(err)
	// log.Println(customer.ToString())

	// // Authentication Usecase
	// usercredRepo := repository.NewUserCredentialRepository(db)
	// auth := usecase.NewAuthenticationUseCase(repo, usercredRepo)
	// // success
	// cust, err := auth.Login("yurham", "12345678")
	// utils.IsError(err)
	// fmt.Println(cust.ToString())
	// // wrong password
	// cust, err := auth.Login("yurham", "87654321")
	// utils.IsError(err)
	// fmt.Println(cust.ToString())
	// // user not found
	// cust, err := auth.Login("siapaya", "12345678")
	// utils.IsError(err)
	// fmt.Println(cust.ToString())

	// // member activation
	// activation := usecase.NewMemberActivationUsecase(repo)
	// // success
	// cust, err := activation.ActivateMember("yurham.afif@google.com")
	// utils.IsError(err)
	// fmt.Println(cust.ToString())
	// // wrong email
	// cust, err := activation.ActivateMember("nicho.afif@google.com")
	// utils.IsError(err)
	// fmt.Println(cust.ToString())

	// // customer balance
	// balance := usecase.NewCustomerBalanceUsecase(repo)
	// // deposit
	// cust, err := balance.Deposit("3921d98e-552f-4d71-8289-8f594a5bcc5a", 20000)
	// utils.IsError(err)
	// fmt.Println(cust.ToString())
	// // deposit cust not found
	// cust, err := balance.Deposit("552f-4d71-8289-8f594a5bcc5a", 20000)
	// utils.IsError(err)
	// fmt.Println(cust.ToString())
	// // withdraw
	// cust, err := balance.Withdraw("3921d98e-552f-4d71-8289-8f594a5bcc5a", 20000)
	// utils.IsError(err)
	// fmt.Println(cust.ToString())
	// // withdraw cust not found
	// cust, err := balance.Withdraw("552f-4d71-8289-8f594a5bcc5a", 20000)
	// utils.IsError(err)
	// fmt.Println(cust.ToString())
	// // withdraw amount exceeding
	// cust, err := balance.Withdraw("3921d98e-552f-4d71-8289-8f594a5bcc5a", 20000000)
	// utils.IsError(err)
	// fmt.Println(cust.ToString())

	// Many to many
	// case1: Insert customer sekaligus product
	// id := utils.GenerateId()
	// password, _ := utils.HashPassword("12345678")
	// customer01 := entity.Customer{
	// 	Id:   id,
	// 	Name: "Yurham Afif",
	// 	Address: []entity.Address{
	// 		{
	// 			StreetName: "Jln. Keputih",
	// 			City:       "Surabaya",
	// 			PostalCode: "60111",
	// 		}, {
	// 			StreetName: "Jln. KTT",
	// 			City:       "Surabaya",
	// 			PostalCode: "60111",
	// 		},
	// 	},
	// 	Balance: 200000,
	// 	UserCredential: entity.UserCredential{
	// 		UserName: "yurham",
	// 		Password: password,
	// 	},
	// 	Email: "yurham.afif@gmail.com",
	// 	Phone: "08912345679",
	// 	Products: []*entity.Product{
	// 		{ProductName: "caca marica"},
	// 		{ProductName: "beng beng"},
	// 	},
	// }
	// err := repo.Create(&customer01)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// case2: Insert new product into enlisted customer
	// cust, err := repo.FindById("1c33cdcf-3880-43c7-b366-86f0236fcd1b")
	// utils.IsError(err)
	// cust.Products = []*entity.Product{
	// 	{ProductName: "cocolatos"},
	// }
	// err = repo.UpdateBy(&cust)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// cust.Products = []*entity.Product{
	// 	{ProductName: "Silver qing"},
	// }
	// err = repo.UpdateByModel(&cust)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// case3: create new customer using exsisting product
	// product1, err := productRepo.FindById(1)
	// utils.IsError(err)
	// fmt.Println(product1.ToString())
	// id := utils.GenerateId()
	// password, _ := utils.HashPassword("12345678")
	// customer2 := entity.Customer{
	// 	Id:   id,
	// 	Name: "Mpus Afif",
	// 	Address: []entity.Address{
	// 		{
	// 			StreetName: "Jln. Gedongombo",
	// 			City:       "Tuban",
	// 			PostalCode: "62281",
	// 		}, {
	// 			StreetName: "Jln. Rinjani",
	// 			City:       "Tuban",
	// 			PostalCode: "62314",
	// 		},
	// 	},
	// 	Balance: 200000,
	// 	UserCredential: entity.UserCredential{
	// 		UserName: "apep",
	// 		Password: password,
	// 	},
	// 	Email: "mpus.afif@gmail.com",
	// 	Phone: "089123456779",
	// 	Products: []*entity.Product{
	// 		&product1,
	// 	},
	// }
	// err = repo.Create(&customer2)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// case4: add enlisted product to enlisted customer (customer_with_products)
	// custExisting, err := repo.FindById("e85dac0e-7858-4b3e-a1c0-118a59605152")
	// utils.IsError(err)
	// prodExisting, err := productRepo.FindById(2)
	// prodExisting2, err := productRepo.FindById(3)
	// utils.IsError(err)
	// custExisting.Products = []*entity.Product{&prodExisting}
	// repo.UpdateByModel(&custExisting)
	// custExisting.Products = append(custExisting.Products, &prodExisting2)
	// repo.UpdateByModel(&custExisting)

	// // delete associated product by customer yet keeping the product in mst_product
	// cust, err := repo.FindById("e85dac0e-7858-4b3e-a1c0-118a59605152")
	// utils.IsError(err)
	// prod, err := productRepo.FindById(2)
	// utils.IsError(err)
	// err = repo.DeleteAssociation(&cust, "Products", &prod)
	// utils.IsError(err)
	
	// // existing customer want to update their related product
	// cust, err := repo.FindFirstWithPreloaded(map[string]interface{}{
	// 	"id": "e85dac0e-7858-4b3e-a1c0-118a59605152",
	// }, "Products")
	// utils.IsError(err)
	// newprod, err := productRepo.FindById(1)
	// utils.IsError(err)
	// fmt.Println(newprod.ToString())
	// var oldprodId uint = 4
	// var newProdSlice []entity.Product
	// for _, prod := range cust.Products {
	// 	if (*prod).ID != oldprodId {
	// 		newProdSlice = append(newProdSlice, *prod)
	// 	}
	// }
	// fmt.Println("newProdSlice before ", newProdSlice)
	// newProdSlice = append(newProdSlice, newprod)
	// fmt.Println("newProdSlice after ", newProdSlice)
	// err = repo.UpdateAssociation(&cust, "Products", newProdSlice)
	// utils.IsError(err)

	// get total product each customer
	// total := usecase.NewCustomerUseCase(repo, productRepo)
	// model, _ := repo.FindAllWithPreload("Products")
	// fmt.Println(model)
	// result, err := total.GetTotalProductEachCustomer(&model)
	// utils.IsError(err)
	// fmt.Println("total product each customer:")
	// fmt.Println(result)
	// get customer with no product
	total := usecase.NewCustomerUseCase(repo, productRepo)
	model, _ := repo.FindAllWithPreload("Products")
	// fmt.Println(model)
	result, err := total.GetTotalCustomerWithNoProduct(&model)
	utils.IsError(err)
	fmt.Println("customer with no product:")
	fmt.Println(result)
}
