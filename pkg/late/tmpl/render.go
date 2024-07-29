package tmpl

import (
	"io"

	"github.com/pkg/errors"

	"github.com/Masterminds/sprig/v3"
	"text/template"
    path "path/filepath"

	log "github.com/kubernetes/klog"
)

func DefaultRenderOptions() *RenderOptions {
	return &RenderOptions{
		DelimiterLeft:  "{{",
		DelimiterRight: "}}",
	}
}

type RenderOptions struct {
	DelimiterLeft  string
	DelimiterRight string
}

func Render(data interface{}, tmpl string, out io.Writer, options *RenderOptions) error {
	t := template.New(path.Base(tmpl))
	log.V(5).Infoln("Use Delimiter options:", options.DelimiterLeft, options.DelimiterRight)
	t = t.Delims(options.DelimiterLeft, options.DelimiterRight)
	t = t.Funcs(sprig.FuncMap())
	//
	var err error
	t, err = t.ParseFiles(tmpl)
	if err != nil {
		return errors.Wrap(err, "Could not parse template")
	}

	log.V(5).Infof("Associated templates: %q", t.Templates())

	err = t.Execute(out, data)
	if err != nil {
		return errors.Wrap(err, "Could not execute template")
	}
	return nil
}
