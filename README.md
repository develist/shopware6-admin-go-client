# Client for Shopware 6 Admin API

**Be aware, this package is currently under heavy development.**

## Create a repository

You have to create an integration in the Shopware administration panel. For more information please refer
https://docs.shopware.com/en/shopware-6-en/settings/system/integrationen.
Beside this token you need the toke URL and the base URL for the API as mentioned in this article
https://shopware.stoplight.io/docs/admin-api. 

    import (
        "context"
        sw6client "github.com/develist/shopware6-admin-go-client"
        "github.com/develist/shopware6-admin-go-client/entity"
        "golang.org/x/oauth2/clientcredentials"
        "net/http"
    )

    func main() {
        config := clientcredentials.Config{
            ClientID:     "<Your Access key ID>",
            ClientSecret: "<Your Secret access key>",
            TokenURL:     "https://<Your Shopware address>/api/oauth/token",
        }

        client = config.Client(context.Background())

        orderRepo := sw6client.CreateEntityRepository[entity.Order]("https://<Your Shopware domain>/api", client)
        // ....

        shippingRepo := sw6client.CreateEntityRepository[entity.ShippingMethod]("https://<Your Shopware domain>/api", client)
        // ....

    }

## Repository usage

### Fetch a specific order

    var order *entity.Order
    var err *entity.ApiErrors
    order, err = orderRepo.GetDetail("000376e68bbd44828d278b3b0df8f97c")

### Fetch a list of entities

    var shippingMethods *[]entity.ShippingMethod
    var err *entity.ApiErrors
    shippingMethods, err = shippingRepo.GetList()

### Search entities

	search := entity.Criteria{Page: 1, Limit: 100,
		Filter: []entity.FilterCriteria{{Type: entity.Range,
					Field: field.OrderDateTime,
					Parameters: map[entity.FilterParameter]string{entity.Gte: "2022-04-23 00:00:00", entity.Lt: "2022-04-23 15:00:00"}}},
		Sort:   []entity.SortCriteria{{Field: field.OrderNumber, Order: entity.Asc}},
		Associations: map[field.Field]entity.Criteria{
			field.BillingAddress: {},
			field.Deliveries:     {},
			field.Transactions: {Associations: map[field.Field]entity.Criteria{
				field.PaymentMethod: {}}},
		}}

	var orders *[]entity.Order
    var err *entity.ApiErrors
    orders, err = orderRepo.Search(search)

### Writing entities

	address := entity.CustomerAddress{FirstName: "MyFirstName",
		LastName:  "MyLastName",
		Street:    "MyStreet",
		City:      "MyCity",
		Zipcode:   "MyZipcode",
		CountryId: "f9edd8bdd916432397a2d4f11b92e424"}

	customer := entity.Customer{FirstName: "MyFirstName",
		LastName:               "MyLastName",
		CustomerNumber:         "1000",
		DefaultBillingAddress:  &address,
		Email:                  "MyEmail",
		DefaultPaymentMethodId: "23ee014007c94583b9c16fbe383b0d10",
		GroupId:                "cfbd5018d38d41d8adca10d94fc8bdd6",
		SalesChannelId:         "46d57c6bf6aa40a9811d88fe8e29d34b"}

    var err *entity.ApiErrors

	err = customerRepo.Create(&customer)
