import { svelte } from '@sveltejs/vite-plugin-svelte'
import glob from 'glob'
import path from 'path'
import sveltePreprocess from 'svelte-preprocess'
import { fileURLToPath } from 'url'
import { defineConfig } from 'vite'
import Inspect from 'vite-plugin-inspect'

const input = Object.fromEntries(
  glob
    .sync('./web_src/{js,css}/*.{js,css}')
    .map((file) => [
      path.relative(
        'web_src',
        file.slice(0, file.length - path.extname(file).length)
      ),
      fileURLToPath(new URL(file, import.meta.url))
    ])
)
console.log(input)

const output = {
  assetFileNames: ({ name, source, type }) => {
    if (name === 'serviceworker') return '[name].[hash].js'
    if (/\.css/.test(name)) return 'css/[name].[hash].css'
    if (/\.js/.test(name)) return '[name].[hash].js'
    if (/\.(eot|ttf|otf|woff)$/.test(name)) return 'fonts/[name].[hash].[ext]'
    if (/\.(png|jpg|jpeg|gif|svg)$/.test(name))
      return 'img/webpack/[name].[hash].[ext]'
    console.log('asset => ', name, type)
    return '[name].[hash].[ext]'
  },
  chunkFileNames: ({ facadeModuleId, name }) => {
    const matches = /monaco.*languages?_.+?_(.+?)_/.exec(facadeModuleId)
    const language = (matches || [])[1]
    // console.log(
    //   'chunk => ',
    //   facadeModuleId,
    //   name,
    //   ', language => ',
    //   language,
    //   matches
    // )
    return `js/chunks/${
      language ? `monaco-language-${language.toLowerCase()}-[hash]` : `[name]`
    }.js`
  },
  entryFileNames: ({ name }) =>
    name === 'js/serviceworker' ? '[name].js' : '[name]-[hash].js'
  //TODO: should production bundle with sourcemap?
  // sourcemap: true
}

export default defineConfig(({ mode }) => {
  return {
    server: {
      port: 4001,
      cors: true
      // cors: {
      //   origin: 'http://localhost:4000',
      //   methods: '*'
      //   // preflightContinue: true,
      //   // optionsSuccessStatus: 204
      // }
    },
    publicDir: './web_src/public',
    build: {
      rollupOptions: {
        input,
        output,
        external: ['jquery', 'react', 'react-dom']
      },
      outDir: './public'
    },
    plugins: [
      svelte({
        preprocess: sveltePreprocess({
          typescript: true,
          postcss: true
        }),
        compilerOptions: {
          dev: mode !== 'production'
        }
      }),
      Inspect()
    ],
    test: {
      include: ['web_src/js/**/*.{test,spec}.{js,ts}']
    },
    clearScreen: false
  }
})
