## Fullstack eCommerce project backend 


    backend: golang
    front: react js
    admin-panel: react js


### Backend tools
    database : postgres 
    cache : redis
    routes:  goFiber 
    read env file : go get github.com/joho/godotenv
    database ORM : GORM

## Database migrations command:

### Create migration table command :
    migrate create -ext sql -dir ./pkg/database/migrations/ -seq init
### Database up command :
    migrate -path ./pkg/database/migrations/ -database 'postgres://postgres:12345@localhost:5432/fullstack_ecommerce?sslmode=disable' up