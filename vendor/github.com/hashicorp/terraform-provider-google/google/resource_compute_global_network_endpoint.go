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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceComputeGlobalNetworkEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeGlobalNetworkEndpointCreate,
		Read:   resourceComputeGlobalNetworkEndpointRead,
		Delete: resourceComputeGlobalNetworkEndpointDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeGlobalNetworkEndpointImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(6 * time.Minute),
			Delete: schema.DefaultTimeout(6 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"global_network_endpoint_group": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareResourceNames,
				Description:      `The global network endpoint group this endpoint is part of.`,
			},
			"port": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: `Port number of the external endpoint.`,
			},
			"fqdn": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `Fully qualified domain name of network endpoint.
This can only be specified when network_endpoint_type of the NEG is INTERNET_FQDN_PORT.`,
				AtLeastOneOf: []string{"fqdn", "ip_address"},
			},
			"ip_address": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `IPv4 address external endpoint.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceComputeGlobalNetworkEndpointCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	portProp, err := expandNestedComputeGlobalNetworkEndpointPort(d.Get("port"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("port"); !isEmptyValue(reflect.ValueOf(portProp)) && (ok || !reflect.DeepEqual(v, portProp)) {
		obj["port"] = portProp
	}
	ipAddressProp, err := expandNestedComputeGlobalNetworkEndpointIpAddress(d.Get("ip_address"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_address"); !isEmptyValue(reflect.ValueOf(ipAddressProp)) && (ok || !reflect.DeepEqual(v, ipAddressProp)) {
		obj["ipAddress"] = ipAddressProp
	}
	fqdnProp, err := expandNestedComputeGlobalNetworkEndpointFqdn(d.Get("fqdn"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fqdn"); !isEmptyValue(reflect.ValueOf(fqdnProp)) && (ok || !reflect.DeepEqual(v, fqdnProp)) {
		obj["fqdn"] = fqdnProp
	}

	obj, err = resourceComputeGlobalNetworkEndpointEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	lockName, err := replaceVars(d, config, "networkEndpoint/{{project}}/{{global_network_endpoint_group}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/networkEndpointGroups/{{global_network_endpoint_group}}/attachNetworkEndpoints")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new GlobalNetworkEndpoint: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating GlobalNetworkEndpoint: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{project}}/{{global_network_endpoint_group}}/{{ip_address}}/{{fqdn}}/{{port}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating GlobalNetworkEndpoint",
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create GlobalNetworkEndpoint: %s", err)
	}

	log.Printf("[DEBUG] Finished creating GlobalNetworkEndpoint %q: %#v", d.Id(), res)

	return resourceComputeGlobalNetworkEndpointRead(d, meta)
}

func resourceComputeGlobalNetworkEndpointRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/networkEndpointGroups/{{global_network_endpoint_group}}/listNetworkEndpoints")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "POST", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeGlobalNetworkEndpoint %q", d.Id()))
	}

	res, err = flattenNestedComputeGlobalNetworkEndpoint(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing ComputeGlobalNetworkEndpoint because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	res, err = resourceComputeGlobalNetworkEndpointDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing ComputeGlobalNetworkEndpoint because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading GlobalNetworkEndpoint: %s", err)
	}

	if err := d.Set("port", flattenNestedComputeGlobalNetworkEndpointPort(res["port"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalNetworkEndpoint: %s", err)
	}
	if err := d.Set("ip_address", flattenNestedComputeGlobalNetworkEndpointIpAddress(res["ipAddress"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalNetworkEndpoint: %s", err)
	}
	if err := d.Set("fqdn", flattenNestedComputeGlobalNetworkEndpointFqdn(res["fqdn"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalNetworkEndpoint: %s", err)
	}

	return nil
}

func resourceComputeGlobalNetworkEndpointDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	lockName, err := replaceVars(d, config, "networkEndpoint/{{project}}/{{global_network_endpoint_group}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/networkEndpointGroups/{{global_network_endpoint_group}}/detachNetworkEndpoints")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	toDelete := make(map[string]interface{})
	portProp, err := expandNestedComputeGlobalNetworkEndpointPort(d.Get("port"), d, config)
	if err != nil {
		return err
	}
	if portProp != "" {
		toDelete["port"] = portProp
	}

	ipAddressProp, err := expandNestedComputeGlobalNetworkEndpointIpAddress(d.Get("ip_address"), d, config)
	if err != nil {
		return err
	}
	if ipAddressProp != "" {
		toDelete["ipAddress"] = ipAddressProp
	}

	fqdnProp, err := expandNestedComputeGlobalNetworkEndpointFqdn(d.Get("fqdn"), d, config)
	if err != nil {
		return err
	}
	if fqdnProp != "" {
		toDelete["fqdn"] = fqdnProp
	}

	obj = map[string]interface{}{
		"networkEndpoints": []map[string]interface{}{toDelete},
	}
	log.Printf("[DEBUG] Deleting GlobalNetworkEndpoint %q", d.Id())

	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "GlobalNetworkEndpoint")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting GlobalNetworkEndpoint",
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting GlobalNetworkEndpoint %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeGlobalNetworkEndpointImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	// FQDN, port and ip_address are optional, so use * instead of + when reading the import id
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]?)/global/networkEndpointGroups/(?P<global_network_endpoint_group>[^/]+)/(?P<ip_address>[^/]+)/(?P<fqdn>[^/]*)/(?P<port>[^/]+)",
		"(?P<project>[^/]+)/(?P<global_network_endpoint_group>[^/]+)/(?P<ip_address>[^/]*)/(?P<fqdn>[^/]*)/(?P<port>[^/]*)",
		"(?P<global_network_endpoint_group>[^/]+)/(?P<ip_address>[^/]*)/(?P<fqdn>[^/]*)/(?P<port>[^/]*)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{project}}/{{global_network_endpoint_group}}/{{ip_address}}/{{fqdn}}/{{port}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNestedComputeGlobalNetworkEndpointPort(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles int given in float64 format
	if floatVal, ok := v.(float64); ok {
		return int(floatVal)
	}
	return v
}

func flattenNestedComputeGlobalNetworkEndpointIpAddress(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNestedComputeGlobalNetworkEndpointFqdn(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandNestedComputeGlobalNetworkEndpointPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNestedComputeGlobalNetworkEndpointIpAddress(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNestedComputeGlobalNetworkEndpointFqdn(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceComputeGlobalNetworkEndpointEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Network Endpoint Group is a URL parameter only, so replace self-link/path with resource name only.
	d.Set("global_network_endpoint_group", GetResourceNameFromSelfLink(d.Get("global_network_endpoint_group").(string)))

	wrappedReq := map[string]interface{}{
		"networkEndpoints": []interface{}{obj},
	}
	return wrappedReq, nil
}

func flattenNestedComputeGlobalNetworkEndpoint(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["items"]
	if !ok || v == nil {
		return nil, nil
	}

	switch v.(type) {
	case []interface{}:
		break
	case map[string]interface{}:
		// Construct list out of single nested resource
		v = []interface{}{v}
	default:
		return nil, fmt.Errorf("expected list or map for value items. Actual value: %v", v)
	}

	_, item, err := resourceComputeGlobalNetworkEndpointFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceComputeGlobalNetworkEndpointFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedFqdn, err := expandNestedComputeGlobalNetworkEndpointFqdn(d.Get("fqdn"), d, meta.(*Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedFqdn := flattenNestedComputeGlobalNetworkEndpointFqdn(expectedFqdn, d, meta.(*Config))
	expectedIpAddress, err := expandNestedComputeGlobalNetworkEndpointIpAddress(d.Get("ip_address"), d, meta.(*Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedIpAddress := flattenNestedComputeGlobalNetworkEndpointIpAddress(expectedIpAddress, d, meta.(*Config))
	expectedPort, err := expandNestedComputeGlobalNetworkEndpointPort(d.Get("port"), d, meta.(*Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedPort := flattenNestedComputeGlobalNetworkEndpointPort(expectedPort, d, meta.(*Config))

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		// Decode list item before comparing.
		item, err := resourceComputeGlobalNetworkEndpointDecoder(d, meta, item)
		if err != nil {
			return -1, nil, err
		}

		itemFqdn := flattenNestedComputeGlobalNetworkEndpointFqdn(item["fqdn"], d, meta.(*Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemFqdn)) && isEmptyValue(reflect.ValueOf(expectedFlattenedFqdn))) && !reflect.DeepEqual(itemFqdn, expectedFlattenedFqdn) {
			log.Printf("[DEBUG] Skipping item with fqdn= %#v, looking for %#v)", itemFqdn, expectedFlattenedFqdn)
			continue
		}
		itemIpAddress := flattenNestedComputeGlobalNetworkEndpointIpAddress(item["ipAddress"], d, meta.(*Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemIpAddress)) && isEmptyValue(reflect.ValueOf(expectedFlattenedIpAddress))) && !reflect.DeepEqual(itemIpAddress, expectedFlattenedIpAddress) {
			log.Printf("[DEBUG] Skipping item with ipAddress= %#v, looking for %#v)", itemIpAddress, expectedFlattenedIpAddress)
			continue
		}
		itemPort := flattenNestedComputeGlobalNetworkEndpointPort(item["port"], d, meta.(*Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemPort)) && isEmptyValue(reflect.ValueOf(expectedFlattenedPort))) && !reflect.DeepEqual(itemPort, expectedFlattenedPort) {
			log.Printf("[DEBUG] Skipping item with port= %#v, looking for %#v)", itemPort, expectedFlattenedPort)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}
func resourceComputeGlobalNetworkEndpointDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	v, ok := res["networkEndpoint"]
	if !ok || v == nil {
		return res, nil
	}

	return v.(map[string]interface{}), nil
}
