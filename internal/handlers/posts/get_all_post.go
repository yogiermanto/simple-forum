package posts

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetAllPost(c *gin.Context) {
	ctx := c.Request.Context()

	pageIndexStr := c.Query("page_index")
	pageIndex, err := strconv.Atoi(pageIndexStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("page_index is not valid").Error()})
		return
	}

	pageSizeStr := c.Query("page_size")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("page_limit is not valid").Error()})
		return
	}
	if pageIndex > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("page_index can't greater than 100").Error()})
		return
	}

	resp, err := h.postSvc.GetALlPost(ctx, pageSize, pageIndex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
