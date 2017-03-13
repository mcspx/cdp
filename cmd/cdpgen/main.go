// The cdpgen command generates the package cdp from the provided protocol definitions.
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"

	"github.com/mafredri/cdp/cmd/cdpgen/proto"
)

// Global constants.
const (
	OptionalPropPrefix = ""
	realEnum           = true
)

var (
	nonPtrMap = make(map[string]bool)
)

func main() {
	var (
		destPkg          string
		browserProtoJSON string
		jsProtoFileJSON  string
	)
	flag.StringVar(&destPkg, "dest-pkg", "", "Destination for generated cdp package (inside $GOPATH)")
	flag.StringVar(&browserProtoJSON, "browser-proto", "./protodef/browser_protocol.json", "Path to browser protocol")
	flag.StringVar(&jsProtoFileJSON, "js-proto", "./protodef/js_protocol.json", "Path to JS protocol")
	flag.Parse()

	if destPkg == "" {
		fmt.Fprintln(os.Stderr, "error: dest-pkg must be set")
		os.Exit(1)
	}

	var protocol, jsProtocol proto.Protocol
	protocolData, err := ioutil.ReadFile(browserProtoJSON)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(protocolData, &protocol)
	if err != nil {
		panic(err)
	}
	jsProtocolData, err := ioutil.ReadFile(jsProtoFileJSON)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsProtocolData, &jsProtocol)
	if err != nil {
		panic(err)
	}

	protocol.Domains = append(protocol.Domains, jsProtocol.Domains...)
	sort.Slice(protocol.Domains, func(i, j int) bool {
		return protocol.Domains[i].Domain < protocol.Domains[j].Domain
	})

	var g, tg, cg, eg Generator

	g.dir = destPkg
	g.pkg = "cdp"
	os.Mkdir(g.path(), 0755)

	tg.pkg = "cdptype"
	tg.dir = path.Join(g.dir, tg.pkg)
	os.Mkdir(tg.path(), 0755)

	cg.pkg = "cdpcmd"
	cg.dir = path.Join(g.dir, cg.pkg)
	os.Mkdir(cg.path(), 0755)

	eg.pkg = "cdpevent"
	eg.dir = path.Join(g.dir, eg.pkg)
	os.Mkdir(eg.path(), 0755)

	g.imports = []string{
		"github.com/mafredri/cdp/rpcc",
		tg.dir, cg.dir, eg.dir,
	}
	cg.imports = []string{tg.dir}
	eg.imports = []string{tg.dir}

	// Define the cdp Client.
	g.PackageHeader()
	g.CdpClient(protocol.Domains)
	g.writeFile("cdp_client.go")

	// Define all CDP command methods.
	cg.PackageHeader()
	cg.CmdType(protocol.Domains)
	cg.writeFile("cmd.go")

	// Define all CDP event methods.
	eg.PackageHeader()
	eg.EventType(protocol.Domains)
	eg.writeFile("event.go")

	for _, d := range protocol.Domains {
		for _, t := range d.Types {
			nam := t.Name(d)
			if isNonPointer(tg.pkg, d, t) {
				nonPtrMap[nam] = true
				nonPtrMap[tg.pkg+"."+nam] = true
			}
		}
	}

	for _, d := range protocol.Domains {
		g.PackageHeader()
		g.Domain(d)

		if len(d.Types) > 0 {
			tg.PackageHeader()
			for _, t := range d.Types {
				tg.DomainType(d, t)
			}
		}

		if len(d.Commands) > 0 {
			cg.PackageHeader()
			for _, c := range d.Commands {
				cg.DomainCmd(d, c)
			}
		}

		if len(d.Events) > 0 {
			eg.PackageHeader()
			for _, e := range d.Events {
				eg.DomainEvent(d, e)
			}
		}

		// Write dom definitions into separate files.
		// domfile := fmt.Sprintf("%s.go", strings.ToLower(d.Domain))
		// tg.writeFile(domfile)
		// cg.writeFile(domfile)
		// eg.writeFile(domfile)
		// g.writeFile(domfile)
	}
	tg.writeFile(tg.pkg + ".go")
	cg.writeFile(cg.pkg + ".go")
	eg.writeFile(eg.pkg + ".go")
	g.writeFile(g.pkg + ".go")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	goimports := exec.CommandContext(ctx, "goimports", "-w", g.path(), tg.path(), cg.path(), eg.path())
	out, err := goimports.CombinedOutput()
	if err != nil {
		log.Printf("%s", out)
		log.Println(err)
		os.Exit(1)
	}

	goinstall := exec.CommandContext(ctx, "go", "install", path.Join(g.path(), "..."))
	out, err = goinstall.CombinedOutput()
	if err != nil {
		log.Printf("%s", out)
		log.Println(err)
		os.Exit(1)
	}
}

// Generator holds the state of the analysis. Primarily used to buffer
// the output for format.Source.
type Generator struct {
	dir        string
	pkg        string
	imports    []string
	buf        bytes.Buffer // Accumulated output.
	hasContent bool
	hasHeader  bool
}

func (g *Generator) path() string {
	return path.Join(os.Getenv("GOPATH"), "src", g.dir)
}

// Printf prints to the Generator buffer.
func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

func (g *Generator) writeFile(f string) {
	fp := path.Join(g.path(), f)
	if !g.hasContent {
		log.Printf("No content, skipping %s...", fp)
		g.clear()
		return
	}
	if g.buf.Len() == 0 {
		log.Printf("Empty buffer, skipping %s...", fp)
		return
	}
	log.Printf("Writing %s...", fp)
	if err := ioutil.WriteFile(fp, g.format(), 0644); err != nil {
		panic(err)
	}
	g.clear()
}

func (g *Generator) clear() {
	g.hasContent = false
	g.hasHeader = false
	g.buf.Truncate(0)
}

// format returns the gofmt-ed contents of the Generator's buffer.
func (g *Generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return g.buf.Bytes()
	}
	return src
}

// CdpClient creates the cdp.Client type.
func (g *Generator) CdpClient(domains []proto.Domain) {
	g.hasContent = true
	var fields, newFields Generator
	for _, d := range domains {
		fields.Printf("\t%s %s\n", d.Name(), d.Type())
		newFields.Printf("\t\t%s: &%sDomain{conn: conn},\n", d.Name(), strings.ToLower(d.Name()))
	}
	g.Printf(`
// Client represents a Chrome Debugging Protocol client that can be used to
// invoke methods or listen to events in every CDP domain. The Client consumes
// a rpcc connection, used to invoke the methods.
type Client struct {
	%s
}

// NewClient returns a new Client.
func NewClient(conn *rpcc.Conn) *Client {
	return &Client{
		%s
	}
}
`, fields.buf.Bytes(), newFields.buf.Bytes())
}

// PackageHeader writes the header for a package.
func (g *Generator) PackageHeader() {
	if g.hasHeader {
		return
	}
	g.hasHeader = true
	g.Printf(`// Code generated by cdpgen; DO NOT EDIT!

package %s

import (
	"context"
	"encoding/json"
	"fmt"

	%s
)
`, g.pkg, quotedImports(g.imports))
}

// Domain defines the entire domain in the cdp package.
func (g *Generator) Domain(d proto.Domain) {
	g.hasContent = true

	var interfaceDef, domDef Generator
	for _, c := range d.Commands {
		request := ""
		reply := "error"
		if len(c.Parameters) > 0 {
			request = ", *cdpcmd." + c.ArgsName(d)
		}
		if len(c.Returns) > 0 {
			reply = fmt.Sprintf("(*cdpcmd.%s, error)", c.ReplyName(d))
		}
		interfaceDef.Printf("\n\t// Command %s\n\t//\n\t// %s\n\t%s(context.Context%s) %s\n", c.Name(), c.Desc(true), c.Name(), request, reply)

		// Implement command on %sDomain.
		invokeArgs := "nil"
		invokeReply := "nil"
		if len(c.Parameters) > 0 {
			request = ", args *cdpcmd." + c.ArgsName(d)
			invokeArgs = "args"
		}
		if len(c.Returns) > 0 {
			reply = fmt.Sprintf("(reply *cdpcmd.%s, err error)", c.ReplyName(d))
		} else {
			reply = "(err error)"
		}
		domDef.Printf("\nfunc (d *%sDomain) %s(ctx context.Context%s) %s {\n", strings.ToLower(d.Name()), c.Name(), request, reply)
		if len(c.Returns) > 0 {
			domDef.Printf("\treply = new(cdpcmd.%s)\n", c.ReplyName(d))
			invokeReply = "reply"
		}
		domDef.Printf("\terr = rpcc.Invoke(ctx, cdpcmd.%s.String(), %s, %s, d.conn)\n", c.CmdName(d, true), invokeArgs, invokeReply)
		domDef.Printf("\treturn\n")
		domDef.Printf("}\n")

		if len(c.Parameters) > 0 {
			domDef.Printf(`
// New%[1]s initializes the arguments for %[4]s.
func New%[1]s(%[2]s) *cdpcmd.%[1]s {
	args := new(cdpcmd.%[1]s)
	%[3]s
	return args
}
`, c.ArgsName(d), c.ArgsSignature(d), c.ArgsAssign("args", d), c.Name())
		}
	}
	for _, e := range d.Events {
		eventClient := fmt.Sprintf("%sClient", e.EventName(d))
		eventClientImpl := strings.ToLower(d.Domain) + "" + e.Name() + "Client"
		interfaceDef.Printf("\n\t// Event %s\n\t//\n\t// %s\n\t%s(context.Context) (cdpevent.%s, error)\n", e.Name(), e.Desc(true), e.Name(), eventClient)

		// Implement event on %sDomain.
		domDef.Printf(`
func (d *%sDomain) %s(ctx context.Context) (cdpevent.%s, error) {
	s, err := rpcc.NewStream(ctx, cdpevent.%s.String(), d.conn)
	if err != nil {
		return nil, err
	}
	return &%s{Stream: s}, nil
}
`, strings.ToLower(d.Name()), e.Name(), eventClient, e.EventName(d), eventClientImpl)

		domDef.Printf(`
// %[4]s implements %[1]s.
type %[4]s struct { rpcc.Stream }

func (c *%[4]s) Recv() (*cdpevent.%[3]s, error) {
	event := new(cdpevent.%[3]s)
	if err := c.RecvMsg(event); err != nil {
		return nil, errors.New("cdp: %[1]s Recv: " + err.Error())
	}
	return event, nil
}
`, eventClient, "", e.ReplyName(d), eventClientImpl)

	}

	g.Printf(`
// The %[1]s domain. %[2]s
type %[1]s interface{%[3]s}

// %[4]sDomain implements the %[1]s domain.
type %[4]sDomain struct{ conn *rpcc.Conn }

%[5]s
`, d.Name(), d.Desc(), interfaceDef.buf.String(), strings.ToLower(d.Name()), domDef.buf.String())
}

// DomainType creates the type definition.
func (g *Generator) DomainType(d proto.Domain, t proto.AnyType) {
	g.hasContent = true
	g.Printf(`
// %[1]s %[2]s
type %[1]s `, t.Name(d), t.Desc())
	switch t.GoType(g.pkg, d) {
	case "struct":
		g.domainTypeStruct(d, t)
	case "enum":
		g.domainTypeEnum(d, t)
	case "RawMessage":
		g.domainTypeRawMessage(d, t)
	default:
		g.Printf(t.GoType(g.pkg, d))
	}
	g.Printf("\n\n")
}

func (g *Generator) printStructProperties(d proto.Domain, name string, props []proto.AnyType, ptrOptional, renameOptional bool) {
	for _, prop := range props {
		jsontag := prop.NameName
		ptype := prop.GoType(g.pkg, d)
		// Make all optional properties into pointers, unless they are slices.
		if prop.Optional {
			isNonPtr := nonPtrMap[ptype]
			if ptrOptional && !isNonPtr && !isNonPointer(g.pkg, d, prop) {
				ptype = "*" + ptype
			}
			jsontag += ",omitempty"
		}

		// Avoid recursive type definitions.
		if ptype == name {
			ptype = "*" + ptype
		}

		exportedName := prop.ExportedName(d)
		if renameOptional && prop.Optional {
			exportedName = OptionalPropPrefix + exportedName
		}

		g.Printf("\t%s %s `json:\"%s\"` // %s\n", exportedName, ptype, jsontag, prop.Desc())
	}
}

func (g *Generator) domainTypeStruct(d proto.Domain, t proto.AnyType) {
	g.Printf("struct{\n")
	g.printStructProperties(d, t.Name(d), t.Properties, true, false)
	g.Printf("}")
}

func (g *Generator) domainTypeRawMessage(d proto.Domain, t proto.AnyType) {
	g.Printf(`[]byte

// MarshalJSON returns m as the JSON encoding of m.
func (m %[1]s) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *%[1]s) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("%[2]s.%[1]s: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[0:0], data...)
	return nil
}

var _ json.Marshaler = (*%[1]s)(nil)
var _ json.Unmarshaler = (*%[1]s)(nil)
`, t.Name(d), g.pkg)
}

func (g *Generator) domainTypeEnum(d proto.Domain, t proto.AnyType) {
	if t.Type != "string" {
		log.Panicf("unknown enum type: %s", t.Type)
	}
	if realEnum {
		name := strings.Title(t.Name(d))
		g.Printf(`int

// %s as enums.
const (
	%sNotSet %s = iota`, name, name, name)
		for _, e := range t.Enum {
			g.Printf("\n\t%s%s", name, e.Name())
		}
		g.Printf(`
)

func (e %s) Valid() bool {
	return e >= 1 && e <= %d
}

func (e %s) String() string {
	switch e {
	case 0:
		return "%sNotSet"`, name, len(t.Enum), name, name)
		for i, e := range t.Enum {
			g.Printf(`
	case %d:
		return "%s"`, i+1, e)
		}
		g.Printf(`
	}
	return fmt.Sprintf("%s(%%d)", e)
}

func (e %s) MarshalJSON() ([]byte, error) {
	if e == 0 {
		return []byte("null"), nil
	}
	return json.Marshal(e.String())
}

func (e *%s) UnmarshalJSON(data []byte) error {
	if data == nil {
		*e = 0
		return nil
	}
	switch string(data) {`, name, name, name)
		for i, e := range t.Enum {
			g.Printf(`
	case "\"%s\"":
		*e = %d`, e, i+1)
		}
		g.Printf(`
	default:
		return fmt.Errorf("bad %s: %%s", data)
	}
	return nil
}`, name)
	} else {
		g.Printf(`string

func (e %[1]s) String() string {
	return string(e)
}

// %[1]s types.
const (
`, t.Name(d))
		for _, e := range t.Enum {
			g.Printf("\t%s %s = %q\n", t.Name(d)+e.Name(), t.Name(d), e)
		}
		g.Printf(")")
	}
}

// CmdType generates the type for CDP methods names.
func (g *Generator) CmdType(doms []proto.Domain) {
	g.hasContent = true
	g.Printf(`
// CmdType is the type for CDP methods names.
type CmdType string

func (c CmdType) String() string {
	return string(c)
}

// Cmd methods.
const (`)
	for _, d := range doms {
		for _, c := range d.Commands {
			g.Printf("\n\t%s CmdType = %q", c.CmdName(d, true), d.Domain+"."+c.NameName)
		}
	}
	g.Printf("\n)\n")
}

// DomainCmd defines the command args and reply.
func (g *Generator) DomainCmd(d proto.Domain, c proto.Command) {
	if len(c.Parameters) > 0 {
		g.hasContent = true
		g.domainCmdArgs(d, c)
	}
	if len(c.Returns) > 0 {
		g.hasContent = true
		g.domainCmdReply(d, c)
	}
}

func (g *Generator) domainCmdArgs(d proto.Domain, c proto.Command) {
	g.Printf(`
// %[1]s contains the arguments for %[2]s.
type %[1]s struct {`, c.ArgsName(d), c.CmdName(d, false))
	g.printStructProperties(d, c.ArgsName(d), c.Parameters, true, true)
	g.Printf("}\n\n")

	for _, arg := range c.Parameters {
		typ := arg.GoType(g.pkg, d)
		isNonPtr := nonPtrMap[typ]
		if !arg.Optional || isNonPtr || isNonPointer(g.pkg, d, arg) {
			continue
		}
		name := arg.Name(d)
		if name == "range" || name == "type" {
			name = name[0 : len(name)-1]
		}
		g.Printf(`
// Set%[1]s sets the %[1]s optional argument. %[6]s
func (a *%[2]s) Set%[1]s(%[3]s %[4]s) *%[2]s {
	a.%[5]s%[1]s = &%[3]s
	return a
}
`, arg.ExportedName(d), c.ArgsName(d), name, arg.GoType("cdp", d), OptionalPropPrefix, arg.Desc())
	}
}

func (g *Generator) domainCmdReply(d proto.Domain, c proto.Command) {
	g.Printf(`
// %[1]s contains the return values for %[2]s.
type %[1]s struct {`, c.ReplyName(d), c.CmdName(d, false))
	g.printStructProperties(d, c.ReplyName(d), c.Returns, true, false)
	g.Printf("}\n\n")
}

// EventType generates the type for CDP event names.
func (g *Generator) EventType(doms []proto.Domain) {
	g.hasContent = true
	g.Printf(`
// EventType is the type for CDP event names.
type EventType string

func (e EventType) String() string {
	return string(e)
}

// Event methods.
const (`)
	for _, d := range doms {
		for _, e := range d.Events {
			g.Printf("\n\t%s EventType = %q", e.EventName(d), d.Domain+"."+e.NameName)
		}
	}
	g.Printf("\n)\n")
}

// DomainEvent defines the event client and reply.
func (g *Generator) DomainEvent(d proto.Domain, e proto.Event) {
	g.hasContent = true
	g.domainEventClient(d, e)
	g.domainEventReply(d, e)
}

func (g *Generator) domainEventClient(d proto.Domain, e proto.Event) {
	eventClient := fmt.Sprintf("%sClient", e.EventName(d))
	g.Printf(`
// %[1]s receives %[2]s events.
type %[1]s interface {
	Recv() (*%[3]s, error)
	rpcc.Stream
}
`, eventClient, e.Name(), e.ReplyName(d))
}

func (g *Generator) domainEventReply(d proto.Domain, e proto.Event) {
	g.Printf(`
// %[1]s %[2]s
type %[1]s struct {`, e.ReplyName(d), e.Desc(false))
	g.printStructProperties(d, e.ReplyName(d), e.Parameters, true, false)
	g.Printf("}\n")
}

func quotedImports(imports []string) string {
	if len(imports) == 0 {
		return ""
	}

	return "\"" + strings.Join(imports, "\"\n\"") + "\""
}

func isNonPointer(pkg string, d proto.Domain, t proto.AnyType) bool {
	typ := t.GoType(pkg, d)
	switch {
	case t.IsEnum():
	case strings.HasPrefix(typ, "[]"):
	case strings.HasPrefix(typ, "map["):
	case typ == "json.RawMessage":
	case typ == "RawMessage":
	case typ == "interface{}":
	default:
		return false
	}
	return true
}