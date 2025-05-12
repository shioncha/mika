package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shioncha/mika/backend/ent"
	"github.com/shioncha/mika/backend/ent/users"
	"github.com/shioncha/mika/backend/internal/auth"
	"github.com/shioncha/mika/backend/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type SignUpRequest struct {
	Email           string `json:"email" binding:"required,email"`
	Name            string `json:"name" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" binding:"required,eqfield=Password"`
}

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type AuthResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func SignUp(c *gin.Context, client *ent.Client) {
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	ctx := c.Request.Context()
	lowerEmail := auth.NormarlizeEmail(req.Email)

	if isExist, err := client.Users.Query().Where(users.EmailEQ(lowerEmail)).Exist(ctx); err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	} else if isExist {
		respondWithError(c, http.StatusConflict, "Email already registered")
		return
	}

	hashedPassword, err := auth.GenerateHashedPassword(req.Password)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	id := utils.GenerateULID()

	tx, err := client.Tx(ctx)
	if err != nil {
		log.Printf("Failed to start transaction: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	_, err = tx.Users.
		Create().
		SetUlid(id).
		SetEmail(lowerEmail).
		SetName(req.Name).
		SetPasswordHash(string(hashedPassword)).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		respondWithError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	respondWithTokens(c, id)
}

func SignIn(c *gin.Context, client *ent.Client) {
	var req SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	ctx := c.Request.Context()
	lowerEmail := auth.NormarlizeEmail(req.Email)

	user, err := client.Users.Query().Where(users.EmailEQ(lowerEmail)).Select(users.FieldPasswordHash, users.FieldUlid).First(ctx)
	if err != nil && ent.IsNotFound(err) {
		respondWithError(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		respondWithError(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	respondWithTokens(c, user.Ulid)
}

func respondWithError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

func respondWithTokens(c *gin.Context, id string) {
	token, _ := auth.GenerateJWT(id)
	refreshToken := "refresh_token"

	c.JSON(200, AuthResponse{
		Token:        token,
		RefreshToken: refreshToken,
	})
}
