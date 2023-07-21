package main

import (
	"aegis_test/apps/cashier"
	"aegis_test/apps/product"
	"aegis_test/apps/supplier"
	"aegis_test/apps/unit"
	"aegis_test/apps/user"
	"aegis_test/apps/welcome"
	"aegis_test/libs"
	"aegis_test/libs/db"

	"github.com/gin-gonic/gin"
)

const (
	pathIndex  string = "/"
	pathById   string = "/id/:id"
	pathBySlug string = "/slug/:slug"
	pathSave   string = "/save"
)

func init() {
	libs.LoadEnv()
}

func main() {
	dbFactory := db.NewDBFactory()

	router := gin.Default()
	router.Use(libs.JWTMiddleware())

	v1 := router.Group("v1")
	v1.GET(pathIndex, welcome.Welcome)

	loginController := user.NewLoginController(dbFactory)
	v1.POST("/login", loginController.Login)

	adminV1 := v1.Group("admin")
	adminUserV1 := adminV1.Group("user")
	adminUnitV1 := adminV1.Group("unit")
	adminSupplierV1 := adminV1.Group("supplier")
	adminProductV1 := adminV1.Group("product")

	cashierV1 := v1.Group("cashier")

	userGetter := user.NewGetterController(dbFactory)
	userSaver := user.NewSaverController(dbFactory)
	unitGetter := unit.NewGetterController(dbFactory)
	unitSaver := unit.NewSaverController(dbFactory)
	supplierGetter := supplier.NewGetterController(dbFactory)
	supplierSaver := supplier.NewSaverController(dbFactory)
	productGetter := product.NewGetterController(dbFactory)
	productSaver := product.NewSaverController(dbFactory)
	cashierSaver := cashier.NewSaverController(dbFactory)

	adminUserV1.GET(pathIndex, userGetter.GetAll)
	adminUserV1.GET(pathById, userGetter.GetById)
	adminUserV1.POST(pathSave, userSaver.Save)

	adminUnitV1.GET(pathIndex, unitGetter.GetAll)
	adminUnitV1.GET(pathById, unitGetter.GetById)
	adminUnitV1.GET(pathBySlug, unitGetter.GetBySlug)
	adminUnitV1.POST(pathSave, unitSaver.Save)

	adminSupplierV1.GET(pathIndex, supplierGetter.GetAll)
	adminSupplierV1.GET(pathById, supplierGetter.GetById)
	adminSupplierV1.GET(pathBySlug, supplierGetter.GetBySlug)
	adminSupplierV1.POST(pathSave, supplierSaver.Save)

	adminProductV1.GET(pathIndex, productGetter.GetAll)
	adminProductV1.GET(pathById, productGetter.GetById)
	adminProductV1.GET(pathBySlug, productGetter.GetBySlug)
	adminProductV1.POST(pathSave, productSaver.Save)

	cashierV1.POST(pathSave, cashierSaver.Save)

	router.Run()
}
