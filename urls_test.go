package giturls

import (
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var tests []*Test

type Test struct {
	in      string
	wantURL *url.URL
	wantStr string // expected result of reserializing the URL; empty means same as "in".
}

func NewTest(in, transport, user, host, path, str, rawquery string) *Test {
	var userinfo *url.Userinfo

	if user != "" {
		if strings.Contains(user, ":") {
			username := strings.Split(user, ":")[0]
			password := strings.Split(user, ":")[1]
			userinfo = url.UserPassword(username, password)
		} else {
			userinfo = url.User(user)
		}
	}
	if str == "" {
		str = in
	}

	return &Test{
		in: in,
		wantURL: &url.URL{
			Scheme:   transport,
			Host:     host,
			Path:     path,
			User:     userinfo,
			RawQuery: rawquery,
		},
		wantStr: str,
	}
}

func init() {
	// https://www.kernel.org/pub/software/scm/git/docs/git-clone.html
	tests = []*Test{
		NewTest(
			"user@host.xz:path/to/repo.git/",
			"ssh", "user", "host.xz", "path/to/repo.git/",
			"ssh://user@host.xz/path/to/repo.git/", "",
		),
		NewTest(
			"host.xz:path/to/repo.git/",
			"ssh", "", "host.xz", "path/to/repo.git/",
			"ssh://host.xz/path/to/repo.git/", "",
		),
		NewTest(
			"host.xz:/path/to/repo.git/",
			"ssh", "", "host.xz", "/path/to/repo.git/",
			"ssh://host.xz/path/to/repo.git/", "",
		),
		NewTest(
			"host.xz:path/to/repo-with_specials.git/",
			"ssh", "", "host.xz", "path/to/repo-with_specials.git/",
			"ssh://host.xz/path/to/repo-with_specials.git/", "",
		),
		NewTest(
			"git://host.xz/path/to/repo.git/",
			"git", "", "host.xz", "/path/to/repo.git/",
			"", "",
		),
		NewTest(
			"git://host.xz:1234/path/to/repo.git/",
			"git", "", "host.xz:1234", "/path/to/repo.git/",
			"", "",
		),
		NewTest(
			"http://host.xz/path/to/repo.git/",
			"http", "", "host.xz", "/path/to/repo.git/",
			"", "",
		),
		NewTest(
			"http://host.xz:1234/path/to/repo.git/",
			"http", "", "host.xz:1234", "/path/to/repo.git/",
			"", "",
		),
		NewTest(
			"https://host.xz/path/to/repo.git/",
			"https", "", "host.xz", "/path/to/repo.git/",
			"", "",
		),
		NewTest(
			"https://host.xz:1234/path/to/repo.git/",
			"https", "", "host.xz:1234", "/path/to/repo.git/",
			"", "",
		),
		NewTest(
			"ftp://host.xz/path/to/repo.git/",
			"ftp", "", "host.xz", "/path/to/repo.git/",
			"", "",
		),
		NewTest(
			"ftp://host.xz:1234/path/to/repo.git/",
			"ftp", "", "host.xz:1234", "/path/to/repo.git/",
			"", "",
		),
		NewTest(
			"ftps://host.xz/path/to/repo.git/",
			"ftps", "", "host.xz", "/path/to/repo.git/",
			"", "",
		),
		NewTest(
			"ftps://host.xz:1234/path/to/repo.git/",
			"ftps", "", "host.xz:1234", "/path/to/repo.git/",
			"", "",
		),
		NewTest(
			"rsync://host.xz/path/to/repo.git/",
			"rsync", "", "host.xz", "/path/to/repo.git/",
			"", "",
		),
		NewTest(
			"ssh://user@host.xz:1234/path/to/repo.git/",
			"ssh", "user", "host.xz:1234", "/path/to/repo.git/",
			"", "",
		),
		NewTest(
			"ssh://host.xz:1234/path/to/repo.git/",
			"ssh", "", "host.xz:1234", "/path/to/repo.git/",
			"", "",
		),
		NewTest(
			"ssh://host.xz/path/to/repo.git/",
			"ssh", "", "host.xz", "/path/to/repo.git/",
			"", "",
		),
		NewTest(
			"git+ssh://host.xz/path/to/repo.git/",
			"git+ssh", "", "host.xz", "/path/to/repo.git/",
			"", "",
		),
		NewTest(
			"/path/to/repo.git/",
			"file", "", "", "/path/to/repo.git/",
			"file:///path/to/repo.git/", "",
		),
		NewTest(
			"file:///path/to/repo.git/",
			"file", "", "", "/path/to/repo.git/",
			"", "",
		),
		// Tests with query strings
		NewTest(
			"https://host.xz/organization/repo.git?ref=",
			"https", "", "host.xz", "/organization/repo.git",
			"", "ref=",
		),
		NewTest(
			"https://host.xz/organization/repo.git?ref=test",
			"https", "", "host.xz", "/organization/repo.git",
			"", "ref=test",
		),
		NewTest(
			"https://host.xz/organization/repo.git?ref=feature/test",
			"https", "", "host.xz", "/organization/repo.git",
			"", "ref=feature/test",
		),
		NewTest(
			"git@host.xz:organization/repo.git?ref=test",
			"ssh", "git", "host.xz", "organization/repo.git",
			"ssh://git@host.xz/organization/repo.git?ref=test", "ref=test",
		),
		NewTest(
			"git@host.xz:organization/repo.git?ref=feature/test",
			"ssh", "git", "host.xz", "organization/repo.git",
			"ssh://git@host.xz/organization/repo.git?ref=feature/test", "ref=feature/test",
		),
		// Tests with user+password and some with query strings
		NewTest(
			"https://user:password@host.xz/organization/repo.git/",
			"https", "user:password", "host.xz", "/organization/repo.git/",
			"", "",
		),
		NewTest(
			"https://user:password@host.xz/organization/repo.git?ref=test",
			"https", "user:password", "host.xz", "/organization/repo.git",
			"", "ref=test",
		),
		NewTest(
			"https://user:password@host.xz/organization/repo.git?ref=feature/test",
			"https", "user:password", "host.xz", "/organization/repo.git",
			"", "ref=feature/test",
		),
		NewTest(
			"user-1234@host.xz:path/to/repo.git/",
			"ssh", "user-1234", "host.xz", "path/to/repo.git/",
			"ssh://user-1234@host.xz/path/to/repo.git/", "",
		),
	}
}

func TestParse(t *testing.T) {
	for _, tt := range tests {
		got, err := Parse(tt.in)
		if err != nil {
			t.Errorf("Parse(%q) = unexpected err %q, want %q", tt.in, err, tt.wantURL)
			continue
		}
		if !reflect.DeepEqual(got, tt.wantURL) {
			t.Errorf("Parse(%q) = %q, want %q", tt.in, got, tt.wantURL)
		}
		str := got.String()
		if str != tt.wantStr {
			t.Errorf("Parse(%q).String() = %q, want %q", tt.in, str, tt.wantStr)
		}
	}
}

func runTimingTest(t *testing.T, url string, shouldError bool) {
	begin := time.Now()

	_, err := ParseScp(url)
	if shouldError {
		assert.Errorf(t, err, "len of %d should trigger error", len(url))
	} else {
		if t == nil {
			panic("t is nil")
		}
		assert.Nilf(t, err, "unexpected error: %v", err)
	}
	elapsed := time.Since(begin)
	t.Logf("url len is %d, function took %+v", len(url), elapsed)
}

// TestRegex tests to see if we have an excessively long URL
func TestRegex(t *testing.T) {

	// First case is 7909 bytes which4yy should still be fast
	long_url := `https://=` + strings.Repeat(`/`, 7900)
	runTimingTest(t, long_url, false)

	// Second case is 20,000 bytes which should be too slow
	long_url = `https://=` + strings.Repeat(`/`, 190000000)
	runTimingTest(t, long_url, true)

	goodURL := `https://stackoverflow.com/q/417142/31319`
	runTimingTest(t, goodURL, false)

	goodURL = `https://kinesis-ergo.com/wp-content/uploads/Advantage360-SmartSet-KB360-Users-Manual-v10-12-22.pdf`
	runTimingTest(t, goodURL, false)
}
