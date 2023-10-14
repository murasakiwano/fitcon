// Code generated by templ@v0.2.334 DO NOT EDIT.

package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func UserIdInvalid(id string) templ.Component {
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
		var_2 := `Matrícula inválida`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1><p class=\"mb-6 text-gray-500 dark:text-gray-300\">")
		if err != nil {
			return err
		}
		var_3 := `O formato da matrícula é inválido. Por favor insira uma matrícula válida.`
		_, err = templBuffer.WriteString(var_3)
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