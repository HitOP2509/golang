package main

import "fmt"

/********** Pointer embedding **********/
// Pointer embedding in Go is conceptually similar to using a foreign key in DB.
// Both mean: “Don’t duplicate the full data. Keep a reference to the original thing.”
// If the original data changes, everyone referring to it sees the updated data.

// DB foreign key -> stores an ID/reference to another row.
// Go pointer     -> stores a memory address/reference to another object.

/********** Value embedding **********/
// Value embedding in Go means the struct keeps its own copy of the data.
// So even if the original data changes later, this struct still has the old copied data.

// This is useful when you want a snapshot-like behavior and do not want the data to be updated when the original data changes.

// Struct embedding / Struct composition
type User struct {
	name    string
	address string
}

type OrderDetails struct {
	id              string
	productId       string
	deliveryAddress string
	status          string
}
type Order struct {
	OrderDetails //Value embedding
	*User        //Pointer embedding
}

// Value receivers methods
// You can use value when just reading it
func (o Order) GetOrderDeliveryInfo() string {
	return fmt.Sprintf("Order status for %s is: %s", o.name, o.status)
}

// Pointer receiver method
// Use pointer receiver when the method needs to modify the original struct.
func (o *Order) UpdateDeliveryStatus(status string) {
	o.status = status
}

func main() {
	user := &User{
		name:    "Rohit Singh",
		address: "Chinarpark, Kolkata - 700059",
	}

	order := OrderDetails{
		id:              "1",
		productId:       "1",
		deliveryAddress: "Chinarpark, Kolkata - 700059",
		status:          "pending",
	}

	newOrder := Order{
		User:         user,
		OrderDetails: order,
	}
	newOrder2 := Order{
		User:         user,
		OrderDetails: order,
	}

	// This changes the original object and reflects on both the objects newOrder and newOrder2
	newOrder2.User.name = "Mr. Rohit Singh"

	//This change reflects inside only newOrder2
	newOrder2.UpdateDeliveryStatus("delivered")

	fmt.Println(newOrder.GetOrderDeliveryInfo())  //Order status for Mr. Rohit Singh is: pending
	fmt.Println(newOrder2.GetOrderDeliveryInfo()) //Order status for Mr. Rohit Singh is: delivered

	// WE CAN SEE THE NAME UPDATED IN BOTH THE OBJECTS DESPITE CHANGING IT IN ONLY 1 BUT DELIVERY STATUS IS DIFFERENT

}
