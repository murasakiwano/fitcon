// Code generated by templ@v0.2.334 DO NOT EDIT.

package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Index() templ.Component {
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
		_, err = templBuffer.WriteString("<!doctype html><html lang=\"en\"><head><title>")
		if err != nil {
			return err
		}
		var_2 := `Metas Fitcon`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</title><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><link href=\"/css/output.css\" rel=\"stylesheet\"><script src=\"/assets/js/htmx.min.js\">")
		if err != nil {
			return err
		}
		var_3 := ``
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script></head><body class=\"bg-[#1f1f28]\"><div id=\"main-container\" class=\"md: container md:mx-auto prose prose-invert lg:prose-xl py-10 min-w-[90%]\"><h1 class=\"text-center\">")
		if err != nil {
			return err
		}
		var_4 := `Metas Fitcon`
		_, err = templBuffer.WriteString(var_4)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1><br>")
		if err != nil {
			return err
		}
		err = getParticipantButton().Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div></body></html>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func getParticipantButton() templ.Component {
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
		_, err = templBuffer.WriteString("<form method=\"GET\" hx-get=\"/users\" hx-swap=\"outerHTML\" class=\"flex lg:flex-row space-x-5 mx-auto max-w-[60%]\"><div><label class=\"relative block basis-3/4\"><span class=\"sr-only\">")
		if err != nil {
			return err
		}
		var_6 := `Search`
		_, err = templBuffer.WriteString(var_6)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</span></label></div><input class=\"placeholder:italic placeholder:text-slate-400 block bg-white\n										dark:bg-slate-800 w-full border border-slate-300\n										dark:border-slate-700 rounded-md py-3 pl-4 pr-3 shadow-sm\n										focus:outline-none focus:border-sky-500 focus:ring-sky-500\n										focus:ring-1 sm:text-sm\" placeholder=\"Insira sua matrícula\" type=\"text\" name=\"matricula\"><button type=\"submit\" class=\"rounded-full bg-sky-600 px-8 text-zinc-100  font-medium basis-1/4 hover:bg-sky-700 transition-all\">")
		if err != nil {
			return err
		}
		var_7 := `Buscar`
		_, err = templBuffer.WriteString(var_7)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button></form>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func UserTable(name string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_8 := templ.GetChildren(ctx)
		if var_8 == nil {
			var_8 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div class=\"space-x-5 mx-2 overflow-x-auto min-w-[100%]\"><table class=\"drop-shadow-2xl bg-white dark:bg-zinc-600 rounded-lg text-center\"><thead class=\"uppercase\"><tr><th class=\"p-5 bg-sky-600 rounded-tl-lg\" rowspan=\"2\">")
		if err != nil {
			return err
		}
		var_9 := `Nome`
		_, err = templBuffer.WriteString(var_9)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th><th class=\"p-5 bg-sky-600\" colspan=\"2\">")
		if err != nil {
			return err
		}
		var_10 := `Meta 1`
		_, err = templBuffer.WriteString(var_10)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th><th class=\"p-5 bg-sky-600 rounded-tr-lg\" colspan=\"3\">")
		if err != nil {
			return err
		}
		var_11 := `Meta 2`
		_, err = templBuffer.WriteString(var_11)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th></tr><tr><th class=\"p-3 bg-sky-700\">")
		if err != nil {
			return err
		}
		var_12 := `Percentual de Gordura`
		_, err = templBuffer.WriteString(var_12)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th><th class=\"p-3 bg-sky-700\">")
		if err != nil {
			return err
		}
		var_13 := `Massa Magra`
		_, err = templBuffer.WriteString(var_13)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th><th class=\"p-3 bg-sky-700\">")
		if err != nil {
			return err
		}
		var_14 := `Gordura Visceral`
		_, err = templBuffer.WriteString(var_14)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th><th class=\"p-3 bg-sky-700\">")
		if err != nil {
			return err
		}
		var_15 := `Gordura Corporal`
		_, err = templBuffer.WriteString(var_15)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th><th class=\"p-3 bg-sky-700\">")
		if err != nil {
			return err
		}
		var_16 := `Massa Magra`
		_, err = templBuffer.WriteString(var_16)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th></tr></thead><tr><td class=\"text-center p-3 bg-zinc-700 font-medium rounded-bl-lg\">")
		if err != nil {
			return err
		}
		var var_17 string = name
		_, err = templBuffer.WriteString(templ.EscapeString(var_17))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</td><td class=\"text-center p-3 bg-zinc-700\">")
		if err != nil {
			return err
		}
		var_18 := `Diminuir 5%`
		_, err = templBuffer.WriteString(var_18)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</td><td class=\"text-center p-3 bg-zinc-700\">")
		if err != nil {
			return err
		}
		var_19 := `*`
		_, err = templBuffer.WriteString(var_19)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</td><td class=\"text-center p-3 bg-zinc-700\">")
		if err != nil {
			return err
		}
		var_20 := `*`
		_, err = templBuffer.WriteString(var_20)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</td><td class=\"text-center p-3 bg-zinc-700\">")
		if err != nil {
			return err
		}
		var_21 := `Aumentar 2 kg`
		_, err = templBuffer.WriteString(var_21)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</td><td class=\"text-center p-3 bg-zinc-700 rounded-br-lg\">")
		if err != nil {
			return err
		}
		var_22 := `*`
		_, err = templBuffer.WriteString(var_22)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</td></tr></table><button hx-get=\"/\" hx-target=\"#main-container\" class=\"font-medium  hover:underline p-0 transition-all\">")
		if err != nil {
			return err
		}
		var_23 := `Voltar`
		_, err = templBuffer.WriteString(var_23)
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