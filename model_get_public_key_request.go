/*
 * QEDIT - Asset Transfers
 *
 * This SDK provides a programmatic way for interacting with QEDIT's _Asset Transfer_ API. The specification definition file is publicly available [in this repository](https://github.com/QED-it/asset_transfers_dev_guide/).
 *
 * API version: 1.7.0
 * Contact: dev@qed-it.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package goqedit

type GetPublicKeyRequest struct {
	// The ID of the Wallet for which to retrieve the public key
	WalletId string `json:"wallet_id"`
}
