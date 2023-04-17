/*
Catalog Governor Service REST API

This is the service to track assets deployed in customer clusters.  NOTE: Catalog Governor Service is an internal tool for the Content-Building Ecosystem team.

API version: ${project.version}
Contact: content-building-ecosystem@vmware.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// Product VMWare Application Catalog Product's information.
type Product struct {
	// Product name from VMWare Application Catalog
	Name string `json:"name"`
	// Product branch from VMWare Application Catalog
	Branch string `json:"branch"`
	// Product version from VMWare Application Catalog
	Version string `json:"version"`
	// Product revision from VMWare Application Catalog
	Revision *string `json:"revision,omitempty"`
	// The date-time which the product was released at
	ReleasedAt time.Time `json:"released_at"`
	// Last release version of product
	LastVersionReleased *string            `json:"last_version_released,omitempty"`
	DeprecationPolicy   *DeprecationPolicy `json:"deprecation_policy,omitempty"`
	NonsupportPolicy    *NonSupportPolicy  `json:"nonsupport_policy,omitempty"`
	// The status of the product in the catalog. Available values are DRAFT, ACTIVE, SCHEDULED_DEPRECATION, DEPRECATION_GRACE_PERIOD, DEPRECATED, NON_SUPPORTED
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Product Product

// NewProduct instantiates a new Product object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct(name string, branch string, version string, releasedAt time.Time) *Product {
	this := Product{}
	this.Name = name
	this.Branch = branch
	this.Version = version
	this.ReleasedAt = releasedAt
	return &this
}

// NewProductWithDefaults instantiates a new Product object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductWithDefaults() *Product {
	this := Product{}
	return &this
}

// GetName returns the Name field value
func (o *Product) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Product) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Product) SetName(v string) {
	o.Name = v
}

// GetBranch returns the Branch field value
func (o *Product) GetBranch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Branch
}

// GetBranchOk returns a tuple with the Branch field value
// and a boolean to check if the value has been set.
func (o *Product) GetBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Branch, true
}

// SetBranch sets field value
func (o *Product) SetBranch(v string) {
	o.Branch = v
}

// GetVersion returns the Version field value
func (o *Product) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Product) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Product) SetVersion(v string) {
	o.Version = v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *Product) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *Product) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *Product) SetRevision(v string) {
	o.Revision = &v
}

// GetReleasedAt returns the ReleasedAt field value
func (o *Product) GetReleasedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ReleasedAt
}

// GetReleasedAtOk returns a tuple with the ReleasedAt field value
// and a boolean to check if the value has been set.
func (o *Product) GetReleasedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleasedAt, true
}

// SetReleasedAt sets field value
func (o *Product) SetReleasedAt(v time.Time) {
	o.ReleasedAt = v
}

// GetLastVersionReleased returns the LastVersionReleased field value if set, zero value otherwise.
func (o *Product) GetLastVersionReleased() string {
	if o == nil || o.LastVersionReleased == nil {
		var ret string
		return ret
	}
	return *o.LastVersionReleased
}

// GetLastVersionReleasedOk returns a tuple with the LastVersionReleased field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetLastVersionReleasedOk() (*string, bool) {
	if o == nil || o.LastVersionReleased == nil {
		return nil, false
	}
	return o.LastVersionReleased, true
}

// HasLastVersionReleased returns a boolean if a field has been set.
func (o *Product) HasLastVersionReleased() bool {
	if o != nil && o.LastVersionReleased != nil {
		return true
	}

	return false
}

// SetLastVersionReleased gets a reference to the given string and assigns it to the LastVersionReleased field.
func (o *Product) SetLastVersionReleased(v string) {
	o.LastVersionReleased = &v
}

// GetDeprecationPolicy returns the DeprecationPolicy field value if set, zero value otherwise.
func (o *Product) GetDeprecationPolicy() DeprecationPolicy {
	if o == nil || o.DeprecationPolicy == nil {
		var ret DeprecationPolicy
		return ret
	}
	return *o.DeprecationPolicy
}

// GetDeprecationPolicyOk returns a tuple with the DeprecationPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetDeprecationPolicyOk() (*DeprecationPolicy, bool) {
	if o == nil || o.DeprecationPolicy == nil {
		return nil, false
	}
	return o.DeprecationPolicy, true
}

// HasDeprecationPolicy returns a boolean if a field has been set.
func (o *Product) HasDeprecationPolicy() bool {
	if o != nil && o.DeprecationPolicy != nil {
		return true
	}

	return false
}

// SetDeprecationPolicy gets a reference to the given DeprecationPolicy and assigns it to the DeprecationPolicy field.
func (o *Product) SetDeprecationPolicy(v DeprecationPolicy) {
	o.DeprecationPolicy = &v
}

// GetNonsupportPolicy returns the NonsupportPolicy field value if set, zero value otherwise.
func (o *Product) GetNonsupportPolicy() NonSupportPolicy {
	if o == nil || o.NonsupportPolicy == nil {
		var ret NonSupportPolicy
		return ret
	}
	return *o.NonsupportPolicy
}

// GetNonsupportPolicyOk returns a tuple with the NonsupportPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetNonsupportPolicyOk() (*NonSupportPolicy, bool) {
	if o == nil || o.NonsupportPolicy == nil {
		return nil, false
	}
	return o.NonsupportPolicy, true
}

// HasNonsupportPolicy returns a boolean if a field has been set.
func (o *Product) HasNonsupportPolicy() bool {
	if o != nil && o.NonsupportPolicy != nil {
		return true
	}

	return false
}

// SetNonsupportPolicy gets a reference to the given NonSupportPolicy and assigns it to the NonsupportPolicy field.
func (o *Product) SetNonsupportPolicy(v NonSupportPolicy) {
	o.NonsupportPolicy = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Product) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Product) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Product) SetStatus(v string) {
	o.Status = &v
}

func (o Product) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["branch"] = o.Branch
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	if true {
		toSerialize["released_at"] = o.ReleasedAt
	}
	if o.LastVersionReleased != nil {
		toSerialize["last_version_released"] = o.LastVersionReleased
	}
	if o.DeprecationPolicy != nil {
		toSerialize["deprecation_policy"] = o.DeprecationPolicy
	}
	if o.NonsupportPolicy != nil {
		toSerialize["nonsupport_policy"] = o.NonsupportPolicy
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Product) UnmarshalJSON(bytes []byte) (err error) {
	varProduct := _Product{}

	if err = json.Unmarshal(bytes, &varProduct); err == nil {
		*o = Product(varProduct)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "branch")
		delete(additionalProperties, "version")
		delete(additionalProperties, "revision")
		delete(additionalProperties, "released_at")
		delete(additionalProperties, "last_version_released")
		delete(additionalProperties, "deprecation_policy")
		delete(additionalProperties, "nonsupport_policy")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProduct struct {
	value *Product
	isSet bool
}

func (v NullableProduct) Get() *Product {
	return v.value
}

func (v *NullableProduct) Set(val *Product) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct(val *Product) *NullableProduct {
	return &NullableProduct{value: val, isSet: true}
}

func (v NullableProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}