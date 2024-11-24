# Bank Merchant API
This project is a simple API for login authentication using Golang with gorilla/mux framework and testify testing library.

# Feature
-Endpoint login with username and password validation.
-Use of routing with gorilla/mux.
-Unit testing to ensure endpoints are functioning correctly.

# Structure folder
BankMerchantAPI/
├── controllers/       # Logic for handlers and routers
├── data/              # Data storage, such as a mock database
├── models/            # Model definition for data
├── routes/            # Additional routing configuration
├── tests/             # Test file (unit test)
├── utils/             # Additional functions or utilities
├── views/             # Logic for response formats 
├── main.go            # Main entry point of the application
├── go.mod             # Go module file for defining dependencies
├── go.sum             # Checksum record of dependencies

# How to Run a Project

# Clone Repository
Clone this repository onto your computer using the following command:
git clone https://github.com/IkrmMrbsy/BankMerchantAPI.git
cd BankMerchantAPI

# Install Dependencies
go mod tidy

# Run the Server
go run main.go

# Contribution
If you would like to contribute to this project, please fork this repository and submit a pull request. All contributions are greatly appreciated!

# License
This project uses the MIT license. You are free to use it, but please respect this license.

