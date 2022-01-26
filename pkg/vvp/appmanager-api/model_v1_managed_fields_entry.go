/*
 * Application Manager API
 *
 * Application Manager APIs to control Apache Flink jobs
 *
 * API version: 2.4.3
 * Contact: platform@ververica.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package appmanagerapi

import (
	"time"
)

// ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of the resource that the fieldset applies to.
type V1ManagedFieldsEntry struct {
	// APIVersion defines the version of this resource that this field set applies to. The format is \"group/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted.
	ApiVersion string `json:"apiVersion,omitempty"`
	// FieldsType is the discriminator for the different fields format and version. There is currently only one possible value: \"FieldsV1\"
	FieldsType string `json:"fieldsType,omitempty"`
	// FieldsV1 holds the first JSON version format as described in the \"FieldsV1\" type.
	FieldsV1 interface{} `json:"fieldsV1,omitempty"`
	// Manager is an identifier of the workflow managing these fields.
	Manager string `json:"manager,omitempty"`
	// Operation is the type of operation which lead to this ManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'.
	Operation string `json:"operation,omitempty"`
	// Time is timestamp of when these fields were set. It should always be empty if Operation is 'Apply'
	Time time.Time `json:"time,omitempty"`
}
