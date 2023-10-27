package goccavenue

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func escapeStr(str string) string {
	// return url.QueryEscape(str)
	return str
}

// Sets the merchant id
func SetMerchantId(merchantId string) {
	ccAvenueMerchantId = merchantId
}

// Gets the merchant id
func GetMerchantId() (merchantId string) {
	return ccAvenueMerchantId
}

// Sets the encryption key
func SetEncryptionKey(encryptionKey string) {
	ccAvenueEncryptionKey = encryptionKey
}

// Gets the encryption key
func GetEncryptionKey() (encryptionKey string) {
	return ccAvenueEncryptionKey
}

// Generates the string in a format acceptable to CCAvenue prior to encryption
func CreatePayload(paymentReq CCAvenueRequest) string {
	var stringsToEncode = []string{
		fmt.Sprintf("merchant_id=%s", escapeStr(paymentReq.MerchantId)),
		fmt.Sprintf("order_id=%s", escapeStr(paymentReq.OrderId)),
		fmt.Sprintf("currency=%s", escapeStr(paymentReq.Currency)),
		fmt.Sprintf("amount=%s", escapeStr(paymentReq.Amount)),
		fmt.Sprintf("redirect_url=%s", escapeStr(paymentReq.RedirectUrl)),
		fmt.Sprintf("cancel_url=%s", escapeStr(paymentReq.CancelUrl)),
		fmt.Sprintf("language=%s", escapeStr(paymentReq.Language)),

		fmt.Sprintf("billing_name=%s", escapeStr(paymentReq.BillingName)),
		fmt.Sprintf("billing_address=%s", escapeStr(paymentReq.BillingAddress)),
		fmt.Sprintf("billing_city=%s", escapeStr(paymentReq.BillingCity)),
		fmt.Sprintf("billing_state=%s", escapeStr(paymentReq.BillingState)),
		fmt.Sprintf("billing_zip=%s", escapeStr(paymentReq.BillingZip)),
		fmt.Sprintf("billing_country=%s", escapeStr(paymentReq.BillingCountry)),
		fmt.Sprintf("billing_tel=%s", escapeStr(paymentReq.BillingTel)),
		fmt.Sprintf("billing_email=%s", escapeStr(paymentReq.BillingEmail)),

		fmt.Sprintf("delivery_name=%s", escapeStr(paymentReq.DeliveryName)),
		fmt.Sprintf("delivery_address=%s", escapeStr(paymentReq.DeliveryAddress)),
		fmt.Sprintf("delivery_city=%s", escapeStr(paymentReq.DeliveryCity)),
		fmt.Sprintf("delivery_state=%s", escapeStr(paymentReq.DeliveryState)),
		fmt.Sprintf("delivery_zip=%s", escapeStr(paymentReq.DeliveryZip)),
		fmt.Sprintf("delivery_country=%s", escapeStr(paymentReq.DeliveryCountry)),
		fmt.Sprintf("delivery_tel=%s", escapeStr(paymentReq.DeliveryTel)),

		fmt.Sprintf("merchant_param1=%s", escapeStr(paymentReq.MerchantParam1)),
		fmt.Sprintf("merchant_param2=%s", escapeStr(paymentReq.MerchantParam2)),
		fmt.Sprintf("merchant_param3=%s", escapeStr(paymentReq.MerchantParam3)),
		fmt.Sprintf("merchant_param4=%s", escapeStr(paymentReq.MerchantParam4)),
		fmt.Sprintf("merchant_param5=%s", escapeStr(paymentReq.MerchantParam5)),

		fmt.Sprintf("integration_type=%s", escapeStr(paymentReq.IntegrationType)),
		fmt.Sprintf("promo_code=%s", escapeStr(paymentReq.PromoCode)),
		fmt.Sprintf("customer_identifier=%s", escapeStr(paymentReq.CustomerIdentifier)),

		fmt.Sprintf("merchant_id=%s", paymentReq.MerchantId),
	}

	return strings.Join(stringsToEncode, "&")
}

// Pads the cipher text so that the length is a multiple of the block size (as per CCAvenue's requirement)
func padPayload(ciphertext string) (paddedString string) {
	padding := (ccAvenueBlockSize - len(ciphertext)%ccAvenueBlockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return ciphertext + string(padtext)
}

// Encrypts plain text string into cipher text string that is recognized by CCAvenue
func EncryptPayload(paymentReq CCAvenueRequest) (string, error) {
	hash := md5.Sum([]byte(GetEncryptionKey()))
	key := hash[:]
	plainText := padPayload(CreatePayload(paymentReq))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, len(plainText))
	mode := cipher.NewCBCEncrypter(block, ccAvenueInitializationVector)
	mode.CryptBlocks(cipherText, []byte(plainText))

	return hex.EncodeToString(cipherText), err
}

// Decrypts cipher text string into plain text string
func DecryptPayload(cipherText string) (decryptedString string, err error) {
	hash := md5.Sum([]byte(GetEncryptionKey()))
	key := hash[:]
	cipherTextDecoded, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	paddedBytes := make([]byte, len(cipherTextDecoded))

	mode := cipher.NewCBCDecrypter(block, ccAvenueInitializationVector)
	mode.CryptBlocks(paddedBytes, cipherTextDecoded)

	paddingLen := paddedBytes[len(paddedBytes)-1]
	if paddingLen < ccAvenueBlockSize {
		// Message has been padded, so trim the last n bytes
		paddedBytes = paddedBytes[:len(paddedBytes)-int(paddingLen)]
	}

	plainText := string(paddedBytes)

	return plainText, nil
}
