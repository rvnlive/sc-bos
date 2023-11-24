/*
Airthings API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: support@airthings.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// checks if the GetLocationSamplesResponseEnriched type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLocationSamplesResponseEnriched{}

// GetLocationSamplesResponseEnriched struct for GetLocationSamplesResponseEnriched
type GetLocationSamplesResponseEnriched struct {
	Devices []DeviceSampleResponseEnriched `json:"devices"`
	Id      string                         `json:"id"`
	Name    string                         `json:"name"`
}

type _GetLocationSamplesResponseEnriched GetLocationSamplesResponseEnriched

// NewGetLocationSamplesResponseEnriched instantiates a new GetLocationSamplesResponseEnriched object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLocationSamplesResponseEnriched(devices []DeviceSampleResponseEnriched, id string, name string) *GetLocationSamplesResponseEnriched {
	this := GetLocationSamplesResponseEnriched{}
	this.Devices = devices
	this.Id = id
	this.Name = name
	return &this
}

// NewGetLocationSamplesResponseEnrichedWithDefaults instantiates a new GetLocationSamplesResponseEnriched object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLocationSamplesResponseEnrichedWithDefaults() *GetLocationSamplesResponseEnriched {
	this := GetLocationSamplesResponseEnriched{}
	return &this
}

// GetDevices returns the Devices field value
func (o *GetLocationSamplesResponseEnriched) GetDevices() []DeviceSampleResponseEnriched {
	if o == nil {
		var ret []DeviceSampleResponseEnriched
		return ret
	}

	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value
// and a boolean to check if the value has been set.
func (o *GetLocationSamplesResponseEnriched) GetDevicesOk() ([]DeviceSampleResponseEnriched, bool) {
	if o == nil {
		return nil, false
	}
	return o.Devices, true
}

// SetDevices sets field value
func (o *GetLocationSamplesResponseEnriched) SetDevices(v []DeviceSampleResponseEnriched) {
	o.Devices = v
}

// GetId returns the Id field value
func (o *GetLocationSamplesResponseEnriched) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetLocationSamplesResponseEnriched) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetLocationSamplesResponseEnriched) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GetLocationSamplesResponseEnriched) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetLocationSamplesResponseEnriched) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetLocationSamplesResponseEnriched) SetName(v string) {
	o.Name = v
}

func (o GetLocationSamplesResponseEnriched) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLocationSamplesResponseEnriched) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["devices"] = o.Devices
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *GetLocationSamplesResponseEnriched) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"devices",
		"id",
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetLocationSamplesResponseEnriched := _GetLocationSamplesResponseEnriched{}

	err = json.Unmarshal(bytes, &varGetLocationSamplesResponseEnriched)

	if err != nil {
		return err
	}

	*o = GetLocationSamplesResponseEnriched(varGetLocationSamplesResponseEnriched)

	return err
}

type NullableGetLocationSamplesResponseEnriched struct {
	value *GetLocationSamplesResponseEnriched
	isSet bool
}

func (v NullableGetLocationSamplesResponseEnriched) Get() *GetLocationSamplesResponseEnriched {
	return v.value
}

func (v *NullableGetLocationSamplesResponseEnriched) Set(val *GetLocationSamplesResponseEnriched) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLocationSamplesResponseEnriched) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLocationSamplesResponseEnriched) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLocationSamplesResponseEnriched(val *GetLocationSamplesResponseEnriched) *NullableGetLocationSamplesResponseEnriched {
	return &NullableGetLocationSamplesResponseEnriched{value: val, isSet: true}
}

func (v NullableGetLocationSamplesResponseEnriched) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLocationSamplesResponseEnriched) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}