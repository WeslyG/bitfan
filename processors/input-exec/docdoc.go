// Code generated by "bitfanDoc "; DO NOT EDIT
package execinput

import "github.com/vjeantet/bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:     "execinput",
  Doc:      "",
  DocShort: "",
  Options:  &doc.ProcessorOptions{
    Doc:     "",
    Options: []*doc.ProcessorOption{
      &doc.ProcessorOption{
        Name:         "Command",
        Alias:        "",
        Doc:          "",
        Required:     false,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Args",
        Alias:        "",
        Doc:          "",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Add_field",
        Alias:        "",
        Doc:          "",
        Required:     false,
        Type:         "hash",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Interval",
        Alias:        "",
        Doc:          "",
        Required:     false,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Codec",
        Alias:        "",
        Doc:          "",
        Required:     false,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Tags",
        Alias:        "",
        Doc:          "",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Type",
        Alias:        "",
        Doc:          "",
        Required:     false,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "",
      },
    },
  },
  Ports: []*doc.ProcessorPort{},
}
}