# Apk User Management System

### Features

- User management
- Mysql database
- Registration, authorization, role manager
- Use middleware for authenticate and permissions
- Router with https://github.com/gin-gonic/gin
- Database orm with https://github.com/go-gorm/gorm/
- Log system with https://github.com/uber-go/zap
- RBAC with https://github.com/harranali/authority
- Dashboard panel with https://github.com/coreui/coreui-free-bootstrap-admin-template

## Requirements
- Mysql DB

## Run
1. At first run Mysql database
2. Put your db config into app/config/consts
3. In project folder run this command
```bash
go run main.go
```
* Database tables is auto migrate