package connparse_test

import (
	"testing"

	"github.com/wavetermdev/waveterm/pkg/remote/connparse"
)

func TestParseURI_WSHWithScheme(t *testing.T) {
	t.Parallel()

	// Test with localhost
	cstr := "wsh://user@localhost:8080/path/to/file"
	c, err := connparse.ParseURI(cstr)
	if err != nil {
		t.Fatalf("failed to parse URI: %v", err)
	}
	expected := "/path/to/file"
	if c.Path != expected {
		t.Fatalf("expected path to be %q, got %q", expected, c.Path)
	}
	expected = "user@localhost:8080"
	if c.Host != expected {
		t.Fatalf("expected host to be %q, got %q", expected, c.Host)
	}
	expected = "user@localhost:8080/path/to/file"
	pathWithHost := c.GetPathWithHost()
	if pathWithHost != expected {
		t.Fatalf("expected path with host to be %q, got %q", expected, pathWithHost)
	}
	expected = "wsh"
	if c.Scheme != expected {
		t.Fatalf("expected scheme to be %q, got %q", expected, c.Scheme)
	}
	if len(c.GetSchemeParts()) != 1 {
		t.Fatalf("expected scheme parts to be 1, got %d", len(c.GetSchemeParts()))
	}

	// Test with an IP address
	cstr = "wsh://user@192.168.0.1:22/path/to/file"
	c, err = connparse.ParseURI(cstr)
	if err != nil {
		t.Fatalf("failed to parse URI: %v", err)
	}
	expected = "/path/to/file"
	if c.Path != expected {
		t.Fatalf("expected path to be %q, got %q", expected, c.Path)
	}
	expected = "user@192.168.0.1:22"
	if c.Host != expected {
		t.Fatalf("expected host to be %q, got %q", expected, c.Host)
	}
	expected = "user@192.168.0.1:22/path/to/file"
	pathWithHost = c.GetPathWithHost()
	if pathWithHost != expected {
		t.Fatalf("expected path with host to be %q, got %q", expected, pathWithHost)
	}
	expected = "wsh"
	if c.GetType() != expected {
		t.Fatalf("expected conn type to be %q, got %q", expected, c.Scheme)
	}
	if len(c.GetSchemeParts()) != 1 {
		t.Fatalf("expected scheme parts to be 1, got %d", len(c.GetSchemeParts()))
	}
	got := c.GetFullURI()
	if got != cstr {
		t.Fatalf("expected full URI to be %q, got %q", cstr, got)
	}
}

func TestParseURI_WSHRemoteShorthand(t *testing.T) {
	t.Parallel()

	// Test with a simple remote path
	cstr := "//conn/path/to/file"
	c, err := connparse.ParseURI(cstr)
	if err != nil {
		t.Fatalf("failed to parse URI: %v", err)
	}
	expected := "/path/to/file"
	if c.Path != expected {
		t.Fatalf("expected path to be %q, got %q", expected, c.Path)
	}
	if c.Host != "conn" {
		t.Fatalf("expected host to be empty, got %q", c.Host)
	}
	expected = "wsh"
	if c.Scheme != expected {
		t.Fatalf("expected scheme to be %q, got %q", expected, c.Scheme)
	}
	expected = "wsh://conn/path/to/file"
	if c.GetFullURI() != expected {
		t.Fatalf("expected full URI to be %q, got %q", expected, c.GetFullURI())
	}

	// Test with a complex remote path
	cstr = "//user@localhost:8080/path/to/file"
	c, err = connparse.ParseURI(cstr)
	if err != nil {
		t.Fatalf("failed to parse URI: %v", err)
	}
	expected = "/path/to/file"
	if c.Path != expected {
		t.Fatalf("expected path to be %q, got %q", expected, c.Path)
	}
	expected = "user@localhost:8080"
	if c.Host != expected {
		t.Fatalf("expected host to be %q, got %q", expected, c.Host)
	}
	expected = "wsh"
	if c.Scheme != expected {
		t.Fatalf("expected scheme to be %q, got %q", expected, c.Scheme)
	}
	expected = "wsh://user@localhost:8080/path/to/file"
	if c.GetFullURI() != expected {
		t.Fatalf("expected full URI to be %q, got %q", expected, c.GetFullURI())
	}

	// Test with an IP address
	cstr = "//user@192.168.0.1:8080/path/to/file"
	c, err = connparse.ParseURI(cstr)
	if err != nil {
		t.Fatalf("failed to parse URI: %v", err)
	}
	expected = "/path/to/file"
	if c.Path != expected {
		t.Fatalf("expected path to be %q, got %q", expected, c.Path)
	}
	expected = "user@192.168.0.1:8080"
	if c.Host != expected {
		t.Fatalf("expected host to be %q, got %q", expected, c.Host)
	}
	expected = "wsh"
	if c.Scheme != expected {
		t.Fatalf("expected scheme to be %q, got %q", expected, c.Scheme)
	}
	expected = "wsh://user@192.168.0.1:8080/path/to/file"
	if c.GetFullURI() != expected {
		t.Fatalf("expected full URI to be %q, got %q", expected, c.GetFullURI())
	}
}

func TestParseURI_WSHCurrentPathShorthand(t *testing.T) {
	t.Parallel()

	// Test with a relative path to home
	cstr := "~/path/to/file"
	c, err := connparse.ParseURI(cstr)
	if err != nil {
		t.Fatalf("failed to parse URI: %v", err)
	}
	expected := "~/path/to/file"
	if c.Path != expected {
		t.Fatalf("expected path to be %q, got %q", expected, c.Path)
	}
	expected = "current"
	if c.Host != expected {
		t.Fatalf("expected host to be %q, got %q", expected, c.Host)
	}
	expected = "wsh"
	if c.Scheme != expected {
		t.Fatalf("expected scheme to be %q, got %q", expected, c.Scheme)
	}
	expected = "wsh://current/~/path/to/file"
	if c.GetFullURI() != expected {
		t.Fatalf("expected full URI to be %q, got %q", expected, c.GetFullURI())
	}

	// Test with a absolute path
	cstr = "/path/to/file"
	c, err = connparse.ParseURI(cstr)
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
	expected = "/path/to/file"
	if c.Path != expected {
		t.Fatalf("expected path to be %q, got %q", expected, c.Path)
	}
	expected = "current"
	if c.Host != expected {
		t.Fatalf("expected host to be %q, got %q", expected, c.Host)
	}
	expected = "wsh"
	if c.Scheme != expected {
		t.Fatalf("expected scheme to be %q, got %q", expected, c.Scheme)
	}
	expected = "wsh://current/path/to/file"
	if c.GetFullURI() != expected {
		t.Fatalf("expected full URI to be %q, got %q", expected, c.GetFullURI())
	}
}

func TestParseURI_WSHLocalShorthand(t *testing.T) {
	t.Parallel()
	cstr := "/~/path/to/file"
	c, err := connparse.ParseURI(cstr)
	if err != nil {
		t.Fatalf("failed to parse URI: %v", err)
	}
	expected := "~/path/to/file"
	if c.Path != expected {
		t.Fatalf("expected path to be %q, got %q", expected, c.Path)
	}
	if c.Host != "local" {
		t.Fatalf("expected host to be empty, got %q", c.Host)
	}
	expected = "wsh"
	if c.Scheme != expected {
		t.Fatalf("expected scheme to be %q, got %q", expected, c.Scheme)
	}
}

func TestParseURI_BasicS3(t *testing.T) {
	t.Parallel()
	cstr := "profile:s3://bucket/path/to/file"
	c, err := connparse.ParseURI(cstr)
	if err != nil {
		t.Fatalf("failed to parse URI: %v", err)
	}
	expected := "/path/to/file"
	if c.Path != expected {
		t.Fatalf("expected path to be %q, got %q", expected, c.Path)
	}
	expected = "bucket"
	if c.Host != expected {
		t.Fatalf("expected host to be %q, got %q", expected, c.Host)
	}
	expected = "bucket/path/to/file"
	pathWithHost := c.GetPathWithHost()
	if pathWithHost != expected {
		t.Fatalf("expected path with host to be %q, got %q", expected, pathWithHost)
	}
	expected = "s3"
	if c.GetType() != expected {
		t.Fatalf("expected conn type to be %q, got %q", expected, c.GetType())
	}
	if len(c.GetSchemeParts()) != 2 {
		t.Fatalf("expected scheme parts to be 2, got %d", len(c.GetSchemeParts()))
	}
}
