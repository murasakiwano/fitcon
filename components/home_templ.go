// Code generated by templ@v0.2.334 DO NOT EDIT.

package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Home() templ.Component {
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
		_, err = templBuffer.WriteString("<div id=\"card\" class=\"container mx-auto\"><div class=\"mx-auto max-w-xl rounded-lg bg-white p-6 shadow-xl dark:bg-gray-700\"><h1 class=\"mb-4 text-2xl font-semibold text-gray-700 dark:text-gray-50\">")
		if err != nil {
			return err
		}
		var_2 := `Insira sua matrícula`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1><form id=\"user-id-form\" method=\"GET\" hx-get=\"/users\" hx-target=\"#card\"><div class=\"mb-4\"><input type=\"text\" name=\"matricula\" placeholder=\"C012345\" class=\"w-full rounded-md border p-2 dark:border-none dark:bg-gray-800 dark:text-gray-300 dark:shadow-sm\"></div><div><button type=\"submit\" class=\"rounded-md bg-orange-600 px-4 py-2 text-white hover:bg-orange-700\">")
		if err != nil {
			return err
		}
		var_3 := `Ver Metas`
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button></div></form></div></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
