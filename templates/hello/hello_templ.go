// Code generated by templ@v0.2.316 DO NOT EDIT.

package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import common "github.com/edwincarlflores/go-htmx/templates/common"

func Hello() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div><p>")
		if err != nil {
			return err
		}
		var_2 := `What's your name?`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><form class=\"flex flex-row space-x-3\" hx-post=\"/hello\" hx-swap=\"innerHTML\" hx-target=\"#content\"><input type=\"text\" name=\"name\" class=\"border border-black\"><button type=\"submit\">")
		if err != nil {
			return err
		}
		var_3 := `Send`
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button></form></div><div id=\"content\"></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func HelloName(name string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_4 := templ.GetChildren(ctx)
		if var_4 == nil {
			var_4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if name == "" {
			_, err = templBuffer.WriteString("<p class=\"text-red-300\">")
			if err != nil {
				return err
			}
			var_5 := `Please enter your name`
			_, err = templBuffer.WriteString(var_5)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</p>")
			if err != nil {
				return err
			}
		} else {
			_, err = templBuffer.WriteString("<p class=\"text-4xl font-bold text-emerald-300\">")
			if err != nil {
				return err
			}
			var_6 := `Hello `
			_, err = templBuffer.WriteString(var_6)
			if err != nil {
				return err
			}
			var var_7 string = name
			_, err = templBuffer.WriteString(templ.EscapeString(var_7))
			if err != nil {
				return err
			}
			var_8 := `!`
			_, err = templBuffer.WriteString(var_8)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</p>")
			if err != nil {
				return err
			}
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func HelloPage() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_9 := templ.GetChildren(ctx)
		if var_9 == nil {
			var_9 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var_10 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			err = Hello().Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = common.Page("Hello").Render(templ.WithChildren(ctx, var_10), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}
