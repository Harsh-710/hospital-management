package schedule

import (
	"fmt"

	"github.com/Harsh-710/hospital-management/models"
)

func getAppointmentIDs(items []models.CartCheckoutItem) ([]int, error) {
	patientIds := make([]int, len(items))
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid number of patient] %d", item.PatientID)
		}

		patientIds[i] = item.PatientID
	}

	return patientIds, nil
}

func checkSchedule(appointments []models.CartCheckoutItem, products map[int]models.Patient) error {
	if len(appointments) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range appointments {
		patient, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d is not available in the store, please refresh your cart", item.ProductID)
		}

		if patient.Quantity < item.Quantity {
			return fmt.Errorf("product %s is not available in the quantity requested", patient.Name)
		}
	}

	return nil
}

func calculateTotalPrice(cartItems []types.CartCheckoutItem, products map[int]types.Product) float64 {
	var total float64

	for _, item := range cartItems {
		product := products[item.ProductID]
		total += product.Price * float64(item.Quantity)
	}

	return total
}

func (h *Handler) createOrder(products []types.Product, cartItems []types.CartCheckoutItem, userID int) (int, float64, error) {
	// create a map of products for easier access
	productsMap := make(map[int]models.Product)
	for _, product := range products {
		productsMap[product.ID] = product
	}

	// check if all products are available
	if err := checkSchedule(cartItems, productsMap); err != nil {
		return 0, 0, err
	}

	// calculate total price
	totalPrice := calculateTotalPrice(cartItems, productsMap)

	// reduce the quantity of products in the store
	for _, item := range cartItems {
		product := productsMap[item.ProductID]
		product.Quantity -= item.Quantity
		h.store.UpdateProduct(product)
	}

	// create order record
	orderID, err := h.orderStore.CreateOrder(models.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "some address", // could fetch address from a user addresses table
	})
	if err != nil {
		return 0, 0, err
	}

	// create order the items records
	for _, item := range cartItems {
		h.orderStore.CreateOrderItem(models.OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productsMap[item.ProductID].Price,
		})
	}

	return orderID, totalPrice, nil
}
