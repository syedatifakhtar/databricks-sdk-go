/*
 * Databricks
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package models

type ScimUser struct {
	Entitlements []Entitlements `json:"entitlements,omitempty"`
	DisplayName  string         `json:"displayName,omitempty"`
	Groups       []Groups       `json:"groups,omitempty"`
	Emails       []Emails       `json:"emails,omitempty"`
	Id           string         `json:"id,omitempty"`
	Name         *Name          `json:"name,omitempty"`
	Active       bool           `json:"active,omitempty"`
	UserName     string         `json:"userName,omitempty"`
}
