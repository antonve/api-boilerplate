# api-boilerplate

## Setup
- Install glide: https://github.com/Masterminds/glide
- Install dependencies
  ```
  $ glide install
  ```
- Make a copy of `dev.yml.example` and save it as dev.yml
  - Uses PostgreSQL
  - Update `connection_string` and `database_name` eg:
    ```
    connection_string: user=anton sslmode=disable dbname=
    database: api_dev
    ```
- Do the same for `test.yml.example`
- Run tests
  ```
  $ go test $(go list ./... | grep -v /vendor/)
  ```
- Run migrations in `dev` `prod`: todo
