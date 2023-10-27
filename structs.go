package goccavenue

// Stores the necessary parameters to initiate a payment request via CCAvenue
type CCAvenueRequest struct {
	MerchantId  string `json:"merchant_id"`
	OrderId     string `json:"order_id"`
	Currency    string `json:"currency"`
	Amount      string `json:"amount"`
	RedirectUrl string `json:"redirect_url"`
	CancelUrl   string `json:"cancel_url"`
	Language    string `json:"language"`

	BillingName    string `json:"billing_name"`
	BillingAddress string `json:"billing_address"`
	BillingCity    string `json:"billing_city"`
	BillingState   string `json:"billing_state"`
	BillingZip     string `json:"billing_zip"`
	BillingCountry string `json:"billing_country"`
	BillingTel     string `json:"billing_tel"`
	BillingEmail   string `json:"billing_email"`

	DeliveryName    string `json:"delivery_name"`
	DeliveryAddress string `json:"delivery_address"`
	DeliveryCity    string `json:"delivery_city"`
	DeliveryState   string `json:"delivery_state"`
	DeliveryZip     string `json:"delivery_zip"`
	DeliveryCountry string `json:"delivery_country"`
	DeliveryTel     string `json:"delivery_tel"`

	MerchantParam1 string `json:"merchant_param1"`
	MerchantParam2 string `json:"merchant_param2"`
	MerchantParam3 string `json:"merchant_param3"`
	MerchantParam4 string `json:"merchant_param4"`
	MerchantParam5 string `json:"merchant_param5"`

	IntegrationType    string `json:"integration_type"`
	PromoCode          string `json:"promo_code"`
	CustomerIdentifier string `json:"customer_identifier"`
}

// Custom type to hold the order status
type OrderStatus int

// Order Statuses
const (
	OrderSuccess OrderStatus = iota + 1
	OrderFailure
	OrderAborted
)

func (status OrderStatus) String() string {
	return [...]string{"Success", "Failure", "Aborted"}[status-1]
}

func strToOrderStatus(str string) OrderStatus {
	if str == "Success" {
		return OrderSuccess
	} else if str == "Failure" {
		return OrderFailure
	} else if str == "Aborted" {
		return OrderAborted
	}
	return OrderFailure
}

// Stores the parameters in a CCAvenue response
type CCAvenueResponse struct {
	OrderId        string      `json:"order_id"`
	TrackingId     string      `json:"tracking_id"`
	BankRefNo      string      `json:"bank_ref_no"`
	OrderStatus    OrderStatus `json:"order_status"`
	FailureMessage string      `json:"failure_message"`
	PaymentMode    string      `json:"Net Banking"`
	CardName       string      `json:"card_name"`

	StatusCode    string `json:"status_code"`
	StatusMessage string `json:"status_message"`

	Currency string `json:"currency"`
	Amount   string `json:"amount"`

	BillingName    string `json:"billing_name"`
	BillingAddress string `json:"billing_address"`
	BillingCity    string `json:"billing_city"`
	BillingState   string `json:"billing_state"`
	BillingZip     string `json:"billing_zip"`
	BillingCountry string `json:"billing_country"`
	BillingTel     string `json:"billing_tel"`
	BillingEmail   string `json:"billing_email"`

	DeliveryName    string `json:"delivery_name"`
	DeliveryAddress string `json:"delivery_address"`
	DeliveryCity    string `json:"delivery_city"`
	DeliveryState   string `json:"delivery_state"`
	DeliveryZip     string `json:"delivery_zip"`
	DeliveryCountry string `json:"delivery_country"`
	DeliveryTel     string `json:"delivery_tel"`

	MerchantParam1 string `json:"merchant_param1"`
	MerchantParam2 string `json:"merchant_param2"`
	MerchantParam3 string `json:"merchant_param3"`
	MerchantParam4 string `json:"merchant_param4"`
	MerchantParam5 string `json:"merchant_param5"`

	Vault     string `json:"vault"`
	OfferType string `json:"offer_type"`
	OfferCode string `json:"offer_code"`

	DiscountValue   string `json:"discount_value"`
	MerchantAmount  string `json:"mer_amount"`
	ECIValue        string `json:"eci_value"`
	Retry           string `json:"retry"`
	ResponseCode    string `json:"response_code"`
	BillingNotes    string `json:"billing_notes"`
	TransactionDate string `json:"trans_date"`

	BinCountry string `json:"bin_country"`
}

// Stores the Merchant ID
var ccAvenueMerchantId = ""

// Stores the Encryption Key as given by CCAvenue
var ccAvenueEncryptionKey = ""

// Stores the block size for encrypting the payload
const ccAvenueBlockSize = 16

// Stores the initialization vector for encrypting the payload
var ccAvenueInitializationVector = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
