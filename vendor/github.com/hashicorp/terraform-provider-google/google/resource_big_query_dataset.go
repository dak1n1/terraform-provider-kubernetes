// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"google.golang.org/api/googleapi"
)

const datasetIdRegexp = `[0-9A-Za-z_]+`

func validateDatasetId(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if !regexp.MustCompile(datasetIdRegexp).MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"%q must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_)", k))
	}

	if len(value) > 1024 {
		errors = append(errors, fmt.Errorf(
			"%q cannot be greater than 1,024 characters", k))
	}

	return
}

func validateDefaultTableExpirationMs(v interface{}, k string) (ws []string, errors []error) {
	value := v.(int)
	if value < 3600000 {
		errors = append(errors, fmt.Errorf("%q cannot be shorter than 3600000 milliseconds (one hour)", k))
	}

	return
}

func resourceBigQueryDataset() *schema.Resource {
	return &schema.Resource{
		Create: resourceBigQueryDatasetCreate,
		Read:   resourceBigQueryDatasetRead,
		Update: resourceBigQueryDatasetUpdate,
		Delete: resourceBigQueryDatasetDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBigQueryDatasetImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"dataset_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateDatasetId,
				Description: `A unique ID for this dataset, without the project name. The ID
must contain only letters (a-z, A-Z), numbers (0-9), or
underscores (_). The maximum length is 1,024 characters.`,
			},

			"access": {
				Type:        schema.TypeSet,
				Computed:    true,
				Optional:    true,
				Description: `An array of objects that define dataset access for one or more entities.`,
				Elem:        bigqueryDatasetAccessSchema(),
				// Default schema.HashSchema is used.
			},
			"default_encryption_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `The default encryption key for all tables in the dataset. Once this property is set,
all newly-created partitioned tables in the dataset will have encryption key set to
this value, unless table creation request (or query) overrides the key.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kms_key_name": {
							Type:     schema.TypeString,
							Required: true,
							Description: `Describes the Cloud KMS encryption key that will be used to protect destination
BigQuery table. The BigQuery Service Account associated with your project requires
access to this encryption key.`,
						},
					},
				},
			},
			"default_partition_expiration_ms": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: `The default partition expiration for all partitioned tables in
the dataset, in milliseconds.


Once this property is set, all newly-created partitioned tables in
the dataset will have an 'expirationMs' property in the 'timePartitioning'
settings set to this value, and changing the value will only
affect new tables, not existing ones. The storage in a partition will
have an expiration time of its partition time plus this value.
Setting this property overrides the use of 'defaultTableExpirationMs'
for partitioned tables: only one of 'defaultTableExpirationMs' and
'defaultPartitionExpirationMs' will be used for any new partitioned
table. If you provide an explicit 'timePartitioning.expirationMs' when
creating or updating a partitioned table, that value takes precedence
over the default partition expiration time indicated by this property.`,
			},
			"default_table_expiration_ms": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validateDefaultTableExpirationMs,
				Description: `The default lifetime of all tables in the dataset, in milliseconds.
The minimum value is 3600000 milliseconds (one hour).


Once this property is set, all newly-created tables in the dataset
will have an 'expirationTime' property set to the creation time plus
the value in this property, and changing the value will only affect
new tables, not existing ones. When the 'expirationTime' for a given
table is reached, that table will be deleted automatically.
If a table's 'expirationTime' is modified or removed before the
table expires, or if you provide an explicit 'expirationTime' when
creating a table, that value takes precedence over the default
expiration time indicated by this property.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A user-friendly description of the dataset`,
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A descriptive name for the dataset`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `The labels associated with this dataset. You can use these to
organize and group your datasets`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The geographic location where the dataset should reside.
See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).


There are two types of locations, regional or multi-regional. A regional
location is a specific geographic place, such as Tokyo, and a multi-regional
location is a large geographic area, such as the United States, that
contains at least two geographic places.


Possible regional values include: 'asia-east1', 'asia-northeast1',
'asia-southeast1', 'australia-southeast1', 'europe-north1',
'europe-west2' and 'us-east4'.


Possible multi-regional values: 'EU' and 'US'.


The default value is multi-regional location 'US'.
Changing this forces a new resource to be created.`,
				Default: "US",
			},
			"creation_time": {
				Type:     schema.TypeInt,
				Computed: true,
				Description: `The time when this dataset was created, in milliseconds since the
epoch.`,
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `A hash of the resource.`,
			},
			"last_modified_time": {
				Type:     schema.TypeInt,
				Computed: true,
				Description: `The date when this dataset or any of its tables was last modified, in
milliseconds since the epoch.`,
			},
			"delete_contents_on_destroy": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func bigqueryDatasetAccessSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `A domain to grant access to. Any users signed in with the
domain specified will be granted the specified access`,
			},
			"group_by_email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An email address of a Google Group to grant access to.`,
			},
			"role": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Describes the rights granted to the user specified by the other
member of the access object. Primitive, Predefined and custom
roles are supported. Predefined roles that have equivalent
primitive roles are swapped by the API to their Primitive
counterparts. See
[official docs](https://cloud.google.com/bigquery/docs/access-control).`,
			},
			"special_group": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `A special group to grant access to. Possible values include:


* 'projectOwners': Owners of the enclosing project.


* 'projectReaders': Readers of the enclosing project.


* 'projectWriters': Writers of the enclosing project.


* 'allAuthenticatedUsers': All authenticated BigQuery users.`,
			},
			"user_by_email": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `An email address of a user to grant access to. For example:
fred@example.com`,
			},
			"view": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `A view from a different dataset to grant access to. Queries
executed against that view will have read access to tables in
this dataset. The role field is not required when this field is
set. If that view is updated by any user, access to the view
needs to be granted again via an update operation.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dataset_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The ID of the dataset containing this table.`,
						},
						"project_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The ID of the project containing this table.`,
						},
						"table_id": {
							Type:     schema.TypeString,
							Required: true,
							Description: `The ID of the table. The ID must contain only letters (a-z,
A-Z), numbers (0-9), or underscores (_). The maximum length
is 1,024 characters.`,
						},
					},
				},
			},
		},
	}
}

func resourceBigQueryDatasetCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	accessProp, err := expandBigQueryDatasetAccess(d.Get("access"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("access"); !isEmptyValue(reflect.ValueOf(accessProp)) && (ok || !reflect.DeepEqual(v, accessProp)) {
		obj["access"] = accessProp
	}
	datasetReferenceProp, err := expandBigQueryDatasetDatasetReference(nil, d, config)
	if err != nil {
		return err
	} else if !isEmptyValue(reflect.ValueOf(datasetReferenceProp)) {
		obj["datasetReference"] = datasetReferenceProp
	}
	defaultTableExpirationMsProp, err := expandBigQueryDatasetDefaultTableExpirationMs(d.Get("default_table_expiration_ms"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_table_expiration_ms"); !isEmptyValue(reflect.ValueOf(defaultTableExpirationMsProp)) && (ok || !reflect.DeepEqual(v, defaultTableExpirationMsProp)) {
		obj["defaultTableExpirationMs"] = defaultTableExpirationMsProp
	}
	defaultPartitionExpirationMsProp, err := expandBigQueryDatasetDefaultPartitionExpirationMs(d.Get("default_partition_expiration_ms"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_partition_expiration_ms"); !isEmptyValue(reflect.ValueOf(defaultPartitionExpirationMsProp)) && (ok || !reflect.DeepEqual(v, defaultPartitionExpirationMsProp)) {
		obj["defaultPartitionExpirationMs"] = defaultPartitionExpirationMsProp
	}
	descriptionProp, err := expandBigQueryDatasetDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	friendlyNameProp, err := expandBigQueryDatasetFriendlyName(d.Get("friendly_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("friendly_name"); !isEmptyValue(reflect.ValueOf(friendlyNameProp)) && (ok || !reflect.DeepEqual(v, friendlyNameProp)) {
		obj["friendlyName"] = friendlyNameProp
	}
	labelsProp, err := expandBigQueryDatasetLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	locationProp, err := expandBigQueryDatasetLocation(d.Get("location"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("location"); !isEmptyValue(reflect.ValueOf(locationProp)) && (ok || !reflect.DeepEqual(v, locationProp)) {
		obj["location"] = locationProp
	}
	defaultEncryptionConfigurationProp, err := expandBigQueryDatasetDefaultEncryptionConfiguration(d.Get("default_encryption_configuration"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_encryption_configuration"); !isEmptyValue(reflect.ValueOf(defaultEncryptionConfigurationProp)) && (ok || !reflect.DeepEqual(v, defaultEncryptionConfigurationProp)) {
		obj["defaultEncryptionConfiguration"] = defaultEncryptionConfigurationProp
	}

	url, err := replaceVars(d, config, "{{BigQueryBasePath}}projects/{{project}}/datasets")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Dataset: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Dataset: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/datasets/{{dataset_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Dataset %q: %#v", d.Id(), res)

	return resourceBigQueryDatasetRead(d, meta)
}

func resourceBigQueryDatasetRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{BigQueryBasePath}}projects/{{project}}/datasets/{{dataset_id}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("BigQueryDataset %q", d.Id()))
	}

	// Explicitly set virtual fields to default values if unset
	if _, ok := d.GetOk("delete_contents_on_destroy"); !ok {
		d.Set("delete_contents_on_destroy", false)
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}

	if err := d.Set("access", flattenBigQueryDatasetAccess(res["access"], d, config)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}
	if err := d.Set("creation_time", flattenBigQueryDatasetCreationTime(res["creationTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if flattenedProp := flattenBigQueryDatasetDatasetReference(res["datasetReference"], d, config); flattenedProp != nil {
		if gerr, ok := flattenedProp.(*googleapi.Error); ok {
			return fmt.Errorf("Error reading Dataset: %s", gerr)
		}
		casted := flattenedProp.([]interface{})[0]
		if casted != nil {
			for k, v := range casted.(map[string]interface{}) {
				d.Set(k, v)
			}
		}
	}
	if err := d.Set("default_table_expiration_ms", flattenBigQueryDatasetDefaultTableExpirationMs(res["defaultTableExpirationMs"], d, config)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}
	if err := d.Set("default_partition_expiration_ms", flattenBigQueryDatasetDefaultPartitionExpirationMs(res["defaultPartitionExpirationMs"], d, config)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}
	if err := d.Set("description", flattenBigQueryDatasetDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}
	if err := d.Set("etag", flattenBigQueryDatasetEtag(res["etag"], d, config)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}
	if err := d.Set("friendly_name", flattenBigQueryDatasetFriendlyName(res["friendlyName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}
	if err := d.Set("labels", flattenBigQueryDatasetLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}
	if err := d.Set("last_modified_time", flattenBigQueryDatasetLastModifiedTime(res["lastModifiedTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}
	if err := d.Set("location", flattenBigQueryDatasetLocation(res["location"], d, config)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}
	if err := d.Set("default_encryption_configuration", flattenBigQueryDatasetDefaultEncryptionConfiguration(res["defaultEncryptionConfiguration"], d, config)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}

	return nil
}

func resourceBigQueryDatasetUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	accessProp, err := expandBigQueryDatasetAccess(d.Get("access"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("access"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, accessProp)) {
		obj["access"] = accessProp
	}
	datasetReferenceProp, err := expandBigQueryDatasetDatasetReference(nil, d, config)
	if err != nil {
		return err
	} else if !isEmptyValue(reflect.ValueOf(datasetReferenceProp)) {
		obj["datasetReference"] = datasetReferenceProp
	}
	defaultTableExpirationMsProp, err := expandBigQueryDatasetDefaultTableExpirationMs(d.Get("default_table_expiration_ms"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_table_expiration_ms"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, defaultTableExpirationMsProp)) {
		obj["defaultTableExpirationMs"] = defaultTableExpirationMsProp
	}
	defaultPartitionExpirationMsProp, err := expandBigQueryDatasetDefaultPartitionExpirationMs(d.Get("default_partition_expiration_ms"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_partition_expiration_ms"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, defaultPartitionExpirationMsProp)) {
		obj["defaultPartitionExpirationMs"] = defaultPartitionExpirationMsProp
	}
	descriptionProp, err := expandBigQueryDatasetDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	friendlyNameProp, err := expandBigQueryDatasetFriendlyName(d.Get("friendly_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("friendly_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, friendlyNameProp)) {
		obj["friendlyName"] = friendlyNameProp
	}
	labelsProp, err := expandBigQueryDatasetLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	locationProp, err := expandBigQueryDatasetLocation(d.Get("location"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("location"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, locationProp)) {
		obj["location"] = locationProp
	}
	defaultEncryptionConfigurationProp, err := expandBigQueryDatasetDefaultEncryptionConfiguration(d.Get("default_encryption_configuration"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_encryption_configuration"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, defaultEncryptionConfigurationProp)) {
		obj["defaultEncryptionConfiguration"] = defaultEncryptionConfigurationProp
	}

	url, err := replaceVars(d, config, "{{BigQueryBasePath}}projects/{{project}}/datasets/{{dataset_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Dataset %q: %#v", d.Id(), obj)
	res, err := sendRequestWithTimeout(config, "PUT", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Dataset %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Dataset %q: %#v", d.Id(), res)
	}

	return resourceBigQueryDatasetRead(d, meta)
}

func resourceBigQueryDatasetDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{BigQueryBasePath}}projects/{{project}}/datasets/{{dataset_id}}?deleteContents={{delete_contents_on_destroy}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Dataset %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Dataset")
	}

	log.Printf("[DEBUG] Finished deleting Dataset %q: %#v", d.Id(), res)
	return nil
}

func resourceBigQueryDatasetImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/datasets/(?P<dataset_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<dataset_id>[^/]+)",
		"(?P<dataset_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/datasets/{{dataset_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Explicitly set virtual fields to default values on import
	d.Set("delete_contents_on_destroy", false)

	return []*schema.ResourceData{d}, nil
}

func flattenBigQueryDatasetAccess(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(bigqueryDatasetAccessSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"domain":         flattenBigQueryDatasetAccessDomain(original["domain"], d, config),
			"group_by_email": flattenBigQueryDatasetAccessGroupByEmail(original["groupByEmail"], d, config),
			"role":           flattenBigQueryDatasetAccessRole(original["role"], d, config),
			"special_group":  flattenBigQueryDatasetAccessSpecialGroup(original["specialGroup"], d, config),
			"user_by_email":  flattenBigQueryDatasetAccessUserByEmail(original["userByEmail"], d, config),
			"view":           flattenBigQueryDatasetAccessView(original["view"], d, config),
		})
	}
	return transformed
}
func flattenBigQueryDatasetAccessDomain(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessGroupByEmail(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessRole(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessSpecialGroup(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessUserByEmail(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessView(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["dataset_id"] =
		flattenBigQueryDatasetAccessViewDatasetId(original["datasetId"], d, config)
	transformed["project_id"] =
		flattenBigQueryDatasetAccessViewProjectId(original["projectId"], d, config)
	transformed["table_id"] =
		flattenBigQueryDatasetAccessViewTableId(original["tableId"], d, config)
	return []interface{}{transformed}
}
func flattenBigQueryDatasetAccessViewDatasetId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessViewProjectId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessViewTableId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigQueryDatasetCreationTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenBigQueryDatasetDatasetReference(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["dataset_id"] =
		flattenBigQueryDatasetDatasetReferenceDatasetId(original["datasetId"], d, config)
	return []interface{}{transformed}
}
func flattenBigQueryDatasetDatasetReferenceDatasetId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigQueryDatasetDefaultTableExpirationMs(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenBigQueryDatasetDefaultPartitionExpirationMs(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenBigQueryDatasetDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigQueryDatasetEtag(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigQueryDatasetFriendlyName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigQueryDatasetLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigQueryDatasetLastModifiedTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

// Older Datasets in BigQuery have no Location set in the API response. This may be an issue when importing
// datasets created before BigQuery was available in multiple zones. We can safely assume that these datasets
// are in the US, as this was the default at the time.
func flattenBigQueryDatasetLocation(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return "US"
	}
	return v
}

func flattenBigQueryDatasetDefaultEncryptionConfiguration(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["kms_key_name"] =
		flattenBigQueryDatasetDefaultEncryptionConfigurationKmsKeyName(original["kmsKeyName"], d, config)
	return []interface{}{transformed}
}
func flattenBigQueryDatasetDefaultEncryptionConfigurationKmsKeyName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandBigQueryDatasetAccess(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDomain, err := expandBigQueryDatasetAccessDomain(original["domain"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDomain); val.IsValid() && !isEmptyValue(val) {
			transformed["domain"] = transformedDomain
		}

		transformedGroupByEmail, err := expandBigQueryDatasetAccessGroupByEmail(original["group_by_email"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGroupByEmail); val.IsValid() && !isEmptyValue(val) {
			transformed["groupByEmail"] = transformedGroupByEmail
		}

		transformedRole, err := expandBigQueryDatasetAccessRole(original["role"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRole); val.IsValid() && !isEmptyValue(val) {
			transformed["role"] = transformedRole
		}

		transformedSpecialGroup, err := expandBigQueryDatasetAccessSpecialGroup(original["special_group"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSpecialGroup); val.IsValid() && !isEmptyValue(val) {
			transformed["specialGroup"] = transformedSpecialGroup
		}

		transformedUserByEmail, err := expandBigQueryDatasetAccessUserByEmail(original["user_by_email"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedUserByEmail); val.IsValid() && !isEmptyValue(val) {
			transformed["userByEmail"] = transformedUserByEmail
		}

		transformedView, err := expandBigQueryDatasetAccessView(original["view"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedView); val.IsValid() && !isEmptyValue(val) {
			transformed["view"] = transformedView
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandBigQueryDatasetAccessDomain(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessGroupByEmail(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessRole(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessSpecialGroup(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessUserByEmail(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessView(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDatasetId, err := expandBigQueryDatasetAccessViewDatasetId(original["dataset_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatasetId); val.IsValid() && !isEmptyValue(val) {
		transformed["datasetId"] = transformedDatasetId
	}

	transformedProjectId, err := expandBigQueryDatasetAccessViewProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !isEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	transformedTableId, err := expandBigQueryDatasetAccessViewTableId(original["table_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTableId); val.IsValid() && !isEmptyValue(val) {
		transformed["tableId"] = transformedTableId
	}

	return transformed, nil
}

func expandBigQueryDatasetAccessViewDatasetId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessViewProjectId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessViewTableId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetDatasetReference(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	transformedDatasetId, err := expandBigQueryDatasetDatasetReferenceDatasetId(d.Get("dataset_id"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatasetId); val.IsValid() && !isEmptyValue(val) {
		transformed["datasetId"] = transformedDatasetId
	}

	return transformed, nil
}

func expandBigQueryDatasetDatasetReferenceDatasetId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetDefaultTableExpirationMs(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetDefaultPartitionExpirationMs(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetFriendlyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandBigQueryDatasetLocation(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetDefaultEncryptionConfiguration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeyName, err := expandBigQueryDatasetDefaultEncryptionConfigurationKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !isEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	return transformed, nil
}

func expandBigQueryDatasetDefaultEncryptionConfigurationKmsKeyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
