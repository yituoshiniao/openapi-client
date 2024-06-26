/*
gin-http API

gin-http服务文档

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the InternalApiDtoAppJwtTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternalApiDtoAppJwtTokenResponse{}

// InternalApiDtoAppJwtTokenResponse struct for InternalApiDtoAppJwtTokenResponse
type InternalApiDtoAppJwtTokenResponse struct {
	// token payload 信息
	JwtPayload *InternalModuleAuthApplicationServiceAuthPayload `json:"jwtPayload,omitempty"`
	// Token jwt token
	Token *string `json:"token,omitempty"`
}

// NewInternalApiDtoAppJwtTokenResponse instantiates a new InternalApiDtoAppJwtTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalApiDtoAppJwtTokenResponse() *InternalApiDtoAppJwtTokenResponse {
	this := InternalApiDtoAppJwtTokenResponse{}
	return &this
}

// NewInternalApiDtoAppJwtTokenResponseWithDefaults instantiates a new InternalApiDtoAppJwtTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalApiDtoAppJwtTokenResponseWithDefaults() *InternalApiDtoAppJwtTokenResponse {
	this := InternalApiDtoAppJwtTokenResponse{}
	return &this
}

// GetJwtPayload returns the JwtPayload field value if set, zero value otherwise.
func (o *InternalApiDtoAppJwtTokenResponse) GetJwtPayload() InternalModuleAuthApplicationServiceAuthPayload {
	if o == nil || IsNil(o.JwtPayload) {
		var ret InternalModuleAuthApplicationServiceAuthPayload
		return ret
	}
	return *o.JwtPayload
}

// GetJwtPayloadOk returns a tuple with the JwtPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalApiDtoAppJwtTokenResponse) GetJwtPayloadOk() (*InternalModuleAuthApplicationServiceAuthPayload, bool) {
	if o == nil || IsNil(o.JwtPayload) {
		return nil, false
	}
	return o.JwtPayload, true
}

// HasJwtPayload returns a boolean if a field has been set.
func (o *InternalApiDtoAppJwtTokenResponse) HasJwtPayload() bool {
	if o != nil && !IsNil(o.JwtPayload) {
		return true
	}

	return false
}

// SetJwtPayload gets a reference to the given InternalModuleAuthApplicationServiceAuthPayload and assigns it to the JwtPayload field.
func (o *InternalApiDtoAppJwtTokenResponse) SetJwtPayload(v InternalModuleAuthApplicationServiceAuthPayload) {
	o.JwtPayload = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *InternalApiDtoAppJwtTokenResponse) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalApiDtoAppJwtTokenResponse) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *InternalApiDtoAppJwtTokenResponse) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *InternalApiDtoAppJwtTokenResponse) SetToken(v string) {
	o.Token = &v
}

func (o InternalApiDtoAppJwtTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalApiDtoAppJwtTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JwtPayload) {
		toSerialize["jwtPayload"] = o.JwtPayload
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableInternalApiDtoAppJwtTokenResponse struct {
	value *InternalApiDtoAppJwtTokenResponse
	isSet bool
}

func (v NullableInternalApiDtoAppJwtTokenResponse) Get() *InternalApiDtoAppJwtTokenResponse {
	return v.value
}

func (v *NullableInternalApiDtoAppJwtTokenResponse) Set(val *InternalApiDtoAppJwtTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalApiDtoAppJwtTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalApiDtoAppJwtTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalApiDtoAppJwtTokenResponse(val *InternalApiDtoAppJwtTokenResponse) *NullableInternalApiDtoAppJwtTokenResponse {
	return &NullableInternalApiDtoAppJwtTokenResponse{value: val, isSet: true}
}

func (v NullableInternalApiDtoAppJwtTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalApiDtoAppJwtTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


