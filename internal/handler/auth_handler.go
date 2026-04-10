package handler

import (
	"database/sql"
	"net/http"
	"os"
	"rogeriods/contact-api/internal/model"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// var jwtKey = []byte("your_secret_key")
var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type AuthHandler struct {
	DB *sql.DB
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{DB: db}
}

// Register handles user registration
func (h *AuthHandler) Register(c *gin.Context) {
	var creds model.LoginRequest

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Hash the user's password using bcrypt
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(creds.Password), 12)

	_, err := h.DB.Exec("INSERT INTO users (username, password_hash) VALUES (?, ?)",
		creds.Username, string(hashedPassword))

	// Handle database errors
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

// Login authenticates a user and returns a JWT token
func (h *AuthHandler) Login(c *gin.Context) {
	var creds model.LoginRequest

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var storedHash string
	var userID int

	err := h.DB.QueryRow("SELECT id, password_hash FROM users WHERE username = ?", creds.Username).
		Scan(&userID, &storedHash)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare provided password with stored hash
	if bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(creds.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Set token expiration time (24 hours)
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create JWT claims (payload)
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtKey)

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}