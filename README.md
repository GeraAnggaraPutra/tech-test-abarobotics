# ABAROBOTICS Technical Test

## Entity Relationship Diagram (ERD)

![ABAROBOTICS ERD](ABAROBOTICS%20ERD.png)

## Local Development

### How to run

1. Copy file `Makefile.example` to `Makefile`
  ```sh
  cp Makefile.example Makefile
  ```

2. Setup PostgreSQL then config database url on `Makefile` 
  ```
  MIGRATE_DB_URL=postgresql://username:password@127.0.0.1:5432/db-name?sslmode=disable
  ```

3. Migrate database
  ```sh
  make migrate.up
  ```

4. Copy file `.env.example` to `.env`
  ```sh
  cp env.example .env
  ```

5. Setup the configuration to `.env`, but you're required to config the PostgreSQL
  ```
  DB_HOST=127.0.0.1
  DB_PORT=5432
  DB_USERNAME=
  DB_PASSWORD=
  DB_SCHEMA=
  ```
  
6. Install dependency go
  ```sh
  make deps
  ```

7. Run app for development
  ```sh
  make run
  ```

## Generate Swagger Documentation

1. Add Swagger annotations to your handler functions. Example:
  ```sh
  // @Summary      Login
  // @Description  Login to system and get access token & refresh token
  // @Tags         Auth
  // @Accept       json
  // @Produce      json
  // @Param        request body payload.LoginRequest true "Login credentials"
  // @Success      200  {object}  kernel.responseDataPayload
  // @Failure      400  {object}  kernel.responseErrorPayload
  // @Router       /auth/login [post]
  ```
  Use these comments above your handler functions to define endpoint descriptions, request/response schemas, and tags for grouping in Swagger UI.

2. Generate Swagger documentation by running
  ```sh
  make swag
  ```
  After running this command, Swagger documentation files will be generated inside the docs directory.

3. Access the documentation in your browser
  ```sh
  http://localhost:8000/docs
  ```
  Make sure the application is running so that the documentation can be accessed.

## API Documentation

[Download Postman API JSON File!](Abarobotics%20Tech%20Test%20API.postman_collection.json)