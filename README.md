# API Court Rental
Golang API that provides endpoints to manage court rental entities.

## 🧰 Installation

1. Install [go](https://go.dev/doc/install)
2. Install [docker compose](https://docs.docker.com/compose/install/)
3. Configure *.env* file using *.env.example*
    ```
    # api
    API_SECRET=yoursecret

    # pg database
    DB_HOST=host.docker.internal
    DB_DRIVER=postgres
    DB_USER=user
    DB_PASSWORD=password
    DB_NAME=court_rental_db
    DB_PORT=5432
    ```
   
## 🚀 Run
* Run the following command on project's root folder:
    ```bash
    $ docker-compose up
    ```

## 📁 Project structure
```
soon...
```

## 🟢 Endpoints

#### /owners
* `POST`: Create a new owner

#### /owners/:id
* `PUT`: Update owner
* `DELETE`: Delete owner

#### /login
* `POST`: User login

## 💻 Author

* [Christopher Moura](https://github.com/ccmoura)