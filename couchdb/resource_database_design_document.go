package couchdb

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/rossmerr/couchdb_go/client/design_documents"
	"github.com/rossmerr/couchdb_go/models"
)

func resourceDesignDocument() *schema.Resource {
	return &schema.Resource{
		CreateContext: designDocumentCreate,
		ReadContext:   designDocumentRead,
		UpdateContext: designDocumentUpdate,
		DeleteContext: designDocumentDelete,

		Schema: map[string]*schema.Schema{
			"database": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Database to associate design with",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the design document",
			},
			"revision": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Revision",
			},
			"language": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "javascript",
				Description: "Language of map/ reduce functions",
			},
			"views": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The views inside the design document (wrap in <<EOF { } EOF)",
				StateFunc: func(i interface{}) string {
					viewsDoc := map[string]interface{}{}
					if err := json.Unmarshal([]byte(i.(string)), &viewsDoc); err != nil {
						return ""
					}
					b, err := json.Marshal(viewsDoc)
					if err != nil {
						return ""
					}
					return string(b)
				},
				DefaultFunc: func() (interface{}, error) {
					return "{}", nil
				},
			},
			"indexes": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The indexes inside the design document (wrap in <<EOF { } EOF)",
				StateFunc: func(i interface{}) string {
					viewsDoc := map[string]interface{}{}
					if err := json.Unmarshal([]byte(i.(string)), &viewsDoc); err != nil {
						return ""
					}
					b, err := json.Marshal(viewsDoc)
					if err != nil {
						return ""
					}
					return string(b)
				},
				DefaultFunc: func() (interface{}, error) {
					return "{}", nil
				},
			},
		},
	}
}

func designDocumentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	client, dd := connectToCouchDB(ctx, meta.(*CouchDBConfiguration))
	if dd != nil {
		return append(diags, *dd)
	}

	dbName := d.Get("database").(string)

	docId := fmt.Sprintf("%s", d.Get("name").(string))

	viewsDoc := map[string]interface{}{}
	if interfaceViews, ok := d.GetOk("views"); ok {
		if err := json.Unmarshal([]byte(interfaceViews.(string)), &viewsDoc); err != nil {
			return AppendDiagnostic(diags, err, "Unable to unmarshal JSON")
		}
	}

	indexesDoc := map[string]interface{}{}
	if interfaceIndexes, ok := d.GetOk("indexes"); ok {
		if err := json.Unmarshal([]byte(interfaceIndexes.(string)), &indexesDoc); err != nil {
			return AppendDiagnostic(diags, err, "Unable to unmarshal JSON")
		}
	}

	designDoc := &models.DesignDoc{
		Language: d.Get("language").(string),
		Views:    viewsDoc,
		Indexes:  indexesDoc,
	}

	params := design_documents.NewDesignDocPutParams().WithDb(dbName).WithBody(designDoc).WithDdoc(docId)
	created, accepted, err := client.DesignDocuments.DesignDocPut(params)
	if err != nil {
		return AppendDiagnostic(diags, err, "Unable to create design doc")
	}

	if created != nil {
		d.Set("revision", strings.Trim(created.ETag, "\""))
	}

	if accepted != nil {
		d.Set("revision", strings.Trim(accepted.ETag, "\""))
	}

	d.SetId(docId)

	return designDocumentRead(ctx, d, meta)
}

func designDocumentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	client, dd := connectToCouchDB(ctx, meta.(*CouchDBConfiguration))
	if dd != nil {
		return append(diags, *dd)
	}
	dbName := d.Get("database").(string)
	docId := fmt.Sprintf("%s", d.Get("name").(string))
	rev := d.Get("revision").(string)

	params := design_documents.NewDesignDocGetParams().WithDb(dbName).WithDdoc(docId).WithRev(&rev)
	ok, err := client.DesignDocuments.DesignDocGet(params)
	if err != nil {
		return AppendDiagnostic(diags, err, "Unable to read Design Document")
	}

	if ok.Payload != nil {
		d.Set("language", ok.Payload.Language)

		if ok.Payload.Views != nil {
			if data, err := json.Marshal(ok.Payload.Views); err == nil {
				d.Set("views", string(data))
			} else {
				d.Set("views", nil)
			}
		}

		if ok.Payload.Indexes != nil {
			if data, err := json.Marshal(ok.Payload.Indexes); err == nil {
				d.Set("indexes", string(data))
			} else {
				d.Set("indexes", nil)
			}
		}
	}

	d.Set("revision", strings.Trim(ok.ETag, "\""))
	return diags
}

func designDocumentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	client, dd := connectToCouchDB(ctx, meta.(*CouchDBConfiguration))
	if dd != nil {
		return append(diags, *dd)
	}

	dbName := d.Get("database").(string)

	viewsDoc := map[string]interface{}{}
	if interfaceViews, ok := d.GetOk("views"); ok {
		if err := json.Unmarshal([]byte(interfaceViews.(string)), &viewsDoc); err != nil {
			return AppendDiagnostic(diags, err, "Unable to unmarshal JSON")
		}
	}

	indexesDoc := map[string]interface{}{}
	if interfaceIndexes, ok := d.GetOk("indexes"); ok {
		if err := json.Unmarshal([]byte(interfaceIndexes.(string)), &indexesDoc); err != nil {
			return AppendDiagnostic(diags, err, "Unable to unmarshal JSON")
		}
	}

	designDoc := &models.DesignDoc{
		Language: d.Get("language").(string),
		Views:    viewsDoc,
		Indexes:  indexesDoc,
	}

	params := design_documents.NewDesignDocPutParams().WithDb(dbName).WithBody(designDoc).WithDdoc(d.Id())
	created, accepted, err := client.DesignDocuments.DesignDocPut(params)
	if err != nil {
		return AppendDiagnostic(diags, err, "Unable to update design doc")
	}

	if created != nil {
		d.Set("revision", strings.Trim(created.ETag, "\""))
	}

	if accepted != nil {
		d.Set("revision", strings.Trim(accepted.ETag, "\""))
	}

	return diags
}

func designDocumentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	client, dd := connectToCouchDB(ctx, meta.(*CouchDBConfiguration))
	if dd != nil {
		return append(diags, *dd)
	}

	dbName := d.Get("database").(string)
	rev := d.Get("revision").(string)

	params := design_documents.NewDesignDocDeleteParams().WithDb(dbName).WithDdoc(d.Id()).WithRev(&rev)
	ok, accepted, err := client.DesignDocuments.DesignDocDelete(params)
	if err != nil {
		return AppendDiagnostic(diags, err, "Unable to delete design doc")
	}

	d.SetId("")
	if ok != nil {
		d.Set("revision", strings.Trim(ok.ETag, "\""))
	}

	if accepted != nil {
		d.Set("revision", strings.Trim(accepted.ETag, "\""))
	}

	return diags
}
