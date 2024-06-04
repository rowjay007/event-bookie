package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/rowjay007/event-bookie/internal/models"
    "github.com/rowjay007/event-bookie/internal/service/payment"
)

type PaymentHandler struct {
    PaymentService *service.PaymentService 
}

func NewPaymentHandler(paymentService *service.PaymentService) *PaymentHandler {
    return &PaymentHandler{PaymentService: paymentService}
}

// CreatePayment godoc
// @Summary Create a new payment
// @Description Create a new payment
// @Tags payments
// @Accept json
// @Produce json
// @Param payment body models.Payment true "Payment information"
// @Success 201 {object} models.Payment
// @Failure 400 {object} gin.H "Payment information is invalid"
// @Failure 500 {object} gin.H "Failed to create payment"
// @Router /api/v1/payments [post]
func (ph *PaymentHandler) CreatePayment(c *gin.Context) {
    var payment models.Payment
    if err := c.ShouldBindJSON(&payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := ph.PaymentService.CreatePayment(&payment); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, payment)
}

// GetAllPayments godoc
// @Summary Get all payments
// @Description Get all payments with optional filtering and sorting
// @Tags payments
// @Produce json
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Param sort query string false "Sort by field"
// @Param order query string false "Order (asc or desc)"
// @Param filter query string false "Filter by field:value"
// @Success 200 {object} gin.H "An object containing payments and total count"
// @Failure 400 {object} gin.H "Invalid query parameters"
// @Failure 500 {object} gin.H "Failed to retrieve payments"
// @Router /api/v1/payments [get]
func (ph *PaymentHandler) GetAllPayments(c *gin.Context) {
    queryParams := c.Request.URL.Query()
    offsetStr := queryParams.Get("offset")
    limitStr := queryParams.Get("limit")
    sort := queryParams.Get("sort")
    order := queryParams.Get("order")
    filter := queryParams.Get("filter")

    var offset, limit int
    var err error

    if offsetStr != "" {
        offset, err = strconv.Atoi(offsetStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
            return
        }
    } else {
        offset = -1
    }

    if limitStr != "" {
        limit, err = strconv.Atoi(limitStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
            return
        }
    } else {
        limit = -1
    }

    params := make(map[string]string)
    for key, values := range queryParams {
        if len(values) > 0 {
            params[key] = values[0]
        }
    }

    payments, total, err := ph.PaymentService.GetAllPayments(params, offset, limit, sort, order, filter)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch payments"})
        return
    }

    response := gin.H{
        "payments": payments,
        "total":    total,
    }

    c.JSON(http.StatusOK, response)
}


// GetPaymentByID godoc
// @Summary Get a payment by ID
// @Description Get a payment by ID
// @Tags payments
// @Produce json
// @Param id path int true "Payment ID"
// @Success 200 {object} models.Payment
// @Failure 400 {object} gin.H "Invalid ID"
// @Failure 404 {object} gin.H "Payment not found"
// @Failure 500 {object} gin.H "Failed to retrieve payment"
// @Router /api/v1/payments/{id} [get]
func (ph *PaymentHandler) GetPaymentByID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    payment, err := ph.PaymentService.GetPaymentByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
        return
    }

    c.JSON(http.StatusOK, payment)
}

// UpdatePayment godoc
// @Summary Update a payment
// @Description Update a payment
// @Tags payments
// @Accept json
// @Produce json
// @Param payment body models.Payment true "Updated payment information"
// @Success 200 {object} models.Payment
// @Failure 400 {object} gin.H "Payment information is invalid"
// @Failure 500 {object} gin.H "Failed to update payment"
// @Router /api/v1/payments/{id} [put]
func (ph *PaymentHandler) UpdatePayment(c *gin.Context) {
    var payment models.Payment
    if err := c.ShouldBindJSON(&payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := ph.PaymentService.UpdatePayment(&payment); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update payment"})
        return
    }

    c.JSON(http.StatusOK, payment)
}

// DeletePayment godoc
// @Summary Delete a payment
// @Description Delete a payment
// @Tags payments
// @Produce json
// @Param id path int true "Payment ID"
// @Success 200 {object} gin.H "message": "Payment deleted successfully"
// @Failure 400 {object} gin.H "Invalid ID"
// @Failure 500 {object} gin.H "Failed to delete payment"
// @Router /api/v1/payments/{id} [delete]
func (ph *PaymentHandler) DeletePayment(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := ph.PaymentService.DeletePayment(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete payment"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Payment deleted successfully"})
}
