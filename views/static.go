package views

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"path"
	"path/filepath"
)

// bindata_read reads the given file from disk. It returns an error on failure.
func bindata_read(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// static_gitkeep reads file data from disk. It returns an error on failure.
func static_gitkeep() (*asset, error) {
	path := "/Users/tattsun/work/src/github.com/tattsun/coopy/static/.gitkeep"
	name := "static/.gitkeep"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_assets_css_gitkeep reads file data from disk. It returns an error on failure.
func static_assets_css_gitkeep() (*asset, error) {
	path := "/Users/tattsun/work/src/github.com/tattsun/coopy/static/assets/css/.gitkeep"
	name := "static/assets/css/.gitkeep"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_assets_img_gitkeep reads file data from disk. It returns an error on failure.
func static_assets_img_gitkeep() (*asset, error) {
	path := "/Users/tattsun/work/src/github.com/tattsun/coopy/static/assets/img/.gitkeep"
	name := "static/assets/img/.gitkeep"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_assets_js_gitkeep reads file data from disk. It returns an error on failure.
func static_assets_js_gitkeep() (*asset, error) {
	path := "/Users/tattsun/work/src/github.com/tattsun/coopy/static/assets/js/.gitkeep"
	name := "static/assets/js/.gitkeep"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_assets_js_login_view_js reads file data from disk. It returns an error on failure.
func static_assets_js_login_view_js() (*asset, error) {
	path := "/Users/tattsun/work/src/github.com/tattsun/coopy/static/assets/js/login_view.js"
	name := "static/assets/js/login_view.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_html_gitkeep reads file data from disk. It returns an error on failure.
func static_html_gitkeep() (*asset, error) {
	path := "/Users/tattsun/work/src/github.com/tattsun/coopy/static/html/.gitkeep"
	name := "static/html/.gitkeep"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_login_haml reads file data from disk. It returns an error on failure.
func static_login_haml() (*asset, error) {
	path := "/Users/tattsun/work/src/github.com/tattsun/coopy/static/login.haml"
	name := "static/login.haml"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_newuser_haml reads file data from disk. It returns an error on failure.
func static_newuser_haml() (*asset, error) {
	path := "/Users/tattsun/work/src/github.com/tattsun/coopy/static/newuser.haml"
	name := "static/newuser.haml"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_test_haml reads file data from disk. It returns an error on failure.
func static_test_haml() (*asset, error) {
	path := "/Users/tattsun/work/src/github.com/tattsun/coopy/static/test.haml"
	name := "static/test.haml"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if (err != nil) {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
var _bindata = map[string]func() (*asset, error){
	"static/.gitkeep": static_gitkeep,
	"static/assets/css/.gitkeep": static_assets_css_gitkeep,
	"static/assets/img/.gitkeep": static_assets_img_gitkeep,
	"static/assets/js/.gitkeep": static_assets_js_gitkeep,
	"static/assets/js/login_view.js": static_assets_js_login_view_js,
	"static/html/.gitkeep": static_html_gitkeep,
	"static/login.haml": static_login_haml,
	"static/newuser.haml": static_newuser_haml,
	"static/test.haml": static_test_haml,
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
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"static": &_bintree_t{nil, map[string]*_bintree_t{
		".gitkeep": &_bintree_t{static_gitkeep, map[string]*_bintree_t{
		}},
		"assets": &_bintree_t{nil, map[string]*_bintree_t{
			"css": &_bintree_t{nil, map[string]*_bintree_t{
				".gitkeep": &_bintree_t{static_assets_css_gitkeep, map[string]*_bintree_t{
				}},
			}},
			"img": &_bintree_t{nil, map[string]*_bintree_t{
				".gitkeep": &_bintree_t{static_assets_img_gitkeep, map[string]*_bintree_t{
				}},
			}},
			"js": &_bintree_t{nil, map[string]*_bintree_t{
				".gitkeep": &_bintree_t{static_assets_js_gitkeep, map[string]*_bintree_t{
				}},
				"login_view.js": &_bintree_t{static_assets_js_login_view_js, map[string]*_bintree_t{
				}},
			}},
		}},
		"html": &_bintree_t{nil, map[string]*_bintree_t{
			".gitkeep": &_bintree_t{static_html_gitkeep, map[string]*_bintree_t{
			}},
		}},
		"login.haml": &_bintree_t{static_login_haml, map[string]*_bintree_t{
		}},
		"newuser.haml": &_bintree_t{static_newuser_haml, map[string]*_bintree_t{
		}},
		"test.haml": &_bintree_t{static_test_haml, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

