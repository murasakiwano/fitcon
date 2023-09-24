package templates

const tmpl = `<!DOCTYPE html>
<html lang="pt-br">
  <head>
    <title>Fitcon</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="https://rsms.me/inter/inter.css">
  </head>
  <style>
	{{ stylesheet }}
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
	{{ table }}
      </thead>
    </table>
  </body>
</html>
`

const tableTemplate = ``
