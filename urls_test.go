package giturls

import (
	"net/url"
	"reflect"
	"testing"
)

var tests []*Test

type Test struct {
	rawurl string
	result *Result
}

type Result struct {
	Transport string
	User string
	Host string
	Path string
}

func NewResult(transport, user, host, path string) *Result {
	return &Result{
		Transport: transport,
		User: user,
		Host: host,
		Path: path,
	}
}

func ResultFromURL(u *url.URL) *Result {
	var user string
	if u.User == nil {
		user = ""
	} else {
		user = u.User.String()
	}
	return NewResult(
		u.Scheme,
		user,
		u.Host,
		u.Path,
	)
}

func init (){
	// https://www.kernel.org/pub/software/scm/git/docs/git-clone.html
	tests = []*Test{
		&Test{
			"user@host.xz:path/to/repo.git/",
			NewResult("ssh", "user", "host.xz", "/path/to/repo.git/"),
		},
		&Test{
			"host.xz:path/to/repo.git/",
			NewResult("ssh", "", "host.xz", "/path/to/repo.git/"),
		},
		&Test{
			"host.xz:path/to/repo-with_specials.git/",
			NewResult("ssh", "", "host.xz", "/path/to/repo-with_specials.git/"),
		},
		&Test{
			"git://host.xz/path/to/repo.git/",
			NewResult("git", "", "host.xz", "/path/to/repo.git/"),
		},
		&Test{
			"git://host.xz:1234/path/to/repo.git/",
			NewResult("git", "", "host.xz:1234", "/path/to/repo.git/"),
		},
		&Test{
			"http://host.xz/path/to/repo.git/",
			NewResult("http", "", "host.xz", "/path/to/repo.git/"),
		},
		&Test{
			"http://host.xz:1234/path/to/repo.git/",
			NewResult("http", "", "host.xz:1234", "/path/to/repo.git/"),
		},
		&Test{
			"https://host.xz/path/to/repo.git/",
			NewResult("https", "", "host.xz", "/path/to/repo.git/"),
		},
		&Test{
			"https://host.xz:1234/path/to/repo.git/",
			NewResult("https", "", "host.xz:1234", "/path/to/repo.git/"),
		},
		&Test{
			"ftp://host.xz/path/to/repo.git/",
			NewResult("ftp", "", "host.xz", "/path/to/repo.git/"),
		},
		&Test{
			"ftp://host.xz:1234/path/to/repo.git/",
			NewResult("ftp", "", "host.xz:1234", "/path/to/repo.git/"),
		},
		&Test{
			"ftps://host.xz/path/to/repo.git/",
			NewResult("ftps", "", "host.xz", "/path/to/repo.git/"),
		},
		&Test{
			"ftps://host.xz:1234/path/to/repo.git/",
			NewResult("ftps", "", "host.xz:1234", "/path/to/repo.git/"),
		},
		&Test{
			"rsync://host.xz/path/to/repo.git/",
			NewResult("rsync", "", "host.xz", "/path/to/repo.git/"),
		},
		&Test{
			"ssh://user@host.xz:1234/path/to/repo.git/",
			NewResult("ssh", "user", "host.xz:1234", "/path/to/repo.git/"),
		},
		&Test{
			"ssh://host.xz:1234/path/to/repo.git/",
			NewResult("ssh", "", "host.xz:1234", "/path/to/repo.git/"),
		},
		&Test{
			"ssh://host.xz/path/to/repo.git/",
			NewResult("ssh", "", "host.xz", "/path/to/repo.git/"),
		},
		&Test{
			"/path/to/repo.git/",
			NewResult("file", "", "", "/path/to/repo.git/"),
		},
		&Test{
			"file:///path/to/repo.git/",
			NewResult("file", "", "", "/path/to/repo.git/"),
		},
	}
}

func TestParse(t *testing.T) {
	for _, tt := range tests {
		url, err := Parse(tt.rawurl)
		if url == nil {
			t.Errorf("Parse(%q) -> nil", tt.rawurl)
			continue
		}
		if err != nil {
			t.Errorf("Parse(%q) -> unexpected err %q", tt.rawurl, err)
			continue
		}

		r := ResultFromURL(url)
		if !reflect.DeepEqual(r, tt.result) {
			t.Errorf("Parse(%q) -> %q, want %q", tt.rawurl, r, tt.result)
		}
	}
}
