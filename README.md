# Bank Merchant API.
Bank Merchant API is a backend application built using the Go (Golang) programming language, designed to facilitate the login process, merchant management, and payment transactions between customers and merchants. 

# Feature
-Login Authentication: Validate customers using email and password.

-Merchant Management: Retrieve merchant information.

-Payment Processing: Make payments linked to specific merchants.

# Structure folder
BankMerchantAPI/
├── controllers/       # Logic for handlers and routers

├── data/              # Data storage, such as a mock database

├── models/            # Model definition for data

├── routes/            # Additional routing configuration

├── utils/             # Additional functions or utilities

├── views/             # Logic for response formats

├── main.go            # Main entry point of the application

├── go.mod             # Go module file for defining dependencies

├── go.sum             # Checksum record of dependencies

# How to Run a Project
Clone this repository onto your computer using the following command:

git clone https://github.com/IkrmMrbsy/BankMerchantAPI.git

cd BankMerchantAPI

# Run the Server
go run main.go

# How to use the API in Postman
# Login Customer
-Endpoint: POST /login

-URL: http://localhost:8080/login

-Enter parameters (email and password)

# Payment Process
-Endpoint: POST /payment

-URL: http://localhost:8080/payment

-Enter parameters(email, amount and merchant_id)

# Logout
-Endpoint: POST /logout

-URL: http://localhost:8080/logout

-Enter parameters (email)

# Contribution
If you would like to contribute to this project, please fork this repository and submit a pull request. All contributions are greatly appreciated!

# License
This project uses the MIT license. You are free to use it, but please respect this license.

