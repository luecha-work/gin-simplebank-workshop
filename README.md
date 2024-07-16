# Simple Bank Project

This project is a simple bank application developed using Go and several libraries and tools to manage database migrations, configurations, testing, and more.

## Table of Contents

1. [Libraries and Tools](#libraries-and-tools)
2. [Project Setup Steps](#project-setup-steps)
3. [Library Details and Installation](#library-details-and-installation)

## Libraries and Tools

The following libraries and tools are used in this project:

1. **Migrate**: A tool for managing database migrations.
2. **SQLC**: Generates Go code from SQL queries.
3. **Gin**: A web framework for building APIs.
4. **Viper**: A configuration solution for Go applications.
5. **Gomock**: A mocking framework for Go.
6. **golang-jwt**: A library for creating and verifying JSON Web Tokens.
7. **PASETO**: A library for creating and verifying Platform-Agnostic SEcure TOkens.

## Project Setup Steps

### Step 1: Install Migrate

Migrate is used to manage database migrations.

#### For Linux

1. Set the version variable:

   ```bash
   version=v4.15.0
   ```

2. Download Migrate:

   ```bash
   curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.linux-amd64.tar.gz -o migrate.tar.gz
   ```

3. Extract the tar file:

   ```bash
   tar xvzf migrate.tar.gz
   ```

4. Move Migrate to `/usr/local/bin`:

   ```bash
   sudo mv migrate /usr/local/bin/migrate
   ```

5. Verify the installation:
   ```bash
   migrate -version
   ```

#### For Windows

1. Download the Migrate binary from [here](https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.windows-amd64.tar.gz).
2. Extract the `migrate.exe` file to a directory of your choice.
3. Add the directory to your system's PATH variable.
4. Verify the installation by opening a command prompt and running:
   ```cmd
   migrate -version
   ```

### Step 2: Create Migration Files

1. Create a migration:
   ```bash
   migrate create -ext sql -dir db/migration -seq init_schema
   ```

### Step 3: Basic Docker Commands

#### For Linux

1. Enter the PostgreSQL container:

   ```bash
   docker exec -it postgres12 /bin/sh
   ```

2. Access PostgreSQL:

   ```bash
   psql simple_bank
   ```

3. Drop the `simple_bank` database:

   ```bash
   dropdb simple_bank
   ```

4. Create the `simple_bank` database:

   ```bash
   docker exec -it postgres12 createdb --username=root --owner=root simple_bank
   ```

5. Access the `simple_bank` database:
   ```bash
   docker exec -it postgres12 psql -U root simple_bank
   ```

#### For Windows

1. Enter the PostgreSQL container:

   ```cmd
   docker exec -it postgres12 cmd
   ```

2. Access PostgreSQL:

   ```cmd
   psql simple_bank
   ```

3. Drop the `simple_bank` database:

   ```cmd
   dropdb simple_bank
   ```

4. Create the `simple_bank` database:

   ```cmd
   docker exec -it postgres12 createdb --username=root --owner=root simple_bank
   ```

5. Access the `simple_bank` database:
   ```cmd
   docker exec -it postgres12 psql -U root simple_bank
   ```

### Step 4: Setup SQLC

SQLC is used to generate Go code from SQL queries.

1. Install SQLC:

   ```bash
   go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
   ```

2. Initialize SQLC:
   ```bash
   sqlc init
   ```

### Step 5: Create go.mod File

`go.mod` is essential for managing dependencies in a Go project.

1. Create the `go.mod` file:
   ```bash
   go mod init github.com/techschool/simple_bank
   ```

### Step 6: Clean Dependencies

1. Add missing dependencies and remove unused ones, then update `go.sum`:
   ```bash
   go mod tidy
   ```

### Additional Tools and Libraries

1. **Gin Web Framework**:

   - Install Gin:
     ```bash
     go get -u github.com/gin-gonic/gin
     ```
   - Install Gin CLI Tool:
     ```bash
     go install github.com/codegangsta/gin@latest
     ```
   - Start the project:
     ```bash
     gin run main.go
     ```
     Or:
     ```bash
     gin --port 8080 run main.go
     ```

2. **Viper**:

   - Install Viper:
     ```bash
     go get github.com/spf13/viper
     ```

3. **Gomock**:

   - Install Gomock:
     ```bash
     go install github.com/golang/mock/mockgen@latest
     go get github.com/golang/mock/mockgen@latest
     ```

4. **golang-jwt**:

   - Install golang-jwt:
     ```bash
     go get github.com/golang-jwt/jwt
     ```

5. **PASETO**:
   - Install PASETO:
     ```bash
     go get github.com/o1egl/paseto
     ```

## Library Details and Installation

### Migrate

Migrate is a tool used for managing database migrations. It helps in applying, reverting, and tracking changes to the database schema.

### SQLC

SQLC generates Go code from SQL queries. This tool ensures that your database queries are type-safe and reduces boilerplate code.

### Gin

Gin is a web framework written in Go. It provides a robust set of features for building RESTful APIs and web services.

### Viper

Viper is a comprehensive configuration solution for Go applications. It supports reading from various configuration file formats and environment variables.

### Gomock

Gomock is a mocking framework for Go. It helps in creating mocks for interfaces, which is useful for unit testing.

### golang-jwt

The `golang-jwt` library is used for creating and verifying JSON Web Tokens (JWT). It simplifies authentication and authorization processes in your application.

### PASETO

PASETO (Platform-Agnostic SEcure TOkens) is a token format designed to be more secure and easier to use than JWT. It addresses common security pitfalls associated with JWT.

## Conclusion

This README provides a comprehensive guide to setting up and using various tools and libraries for the Simple Bank project. Follow the steps carefully to ensure a smooth setup process.

### Docker

Docker create network เป็นการสร้างเครื่อข่ายเชื่อมต่อคอนเทนเนอร์หลายตัวกับเครือข่าย

1. Create network with docker:

   ```bash
   docker network create bank-network
   ```

2. network connect multi docker:

   ```bash
   docker network connect bank-network postgres12
   ```

3. ตรวจสอบรายละเอียดของเครือข่าย Docker:

   ```bash
   docker network inspect bank-network
   ```

Use Docker for deploy project

1. Create Docker image with Dockerfile:

   ```bash
   docker build -t simplebank:latest .
   ```

2. Check Docker image:

   ```bash
   docker images
   ```

3. Start Docker image with Dockerfile:

   ```bash
   docker run --name simplebank -p 8080:8080 -e GIN_MODE=release simplebank:latest

   Or use Network

   docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e "DB_SOURCE=postgresql://root:secret@postgres12:5432/simple_bank?sslmode=disable" simplebank:lates
   ```

   **คำอธิบาย**:

- **`--name simplebank`**: กำหนดชื่อให้กับคอนเทนเนอร์
- **`-p 8080:8080`**: แมพพอร์ต 8080 บนเครื่องของคุณกับพอร์ต 8080 ในคอนเทนเนอร์
- **`-e GIN_MODE=release`**: ตั้งค่าตัวแปรสภาพแวดล้อม `GIN_MODE` ให้เป็น `release`
- **`postgres12`**: Database container name
- **`--network bank-network`**: Use Network is bank-network

### Basic docker cmd

- docker container inspect container_name => ตรวจสอบรายละเอียดของ container
- docker network inspect container_name => ตรวจสอบรายละเอียดเครือข่ายของ container

# wait-for.sh

`wait-for.sh` เป็นสคริปต์ Bash ที่ช่วยในการรอให้บริการหรือทรัพยากรภายนอก (เช่น ฐานข้อมูล) พร้อมใช้งานก่อนที่จะเริ่มทำงานของแอปพลิเคชัน โดยเฉพาะในบริบทของการพัฒนาหรือการใช้งาน Docker คอนเทนเนอร์

## วิธีการใช้งาน

คุณสามารถใช้ `wait-for.sh` เพื่อรอให้บริการที่ระบุพร้อมใช้งาน จากนั้นจึงดำเนินการคำสั่งที่ต้องการ

### ตัวอย่างการใช้งาน

สมมติว่าคุณต้องการรอให้บริการ PostgreSQL บนโฮสต์ `database` พอร์ต `5432` พร้อมใช้งานก่อนที่จึงจะดำเนินการคำสั่งต่อไป คุณสามารถใช้สคริปต์ดังนี้:

```sh
./wait-for.sh database:5432 -- echo "Database is up"
```

# AWS ECR (Amazon Elastic Container Registry)

คือบริการจาก AWS ที่ใช้สำหรับเก็บและจัดการ `container images` สำหรับ `Docker` โดยมีคุณสมบัติดังนี้

```bash
`คุณสมบัติหลัก`
1. การเก็บรักษา: ECR ช่วยให้คุณสามารถเก็บ container images ได้อย่างปลอดภัย พร้อมการเข้ารหัสข้อมูลทั้งที่เก็บและส่งข้อมูล

2. การจัดการ: ECR มีเครื่องมือในการจัดการและเวอร์ชันของ images รวมถึงการจัดการการเข้าถึง (IAM) เพื่อควบคุมว่าใครสามารถเข้าถึงและจัดการ images ได้

3. การบูรณาการ: ECR ทำงานร่วมกับบริการอื่น ๆ ของ AWS ได้ดี เช่น Amazon ECS (Elastic Container Service) และ EKS (Elastic Kubernetes Service) เพื่อให้คุณสามารถใช้ container images ได้ง่ายขึ้นในแอปพลิเคชันของคุณ

4. ประสิทธิภาพ: ECR มีความสามารถในการทำ cache และดึง images อย่างรวดเร็ว ช่วยลดเวลาในการ deployment
```

### การใช้งาน

การสร้างและ push images: คุณสามารถสร้าง Docker images และ push ขึ้นไปยัง ECR เพื่อเก็บไว้
การดึง images: เมื่อคุณต้องการใช้งาน container images ในการ deploy สามารถดึง images จาก ECR ได้โดยตรง
