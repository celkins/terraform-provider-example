package example

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
)

type Project struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
}

func resourceProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceProjectCreate,
		Update: resourceProjectUpdate,
		Read:   resourceProjectRead,
		Delete: resourceProjectDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceProjectCreate(d *schema.ResourceData, m interface{}) error {

	client := m.(*ExampleClient)
	proj := newProjectFromResource(d)

	bytedata, err := json.Marshal(proj)

	if err != nil {
		return err
	}

	_, err = client.Post(fmt.Sprintf("projects/%s",
		d.Get("name").(string),
	), bytes.NewBuffer(bytedata))

	if err != nil {
		return err
	}

	d.SetId(string(fmt.Sprintf("%s", d.Get("name").(string))))

	return resourceProjectRead(d, m)
}

func newProjectFromResource(d *schema.ResourceData) *Project {
	proj := &Project{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	return proj
}

func resourceProjectRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*ExampleClient)
	proj_req, _ := client.Get(fmt.Sprintf("projects/%s",
		d.Get("name").(string),
	))

	if proj_req.StatusCode == 200 {

		var proj Project

		body, readerr := ioutil.ReadAll(proj_req.Body)
		if readerr != nil {
			return readerr
		}

		decodeerr := json.Unmarshal(body, &proj)
		if decodeerr != nil {
			return decodeerr
		}

		d.Set("name", proj.Name)
		d.Set("description", proj.Description)
	}

	return nil
}

func resourceProjectUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*ExampleClient)
	proj := newProjectFromResource(d)

	var jsonbuffer []byte

	jsonpayload := bytes.NewBuffer(jsonbuffer)
	enc := json.NewEncoder(jsonpayload)
	enc.Encode(proj)

	_, err := client.Put(fmt.Sprintf("projects/%s",
		d.Get("name").(string),
	), jsonpayload)

	if err != nil {
		return err
	}

	return resourceProjectRead(d, m)
}

func resourceProjectDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*ExampleClient)
	_, err := client.Delete(fmt.Sprintf("projects/%s",
		d.Get("name").(string),
	))

	return err
}
