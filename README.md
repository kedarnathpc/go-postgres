# Dockerised Go and Postgres CRUD App 

This Go application demonstrates how to implement database connectivity of Go with Postgres and also use it with docker images.
It provides stocks creation, deletion, stock retrieval, stock updation and retrieval of all stocks using builtin sql libraries.
Other packages required are github.com/gorilla/mux and github.com/joho/godotenv along with github.com/lib/pq (postgres drives).

## Features

- Stock registration
- Stock deletion
- Stock retrieval
- Stock updation
- All stocks retrieval

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go (Golang) installed on your system.
- Docker installed and setup with required permissions.
- Familiarity with Docker and basic Go programming.

## Installation and Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/kedarnathpc/go-postgres.git
  
2. Change the directory:
   
   ```bash
   cd go-postgres

3. Build the docker-compose file:

    ```bash
    docker-compose build

3. Run the docker-compose file:

    ```bash
    docker-compose up

4. Your Go CRUD app should be running on http://localhost:8080.

## Usage

To register a new stock, make a POST request to /api/newstock with the following JSON data:

    {
      "name": "Stock Name",
      "price": "Stock Price",
      "company": "Company Name"
    }

To delete a stock, make a DELETE request to /deletestock/{id}

To retrive a stock, make a GET request to /stock/{id}

To update a stock, make a PUT request to /stock/{it}

To retrive all stocks, make a GET request to /stock

## Contributing 

Fork the repository.
Create a new branch for your feature: git checkout -b feature-name.
Make your changes and commit them: git commit -m 'Add feature'.
Push to the branch: git push origin feature-name.
Create a pull request.
  