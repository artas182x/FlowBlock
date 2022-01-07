package controllers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/artas182x/hyperledger-fabric-master-thesis/backend/models"
	"github.com/artas182x/hyperledger-fabric-master-thesis/backend/services"
	"github.com/gin-gonic/gin"
)

// @Summary Login
// @Schemes
// @Accept json
// @Produce json
// @Success 200
// @Tags login
// @Param login body models.Login true "User data"
// @Router /login [post]
func Authenticate(c *gin.Context) (*models.User, error) {
	var loginVals models.Login
	if err := c.ShouldBind(&loginVals); err != nil {
		return nil, jwt.ErrMissingLoginValues
	}

	_, err := services.GetNetwork(loginVals)

	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	user, err := services.LoginToUser(loginVals)

	if err != nil {
		return nil, err
	}

	return user, nil
}
