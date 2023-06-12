package handler

import (
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w2-4/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-4/internal/section"
	"github.com/extmatperez/meli_bootcamp_go_w2-4/pkg/web"
	"github.com/gin-gonic/gin"
)

type Section struct {
	sectionService section.Service
}

func NewSection(s section.Service) *Section {
	return &Section{
		sectionService: s,
	}
}

func (s *Section) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		sections, err := s.sectionService.GetAll(c.Request.Context())
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Response(c, http.StatusOK, sections)
	}
}

func (s *Section) Get() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (s *Section) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto domain.CreateSection
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if dto.CurrentCapacity == 0 {
			web.Error(c, http.StatusBadRequest, "section number invalid")
			return
		}

		if dto.CurrentCapacity == 0 {
			web.Error(c, http.StatusBadRequest, "current temperature invalid")
			return
		}

		if dto.CurrentCapacity == 0 {
			web.Error(c, http.StatusBadRequest, "minimum temperature invalid")
			return
		}

		if dto.CurrentCapacity == 0 {
			web.Error(c, http.StatusBadRequest, "current capacity invalid")
			return
		}

		if dto.CurrentCapacity == 0 {
			web.Error(c, http.StatusBadRequest, "minimum capacity invalid")
			return

		}

		if dto.CurrentCapacity == 0 {
			web.Error(c, http.StatusBadRequest, "maximum capacity invalid")
			return
		}

		if dto.CurrentCapacity == 0 {
			web.Error(c, http.StatusBadRequest, "warehouse id invalid")
			return
		}

		if dto.CurrentCapacity == 0 {
			web.Error(c, http.StatusBadRequest, "product type id invalid")
			return
		}

		sec, err := s.sectionService.Save(c, dto)
		if err != nil {
			if err == section.ErrInvalidSectionNumber {
				web.Error(c, http.StatusConflict, err.Error())
			} else {
				web.Error(c, http.StatusInternalServerError, err.Error())
			}
			return
		}
		web.Response(c, http.StatusCreated, sec)

	}
}

func (s *Section) Update() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (s *Section) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
