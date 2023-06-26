package handler

import (
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-4/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-4/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-4/pkg/web"
	"github.com/gin-gonic/gin"
)

type Buyer struct {
	buyerService buyer.Service
}

func NewBuyer(b buyer.Service) *Buyer {
	return &Buyer{
		buyerService: b,
	}
}

// Buyer GoDoc
//
//	@Summary		Get a buyer by ID
//	@Description	Get a buyer by ID
//	@Tags			Buyers
//	@Param			id	path		int	true	"Buyer ID"
//	@Success		200	{object}	domain.Buyer
//	@Failure		400	{string}	string	"Invalid ID"
//	@Failure		404	{string}	string	"Buyer not found"
//	@Router			/api/v1/buyers/{id} [get]
func (b *Buyer) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "invalid ID")
			return
		}
		buyer, err := b.buyerService.Get(c, id)
		if err != nil {
			web.Error(c, http.StatusNotFound, "buyer not found")
			return
		}
		web.Success(c, http.StatusOK, buyer)
	}
}

// Buyer GoDoc
//
//	@Summary		Delete a buyer by ID
//	@Description	Delete a buyer by ID
//	@Tags			Buyers
//	@Param			id	path		int		true	"Buyer ID"
//	@Success		200	{string}	string	"Buyer deleted"
//	@Failure		400	{string}	string	"Invalid ID"
//	@Failure		404	{string}	string	"Buyer not found"
//	@Router			/api/v1/buyers/{id} [delete]
func (b *Buyer) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "invalid ID")
			return
		}
		err = b.buyerService.Delete(c, id)
		if err != nil {
			web.Error(c, http.StatusNotFound, "buyer not found")
			return
		}
		web.Success(c, http.StatusOK, "buyer deleted")
	}
}

// Buyer GoDoc
//
//	@Summary		Get all buyers
//	@Description	Get all buyers
//	@Tags			Buyers
//	@Success		200	{array}		domain.Buyer
//	@Failure		500	{string}	string	"Buyer not found"
//	@Failure		204	{string}	string	"No buyers found"
//	@Router			/api/v1/buyers [get]
func (b *Buyer) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		buyers, err := b.buyerService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "buyer not found")
			return
		}
		if len(buyers) == 0 {
			web.Success(c, http.StatusNoContent, buyers)
			return
		}
		web.Success(c, http.StatusOK, buyers)
	}
}

// Buyer GoDoc
//
//	@Summary		Create a new buyer
//	@Description	Create a new buyer
//	@Tags			Buyers
//	@Accept			json
//	@Param			buyer	body		domain.BuyerCreate	true	"Buyer object"
//	@Success		201		{object}	domain.Buyer
//	@Failure		422		{string}	string	"Buyer not created"
//	@Router			/api/v1/buyers [post]
func (b *Buyer) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var buyer domain.BuyerCreate
		if err := c.ShouldBindJSON(&buyer); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "buyer not created")
			return
		}
		buyerF, err := b.buyerService.Create(c.Request.Context(), buyer)
		if err != nil {
			web.Error(c, http.StatusConflict, "buyer not created")
			return
		}
		web.Response(c, http.StatusCreated, buyerF)
	}
}

// Buyer GoDoc
//
//	@Summary		Update a buyer by ID
//	@Description	Update a buyer by ID
//	@Tags			Buyers
//	@Accept			json
//	@Param			id		path		int				true	"Buyer ID"
//	@Param			buyer	body		domain.Buyer	true	"Buyer object"
//	@Success		200		{object}	domain.Buyer
//	@Failure		400		{string}	string	"Invalid ID"
//	@Failure		404		{string}	string	"Buyer not updated"
//	@Router			/api/v1/buyers/{id} [patch]
func (b *Buyer) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var buyer domain.Buyer
		if err := c.ShouldBindJSON(&buyer); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "buyer not created")
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "invalid ID")
			return
		}
		buyerUpdated, err := b.buyerService.Update(c, buyer, id)
		if err != nil {
			web.Error(c, http.StatusNotFound, "buyer not updated")
			return
		}
		web.Success(c, http.StatusOK, buyerUpdated)
	}
}
