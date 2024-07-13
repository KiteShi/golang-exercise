# golang-exercise

## Prerequisites

Before starting, ensure you have the following installed and set up on your machine:

1. **Docker**: Check if Docker is installed:
   ```bash
   docker --version
   ```
   If not installed, download from here: https://docs.docker.com/get-docker/

2. **Docker Compose**: Check if Docker Compose is installed:
   ```bash
   docker-compose --version
   ```
   If not installed, follow instructions here: https://docs.docker.com/compose/install/

## Setup

1. **Clone the repository**:
    ```bash
    git clone https://github.com/KiteShi/golang-exercise.git
    cd golang-exercise
    ```

2. **Configure the `.env` file**:
   You can modify the existing `.env` file provided in the repository or leave it unchanged for testing purposes. Ensure it includes the necessary environment variables:
   ```dotenv
   JWT_KEY=your_jwt_secret_key
   ADMINS="your_admin_1:your_password_1,your_admin_2:your_password_2"
   POSTGRES_DB=your_database_name
   POSTGRES_USER=your_database_user
   POSTGRES_PASSWORD=your_database_password
   ```
   These variables are used for configuring JWT secret, admin users, PostgreSQL database connection details, and credentials.

3. **Build and run the services**:
    ```bash
    docker-compose up --build
    ```
## Usage

Here is an example of how to use this project with the Postman tool, which you can download from here https://www.postman.com/downloads/

1. **Install and run Postman on your machine**
   
2. **Import collection from repository to Postman**:
   - Go to **File** -> **Import...** -> **Create Workspace** (Or choose the existing one). Then press **Continue to Workspace** button, and select the collection file from the repository.
   - Postman collection has name `Golang-exercise.postman_collection.json`.
3. **Use request explanation for Postman provided below**

## Requests Explanation and Examples

- **Login POST :8080/login**
  Route for user authentication.
  Note: here as an example credentials from default .env file; ***if you have changed it then provide your own credentials***.
  
  **Body (raw):**
  ```json
  {
    "username": "john_doe",
    "password": "X397x@n"
  }
  ```
  
  The response to your request will be:
  ```json
  {
    "token": "your_jwt_token"
  }
  ```
  
  ***Copy `your_jwt_token`*** . You will need it for subsequent requests (excluding the Get company request).

- **Create company POST :8080/**
  Creates a new company entry.
 ***Note: only for authenticated users***
  In the "Authorization" tab, select Auth type "Bearer token" and provide `your_jwt_token` in the Token field.
  
  **Body (raw):**
  ```json
  {
    "name": "Company",
    "amount_of_employees": 1234,
    "registered": true,
    "type": "NonProfit"
  }
  ```
  
  The response will include the company structure with a generated `id` field. ***Copy `id`;*** you will need it for next requests.

- **Get company by ID GET :8080/{{uuid}}**
  Retrieves company information by its `id`.
  
  ***Replace `{{uuid}}`*** in the URL with the company `id` obtained from a previous request.
   In response you will get company info.

- **Patch company PATCH :8080/{{uuid}}**
  Updates an existing company entry.
  
  ***Replace `{{uuid}}`*** in the URL with the company `id` obtained from a previous request.
  
   ***Note: only for authenticated users***
  In the "Authorization" tab, select Auth type "Bearer token" and provide `your_jwt_token` in the Token field.
  
  **Body (raw):**
  ```json
  {
    "name": "New name",
    "amount_of_employees": 7,
    "registered": false,
    "type": "Cooperative"
  }
  ```
  
  Note: you don't need to provide all company structure fields.
  The response will include the updated company structure.

- **Delete company DELETE :8080/{{uuid}}**
  Deletes a company by its UUID.
  
  ***Note: only for authenticated users***
  In the "Authorization" tab, select Auth type "Bearer token" and provide `your_jwt_token` in the Token field.
  
  ***Replace `{{uuid}}`*** in the URL with the company UUID obtained from a previous request.

 This operation will completely remove the company from the database. ***Please note that this action is irreversible.***
 
In response you will get `HTTP status code 204 No Content.`

## Stopping the Services

To stop the running services:
```bash
docker-compose down
```
or

```bash
sudo docker-compose down
```

## Additional Commands

- **Running the Linter**:
    ```bash
    docker-compose run linter
    ```
    or
    ```bash
    sudo docker-compose run linter
    ```

- **Rebuild the services**:
    ```bash
    docker-compose up --build
    ```
    or 
    ```bash
    sudo docker-compose up --build
    ```

- **Remove all containers and volumes**:
    ```bash
    docker-compose down -v
    ```
    or
    ```bash
    sudo docker-compose down -v
    ```
