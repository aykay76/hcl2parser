package parser

// TODO: add models for different parsing methods; e.g. one that just grabs ranges so that
// i can know where in a file something is

import (
	"github.com/hashicorp/hcl2/hcl"
	"github.com/zclconf/go-cty/cty"
)

// Variable : Terraform variable definition
type Variable struct {
	Name         string         `hcl:"name,label"`
	DeclaredType hcl.Expression `hcl:"type,attr"`
	Default      *cty.Value     `hcl:"default,attr"`
	Description  hcl.Expression `hcl:"description,attr"`
	Sensitive    *bool          `hcl:"sensitive,attr"`
}

// Local : Terraform local variables definition
type Local struct {
	Remain hcl.Attributes `hcl:",remain"`
}

// Terraform : Terraform configuration definition
type Terraform struct {
	RequiredVersion string         `hcl:"required_version,attr"`
	Remain          hcl.Attributes `hcl:",remain"`
}

// Output : Terraform output definition
type Output struct {
	Name        string         `hcl:"name,label"`
	ValueExpr   hcl.Expression `hcl:"value,attr"`
	DependsOn   hcl.Expression `hcl:"depends_on,attr"`
	Description *string        `hcl:"description,attr"`
	Sensitive   *bool          `hcl:"sensitive,attr"`
}

// Provider : Terraform provider definition
type Provider struct {
	Name    string         `hcl:"name,label"`
	Alias   *string        `hcl:"alias,attr"`
	Version *string        `hcl:"version,optional"`
	Remain  hcl.Attributes `hcl:",remain"`
}

// Resource : Terraform resource definition
type Resource struct {
	Type string `hcl:"type,label"`
	Name string `hcl:"name,label"`

	// Not using these, but could if of interest (would need testing)
	// CountExpr hcl2.Expression `hcl:"count,attr"`
	// Provider  *string         `hcl:"provider,attr"`
	// DependsOn *[]string       `hcl:"depends_on,attr"`

	// Lifecycle    *resourceLifecycle `hcl:"lifecycle,block"`
	// Provisioners []provisioner      `hcl:"provisioner,block"`
	// Connection   *connection        `hcl:"connection,block"`

	Remain hcl.Body `hcl:",remain"`
}

// Data : Terraform data provider definition
type Data struct {
	Type string `hcl:"type,label"`
	Name string `hcl:"name,label"`

	// Not using these, but could if of interest (would need testing)
	// CountExpr hcl2.Expression `hcl:"count,attr"`
	// Provider  *string         `hcl:"provider,attr"`
	// DependsOn *[]string       `hcl:"depends_on,attr"`

	Remain hcl.Body `hcl:",remain"`
}

// Module : Terraform module definition
type Module struct {
	Name      string         `hcl:"name,label"`
	Source    string         `hcl:"source,attr"`
	Version   *string        `hcl:"version,attr"`
	Providers hcl.Expression `hcl:"providers,attr"`
	Remain    hcl.Body       `hcl:",remain"`
}

// Root : the root of the HCL config tree
type Root struct {
	Variables []*Variable  `hcl:"variable,block"`
	Resources []*Resource  `hcl:"resource,block"`
	Data      []*Data      `hcl:"data,block"`
	Modules   []*Module    `hcl:"module,block"`
	Outputs   []*Output    `hcl:"output,block"`
	Providers []*Provider  `hcl:"provider,block"`
	Locals    []*Local     `hcl:"locals,block"`
	Terraform []*Terraform `hcl:"terraform,block"`
}

// FileData : placeholder for file content where referenced by expression ranges
type FileData struct {
	Filename string
	Content  []byte
}
