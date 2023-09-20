package main

const tmpl = `<!DOCTYPE html>
<html lang="pt-br">
  <head>
    <title>Fitcon</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="https://rsms.me/inter/inter.css">
  </head>
  <style>
   /*! tailwindcss v3.3.3 | MIT License | https://tailwindcss.com*/*,:after,:before{box-sizing:border-box;border:0 solid #e5e7eb}:after,:before{--tw-content:""}html{line-height:1.5;-webkit-text-size-adjust:100%;-moz-tab-size:4;-o-tab-size:4;tab-size:4;font-family:ui-sans-serif,system-ui,-apple-system,BlinkMacSystemFont,Segoe UI,Roboto,Helvetica Neue,Arial,Noto Sans,sans-serif,Apple Color Emoji,Segoe UI Emoji,Segoe UI Symbol,Noto Color Emoji;font-feature-settings:normal;font-variation-settings:normal}body{margin:0;line-height:inherit}hr{height:0;color:inherit;border-top-width:1px}abbr:where([title]){-webkit-text-decoration:underline dotted;text-decoration:underline dotted}h1,h2,h3,h4,h5,h6{font-size:inherit;font-weight:inherit}a{color:inherit;text-decoration:inherit}b,strong{font-weight:bolder}code,kbd,pre,samp{font-family:ui-monospace,SFMono-Regular,Menlo,Monaco,Consolas,Liberation Mono,Courier New,monospace;font-size:1em}small{font-size:80%}sub,sup{font-size:75%;line-height:0;position:relative;vertical-align:initial}sub{bottom:-.25em}sup{top:-.5em}table{text-indent:0;border-color:inherit;border-collapse:collapse}button,input,optgroup,select,textarea{font-family:inherit;font-feature-settings:inherit;font-variation-settings:inherit;font-size:100%;font-weight:inherit;line-height:inherit;color:inherit;margin:0;padding:0}button,select{text-transform:none}[type=button],[type=reset],[type=submit],button{-webkit-appearance:button;background-color:initial;background-image:none}:-moz-focusring{outline:auto}:-moz-ui-invalid{box-shadow:none}progress{vertical-align:initial}::-webkit-inner-spin-button,::-webkit-outer-spin-button{height:auto}[type=search]{-webkit-appearance:textfield;outline-offset:-2px}::-webkit-search-decoration{-webkit-appearance:none}::-webkit-file-upload-button{-webkit-appearance:button;font:inherit}summary{display:list-item}blockquote,dd,dl,figure,h1,h2,h3,h4,h5,h6,hr,p,pre{margin:0}fieldset{margin:0}fieldset,legend{padding:0}menu,ol,ul{list-style:none;margin:0;padding:0}dialog{padding:0}textarea{resize:vertical}input::-moz-placeholder,textarea::-moz-placeholder{opacity:1;color:#9ca3af}input::placeholder,textarea::placeholder{opacity:1;color:#9ca3af}[role=button],button{cursor:pointer}:disabled{cursor:default}audio,canvas,embed,iframe,img,object,svg,video{display:block;vertical-align:middle}img,video{max-width:100%;height:auto}[hidden]{display:none}*,::backdrop,:after,:before{--tw-border-spacing-x:0;--tw-border-spacing-y:0;--tw-translate-x:0;--tw-translate-y:0;--tw-rotate:0;--tw-skew-x:0;--tw-skew-y:0;--tw-scale-x:1;--tw-scale-y:1;--tw-pan-x: ;--tw-pan-y: ;--tw-pinch-zoom: ;--tw-scroll-snap-strictness:proximity;--tw-gradient-from-position: ;--tw-gradient-via-position: ;--tw-gradient-to-position: ;--tw-ordinal: ;--tw-slashed-zero: ;--tw-numeric-figure: ;--tw-numeric-spacing: ;--tw-numeric-fraction: ;--tw-ring-inset: ;--tw-ring-offset-width:0px;--tw-ring-offset-color:#fff;--tw-ring-color:#3b82f680;--tw-ring-offset-shadow:0 0 #0000;--tw-ring-shadow:0 0 #0000;--tw-shadow:0 0 #0000;--tw-shadow-colored:0 0 #0000;--tw-blur: ;--tw-brightness: ;--tw-contrast: ;--tw-grayscale: ;--tw-hue-rotate: ;--tw-invert: ;--tw-saturate: ;--tw-sepia: ;--tw-drop-shadow: ;--tw-backdrop-blur: ;--tw-backdrop-brightness: ;--tw-backdrop-contrast: ;--tw-backdrop-grayscale: ;--tw-backdrop-hue-rotate: ;--tw-backdrop-invert: ;--tw-backdrop-opacity: ;--tw-backdrop-saturate: ;--tw-backdrop-sepia: }.mx-auto{margin-left:auto;margin-right:auto}.inline{display:inline}.flex{display:flex}.table{display:table}.h-48{height:12rem}.w-48{width:12rem}.max-w-sm{max-width:24rem}.shrink-0{flex-shrink:0}.table-auto{table-layout:auto}.items-center{align-items:center}.space-x-4>:not([hidden])~:not([hidden]){--tw-space-x-reverse:0;margin-right:calc(1rem*var(--tw-space-x-reverse));margin-left:calc(1rem*(1 - var(--tw-space-x-reverse)))}.rounded-lg{border-radius:.5rem}.rounded-xl{border-radius:.75rem}.rounded-bl-lg{border-bottom-left-radius:.5rem}.rounded-br-lg{border-bottom-right-radius:.5rem}.rounded-tl-lg{border-top-left-radius:.5rem}.rounded-tr-lg{border-top-right-radius:.5rem}.bg-sky-600{--tw-bg-opacity:1;background-color:rgb(2 132 199/var(--tw-bg-opacity))}.bg-sky-700{--tw-bg-opacity:1;background-color:rgb(3 105 161/var(--tw-bg-opacity))}.bg-white{--tw-bg-opacity:1;background-color:rgb(255 255 255/var(--tw-bg-opacity))}.bg-zinc-100{--tw-bg-opacity:1;background-color:rgb(244 244 245/var(--tw-bg-opacity))}.bg-zinc-700{--tw-bg-opacity:1;background-color:rgb(63 63 70/var(--tw-bg-opacity))}.p-0{padding:0}.p-3{padding:.75rem}.p-5{padding:1.25rem}.p-6{padding:1.5rem}.text-center{text-align:center}.text-3xl{font-size:1.875rem;line-height:2.25rem}.font-medium{font-weight:500}.uppercase{text-transform:uppercase}.text-black{--tw-text-opacity:1;color:rgb(0 0 0/var(--tw-text-opacity))}.drop-shadow-2xl{--tw-drop-shadow:drop-shadow(0 25px 25px #00000026);filter:var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)}@media (prefers-color-scheme:dark){.dark\:bg-zinc-600{--tw-bg-opacity:1;background-color:rgb(82 82 91/var(--tw-bg-opacity))}.dark\:bg-zinc-800{--tw-bg-opacity:1;background-color:rgb(39 39 42/var(--tw-bg-opacity))}.dark\:text-zinc-100{--tw-text-opacity:1;color:rgb(244 244 245/var(--tw-text-opacity))}.dark\:text-zinc-200{--tw-text-opacity:1;color:rgb(228 228 231/var(--tw-text-opacity))}}
  </style>
  <script>
    if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
      document.documentElement.classList.add('dark')
    } else {
      document.documentElement.classList.remove('dark')
    }
  </script>
  <body class="bg-zinc-100 dark:bg-zinc-800 text-black dark:text-zinc-200">
    <div class="p-6 max-w-sm mx-auto bg-zinc-100 dark:bg-zinc-800 dark:text-zinc-100 rounded-xl items-center space-x-4 flex items-center">
      <div class="shrink-0">
        <img class="h-48 w-48 p-0" src="./img/logo.png" alt="Fitcon Logo">
      </div>
      <div>
        <div class="text-3xl font-medium">Metas</div>
      </div>
    </div>
    <table class="mx-auto drop-shadow-2xl table-auto bg-white dark:bg-zinc-600 rounded-lg dark:text-zinc-100">
      <thead class="uppercase">
        <tr>
          <th class="p-5 bg-sky-600 rounded-tl-lg" rowspan="2">Nome</th>
          <th class="p-5 bg-sky-600" colspan="2">Meta 1</th>
          <th class="p-5 bg-sky-600 rounded-tr-lg" colspan="3">Meta 2</th>
        </tr>
        <tr>
          <th class="p-3 bg-sky-700">Percentual de Gordura</th>
          <th class="p-3 bg-sky-700">Massa Magra</th>
          <th class="p-3 bg-sky-700">Gordura Visceral</th>
          <th class="p-3 bg-sky-700">Gordura Corporal</th>
          <th class="p-3 bg-sky-700">Massa Magra</th>
        </tr>
      </thead>
      <tr>
        <td class="text-center p-3 bg-zinc-700 font-medium rounded-bl-lg">{{ .Name }}</td>
        <td class="text-center p-3 bg-zinc-700">{{ .Meta1.FatPercentage }}</td>
        <td class="text-center p-3 bg-zinc-700">{{ .Meta1.LeanMass }}</td>
        <td class="text-center p-3 bg-zinc-700">{{ .Meta2.VisceralFat }}</td>
        <td class="text-center p-3 bg-zinc-700">{{ .Meta2.FatPercentage }}</td>
        <td class="text-center p-3 bg-zinc-700 rounded-br-lg">{{ .Meta2.LeanMass }}</td>
      </tr>
    </table>
  </body>
</html>
`
