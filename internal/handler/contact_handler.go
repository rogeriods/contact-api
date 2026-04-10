package handler

import (
	"database/sql"
	"net/http"
	"rogeriods/contact-api/internal/model"

	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	DB *sql.DB
}

func NewContactHandler(db *sql.DB) *ContactHandler {
	return &ContactHandler{DB: db}
}

// Get all contacts by user logged
func (h *ContactHandler) GetContacts(c *gin.Context) {
	// Get userID in Gin context
	userID := c.GetInt("userID")

	rows, err := h.DB.Query("SELECT id, name, phone FROM contacts WHERE user_id = ?", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}
	defer rows.Close()

	var contacts []model.Contact

	for rows.Next() {
		var ct model.Contact
		rows.Scan(&ct.ID, &ct.Name, &ct.Phone)
		contacts = append(contacts, ct)
	}

	if contacts == nil {
		contacts = []model.Contact{}
	}

	c.JSON(http.StatusOK, contacts)
}

// Create new contact
func (h *ContactHandler) Create(c *gin.Context) {
	// Get userID in Gin context
	userID := c.GetInt("userID")

	var contact model.Contact

	if c.ShouldBindJSON(&contact) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := h.DB.Exec("INSERT INTO contacts (user_id, name, phone) VALUES (?, ?, ?)",
		userID, contact.Name, contact.Phone)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Insert failed"})
		return
	}

	c.Status(http.StatusCreated)
}

// Get contact by ID and user logged
func (h *ContactHandler) GetByID(c *gin.Context) {
	// Get userID in Gin context
	userID := c.GetInt("userID")
	id := c.Param("id")

	var contact model.Contact

	err := h.DB.QueryRow(
		"SELECT id, name, phone FROM contacts WHERE id = ? AND user_id = ?",
		id, userID,
	).Scan(&contact.ID, &contact.Name, &contact.Phone)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	c.JSON(http.StatusOK, contact)
}

// Update contact by ID and user logged
func (h *ContactHandler) Update(c *gin.Context) {
	// Get userID in Gin context
	userID := c.GetInt("userID")
	id := c.Param("id")

	var contact model.Contact
	c.ShouldBindJSON(&contact)

	res, err := h.DB.Exec(
		"UPDATE contacts SET name = ?, phone = ? WHERE id = ? AND user_id = ?",
		contact.Name, contact.Phone, id, userID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated"})
}

// Delete contact by ID and user logged
func (h *ContactHandler) Delete(c *gin.Context) {
	// Get userID in Gin context
	userID := c.GetInt("userID")
	id := c.Param("id")

	res, err := h.DB.Exec("DELETE FROM contacts WHERE id = ? AND user_id = ?", id, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	c.Status(http.StatusNoContent)
}