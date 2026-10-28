package mailerlite

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const ecommerceShopEndpoint = "/ecommerce/shops"

// EcommerceService groups the e-commerce services.
type EcommerceService struct {
	Shop           ShopService
	Product        ProductService
	ProductVariant ProductVariantService
	Category       CategoryService
	Customer       CustomerService
	Order          OrderService
	Cart           CartService
	CartItem       CartItemService
	Import         EcommerceImportService
	Setting        EcommerceSettingService
}

type ShopService interface {
	List(ctx context.Context, options *ListEcommerceOptions) (*RootShops, *Response, error)
	Get(ctx context.Context, shopID string) (*RootShop, *Response, error)
	Create(ctx context.Context, shop *ShopParams) (*RootShop, *Response, error)
	Update(ctx context.Context, shopID string, shop *ShopParams) (*RootShop, *Response, error)
	Delete(ctx context.Context, shopID string) (*Response, error)
	Stats(ctx context.Context, shopID string, options *ShopStatsOptions) (*RootShopStats, *Response, error)
}

type ProductService interface {
	List(ctx context.Context, shopID string, options *ListProductOptions) (*RootProducts, *Response, error)
	Get(ctx context.Context, shopID, productID string) (*RootProduct, *Response, error)
	Create(ctx context.Context, shopID string, product *ProductParams) (*RootProduct, *Response, error)
	Update(ctx context.Context, shopID, productID string, product *ProductParams) (*RootProduct, *Response, error)
	Delete(ctx context.Context, shopID, productID string) (*Response, error)
	ChangeStatus(ctx context.Context, shopID, productID string, enabled bool) (*RootProduct, *Response, error)
	AssignCategories(ctx context.Context, shopID, productID string, categoryIDs []string, replace bool) (*RootProduct, *Response, error)
}

type ProductVariantService interface {
	List(ctx context.Context, shopID, productID string) (*RootProductVariants, *Response, error)
	Sync(ctx context.Context, shopID, productID string, variants []ProductVariantParams) (*RootProduct, *Response, error)
	Update(ctx context.Context, shopID, productID, variantID string, variant *ProductVariantParams) (*RootProduct, *Response, error)
	Delete(ctx context.Context, shopID, productID, variantID string) (*Response, error)
}

type CategoryService interface {
	List(ctx context.Context, shopID string, options *ListEcommerceOptions) (*RootCategories, *Response, error)
	Get(ctx context.Context, shopID, categoryID string) (*RootCategory, *Response, error)
	Create(ctx context.Context, shopID string, category *CategoryParams) (*RootCategory, *Response, error)
	Update(ctx context.Context, shopID, categoryID string, category *CategoryParams) (*RootCategory, *Response, error)
	Delete(ctx context.Context, shopID, categoryID string) (*Response, error)
	Products(ctx context.Context, shopID, categoryID string, options *ListEcommerceOptions) (*RootProducts, *Response, error)
	AssignProduct(ctx context.Context, shopID, categoryID, productID string) (*Response, error)
	UnassignProduct(ctx context.Context, shopID, categoryID, productID string) (*Response, error)
}

type CustomerService interface {
	List(ctx context.Context, shopID string, options *ListEcommerceOptions) (*RootCustomers, *Response, error)
	Get(ctx context.Context, shopID, customerID string) (*RootCustomer, *Response, error)
	Create(ctx context.Context, shopID string, customer *CustomerParams) (*RootCustomer, *Response, error)
	Update(ctx context.Context, shopID, customerID string, customer *CustomerParams) (*RootCustomer, *Response, error)
	Delete(ctx context.Context, shopID, customerID string) (*Response, error)
}

type OrderService interface {
	List(ctx context.Context, shopID string, options *ListEcommerceOptions) (*RootOrders, *Response, error)
	Get(ctx context.Context, shopID, orderID string) (*RootOrder, *Response, error)
	Create(ctx context.Context, shopID string, order *CreateOrderParams) (*RootOrder, *Response, error)
	Update(ctx context.Context, shopID, orderID string, order *UpdateOrderParams) (*RootOrder, *Response, error)
	Delete(ctx context.Context, shopID, orderID string) (*Response, error)
}

type CartService interface {
	List(ctx context.Context, shopID string, options *ListEcommerceOptions) (*RootCarts, *Response, error)
	Get(ctx context.Context, shopID, cartID string) (*RootCart, *Response, error)
	Update(ctx context.Context, shopID, cartID string, cart *UpdateCartParams) (*RootCart, *Response, error)
}

type CartItemService interface {
	List(ctx context.Context, shopID, cartID string, options *ListEcommerceOptions) (*RootCartItems, *Response, error)
	Get(ctx context.Context, shopID, cartID, itemID string) (*RootCartItem, *Response, error)
	Create(ctx context.Context, shopID, cartID string, item *CartItemParams) (*RootCartItem, *Response, error)
	Update(ctx context.Context, shopID, cartID, itemID string, item *CartItemParams) (*RootCartItem, *Response, error)
	Delete(ctx context.Context, shopID, cartID, itemID string) (*Response, error)
}

type EcommerceImportService interface {
	Products(ctx context.Context, shopID string, products []ProductParams) (*Response, error)
	Categories(ctx context.Context, shopID string, categories []CategoryParams) (*Response, error)
	Customers(ctx context.Context, shopID string, customers []CustomerParams) (*Response, error)
	Orders(ctx context.Context, shopID string, orders []CreateOrderParams) (*Response, error)
}

type EcommerceSettingService interface {
	Get(ctx context.Context) (*RootEcommerceSettings, *Response, error)
}

type RootShop struct {
	Data Shop `json:"data"`
}

type RootProduct struct {
	Data Product `json:"data"`
}

type RootCategory struct {
	Data Category `json:"data"`
}

type RootCustomer struct {
	Data Customer `json:"data"`
}

type RootOrder struct {
	Data Order `json:"data"`
}

type RootCart struct {
	Data Cart `json:"data"`
}

type RootCartItem struct {
	Data CartItem `json:"data"`
}

type RootEcommerceSettings struct {
	Data EcommerceSettings `json:"data"`
}

type RootShopStats struct {
	Data ShopStats `json:"data"`
}

type RootEcommerceReport struct {
	Data EcommerceReport `json:"data"`
}

type RootShops struct {
	Data  []Shop `json:"data"`
	Links Links  `json:"links"`
	Meta  Meta   `json:"meta"`
}

type RootProducts struct {
	Data  []Product `json:"data"`
	Links Links     `json:"links"`
	Meta  Meta      `json:"meta"`
}

type RootProductVariants struct {
	Data  []ProductVariant `json:"data"`
	Links Links            `json:"links"`
	Meta  Meta             `json:"meta"`
}

type RootCategories struct {
	Data  []Category `json:"data"`
	Links Links      `json:"links"`
	Meta  Meta       `json:"meta"`
}

type RootCustomers struct {
	Data  []Customer `json:"data"`
	Links Links      `json:"links"`
	Meta  Meta       `json:"meta"`
}

type RootOrders struct {
	Data  []Order `json:"data"`
	Links Links   `json:"links"`
	Meta  Meta    `json:"meta"`
}

type RootCarts struct {
	Data  []Cart `json:"data"`
	Links Links  `json:"links"`
	Meta  Meta   `json:"meta"`
}

type RootCartItems struct {
	Data  []CartItem `json:"data"`
	Links Links      `json:"links"`
	Meta  Meta       `json:"meta"`
}

type RootEcommerceReportOrders struct {
	Data  []EcommerceReportOrder `json:"data"`
	Links Links                  `json:"links"`
	Meta  Meta                   `json:"meta"`
}

type RootAutomationEcommerceOrders struct {
	Data  []map[string]interface{} `json:"data"`
	Links Links                    `json:"links"`
	Meta  Meta                     `json:"meta"`
}

// Amount is a money value the API may send as a JSON number or a numeric string.
type Amount float64

func (a *Amount) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		if s == "" {
			*a = 0
			return nil
		}
		f, err := strconv.ParseFloat(s, 64)
		*a = Amount(f)
		return err
	}

	var f *float64
	if err := json.Unmarshal(b, &f); err != nil {
		return err
	}
	if f != nil {
		*a = Amount(*f)
	}
	return nil
}

type Shop struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Url               string `json:"url"`
	Currency          string `json:"currency"`
	CurrencySign      string `json:"currency_sign"`
	Platform          string `json:"platform"`
	Enabled           bool   `json:"enabled"`
	EnablePopups      bool   `json:"enable_popups"`
	EnableResubscribe bool   `json:"enable_resubscribe"`
	Group             *Group `json:"group"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type Product struct {
	ID                     string           `json:"id"`
	ResourceID             string           `json:"resource_id"`
	Name                   string           `json:"name"`
	Price                  Amount           `json:"price"`
	Currency               string           `json:"currency"`
	Url                    string           `json:"url"`
	Image                  string           `json:"image"`
	ShortDescription       string           `json:"short_description"`
	Description            string           `json:"description"`
	Type                   string           `json:"type"`
	Enabled                bool             `json:"enabled"`
	ExcludeFromAutomations bool             `json:"exclude_from_automations"`
	Categories             []Category       `json:"categories"`
	Variants               []ProductVariant `json:"variants"`
	CreatedAt              string           `json:"created_at"`
}

type ProductVariant struct {
	ID         string      `json:"id"`
	ResourceID string      `json:"resource_id"`
	Name       string      `json:"name"`
	Attributes interface{} `json:"attributes"`
	Price      Amount      `json:"price"`
	PriceType  string      `json:"price_type"`
	ImageUrl   string      `json:"image_url"`
}

type Category struct {
	ID                     string `json:"id"`
	ResourceID             string `json:"resource_id"`
	Name                   string `json:"name"`
	ExcludeFromAutomations bool   `json:"exclude_from_automations"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`
}

type Customer struct {
	ID               string      `json:"id"`
	ResourceID       string      `json:"resource_id"`
	Email            string      `json:"email"`
	AcceptsMarketing bool        `json:"accepts_marketing"`
	TotalSpent       Amount      `json:"total_spent"`
	OrdersCount      int         `json:"orders_count"`
	Subscriber       *Subscriber `json:"subscriber"`
}

type Order struct {
	ID          string            `json:"id"`
	ResourceID  string            `json:"resource_id"`
	Customer    Customer          `json:"customer"`
	Cart        Cart              `json:"cart"`
	Total       Amount            `json:"total"`
	Status      string            `json:"status"`
	Stage       string            `json:"stage"`
	CreatedAt   string            `json:"created_at"`
	UpdatedAt   string            `json:"updated_at"`
	CompletedAt string            `json:"completed_at"`
	Attribution *OrderAttribution `json:"attribution"`
}

// OrderAttribution is the campaign or automation email an order's revenue is attributed to.
type OrderAttribution struct {
	EmailID string `json:"email_id"`
	Type    string `json:"type"`
}

type Cart struct {
	ID          string     `json:"id"`
	ResourceID  string     `json:"resource_id"`
	CheckoutUrl string     `json:"checkout_url"`
	CartTotal   Amount     `json:"cart_total"`
	Items       []CartItem `json:"items"`
}

type CartItem struct {
	ID                        string  `json:"id"`
	ResourceID                string  `json:"resource_id"`
	Product                   Product `json:"product"`
	EcommerceProductVariantID string  `json:"ecommerce_product_variant_id"`
	Variant                   string  `json:"variant"`
	Quantity                  int     `json:"quantity"`
	Price                     Amount  `json:"price"`
}

type EcommerceSettings struct {
	TrackingInCampaigns   bool `json:"ecommerce_tracking_in_campaigns"`
	TrackingInAutomations bool `json:"ecommerce_tracking_in_automations"`
	AttributionWindow     int  `json:"ecommerce_attribution_window"`
}

type EcommerceStats struct {
	OrdersCount       int     `json:"orders_count"`
	Revenue           Amount  `json:"revenue"`
	AverageOrderValue Amount  `json:"average_order_value"`
	ConversionRate    *Amount `json:"conversion_rate,omitempty"`
}

type EcommerceReport struct {
	Ecommerce           EcommerceStats         `json:"ecommerce"`
	BestPerformingEmail map[string]interface{} `json:"best_performing_email,omitempty"`
}

type ShopStats struct {
	Ecommerce EcommerceStats           `json:"ecommerce"`
	Series    []map[string]interface{} `json:"series"`
}

type EcommerceReportOrder struct {
	Subscriber struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	} `json:"subscriber"`
	Total     Amount `json:"total"`
	UpdatedAt string `json:"updated_at"`
	Order     string `json:"order"`
}

// ListEcommerceOptions - modifies the behavior of the e-commerce List methods
type ListEcommerceOptions struct {
	Filters *[]Filter `json:"filters,omitempty"`
	Sort    string    `url:"sort,omitempty"`
	Page    int       `url:"page,omitempty"`
	Limit   int       `url:"limit,omitempty"`
}

// ListProductOptions - modifies the behavior of ProductService.List method
type ListProductOptions = ListEcommerceOptions

// EcommerceReportOptions - modifies the behavior of the campaign and subscriber e-commerce reports
type EcommerceReportOptions struct {
	ShopID string `url:"shop_id"`
	Page   int    `url:"page,omitempty"`
	Limit  int    `url:"limit,omitempty"`
}

// AutomationEcommerceOptions - modifies the behavior of the automation e-commerce stats
type AutomationEcommerceOptions struct {
	ShopID    string `url:"shop_id,omitempty"`
	StartDate string `url:"start_date,omitempty"`
	EndDate   string `url:"end_date,omitempty"`
}

// AutomationEcommerceOrderOptions - modifies the behavior of AutomationService.EcommerceOrders method
type AutomationEcommerceOrderOptions struct {
	TimeFrame       string `url:"time_frame"`
	TimeFrameStart  string `url:"time_frame_start,omitempty"`
	TimeFrameEnd    string `url:"time_frame_end,omitempty"`
	EcommerceShopID string `url:"ecommerce_shop_id,omitempty"`
	Page            int    `url:"page,omitempty"`
	Limit           int    `url:"limit,omitempty"`
}

type ShopStatsOptions struct {
	StartDate string `url:"start_date"`
	EndDate   string `url:"end_date"`
}

type ShopParams struct {
	Name         string `json:"name,omitempty"`
	Url          string `json:"url,omitempty"`
	Currency     string `json:"currency,omitempty"`
	Platform     string `json:"platform,omitempty"`
	GroupID      string `json:"group_id,omitempty"`
	EnablePopups *bool  `json:"enable_popups,omitempty"`
	Enabled      *bool  `json:"enabled,omitempty"`
	AccessData   string `json:"access_data,omitempty"`
}

type ProductParams struct {
	Name                   string                 `json:"name,omitempty"`
	ResourceID             string                 `json:"resource_id,omitempty"`
	Price                  *float64               `json:"price,omitempty"`
	Url                    string                 `json:"url,omitempty"`
	Image                  string                 `json:"image,omitempty"`
	ShortDescription       string                 `json:"short_description,omitempty"`
	Description            string                 `json:"description,omitempty"`
	ExcludeFromAutomations *bool                  `json:"exclude_from_automations,omitempty"`
	Enabled                *bool                  `json:"enabled,omitempty"`
	Type                   string                 `json:"type,omitempty"`
	Categories             []string               `json:"categories,omitempty"`
	Variants               []ProductVariantParams `json:"variants,omitempty"`
}

type ProductVariantParams struct {
	Name       string            `json:"name"`
	Attributes map[string]string `json:"attributes"`
	PriceType  string            `json:"price_type,omitempty"`
	Price      *float64          `json:"price,omitempty"`
	ResourceID string            `json:"resource_id,omitempty"`
	ImageUrl   string            `json:"image_url,omitempty"`
}

type CategoryParams struct {
	Name                   string `json:"name,omitempty"`
	ResourceID             string `json:"resource_id,omitempty"`
	ExcludeFromAutomations *bool  `json:"exclude_from_automations,omitempty"`
}

type CustomerParams struct {
	Email            string                 `json:"email,omitempty"`
	Name             string                 `json:"name,omitempty"`
	LastName         string                 `json:"last_name,omitempty"`
	AcceptsMarketing *bool                  `json:"accepts_marketing,omitempty"`
	TotalSpent       *float64               `json:"total_spent,omitempty"`
	OrdersCount      *int                   `json:"orders_count,omitempty"`
	CreateSubscriber *bool                  `json:"create_subscriber,omitempty"`
	ResourceID       string                 `json:"resource_id,omitempty"`
	SubscriberFields map[string]interface{} `json:"subscriber_fields,omitempty"`
}

type CartItemParams struct {
	EcommerceProductID        string   `json:"ecommerce_product_id,omitempty"`
	ProductResourceID         string   `json:"product_resource_id,omitempty"`
	EcommerceProductVariantID string   `json:"ecommerce_product_variant_id,omitempty"`
	Variant                   string   `json:"variant,omitempty"`
	Quantity                  int      `json:"quantity,omitempty"`
	Price                     *float64 `json:"price,omitempty"`
	ResourceID                string   `json:"resource_id,omitempty"`
}

type CreateOrderParams struct {
	Customer           CustomerParams  `json:"customer"`
	Cart               OrderCartParams `json:"cart"`
	Status             string          `json:"status,omitempty"`
	Stage              string          `json:"stage,omitempty"`
	TotalPrice         *float64        `json:"total_price,omitempty"`
	ResourceID         string          `json:"resource_id,omitempty"`
	CreatedAt          string          `json:"created_at,omitempty"`
	WithoutAutomations bool            `json:"without_automations,omitempty"`
}

type OrderCartParams struct {
	Items       []CartItemParams `json:"items"`
	CheckoutUrl string           `json:"checkout_url,omitempty"`
	CartTotal   *float64         `json:"cart_total,omitempty"`
	ResourceID  string           `json:"resource_id,omitempty"`
}

type UpdateOrderParams struct {
	Status             string   `json:"status,omitempty"`
	Stage              string   `json:"stage,omitempty"`
	TotalPrice         *float64 `json:"total_price,omitempty"`
	WithoutAutomations bool     `json:"without_automations,omitempty"`
}

type UpdateCartParams struct {
	CheckoutUrl         string   `json:"checkout_url,omitempty"`
	CartTotal           *float64 `json:"cart_total,omitempty"`
	EcommerceCustomerID string   `json:"ecommerce_customer_id,omitempty"`
	ResourceID          string   `json:"resource_id,omitempty"`
}

func newEcommerceService(common *service) *EcommerceService {
	return &EcommerceService{
		Shop:           &shopService{common},
		Product:        &productService{common},
		ProductVariant: &productVariantService{common},
		Category:       &categoryService{common},
		Customer:       &customerService{common},
		Order:          &orderService{common},
		Cart:           &cartService{common},
		CartItem:       &cartItemService{common},
		Import:         &ecommerceImportService{common},
		Setting:        &ecommerceSettingService{common},
	}
}

func shopPath(shopID string, format string, args ...interface{}) string {
	return fmt.Sprintf("%s/%s", ecommerceShopEndpoint, shopID) + fmt.Sprintf(format, args...)
}

func send(ctx context.Context, s *service, method, path string, body interface{}, v interface{}) (*Response, error) {
	req, err := s.client.newRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, v)
}

func sendNoContent(ctx context.Context, s *service, method, path string, body interface{}) (*Response, error) {
	req, err := s.client.newRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}

type shopService struct{ *service }

func (s *shopService) List(ctx context.Context, options *ListEcommerceOptions) (*RootShops, *Response, error) {
	root := new(RootShops)
	res, err := send(ctx, s.service, http.MethodGet, ecommerceShopEndpoint, options, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *shopService) Get(ctx context.Context, shopID string) (*RootShop, *Response, error) {
	root := new(RootShop)
	res, err := send(ctx, s.service, http.MethodGet, shopPath(shopID, ""), nil, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *shopService) Create(ctx context.Context, shop *ShopParams) (*RootShop, *Response, error) {
	root := new(RootShop)
	res, err := send(ctx, s.service, http.MethodPost, ecommerceShopEndpoint, shop, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *shopService) Update(ctx context.Context, shopID string, shop *ShopParams) (*RootShop, *Response, error) {
	root := new(RootShop)
	res, err := send(ctx, s.service, http.MethodPut, shopPath(shopID, ""), shop, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *shopService) Delete(ctx context.Context, shopID string) (*Response, error) {
	return sendNoContent(ctx, s.service, http.MethodDelete, shopPath(shopID, ""), nil)
}

func (s *shopService) Stats(ctx context.Context, shopID string, options *ShopStatsOptions) (*RootShopStats, *Response, error) {
	root := new(RootShopStats)
	res, err := send(ctx, s.service, http.MethodGet, shopPath(shopID, "/reports/stats"), options, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

type productService struct{ *service }

func (s *productService) List(ctx context.Context, shopID string, options *ListProductOptions) (*RootProducts, *Response, error) {
	root := new(RootProducts)
	res, err := send(ctx, s.service, http.MethodGet, shopPath(shopID, "/products"), options, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *productService) Get(ctx context.Context, shopID, productID string) (*RootProduct, *Response, error) {
	root := new(RootProduct)
	res, err := send(ctx, s.service, http.MethodGet, shopPath(shopID, "/products/%s", productID), nil, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *productService) Create(ctx context.Context, shopID string, product *ProductParams) (*RootProduct, *Response, error) {
	root := new(RootProduct)
	res, err := send(ctx, s.service, http.MethodPost, shopPath(shopID, "/products"), product, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *productService) Update(ctx context.Context, shopID, productID string, product *ProductParams) (*RootProduct, *Response, error) {
	root := new(RootProduct)
	res, err := send(ctx, s.service, http.MethodPut, shopPath(shopID, "/products/%s", productID), product, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *productService) Delete(ctx context.Context, shopID, productID string) (*Response, error) {
	return sendNoContent(ctx, s.service, http.MethodDelete, shopPath(shopID, "/products/%s", productID), nil)
}

func (s *productService) ChangeStatus(ctx context.Context, shopID, productID string, enabled bool) (*RootProduct, *Response, error) {
	body := map[string]bool{"status": enabled}
	root := new(RootProduct)
	res, err := send(ctx, s.service, http.MethodPut, shopPath(shopID, "/products/%s/change-status", productID), body, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *productService) AssignCategories(ctx context.Context, shopID, productID string, categoryIDs []string, replace bool) (*RootProduct, *Response, error) {
	body := map[string]interface{}{"categories": categoryIDs, "replace": replace}
	root := new(RootProduct)
	res, err := send(ctx, s.service, http.MethodPut, shopPath(shopID, "/products/%s/categories/multiple", productID), body, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

type productVariantService struct{ *service }

func (s *productVariantService) List(ctx context.Context, shopID, productID string) (*RootProductVariants, *Response, error) {
	root := new(RootProductVariants)
	res, err := send(ctx, s.service, http.MethodGet, shopPath(shopID, "/products/%s/variants", productID), nil, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *productVariantService) Sync(ctx context.Context, shopID, productID string, variants []ProductVariantParams) (*RootProduct, *Response, error) {
	root := new(RootProduct)
	res, err := send(ctx, s.service, http.MethodPost, shopPath(shopID, "/products/%s/variants", productID), variants, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *productVariantService) Update(ctx context.Context, shopID, productID, variantID string, variant *ProductVariantParams) (*RootProduct, *Response, error) {
	root := new(RootProduct)
	res, err := send(ctx, s.service, http.MethodPut, shopPath(shopID, "/products/%s/variants/%s", productID, variantID), variant, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *productVariantService) Delete(ctx context.Context, shopID, productID, variantID string) (*Response, error) {
	return sendNoContent(ctx, s.service, http.MethodDelete, shopPath(shopID, "/products/%s/variants/%s", productID, variantID), nil)
}

type categoryService struct{ *service }

func (s *categoryService) List(ctx context.Context, shopID string, options *ListEcommerceOptions) (*RootCategories, *Response, error) {
	root := new(RootCategories)
	res, err := send(ctx, s.service, http.MethodGet, shopPath(shopID, "/categories"), options, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *categoryService) Get(ctx context.Context, shopID, categoryID string) (*RootCategory, *Response, error) {
	root := new(RootCategory)
	res, err := send(ctx, s.service, http.MethodGet, shopPath(shopID, "/categories/%s", categoryID), nil, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *categoryService) Create(ctx context.Context, shopID string, category *CategoryParams) (*RootCategory, *Response, error) {
	root := new(RootCategory)
	res, err := send(ctx, s.service, http.MethodPost, shopPath(shopID, "/categories"), category, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *categoryService) Update(ctx context.Context, shopID, categoryID string, category *CategoryParams) (*RootCategory, *Response, error) {
	root := new(RootCategory)
	res, err := send(ctx, s.service, http.MethodPut, shopPath(shopID, "/categories/%s", categoryID), category, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *categoryService) Delete(ctx context.Context, shopID, categoryID string) (*Response, error) {
	return sendNoContent(ctx, s.service, http.MethodDelete, shopPath(shopID, "/categories/%s", categoryID), nil)
}

func (s *categoryService) Products(ctx context.Context, shopID, categoryID string, options *ListEcommerceOptions) (*RootProducts, *Response, error) {
	root := new(RootProducts)
	res, err := send(ctx, s.service, http.MethodGet, shopPath(shopID, "/categories/%s/products", categoryID), options, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *categoryService) AssignProduct(ctx context.Context, shopID, categoryID, productID string) (*Response, error) {
	return sendNoContent(ctx, s.service, http.MethodPut, shopPath(shopID, "/categories/%s/products/%s", categoryID, productID), nil)
}

func (s *categoryService) UnassignProduct(ctx context.Context, shopID, categoryID, productID string) (*Response, error) {
	return sendNoContent(ctx, s.service, http.MethodDelete, shopPath(shopID, "/categories/%s/products/%s", categoryID, productID), nil)
}

type customerService struct{ *service }

func (s *customerService) List(ctx context.Context, shopID string, options *ListEcommerceOptions) (*RootCustomers, *Response, error) {
	root := new(RootCustomers)
	res, err := send(ctx, s.service, http.MethodGet, shopPath(shopID, "/customers"), options, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *customerService) Get(ctx context.Context, shopID, customerID string) (*RootCustomer, *Response, error) {
	root := new(RootCustomer)
	res, err := send(ctx, s.service, http.MethodGet, shopPath(shopID, "/customers/%s", customerID), nil, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *customerService) Create(ctx context.Context, shopID string, customer *CustomerParams) (*RootCustomer, *Response, error) {
	root := new(RootCustomer)
	res, err := send(ctx, s.service, http.MethodPost, shopPath(shopID, "/customers"), customer, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *customerService) Update(ctx context.Context, shopID, customerID string, customer *CustomerParams) (*RootCustomer, *Response, error) {
	root := new(RootCustomer)
	res, err := send(ctx, s.service, http.MethodPut, shopPath(shopID, "/customers/%s", customerID), customer, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *customerService) Delete(ctx context.Context, shopID, customerID string) (*Response, error) {
	return sendNoContent(ctx, s.service, http.MethodDelete, shopPath(shopID, "/customers/%s", customerID), nil)
}

type orderService struct{ *service }

func (s *orderService) List(ctx context.Context, shopID string, options *ListEcommerceOptions) (*RootOrders, *Response, error) {
	root := new(RootOrders)
	res, err := send(ctx, s.service, http.MethodGet, shopPath(shopID, "/orders"), options, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *orderService) Get(ctx context.Context, shopID, orderID string) (*RootOrder, *Response, error) {
	root := new(RootOrder)
	res, err := send(ctx, s.service, http.MethodGet, shopPath(shopID, "/orders/%s", orderID), nil, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *orderService) Create(ctx context.Context, shopID string, order *CreateOrderParams) (*RootOrder, *Response, error) {
	root := new(RootOrder)
	res, err := send(ctx, s.service, http.MethodPost, shopPath(shopID, "/orders"), order, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *orderService) Update(ctx context.Context, shopID, orderID string, order *UpdateOrderParams) (*RootOrder, *Response, error) {
	root := new(RootOrder)
	res, err := send(ctx, s.service, http.MethodPut, shopPath(shopID, "/orders/%s", orderID), order, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *orderService) Delete(ctx context.Context, shopID, orderID string) (*Response, error) {
	return sendNoContent(ctx, s.service, http.MethodDelete, shopPath(shopID, "/orders/%s", orderID), nil)
}

type cartService struct{ *service }

func (s *cartService) List(ctx context.Context, shopID string, options *ListEcommerceOptions) (*RootCarts, *Response, error) {
	root := new(RootCarts)
	res, err := send(ctx, s.service, http.MethodGet, shopPath(shopID, "/carts"), options, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *cartService) Get(ctx context.Context, shopID, cartID string) (*RootCart, *Response, error) {
	root := new(RootCart)
	res, err := send(ctx, s.service, http.MethodGet, shopPath(shopID, "/carts/%s", cartID), nil, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *cartService) Update(ctx context.Context, shopID, cartID string, cart *UpdateCartParams) (*RootCart, *Response, error) {
	root := new(RootCart)
	res, err := send(ctx, s.service, http.MethodPut, shopPath(shopID, "/carts/%s", cartID), cart, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

type cartItemService struct{ *service }

func (s *cartItemService) List(ctx context.Context, shopID, cartID string, options *ListEcommerceOptions) (*RootCartItems, *Response, error) {
	root := new(RootCartItems)
	res, err := send(ctx, s.service, http.MethodGet, shopPath(shopID, "/carts/%s/items", cartID), options, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *cartItemService) Get(ctx context.Context, shopID, cartID, itemID string) (*RootCartItem, *Response, error) {
	root := new(RootCartItem)
	res, err := send(ctx, s.service, http.MethodGet, shopPath(shopID, "/carts/%s/items/%s", cartID, itemID), nil, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *cartItemService) Create(ctx context.Context, shopID, cartID string, item *CartItemParams) (*RootCartItem, *Response, error) {
	root := new(RootCartItem)
	res, err := send(ctx, s.service, http.MethodPost, shopPath(shopID, "/carts/%s/items", cartID), item, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *cartItemService) Update(ctx context.Context, shopID, cartID, itemID string, item *CartItemParams) (*RootCartItem, *Response, error) {
	root := new(RootCartItem)
	res, err := send(ctx, s.service, http.MethodPut, shopPath(shopID, "/carts/%s/items/%s", cartID, itemID), item, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *cartItemService) Delete(ctx context.Context, shopID, cartID, itemID string) (*Response, error) {
	return sendNoContent(ctx, s.service, http.MethodDelete, shopPath(shopID, "/carts/%s/items/%s", cartID, itemID), nil)
}

type ecommerceImportService struct{ *service }

func (s *ecommerceImportService) Products(ctx context.Context, shopID string, products []ProductParams) (*Response, error) {
	return sendNoContent(ctx, s.service, http.MethodPost, shopPath(shopID, "/products/import"), products)
}

func (s *ecommerceImportService) Categories(ctx context.Context, shopID string, categories []CategoryParams) (*Response, error) {
	return sendNoContent(ctx, s.service, http.MethodPost, shopPath(shopID, "/categories/import"), categories)
}

func (s *ecommerceImportService) Customers(ctx context.Context, shopID string, customers []CustomerParams) (*Response, error) {
	return sendNoContent(ctx, s.service, http.MethodPost, shopPath(shopID, "/customers/import"), customers)
}

func (s *ecommerceImportService) Orders(ctx context.Context, shopID string, orders []CreateOrderParams) (*Response, error) {
	body := map[string][]CreateOrderParams{"orders": orders}
	return sendNoContent(ctx, s.service, http.MethodPost, shopPath(shopID, "/orders/import"), body)
}

type ecommerceSettingService struct{ *service }

func (s *ecommerceSettingService) Get(ctx context.Context) (*RootEcommerceSettings, *Response, error) {
	root := new(RootEcommerceSettings)
	res, err := send(ctx, s.service, http.MethodGet, "/ecommerce/settings", nil, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
