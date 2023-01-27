// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFirebaseAndroidApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirebaseAndroidAppCreate,
		Read:   resourceFirebaseAndroidAppRead,
		Update: resourceFirebaseAndroidAppUpdate,
		Delete: resourceFirebaseAndroidAppDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFirebaseAndroidAppImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The user-assigned display name of the AndroidApp.`,
			},
			"deletion_policy": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `(Optional) Set to 'ABANDON' to allow the AndroidApp to be untracked from terraform state
rather than deleted upon 'terraform destroy'. This is useful because the AndroidApp may be
serving traffic. Set to 'DELETE' to delete the AndroidApp. Default to 'DELETE'.`,
				Default: "DELETE",
			},
			"package_name": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Immutable. The canonical package name of the Android app as would appear in the Google Play
Developer Console.`,
			},
			"sha1_hashes": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `The SHA1 certificate hashes for the AndroidApp.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sha256_hashes": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `The SHA256 certificate hashes for the AndroidApp.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"app_id": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The globally unique, Firebase-assigned identifier of the AndroidApp.
This identifier should be treated as an opaque token, as the data format is not specified.`,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `This checksum is computed by the server based on the value of other fields, and it may be sent
with update requests to ensure the client has an up-to-date value before proceeding.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The fully qualified resource name of the AndroidApp, for example:
projects/projectId/androidApps/appId`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceFirebaseAndroidAppCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandFirebaseAndroidAppDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	packageNameProp, err := expandFirebaseAndroidAppPackageName(d.Get("package_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("package_name"); !isEmptyValue(reflect.ValueOf(packageNameProp)) && (ok || !reflect.DeepEqual(v, packageNameProp)) {
		obj["packageName"] = packageNameProp
	}
	sha1HashesProp, err := expandFirebaseAndroidAppSha1Hashes(d.Get("sha1_hashes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("sha1_hashes"); !isEmptyValue(reflect.ValueOf(sha1HashesProp)) && (ok || !reflect.DeepEqual(v, sha1HashesProp)) {
		obj["sha1Hashes"] = sha1HashesProp
	}
	sha256HashesProp, err := expandFirebaseAndroidAppSha256Hashes(d.Get("sha256_hashes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("sha256_hashes"); !isEmptyValue(reflect.ValueOf(sha256HashesProp)) && (ok || !reflect.DeepEqual(v, sha256HashesProp)) {
		obj["sha256Hashes"] = sha256HashesProp
	}
	etagProp, err := expandFirebaseAndroidAppEtag(d.Get("etag"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("etag"); !isEmptyValue(reflect.ValueOf(etagProp)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}

	url, err := replaceVars(d, config, "{{FirebaseBasePath}}projects/{{project}}/androidApps")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new AndroidApp: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AndroidApp: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating AndroidApp: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = firebaseOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating AndroidApp", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create AndroidApp: %s", err)
	}

	if err := d.Set("name", flattenFirebaseAndroidAppName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating AndroidApp %q: %#v", d.Id(), res)

	return resourceFirebaseAndroidAppRead(d, meta)
}

func resourceFirebaseAndroidAppRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{FirebaseBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AndroidApp: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("FirebaseAndroidApp %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading AndroidApp: %s", err)
	}

	if err := d.Set("name", flattenFirebaseAndroidAppName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading AndroidApp: %s", err)
	}
	if err := d.Set("display_name", flattenFirebaseAndroidAppDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading AndroidApp: %s", err)
	}
	if err := d.Set("app_id", flattenFirebaseAndroidAppAppId(res["appId"], d, config)); err != nil {
		return fmt.Errorf("Error reading AndroidApp: %s", err)
	}
	if err := d.Set("package_name", flattenFirebaseAndroidAppPackageName(res["packageName"], d, config)); err != nil {
		return fmt.Errorf("Error reading AndroidApp: %s", err)
	}
	if err := d.Set("sha1_hashes", flattenFirebaseAndroidAppSha1Hashes(res["sha1Hashes"], d, config)); err != nil {
		return fmt.Errorf("Error reading AndroidApp: %s", err)
	}
	if err := d.Set("sha256_hashes", flattenFirebaseAndroidAppSha256Hashes(res["sha256Hashes"], d, config)); err != nil {
		return fmt.Errorf("Error reading AndroidApp: %s", err)
	}
	if err := d.Set("etag", flattenFirebaseAndroidAppEtag(res["etag"], d, config)); err != nil {
		return fmt.Errorf("Error reading AndroidApp: %s", err)
	}

	return nil
}

func resourceFirebaseAndroidAppUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AndroidApp: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandFirebaseAndroidAppDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	packageNameProp, err := expandFirebaseAndroidAppPackageName(d.Get("package_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("package_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, packageNameProp)) {
		obj["packageName"] = packageNameProp
	}
	sha1HashesProp, err := expandFirebaseAndroidAppSha1Hashes(d.Get("sha1_hashes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("sha1_hashes"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sha1HashesProp)) {
		obj["sha1Hashes"] = sha1HashesProp
	}
	sha256HashesProp, err := expandFirebaseAndroidAppSha256Hashes(d.Get("sha256_hashes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("sha256_hashes"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sha256HashesProp)) {
		obj["sha256Hashes"] = sha256HashesProp
	}
	etagProp, err := expandFirebaseAndroidAppEtag(d.Get("etag"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("etag"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}

	url, err := replaceVars(d, config, "{{FirebaseBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating AndroidApp %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("package_name") {
		updateMask = append(updateMask, "packageName")
	}

	if d.HasChange("sha1_hashes") {
		updateMask = append(updateMask, "sha1Hashes")
	}

	if d.HasChange("sha256_hashes") {
		updateMask = append(updateMask, "sha256Hashes")
	}

	if d.HasChange("etag") {
		updateMask = append(updateMask, "etag")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating AndroidApp %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating AndroidApp %q: %#v", d.Id(), res)
	}

	return resourceFirebaseAndroidAppRead(d, meta)
}

func resourceFirebaseAndroidAppDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	// Handwritten
	obj := make(map[string]interface{})
	if d.Get("deletion_policy") == "DELETE" {
		obj["immediate"] = true
	} else {
		fmt.Printf("Skip deleting App %q due to deletion_policy: %q\n", d.Id(), d.Get("deletion_policy"))
		return nil
	}
	// End of Handwritten
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for App: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{FirebaseBasePath}}{{name}}:remove")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting App %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "App")
	}

	err = firebaseOperationWaitTime(
		config, res, project, "Deleting App", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting App %q: %#v", d.Id(), res)

	// This is useful if the Delete operation returns before the Get operation
	// during post-test destroy shows the completed state of the resource.
	time.Sleep(5 * time.Second)

	return nil
}

func resourceFirebaseAndroidAppImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<project>[^ ]+) (?P<name>[^ ]+)", "(?P<name>[^ ]+)"}, d, config); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func flattenFirebaseAndroidAppName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseAndroidAppDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseAndroidAppAppId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseAndroidAppPackageName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseAndroidAppSha1Hashes(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseAndroidAppSha256Hashes(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseAndroidAppEtag(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandFirebaseAndroidAppDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseAndroidAppPackageName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseAndroidAppSha1Hashes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseAndroidAppSha256Hashes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseAndroidAppEtag(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}