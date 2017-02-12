package whitelist

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vjeantet/bitfan/processors/doc"
	"github.com/vjeantet/bitfan/processors/testutils"
)

func TestNew(t *testing.T) {
	p := New()
	_, ok := p.(*processor)
	assert.Equal(t, ok, true, "New() should return a processor")
}
func TestDoc(t *testing.T) {
	assert.IsType(t, &doc.Processor{}, New().(*processor).Doc())
}

func TestReceiveMatch(t *testing.T) {
	p := New().(*processor)
	ctx := testutils.NewProcessorContext()
	p.Configure(
		ctx,
		map[string]interface{}{
			"Compare_Field": "message",
			"list":          []string{"val1", "val2"},
		},
	)

	p.Receive(testutils.NewPacket("test", nil))
	p.Receive(testutils.NewPacket("fqsdf", nil))
	p.Receive(testutils.NewPacket("valo", nil))
	p.Receive(testutils.NewPacket("al1", nil))
	p.Receive(testutils.NewPacket("val1", nil))
	p.Receive(testutils.NewPacket("val3", nil))
	p.Receive(testutils.NewPacket("val2", nil))
	assert.Equal(t, 5, ctx.SentPacketsCount(0), "5 events should pass")
}