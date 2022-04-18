## Picus Security Golang Backend Web Development Bootcamp Final Project
In this application, the member with admin authority will be able to create Product or Product Category. Customers will be able to purchase existing products and view their past orders.Using these we aim to develop a basket service.

## Requests of application

1. You should get the required information from the user and create a user in the database and return JWT token in response.(Sign-up)
2. Users registered in the database must log in to the system with their email and password, if both information is correct, you must create a JWT token and return to the user.(Login)
3. With this endpoint, only users in the admin role should create a new category by uploading a CSV file.(Creat Bulk Category)
4. All active and not deleted categories in the database should be listed.(List Category)
5. Users who are logged into the system and whose Token period has not expired can add their products to the basket.(Add To Cart)
6. Users can list the products added to their cart.(ListCartItems)
7. Users can update the number of products added to their basket or remove the product from the basket.(Update/Delete Cart Items)
8. Users can create an order with the products they add to their cart.(Complete Order)
9. Customers can view their past orders.(List Orders)
10. If the customer's order date has not passed 14 days, the customer can cancel their order. If 14 days have passed after the order creation date, the cancellation request should be invalid.(Cancel Order)
11. Users in the admin role should be able to create individual products for the system.(Create Product)
12. Users should be able to list all products. No role control here.(List Product)
13. Users should be able to search within products. In the search section, fields such as product name, stockcode, etc. can be searched.(Search Product)
14. Users in the admin role can delete products.(Delete Product)
15. Users in the admin role can update the product.(Update Product)

 
## Used Handle Functions

### for user
```
			POST("/login", user.Login)
			POST("/signup", user.SignUp)
```
These functions allow us to create a RESTful API with the users we define in the application.

### for category
```
			GET("", category.GetAllCategories)
			GET(":id", category.GetCategory)
			POST("", authMw.CheckRole("admin"), category.CreateCategory)
			DELETE(":id", authMw.CheckRole("admin"), category.DeleteCategory)
			POST("/bulk", authMw.CheckRole("admin"), category.CreateBulkCategory)
```
These functions allow us to create a RESTful API with the categories we define in the application.

### for product
```
			GET("", product.GetAllProducts)
			GET(":id", product.GetProduct)
			GET("/search/:query", product.SearchProducts)
			POST("", authMw.CheckRole("admin"), product.CreateProduct)
			PUT(":id", authMw.CheckRole("admin"), product.UpdateProduct)
			DELETE(":id", authMw.CheckRole("admin"), product.DeleteProduct)
```
These functions allow us to create a RESTful API with the products we define in the application.

### for cart
```
			GET("", cart.GetCart)
			POST("", cart.AddCartItem)
			PUT("", cart.UpdateCartItem)
			DELETE("", cart.RemoveCartItem)
```
These functions allow us to create a RESTful API with the carts we define in the application.

### for order
```
			GET("", order.GetAllOrders)
			POST("", order.CreateOrder)
			PATCH(":id/cancel", order.CancelOrder)
```
These functions allow us to create a RESTful API with the carts we define in the application.

## Used Swagger Document
You can see the whole operation of the application and what has been done here.

Swagger file: https://app.swaggerhub.com/apis/kadirgonen/kadir-gonen_picus_api/1.0#/

## Used Pkg Packages

* database_handler: For database connection
* pagination: For paginate objects
* logger: For see errors as messages
* httpserver: For server processes

## Used Files

* main.go: Main application file.
* Config: Configuration files.
* Repositories: CRUD operations handling.
* Models: It is our domain data.
* Handlers: This files , they receive the request from the user, they ask the services to perform an action for them on the database.
* Services: contains some business logic for each model, and for authorization.
* Middlewares: it contains middlewares(golang functions) that are triggered before the controller action, for example, a middleware which reads the request looking for the Jwt token and trying to authenticate the user before forwarding the request to the corresponding controller action.
* PKG: it contains the packages that are used by the application.

## Requirements

* Go Language
* Git
* Go Module
* GORM
* Gin
* Postgres
* JWT
* Viper
* Swagger
* Zerolog

## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration
* etc

## License
[MIT](https://choosealicense.com/licenses/mit/)
