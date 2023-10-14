// Code generated by templ@v0.2.334 DO NOT EDIT.

package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CreateUser() templ.Component {
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
		_, err = templBuffer.WriteString("<div class=\"container mx-auto\"><div class=\"mx-auto max-w-xl rounded-lg bg-white p-6 shadow-xl dark:bg-gray-700\"><h1 class=\"mb-4 text-2xl font-semibold text-gray-700 dark:text-gray-50\">")
		if err != nil {
			return err
		}
		var_2 := `Create New User`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1><form id=\"create-user-form\" action=\"users\" method=\"POST\"><div class=\"mb-4\"><label for=\"fullName\" class=\"block text-sm font-medium text-gray-600 dark:text-gray-100\">")
		if err != nil {
			return err
		}
		var_3 := `Nome Completo`
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label><input type=\"text\" id=\"fullName\" name=\"fullName\" placeholder=\"John Doe\" class=\"mt-1 w-full rounded-md border p-2 dark:bg-gray-800 dark:text-gray-300 dark:shadow-sm\" required></div><div class=\"mb-4\"><label for=\"matricula\" class=\"block text-sm font-medium text-gray-600 dark:text-gray-100\">")
		if err != nil {
			return err
		}
		var_4 := `Matrícula`
		_, err = templBuffer.WriteString(var_4)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label><input type=\"text\" id=\"matricula\" name=\"matricula\" placeholder=\"C012345\" class=\"mt-1 w-full rounded-md border p-2 dark:bg-gray-800 dark:text-gray-300 dark:shadow-sm\" required></div><div class=\"mb-4\"><label for=\"goal1FatPercentage\" class=\"block text-sm font-medium text-gray-600 dark:text-gray-100\">")
		if err != nil {
			return err
		}
		var_5 := `Meta 1 - Percentual de Gordura`
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label><input type=\"text\" id=\"goal1FatPercentage\" name=\"goal1FatPercentage\" placeholder=\"Meta\" class=\"mt-1 w-full rounded-md border p-2 dark:bg-gray-800 dark:text-gray-300 dark:shadow-sm\" required></div><div class=\"mb-4\"><label for=\"goal1LeanMass\" class=\"block text-sm font-medium text-gray-600 dark:text-gray-100\">")
		if err != nil {
			return err
		}
		var_6 := `Meta 1 - Massa Magra`
		_, err = templBuffer.WriteString(var_6)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label><input type=\"text\" id=\"goal1LeanMass\" name=\"goal1LeanMass\" placeholder=\"Meta\" class=\"mt-1 w-full rounded-md border p-2 dark:bg-gray-800 dark:text-gray-300 dark:shadow-sm\" required></div><div class=\"mb-4\"><label for=\"goal2FatPercentage\" class=\"block text-sm font-medium text-gray-600 dark:text-gray-100\">")
		if err != nil {
			return err
		}
		var_7 := `Meta 2 - Percentual de Gordura`
		_, err = templBuffer.WriteString(var_7)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label><input type=\"text\" id=\"goal2FatPercentage\" name=\"goal2FatPercentage\" placeholder=\"Meta\" class=\"mt-1 w-full rounded-md border p-2 dark:bg-gray-800 dark:text-gray-300 dark:shadow-sm\" required></div><div class=\"mb-4\"><label for=\"goal2LeanMass\" class=\"block text-sm font-medium text-gray-600 dark:text-gray-100\">")
		if err != nil {
			return err
		}
		var_8 := `Meta 2 - Massa Magra`
		_, err = templBuffer.WriteString(var_8)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label><input type=\"text\" id=\"goal2LeanMass\" name=\"goal2LeanMass\" placeholder=\"Meta\" class=\"mt-1 w-full rounded-md border p-2 dark:bg-gray-800 dark:text-gray-300 dark:shadow-sm\" required></div><div class=\"mb-4\"><label for=\"goal2VisceralFat\" class=\"block text-sm font-medium text-gray-600 dark:text-gray-100\">")
		if err != nil {
			return err
		}
		var_9 := `Meta 2 - Gordura Visceral`
		_, err = templBuffer.WriteString(var_9)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label><input type=\"text\" id=\"goal2VisceralFat\" name=\"goal2VisceralFat\" placeholder=\"Meta\" class=\"mt-1 w-full rounded-md border p-2 dark:bg-gray-800 dark:text-gray-300 dark:shadow-sm\" required></div><div class=\"mb-4\"><label for=\"teamName\" class=\"block text-sm font-medium text-gray-600 dark:text-gray-100\">")
		if err != nil {
			return err
		}
		var_10 := `Equipe`
		_, err = templBuffer.WriteString(var_10)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label><input type=\"text\" id=\"teamName\" name=\"teamName\" placeholder=\"Nome da Equipe\" class=\"mt-1 w-full rounded-md border p-2 dark:bg-gray-800 dark:text-gray-300 dark:shadow-sm\" required></div><div class=\"mb-4\"><label for=\"teamNumber\" class=\"block text-sm font-medium text-gray-600 dark:text-gray-100\">")
		if err != nil {
			return err
		}
		var_11 := `Número da Equipe`
		_, err = templBuffer.WriteString(var_11)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label><input type=\"text\" id=\"teamNumber\" name=\"teamNumber\" placeholder=\"Número da Equipe\" class=\"mt-1 w-full rounded-md border p-2 dark:bg-gray-800 dark:text-gray-300 dark:shadow-sm\" required></div><div class=\"mt-6\"><button type=\"submit\" class=\"rounded-md bg-orange-600 px-4 py-2 text-white hover:bg-orange-700\">")
		if err != nil {
			return err
		}
		var_12 := `Criar Participante`
		_, err = templBuffer.WriteString(var_12)
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