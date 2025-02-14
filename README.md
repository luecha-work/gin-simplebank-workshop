# Simple Bank Project

```
This project is a simple bank application developed using Go and several libraries and tools to manage database migrations, configurations, testing, and more.
```

# Table of Contents

1. [Libraries and Tools](#libraries-and-tools)
2. [Project Setup Steps](#project-setup-steps)
3. [Library Details and Installation](#library-details-and-installation)

# Libraries and Tools

The following libraries and tools are used in this project:

1. **Migrate**: A tool for managing database migrations.
2. **SQLC**: Generates Go code from SQL queries.
3. **Gin**: A web framework for building APIs.
4. **Viper**: A configuration solution for Go applications.
5. **Gomock**: A mocking framework for Go.
6. **golang-jwt**: A library for creating and verifying JSON Web Tokens.
7. **PASETO**: A library for creating and verifying Platform-Agnostic SEcure TOkens.

# Project Setup Steps

## Step 1: Install Migrate

Migrate is used to manage database migrations.

### For Linux

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

### For Windows

1. Download the Migrate binary from [here](https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.windows-amd64.tar.gz).
2. Extract the `migrate.exe` file to a directory of your choice.
3. Add the directory to your system's PATH variable.
4. Verify the installation by opening a command prompt and running:
   ```cmd
   migrate -version
   ```

## Step 2: Create Migration Files

1. Create a migration:
   ```bash
   migrate create -ext sql -dir db/migration -seq init_schema
   ```

## Step 3: Basic Docker Commands

### For Linux

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
   docker exec -it postgres12 createdb --username=wginroot --owner=wginroot simple_bank
   ```

5. Access the `simple_bank` database:
   ```bash
   docker exec -it postgres12 psql -U wginroot simple_bank
   ```

### For Windows

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
   docker exec -it postgres12 createdb --username=wginroot --owner=wginroot simple_bank
   ```

5. Access the `simple_bank` database:
   ```cmd
   docker exec -it postgres12 psql -U wginroot simple_bank
   ```

## Step 4: Setup SQLC

SQLC is used to generate Go code from SQL queries.

1. Install SQLC:

   ```bash
   go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
   ```

2. Initialize SQLC:
   ```bash
   sqlc init
   ```

## Step 5: Create go.mod File

`go.mod` is essential for managing dependencies in a Go project.

1. Create the `go.mod` file:
   ```bash
   go mod init github.com/techschool/simple_bank
   ```

## Step 6: Clean Dependencies

1. Add missing dependencies and remove unused ones, then update `go.sum`:
   ```bash
   go mod tidy
   ```

## Additional Tools and Libraries

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

   ```
   mockgen เป็นเครื่องมือในภาษา Go ที่ใช้ในการสร้าง mock objects สำหรับการทดสอบโค้ด โดยเฉพาะในการทดสอบ unit test
   ```

   - Install Gomock:

     ```bash
     go install github.com/golang/mock/mockgen@latest

     or

     go get github.com/golang/mock/mockgen@latest
     ```

   - set up for ubuntu

   ```
   export PATH=$PATH:$(go env GOPATH)/bin
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

Migrate is a tool used for managing database migrations. It helps in applying, reverting, and tracking changes to the database schema.

**SQLC**

SQLC generates Go code from SQL queries. This tool ensures that your database queries are type-safe and reduces boilerplate code.

**Gin**

Gin is a web framework written in Go. It provides a robust set of features for building RESTful APIs and web services.

**Viper**

Viper is a comprehensive configuration solution for Go applications. It supports reading from various configuration file formats and environment variables.

**Gomock**

Gomock is a mocking framework for Go. It helps in creating mocks for interfaces, which is useful for unit testing.

**golang-jwt**

The `golang-jwt` library is used for creating and verifying JSON Web Tokens (JWT). It simplifies authentication and authorization processes in your application.

**PASETO**

PASETO (Platform-Agnostic SEcure TOkens) is a token format designed to be more secure and easier to use than JWT. It addresses common security pitfalls associated with JWT.

**Conclusion**

This README provides a comprehensive guide to setting up and using various tools and libraries for the Simple Bank project. Follow the steps carefully to ensure a smooth setup process.

**Docker**

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

   docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e "DB_SOURCE=postgresql://wginroot:secret@postgres12:5432/simple_bank?sslmode=disable" simplebank:lates
   ```

   **คำอธิบาย**:

- **`--name simplebank`**: กำหนดชื่อให้กับคอนเทนเนอร์
- **`-p 8080:8080`**: แมพพอร์ต 8080 บนเครื่องของคุณกับพอร์ต 8080 ในคอนเทนเนอร์
- **`-e GIN_MODE=release`**: ตั้งค่าตัวแปรสภาพแวดล้อม `GIN_MODE` ให้เป็น `release`
- **`postgres12`**: Database container name
- **`--network bank-network`**: Use Network is bank-network

# Basic docker cmd

- docker container inspect container_name => ตรวจสอบรายละเอียดของ container
- docker network inspect container_name => ตรวจสอบรายละเอียดเครือข่ายของ container

# wait-for.sh

`wait-for.sh` เป็นสคริปต์ Bash ที่ช่วยในการรอให้บริการหรือทรัพยากรภายนอก (เช่น ฐานข้อมูล) พร้อมใช้งานก่อนที่จะเริ่มทำงานของแอปพลิเคชัน โดยเฉพาะในบริบทของการพัฒนาหรือการใช้งาน Docker คอนเทนเนอร์

## วิธีการใช้งาน

คุณสามารถใช้ `wait-for.sh` เพื่อรอให้บริการที่ระบุพร้อมใช้งาน จากนั้นจึงดำเนินการคำสั่งที่ต้องการ

## ตัวอย่างการใช้งาน

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

## การใช้งาน

การสร้างและ push images: คุณสามารถสร้าง Docker images และ push ขึ้นไปยัง ECR เพื่อเก็บไว้
การดึง images: เมื่อคุณต้องการใช้งาน container images ในการ deploy สามารถดึง images จาก ECR ได้โดยตรง

# AWS IAM (Identity and Access Management)

```
เป็นบริการของ AWS ที่ช่วยจัดการการเข้าถึงทรัพยากร AWS อย่างปลอดภัย โดยผู้ใช้สามารถควบคุมว่าใครสามารถทำอะไรกับทรัพยากรใน AWS ได้บ้าง IAM ช่วยให้คุณสามารถ
```

# Pull image form aws ECR

เพิ่มเติมที่ [ที่นี่](https://docs.aws.amazon.com/cli/latest/reference/ecr/get-login-password.html)

```
`login aws`

- get password: aws ecr get-login-password
  or
- aws ecr get-login-password | docker login --username AWS --password-stdin 325274678571.dkr.ecr.ap-southeast-2.amazonaws.com
```

# Amazon Elastic Kubernetes Service (EKS)

```
Amazon Elastic Kubernetes Service (EKS) คือบริการที่มีการจัดการจาก Amazon Web Services (AWS) สำหรับการใช้งาน `Kubernetes` ซึ่งเป็นแพลตฟอร์มโอเพ่นซอร์สสำหรับการจัดการ containerized applications. EKS ช่วยให้ผู้ใช้สามารถเริ่มต้นและบริหารจัดการ `Kubernetes` บน AWS ได้ง่ายขึ้นโดยที่ไม่ต้องตั้งค่าและบริหารจัดการ control plane ของ `Kubernetes` เอง

- Install kubernetes to local form `https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/`

- Connect kube and aws with `aws eks update-kubeconfig --name simple-bank --region ap-southeast-2`

- Check config detail with `cat ~/.kube/config`
```

# Kubernetes เบื่องต้น

```
คำสั่ง aws sts get-caller-identity ใช้เพื่อตรวจสอบและแสดงข้อมูลเกี่ยวกับ IAM identity (ผู้ใช้หรือ role) ที่กำลังใช้งาน AWS CLI หรือ SDK ในขณะนั้น
```

- aws sts get-caller-identity

```
คำสั่งนี้ใช้เพื่อแสดงข้อมูลเกี่ยวกับ Kubernetes cluster ที่คุณเชื่อมต่ออยู่
```

- kubectl cluster-info

```
คำสั่งนี้ใช้เพื่อแสดงรายการของ pods ที่กำลังรันอยู่ใน Kubernetes cluster
```

- kubectl get pods

# K9S

```
K9s เป็นเครื่องมือที่ใช้สำหรับการจัดการและมอนิเตอร์ Kubernetes clusters ผ่านทาง command-line interface (CLI) ที่มีลักษณะของ UI แบบ text-based ซึ่งช่วยให้การจัดการและตรวจสอบ Kubernetes cluster ง่ายและรวดเร็วขึ้น
```

## Install On Ubuntu

- curl -sS https://webinstall.dev/k9s | bash
- k9s version

## Use K9S

- k9s

# Generate DB Docs from DBML Code

```
โดยใช้เครื่องมือหลายตัวที่เกี่ยวข้องกับการสร้างเอกสารฐานข้อมูล (DB Docs) จากโค้ด DBML (Database Markup Language) ซึ่งในภาพประกอบด้วย

- DBML
เป็นภาษาที่ใช้ในการนิยามโครงสร้างฐานข้อมูลเชิงสัมพันธ์ (Relational Database) ทำให้การออกแบบและจัดการฐานข้อมูลง่ายขึ้นผ่านการเขียนโค้ด

- Dbdocs
เครื่องมือที่ช่วยสร้างเอกสารจากโค้ด DBML โดยแปลงข้อมูลโครงสร้างฐานข้อมูลให้เป็นเอกสารที่เข้าใจง่ายและสามารถแชร์ได้

- Dbdiagram
ใช้สำหรับการสร้างและแสดงภาพ (diagram) ของโครงสร้างฐานข้อมูลที่ช่วยให้การสื่อสารระหว่างทีมพัฒนาเป็นไปอย่างมีประสิทธิภาพ
```

# gRPC & protobuf

```
-> https://grpc.io/docs/languages/go/quickstart/
-> https://grpc.io/docs/protoc-installation/

gRPC และ Protocol Buffers (protobuf) เป็นเทคโนโลยีที่ช่วยในการสื่อสารระหว่างบริการ (services) ในระบบ distributed systems หรือไมโครเซอร์วิส


`วิธีการทำงานร่วมกันระหว่าง gRPC & protobuf`

1. เขียนไฟล์ .proto: คุณจะสร้างไฟล์ที่บอกว่าโปรแกรมควรจะส่งข้อมูลแบบไหนและสื่อสารกันอย่างไร
2. คอมไพล์ไฟล์ .proto: ใช้เครื่องมือเพื่อแปลงไฟล์นี้เป็นโค้ดที่โปรแกรมของคุณสามารถใช้ได้
3. พัฒนาโปรแกรม: เขียนโปรแกรมที่ใช้โค้ดจากการคอมไพล์เพื่อเรียกใช้ฟังก์ชันที่อยู่บนเซิร์ฟเวอร์หรือรับข้อมูลจากเซิร์ฟเวอร์

- gRPC: เหมือนการโทรศัพท์ไปถามคำตอบจากเพื่อน
- protobuf: เหมือนการเขียนจดหมายที่สั้นและชัดเจน เพื่อให้เพื่อนอ่านเข้าใจและตอบกลับได้เร็ว
```

# ktr0731/evans

เพิ่มเติมที่ [ที่นี่](https://github.com/ktr0731/evans)

```
ป็นโครงการโอเพนซอร์สที่พัฒนาโดยผู้ใช้ GitHub ชื่อ ktr0731 ซึ่ง Evans เป็นเครื่องมือที่ใช้สำหรับการโต้ตอบกับ gRPC APIs ผ่าน CLI (Command Line Interface) หรือที่เรียกได้ว่าเป็น gRPC client ในรูปแบบ command-line

Install ktr0731/evans

1. ติดตั้ง Evans ผ่าน Go:

go install github.com/ktr0731/evans@latest

2. เพิ่ม Go binary path ไปยัง $PATH:

- ls $(go env GOPATH)/bin/evans
- nano ~/.bashrc
- export PATH=$PATH:$(go env GOPATH)/bin
- source ~/.bashrc

3. ตรวจสอบการติดตั้ง:

evans --version
```

# gRPCurl Guide

`grpcurl` เป็นเครื่องมือ command-line ที่ช่วยทดสอบ gRPC services โดยไม่ต้องเขียน client code ช่วยให้การเรียกใช้ services และการ debug ทำได้สะดวกผ่านการส่งและรับข้อมูลในรูปแบบ JSON

## การติดตั้ง

1. ดาวน์โหลด `grpcurl` จาก [ที่นี่](https://github.com/fullstorydev/grpcurl/releases)
2. ตรวจสอบการติดตั้งด้วยคำสั่ง:
   ```bash
   grpcurl --version
   ```

## การใช้งาน

1. List services บน gRPC Server ใช้คำสั่งนี้เพื่อแสดงรายชื่อ services ที่ gRPC server ให้บริการ:

```
grpcurl -plaintext localhost:50051 list
```

2. List methods ของ service:

```
grpcurl -plaintext localhost:50051 list pb.SimpleBank
```

3. เรียกใช้งาน method บน gRPC Server

```
grpcurl -plaintext -d '{"username": "test_user", "password": "test_pass"}' localhost:50051 pb.SimpleBank.LoginUser
```

4. เรียกใช้ method ที่ไม่ต้องการข้อมูล

```
grpcurl -plaintext localhost:50051 pb.SimpleBank.GetUser
```

6.5. ใช้งานกับ TLS (Secure Connection)

```
grpcurl -insecure -d '{"username": "test_user", "password": "test_pass"}' localhost:50051 pb.SimpleBank.LoginUser
```

# gRPC-Gateway

เพิ่มเติมที่ [ที่นี่](https://github.com/grpc-ecosystem/grpc-gateway)

```
เป็นเครื่องมือในระบบ gRPC ที่ช่วยให้สามารถสร้าง RESTful HTTP API โดยทำการแปลงการเรียก gRPC ไปเป็นการเรียกผ่าน HTTP/1.1 และใช้ JSON ในการสื่อสารได้โดยอัตโนมัติ ซึ่งช่วยให้ผู้พัฒนาสามารถรองรับการสื่อสารทั้งแบบ gRPC และ RESTful API ได้ในเวลาเดียวกัน โดยไม่ต้องเขียนโค้ดสองชุดแยกกัน
```

## การทำงานของ grpc-gateway

```
1. `API การแปลง:` grpc-gateway ทำหน้าที่เป็น proxy ระหว่าง HTTP และ gRPC โดยมันจะรับคำร้องจาก HTTP (REST API) ที่ส่งมาเป็น JSON และแปลงคำร้องเหล่านั้นให้เป็นการเรียก gRPC ภายใน
2. `การเชื่อมโยงระหว่าง gRPC กับ HTTP:` เราจะใช้ไฟล์ .proto (ไฟล์ที่กำหนดโครงสร้างของ gRPC service) เพิ่มคำสั่งที่บอกว่า RPC method แต่ละตัวจะถูกแมปกับ HTTP method (GET, POST, PUT, DELETE) และ endpoint ใด
3. `การตอบกลับ:` เมื่อ gRPC ส่งผลลัพธ์กลับมา มันจะถูกแปลงเป็น JSON และส่งกลับไปยัง client ผ่าน HTTP
```

## Set up in gi framework

clone project [ที่นี่](https://github.com/grpc-ecosystem/grpc-gateway)

```
1. move file to proto/google/api
   `
   google/api/annotations.proto
   google/api/field_behavior.proto
   google/api/http.proto
   google/api/httpbody.proto
   `

2. add option in your service `SimpleBank`
  `
   option (google.api.http) = {
      post: "/v1/create_user"
      body: "*"
   };
  `
3. add `--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative` in proto script
4. run make proto
```

# OpenAPI Swagger

clone project [ที่นี่](https://github.com/grpc-ecosystem/grpc-gateway)

## 1. Setup Swagger

```
1. move file *.proto in Folder protoc-gen-openapiv2/options/ to your project proto/protoc-gen-openapiv2/options
2. create option in service_simple_bank.proto
3. add `--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank` in proto script
4. run make proto
```

## 2. Create Swagger UI

clone project [ที่นี่](https://github.com/swagger-api/swagger-ui)

```
1. move all file in folder dist to your folder `doc/swagger`
2. chang url in file swagger-initializer.js to your swagger file json `(doc/swagger/simple_bank.swagger.json)`
3. add config runGatewayServer at mark TODO: Setup Swagger in main.go file
4. go to swagger with `http://localhost:8080/swagger`
```

# zerolog

เพิ่มเติมที่ [ที่นี่](https://github.com/rs/zerolog)

```
สร้าง log ในระบบของ golang ด้วย zerolog
```
