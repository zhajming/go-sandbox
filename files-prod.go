package sandbox

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// bindata_read reads the given file from disk. It returns an error on failure.
func bindata_read(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

// static_css_app_css reads file data from disk. It returns an error on failure.
func static_css_app_css() ([]byte, error) {
	return bindata_read(
		"/Users/jpillora/Code/Go/src/github.com/jpillora/go-sandbox/static/css/app.css",
		"static/css/app.css",
	)
}

// static_gruntsource_json reads file data from disk. It returns an error on failure.
func static_gruntsource_json() ([]byte, error) {
	return bindata_read(
		"/Users/jpillora/Code/Go/src/github.com/jpillora/go-sandbox/static/Gruntsource.json",
		"static/Gruntsource.json",
	)
}

// static_index_html reads file data from disk. It returns an error on failure.
func static_index_html() ([]byte, error) {
	return bindata_read(
		"/Users/jpillora/Code/Go/src/github.com/jpillora/go-sandbox/static/index.html",
		"static/index.html",
	)
}

// static_js_app_js reads file data from disk. It returns an error on failure.
func static_js_app_js() ([]byte, error) {
	return bindata_read(
		"/Users/jpillora/Code/Go/src/github.com/jpillora/go-sandbox/static/js/app.js",
		"static/js/app.js",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	f, ok := _bindata[cannonicalName]
	if ok {
		return f()
	}
	dir, _ := os.Getwd()
	return ioutil.ReadFile(dir + name)
	// return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"static/css/app.css":      static_css_app_css,
	"static/Gruntsource.json": static_gruntsource_json,
	"static/index.html":       static_index_html,
	"static/js/app.js":        static_js_app_js,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"static": &_bintree_t{nil, map[string]*_bintree_t{
		"css": &_bintree_t{nil, map[string]*_bintree_t{
			"app.css": &_bintree_t{static_css_app_css, map[string]*_bintree_t{}},
		}},
		"Gruntsource.json": &_bintree_t{static_gruntsource_json, map[string]*_bintree_t{}},
		"index.html":       &_bintree_t{static_index_html, map[string]*_bintree_t{}},
		"js": &_bintree_t{nil, map[string]*_bintree_t{
			"app.js": &_bintree_t{static_js_app_js, map[string]*_bintree_t{}},
		}},
	}},
}}