// Code generated by "bitfanDoc "; DO NOT EDIT
package unixinput

import "github.com/vjeantet/bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:     "unixinput",
  Doc:      "",
  DocShort: "",
  Options:  &doc.ProcessorOptions{
    Doc:     "",
    Options: []*doc.ProcessorOption{
      &doc.ProcessorOption{
        Name:         "AddField",
        Alias:        "add_field",
        Doc:          "If this filter is successful, add any arbitrary fields to this event.",
        Required:     false,
        Type:         "hash",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "DataTimeout",
        Alias:        "data_timeout",
        Doc:          "The read timeout in seconds. If a particular connection is idle for more than this timeout period, we will assume it is dead and close it.\nIf you never want to timeout, use 0.\nDefault value is 0",
        Required:     false,
        Type:         "time.Duration",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "ForceUnlink",
        Alias:        "force_unlink",
        Doc:          "Remove socket file in case of EADDRINUSE failure\nDefault value is false",
        Required:     false,
        Type:         "bool",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Mode",
        Alias:        "mode",
        Doc:          "Mode to operate in. server listens for client connections, client connects to a server.\nValue can be any of: \"server\", \"client\"\nDefault value is \"server\"",
        Required:     false,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Path",
        Alias:        "path",
        Doc:          "When mode is server, the path to listen on. When mode is client, the path to connect to.",
        Required:     true,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Tags",
        Alias:        "tags",
        Doc:          "If this filter is successful, add arbitrary tags to the event. Tags can be dynamic\nand include parts of the event using the %{field} syntax.",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Type",
        Alias:        "type",
        Doc:          "Add a type field to all events handled by this input",
        Required:     false,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Codec",
        Alias:        "codec",
        Doc:          "The codec used for input data. Input codecs are a convenient method for decoding\nyour data before it enters the input, without needing a separate filter in your bitfan pipeline",
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