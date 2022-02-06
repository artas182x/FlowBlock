package main

import (
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/artas182x/hyperledger-fabric-master-thesis/backend/controllers"
	docs "github.com/artas182x/hyperledger-fabric-master-thesis/backend/docs"
	"github.com/artas182x/hyperledger-fabric-master-thesis/backend/models"
	"github.com/artas182x/hyperledger-fabric-master-thesis/backend/services"
	"github.com/artas182x/hyperledger-fabric-master-thesis/backend/vars"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	port := os.Getenv("PORT")
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	if port == "" {
		port = "8000"
	}

	docs.SwaggerInfo.BasePath = "/api"

	services.InitCelery()

	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: vars.IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					"certificate":    v.Login.Certificate,
					"privateKey":     v.Login.PrivateKey,
					"mspID":          v.Login.MspID,
					"roles":          v.Roles,
					vars.IdentityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			login := models.Login{
				Certificate: claims["certificate"].(string),
				PrivateKey:  claims["privateKey"].(string),
				MspID:       claims["mspID"].(string),
			}

			rolesInterface := claims["roles"].([]interface{})

			roles := make([]string, len(rolesInterface))

			for i := 0; i < len(roles); i++ {
				roles[i] = rolesInterface[i].(string)
			}

			return &models.User{
				UserName: claims[vars.IdentityKey].(string),
				Login:    login,
				Roles:    roles,
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			user, err := controllers.Authenticate(c)
			if err != nil {
				return nil, err
			}
			return user, nil

		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// TODO Check cert here
			/*if v, ok := data.(*User); ok && v.UserName == "valid" {
				return true
			}

			return false*/
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	apiGroup := r.Group("/api")

	apiGroup.POST("/login", authMiddleware.LoginHandler)
	apiGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	apiV1 := apiGroup.Group("v1")

	// Refresh time can be longer than token timeout
	apiV1.GET("/refresh_token", authMiddleware.RefreshHandler)
	apiV1.Use(authMiddleware.MiddlewareFunc())
	{
		apiV1.GET("/computation/availablemethods/:chaincode_name", controllers.GetAvailableMethods)
		apiV1.POST("/computation/requesttoken", controllers.RequestToken)
		apiV1.GET("/computation/usertokens", controllers.ReadUserTokens)
		apiV1.GET("/computation/token/:token_id", controllers.ReadToken)
		apiV1.POST("/computation/token/:token_id/start", controllers.StartComputation)
		apiV1.GET("/computation/queue", controllers.GetQueue)
	}

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}

	services.DeinitCelery()
}
