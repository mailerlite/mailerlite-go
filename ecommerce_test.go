package mailerlite_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/mailerlite/mailerlite-go"
	"github.com/stretchr/testify/assert"
)

const shopURL = "https://connect.mailerlite.com/api/ecommerce/shops/1"

func TestEcommerceRequests(t *testing.T) {
	ctx := context.TODO()
	price := 19.99
	variant := mailerlite.ProductVariantParams{Name: "Large", Attributes: map[string]string{"size": "L"}}

	cases := []struct {
		name   string
		call   func(c *mailerlite.Client) error
		method string
		url    string
		body   string
	}{
		{"shop list", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.Shop.List(ctx, &mailerlite.ListEcommerceOptions{Limit: 10})
			return err
		}, http.MethodGet, "https://connect.mailerlite.com/api/ecommerce/shops?limit=10", ""},
		{"shop get", func(c *mailerlite.Client) error { _, _, err := c.Ecommerce.Shop.Get(ctx, "1"); return err }, http.MethodGet, shopURL, ""},
		{"shop create", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.Shop.Create(ctx, &mailerlite.ShopParams{Name: "Shop", Url: "https://shop.test", Currency: "EUR"})
			return err
		}, http.MethodPost, "https://connect.mailerlite.com/api/ecommerce/shops", `{"name":"Shop","url":"https://shop.test","currency":"EUR"}`},
		{"shop update", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.Shop.Update(ctx, "1", &mailerlite.ShopParams{Name: "New"})
			return err
		}, http.MethodPut, shopURL, `{"name":"New"}`},
		{"shop delete", func(c *mailerlite.Client) error { _, err := c.Ecommerce.Shop.Delete(ctx, "1"); return err }, http.MethodDelete, shopURL, ""},
		{"shop stats", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.Shop.Stats(ctx, "1", &mailerlite.ShopStatsOptions{StartDate: "2026-10-01", EndDate: "2026-10-31"})
			return err
		}, http.MethodGet, shopURL + "/reports/stats?end_date=2026-10-31&start_date=2026-10-01", ""},
		{"settings get", func(c *mailerlite.Client) error { _, _, err := c.Ecommerce.Setting.Get(ctx); return err }, http.MethodGet, "https://connect.mailerlite.com/api/ecommerce/settings", ""},

		{"product list", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.Product.List(ctx, "1", &mailerlite.ListProductOptions{
				Filters: &[]mailerlite.Filter{{Name: "enabled", Value: true}},
				Sort:    "-name",
			})
			return err
		}, http.MethodGet, shopURL + "/products?filter%5Benabled%5D=true&sort=-name", ""},
		{"product get", func(c *mailerlite.Client) error { _, _, err := c.Ecommerce.Product.Get(ctx, "1", "7"); return err }, http.MethodGet, shopURL + "/products/7", ""},
		{"product create", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.Product.Create(ctx, "1", &mailerlite.ProductParams{Name: "Hat", Price: &price})
			return err
		}, http.MethodPost, shopURL + "/products", `{"name":"Hat","price":19.99}`},
		{"product update", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.Product.Update(ctx, "1", "7", &mailerlite.ProductParams{Name: "Cap"})
			return err
		}, http.MethodPut, shopURL + "/products/7", `{"name":"Cap"}`},
		{"product delete", func(c *mailerlite.Client) error { _, err := c.Ecommerce.Product.Delete(ctx, "1", "7"); return err }, http.MethodDelete, shopURL + "/products/7", ""},
		{"product change status", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.Product.ChangeStatus(ctx, "1", "7", false)
			return err
		}, http.MethodPut, shopURL + "/products/7/change-status", `{"status":false}`},
		{"product assign categories", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.Product.AssignCategories(ctx, "1", "7", []string{"5"}, true)
			return err
		}, http.MethodPut, shopURL + "/products/7/categories/multiple", `{"categories":["5"],"replace":true}`},

		{"variant list", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.ProductVariant.List(ctx, "1", "7")
			return err
		}, http.MethodGet, shopURL + "/products/7/variants", ""},
		{"variant sync", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.ProductVariant.Sync(ctx, "1", "7", []mailerlite.ProductVariantParams{variant})
			return err
		}, http.MethodPost, shopURL + "/products/7/variants", `[{"name":"Large","attributes":{"size":"L"}}]`},
		{"variant update", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.ProductVariant.Update(ctx, "1", "7", "3", &variant)
			return err
		}, http.MethodPut, shopURL + "/products/7/variants/3", `{"name":"Large","attributes":{"size":"L"}}`},
		{"variant delete", func(c *mailerlite.Client) error {
			_, err := c.Ecommerce.ProductVariant.Delete(ctx, "1", "7", "3")
			return err
		}, http.MethodDelete, shopURL + "/products/7/variants/3", ""},

		{"category list", func(c *mailerlite.Client) error { _, _, err := c.Ecommerce.Category.List(ctx, "1", nil); return err }, http.MethodGet, shopURL + "/categories", ""},
		{"category get", func(c *mailerlite.Client) error { _, _, err := c.Ecommerce.Category.Get(ctx, "1", "5"); return err }, http.MethodGet, shopURL + "/categories/5", ""},
		{"category create", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.Category.Create(ctx, "1", &mailerlite.CategoryParams{Name: "Hats"})
			return err
		}, http.MethodPost, shopURL + "/categories", `{"name":"Hats"}`},
		{"category update", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.Category.Update(ctx, "1", "5", &mailerlite.CategoryParams{Name: "Caps"})
			return err
		}, http.MethodPut, shopURL + "/categories/5", `{"name":"Caps"}`},
		{"category delete", func(c *mailerlite.Client) error { _, err := c.Ecommerce.Category.Delete(ctx, "1", "5"); return err }, http.MethodDelete, shopURL + "/categories/5", ""},
		{"category products", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.Category.Products(ctx, "1", "5", nil)
			return err
		}, http.MethodGet, shopURL + "/categories/5/products", ""},
		{"category assign product", func(c *mailerlite.Client) error {
			_, err := c.Ecommerce.Category.AssignProduct(ctx, "1", "5", "7")
			return err
		}, http.MethodPut, shopURL + "/categories/5/products/7", ""},
		{"category unassign product", func(c *mailerlite.Client) error {
			_, err := c.Ecommerce.Category.UnassignProduct(ctx, "1", "5", "7")
			return err
		}, http.MethodDelete, shopURL + "/categories/5/products/7", ""},

		{"customer list", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.Customer.List(ctx, "1", &mailerlite.ListEcommerceOptions{Filters: &[]mailerlite.Filter{{Name: "email", Value: "a@b.test"}}})
			return err
		}, http.MethodGet, shopURL + "/customers?filter%5Bemail%5D=a%40b.test", ""},
		{"customer get", func(c *mailerlite.Client) error { _, _, err := c.Ecommerce.Customer.Get(ctx, "1", "3"); return err }, http.MethodGet, shopURL + "/customers/3", ""},
		{"customer create", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.Customer.Create(ctx, "1", &mailerlite.CustomerParams{Email: "a@b.test"})
			return err
		}, http.MethodPost, shopURL + "/customers", `{"email":"a@b.test"}`},
		{"customer update", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.Customer.Update(ctx, "1", "3", &mailerlite.CustomerParams{Name: "Jane"})
			return err
		}, http.MethodPut, shopURL + "/customers/3", `{"name":"Jane"}`},
		{"customer delete", func(c *mailerlite.Client) error { _, err := c.Ecommerce.Customer.Delete(ctx, "1", "3"); return err }, http.MethodDelete, shopURL + "/customers/3", ""},

		{"order list", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.Order.List(ctx, "1", &mailerlite.ListEcommerceOptions{Filters: &[]mailerlite.Filter{{Name: "email_id", Value: "55"}}})
			return err
		}, http.MethodGet, shopURL + "/orders?filter%5Bemail_id%5D=55", ""},
		{"order get", func(c *mailerlite.Client) error { _, _, err := c.Ecommerce.Order.Get(ctx, "1", "9"); return err }, http.MethodGet, shopURL + "/orders/9", ""},
		{"order update", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.Order.Update(ctx, "1", "9", &mailerlite.UpdateOrderParams{Status: "complete"})
			return err
		}, http.MethodPut, shopURL + "/orders/9", `{"status":"complete"}`},
		{"order delete", func(c *mailerlite.Client) error { _, err := c.Ecommerce.Order.Delete(ctx, "1", "9"); return err }, http.MethodDelete, shopURL + "/orders/9", ""},

		{"cart list", func(c *mailerlite.Client) error { _, _, err := c.Ecommerce.Cart.List(ctx, "1", nil); return err }, http.MethodGet, shopURL + "/carts", ""},
		{"cart get", func(c *mailerlite.Client) error { _, _, err := c.Ecommerce.Cart.Get(ctx, "1", "4"); return err }, http.MethodGet, shopURL + "/carts/4", ""},
		{"cart update", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.Cart.Update(ctx, "1", "4", &mailerlite.UpdateCartParams{CheckoutUrl: "https://shop.test/cart"})
			return err
		}, http.MethodPut, shopURL + "/carts/4", `{"checkout_url":"https://shop.test/cart"}`},
		{"cart item list", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.CartItem.List(ctx, "1", "4", nil)
			return err
		}, http.MethodGet, shopURL + "/carts/4/items", ""},
		{"cart item get", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.CartItem.Get(ctx, "1", "4", "8")
			return err
		}, http.MethodGet, shopURL + "/carts/4/items/8", ""},
		{"cart item create", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.CartItem.Create(ctx, "1", "4", &mailerlite.CartItemParams{EcommerceProductID: "7"})
			return err
		}, http.MethodPost, shopURL + "/carts/4/items", `{"ecommerce_product_id":"7"}`},
		{"cart item update", func(c *mailerlite.Client) error {
			_, _, err := c.Ecommerce.CartItem.Update(ctx, "1", "4", "8", &mailerlite.CartItemParams{Quantity: 2})
			return err
		}, http.MethodPut, shopURL + "/carts/4/items/8", `{"quantity":2}`},
		{"cart item delete", func(c *mailerlite.Client) error {
			_, err := c.Ecommerce.CartItem.Delete(ctx, "1", "4", "8")
			return err
		}, http.MethodDelete, shopURL + "/carts/4/items/8", ""},

		{"import products", func(c *mailerlite.Client) error {
			_, err := c.Ecommerce.Import.Products(ctx, "1", []mailerlite.ProductParams{{Name: "Hat"}})
			return err
		}, http.MethodPost, shopURL + "/products/import", `[{"name":"Hat"}]`},
		{"import categories", func(c *mailerlite.Client) error {
			_, err := c.Ecommerce.Import.Categories(ctx, "1", []mailerlite.CategoryParams{{Name: "Hats"}})
			return err
		}, http.MethodPost, shopURL + "/categories/import", `[{"name":"Hats"}]`},
		{"import customers", func(c *mailerlite.Client) error {
			_, err := c.Ecommerce.Import.Customers(ctx, "1", []mailerlite.CustomerParams{{Email: "a@b.test"}})
			return err
		}, http.MethodPost, shopURL + "/customers/import", `[{"email":"a@b.test"}]`},
		{"import orders", func(c *mailerlite.Client) error {
			_, err := c.Ecommerce.Import.Orders(ctx, "1", []mailerlite.CreateOrderParams{{
				Customer: mailerlite.CustomerParams{Email: "a@b.test"},
				Cart:     mailerlite.OrderCartParams{Items: []mailerlite.CartItemParams{{EcommerceProductID: "7"}}},
			}})
			return err
		}, http.MethodPost, shopURL + "/orders/import", `{"orders":[{"customer":{"email":"a@b.test"},"cart":{"items":[{"ecommerce_product_id":"7"}]}}]}`},

		{"campaign ecommerce stats", func(c *mailerlite.Client) error {
			_, _, err := c.Campaign.EcommerceStats(ctx, "2", &mailerlite.EcommerceReportOptions{ShopID: "1"})
			return err
		}, http.MethodGet, "https://connect.mailerlite.com/api/campaigns/2/reports/ecommerce?shop_id=1", ""},
		{"campaign ecommerce orders", func(c *mailerlite.Client) error {
			_, _, err := c.Campaign.EcommerceOrders(ctx, "2", &mailerlite.EcommerceReportOptions{ShopID: "1", Page: 2})
			return err
		}, http.MethodGet, "https://connect.mailerlite.com/api/campaigns/2/reports/ecommerce/orders?page=2&shop_id=1", ""},
		{"subscriber ecommerce stats", func(c *mailerlite.Client) error {
			_, _, err := c.Subscriber.EcommerceStats(ctx, "6", &mailerlite.EcommerceReportOptions{ShopID: "1"})
			return err
		}, http.MethodGet, "https://connect.mailerlite.com/api/subscribers/6/reports/ecommerce?shop_id=1", ""},
		{"subscriber ecommerce orders", func(c *mailerlite.Client) error {
			_, _, err := c.Subscriber.EcommerceOrders(ctx, "6", &mailerlite.EcommerceReportOptions{ShopID: "1"})
			return err
		}, http.MethodGet, "https://connect.mailerlite.com/api/subscribers/6/reports/ecommerce/orders?shop_id=1", ""},
		{"automation ecommerce stats", func(c *mailerlite.Client) error {
			_, _, err := c.Automation.EcommerceStats(ctx, "8", nil)
			return err
		}, http.MethodGet, "https://connect.mailerlite.com/api/automations/8/monthly-totals/ecommerce", ""},
		{"automation email ecommerce stats", func(c *mailerlite.Client) error {
			_, _, err := c.Automation.EmailEcommerceStats(ctx, "8", "9", &mailerlite.AutomationEcommerceOptions{ShopID: "1"})
			return err
		}, http.MethodGet, "https://connect.mailerlite.com/api/automations/8/emails/9/ecommerce-overview-stats?shop_id=1", ""},
		{"automation ecommerce orders", func(c *mailerlite.Client) error {
			_, _, err := c.Automation.EcommerceOrders(ctx, "8", &mailerlite.AutomationEcommerceOrderOptions{TimeFrame: "all_time"})
			return err
		}, http.MethodGet, "https://connect.mailerlite.com/api/automations/8/order-activity?time_frame=all_time", ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client := mailerlite.NewClient(testKey)
			client.SetHttpClient(NewTestClient(func(req *http.Request) *http.Response {
				assert.Equal(t, tc.method, req.Method)
				assert.Equal(t, tc.url, req.URL.String())
				if tc.body != "" {
					body, _ := io.ReadAll(req.Body)
					assert.JSONEq(t, tc.body, string(body))
				}
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString(`{}`)),
				}
			}))

			assert.NoError(t, tc.call(client))
		})
	}
}

func TestAmountDecodesNumbersAndStrings(t *testing.T) {
	var product mailerlite.Product

	assert.NoError(t, json.Unmarshal([]byte(`{"price":"19.99"}`), &product))
	assert.Equal(t, mailerlite.Amount(19.99), product.Price)

	assert.NoError(t, json.Unmarshal([]byte(`{"price":24.5}`), &product))
	assert.Equal(t, mailerlite.Amount(24.5), product.Price)

	assert.NoError(t, json.Unmarshal([]byte(`{"price":null}`), &product))
}
