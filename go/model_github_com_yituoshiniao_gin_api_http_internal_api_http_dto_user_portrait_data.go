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

// checks if the GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData{}

// GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData struct for GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData
type GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData struct {
	// 国家
	Country *string `json:"country,omitempty"`
	// 上次登陆时间
	LastLogin *string `json:"last_login,omitempty"`
	// 用户id
	UserId *string `json:"user_id,omitempty"`
	// 是否为VIP，0/1
	VipInfo *int32 `json:"vip_info,omitempty"`
}

// NewGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData instantiates a new GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData() *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData {
	this := GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData{}
	return &this
}

// NewGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitDataWithDefaults instantiates a new GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitDataWithDefaults() *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData {
	this := GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) SetCountry(v string) {
	o.Country = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise.
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) GetLastLogin() string {
	if o == nil || IsNil(o.LastLogin) {
		var ret string
		return ret
	}
	return *o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) GetLastLoginOk() (*string, bool) {
	if o == nil || IsNil(o.LastLogin) {
		return nil, false
	}
	return o.LastLogin, true
}

// HasLastLogin returns a boolean if a field has been set.
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) HasLastLogin() bool {
	if o != nil && !IsNil(o.LastLogin) {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given string and assigns it to the LastLogin field.
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) SetLastLogin(v string) {
	o.LastLogin = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) SetUserId(v string) {
	o.UserId = &v
}

// GetVipInfo returns the VipInfo field value if set, zero value otherwise.
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) GetVipInfo() int32 {
	if o == nil || IsNil(o.VipInfo) {
		var ret int32
		return ret
	}
	return *o.VipInfo
}

// GetVipInfoOk returns a tuple with the VipInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) GetVipInfoOk() (*int32, bool) {
	if o == nil || IsNil(o.VipInfo) {
		return nil, false
	}
	return o.VipInfo, true
}

// HasVipInfo returns a boolean if a field has been set.
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) HasVipInfo() bool {
	if o != nil && !IsNil(o.VipInfo) {
		return true
	}

	return false
}

// SetVipInfo gets a reference to the given int32 and assigns it to the VipInfo field.
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) SetVipInfo(v int32) {
	o.VipInfo = &v
}

func (o GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.LastLogin) {
		toSerialize["last_login"] = o.LastLogin
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.VipInfo) {
		toSerialize["vip_info"] = o.VipInfo
	}
	return toSerialize, nil
}

type NullableGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData struct {
	value *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData
	isSet bool
}

func (v NullableGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) Get() *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData {
	return v.value
}

func (v *NullableGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) Set(val *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData(val *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) *NullableGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData {
	return &NullableGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData{value: val, isSet: true}
}

func (v NullableGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
