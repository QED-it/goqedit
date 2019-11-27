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

// The details and Proof attached to a confidential Issuance; null if the Issuance was public
type AnalyticsConfidentialIssuanceDescription struct {
	// The commitment to both the Asset Type and amount of the issued Note
	Casset string `json:"casset,omitempty"`
	// A `Namespace` describes what Asset IDs can be issued in an Issuance Rule. It is a string in the same format as `AssetId`. Additionally, if it ends with a wildcard character `*`, then the namespace covers all asset IDs with the namespace as a prefix. Without a final wildcard, the namespace covers exactly one asset ID. Example: The namespace `currencies.dollar` covers only this exact asset type, while `currencies.*` covers all asset types that start with `currencies.`.
	Namespace string `json:"namespace,omitempty"`
	// The Proof that the issued Asset Type indeed matches the reported Rule
	Zkproof string `json:"zkproof,omitempty"`
}
