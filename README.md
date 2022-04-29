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

        orderRepo = sw6client.CreateRepository[entity.Order]("https://<Your Shopware address>/api", client)
        // ....

        shippingRepo = sw6client.CreateRepository[entity.ShippingMethod]("https://<Your Shopware address>/api", client)
        // ....

    }

## Repository usage

At the moment only read access is implemented

### Fetch a specific order

    var order *entity.Order
    var err error
    order, err = orderRepo.GetDetail("000376e68bbd44828d278b3b0df8f97c")

### Fetch a list of entities

    var shippingMethods *[]entity.ShippingMethod
    var err error
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
    var err error
    orders, err = orderRepo.Search(search)
