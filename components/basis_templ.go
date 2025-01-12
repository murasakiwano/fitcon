// Code generated by templ@v0.2.334 DO NOT EDIT.

package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func InfoPage(title, message string) templ.Component {
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
		_, err = templBuffer.WriteString("<div id=\"card\" class=\"mx-auto max-w-xl rounded-lg bg-gray-100 p-6 text-center shadow-xl dark:bg-gray-700\" hx-swap=\"outerHTML\"><h1 class=\"mb-4 text-2xl font-semibold dark:text-gray-50\">")
		if err != nil {
			return err
		}
		var var_2 string = title
		_, err = templBuffer.WriteString(templ.EscapeString(var_2))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1><p class=\"mb-6 text-gray-500 dark:text-gray-300\">")
		if err != nil {
			return err
		}
		var var_3 string = message
		_, err = templBuffer.WriteString(templ.EscapeString(var_3))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><a href=\"/\" hx-target=\"card\" class=\"text-orange-600 hover:underline\">")
		if err != nil {
			return err
		}
		var_4 := `Retornar à página inicial`
		_, err = templBuffer.WriteString(var_4)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func FormInput(forInput, label, inputType, id, name, placeholder string, required bool) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_5 := templ.GetChildren(ctx)
		if var_5 == nil {
			var_5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div class=\"mb-4\"><label for=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(forInput))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" class=\"block text-sm font-medium text-gray-600 dark:text-gray-100\">")
		if err != nil {
			return err
		}
		var var_6 string = label
		_, err = templBuffer.WriteString(templ.EscapeString(var_6))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label><input type=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(inputType))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" id=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(id))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" name=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(name))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" placeholder=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(placeholder))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" class=\"mt-1 w-full rounded-md border p-2 dark:bg-gray-800 dark:text-gray-300 dark:shadow-sm\"")
		if err != nil {
			return err
		}
		if required {
			_, err = templBuffer.WriteString(" required")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func Button(title string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_7 := templ.GetChildren(ctx)
		if var_7 == nil {
			var_7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div><button type=\"submit\" class=\"rounded-md bg-orange-600 px-4 py-2 text-white hover:bg-orange-700\">")
		if err != nil {
			return err
		}
		var var_8 string = title
		_, err = templBuffer.WriteString(templ.EscapeString(var_8))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func Container(comp templ.Component) templ.Component {
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
		_, err = templBuffer.WriteString("<div class=\"container mx-auto\"><div class=\"mx-auto max-w-xl rounded-lg bg-white p-6 shadow-xl dark:bg-gray-700\">")
		if err != nil {
			return err
		}
		err = comp.Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
