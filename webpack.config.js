import { EsbuildPlugin } from 'esbuild-loader'
import MiniCssExtractPlugin from 'mini-css-extract-plugin'
import MonacoWebpackPlugin from 'monaco-editor-webpack-plugin'
import path from 'node:path'
import sveltePreprocess from 'svelte-preprocess'
import webpack from 'webpack'
import { globSync } from 'glob'
import fs from 'node:fs'
import { fileURLToPath } from 'node:url'
import SpeedMeasurePlugin from 'speed-measure-webpack-plugin'

const smp = new SpeedMeasurePlugin({
  disable: true,
  outputFormat: 'humanVerbose',
  loaderTopFiles: 10
  // compareLoadersBuild: {
  //   filePath: './build.json'
  // }
})

const themes = {}
const filterCssImport = (url, ...args) => {
  const cssFile = args[1] || args[0] // resourcePath is 2nd argument for url and 3rd for import
  const importedFile = url.replace(/[?#].+/, '').toLowerCase()

  if (cssFile.includes('fomantic')) {
    if (/brand-icons/.test(importedFile)) return false
    if (/(eot|ttf|otf|woff|svg)$/.test(importedFile)) return false
  }

  if (cssFile.includes('katex') && /(ttf|woff)$/.test(importedFile)) {
    return false
  }

  return true
}

const isProduction = process.env.NODE_ENV === 'production'
console.log('isProduction:', isProduction)
const isBeta = process.env.NODE_ENV === 'beta'
console.log('isBeta:', isBeta)
const isDevelopment = process.env.NODE_ENV === 'development'
console.log('isDevelopment:', isDevelopment)

let entry
if (isDevelopment) {
  entry = globSync('./web_src/{js,css}/*.{js,css}').reduce((acc, path) => {
    const name = path
      .replace(/\.(css|js)$/, '')
    if (!(name in acc)) acc[name] = []
    acc[name].push('./' + path)
    return acc
  }, {})
} else {
  //TODO: compile all js into a single file maybe lead to some uncatching error
  const files = globSync('./web_src/js/*.js').reduce((acc, filepath) => {
    const file = path.relative('web_src/js', filepath)
    if (file === 'main.js' || file === 'serviceworker.js') {
      return acc
    }
    acc.push('./' + file)
    return acc
  }, [])

  const imports = files.map((file) => `import '${file}'`).join('\n')
  fs.writeFileSync('./web_src/js/main.js', imports)

  entry = {
    index: [
      './web_src/js/main.js',
      './web_src/css/index.css'
    ],
    main: ['./web_src/css/modules/base/main.css'],
    swagger: [
      './web_src/js/standalone/swagger.js',
      './web_src/css/standalone/swagger.css'
    ],
    serviceworker: ['./web_src/js/serviceworker.js'],
    'eventsource.sharedworker': [
      './web_src/js/modules/eventsource.sharedworker.js'
    ],
    ...themes
  }
}
console.log(entry)

export default smp.wrap({
  mode: isProduction ? 'production' : 'development',
  entry,
  devtool: isProduction ? false : 'eval-cheap-module-source-map',
  devServer: {
    headers: {
      'Access-Control-Allow-Origin': 'http://localhost:4000',
      'Access-Control-Allow-Headers': '*'
    },
    hot: true,
    compress: true,
    port: 4001
  },
  output: {
    path: fileURLToPath(new URL('public', import.meta.url)),
    filename: ({ chunk }) => {
      // serviceworker can only manage assets below it's script's directory so
      // we have to put it in / instead of /js/
      return chunk.name === 'serviceworker' ? '[name].js' : 'js/[name].js'
    },
    chunkFilename: ({ chunk }) => {
      const language = (/monaco.*languages?_.+?_(.+?)_/.exec(chunk.id) || [])[1]
      // return `js/${language ? `monaco-language-${language.toLowerCase()}` : `[name]`}.[contenthash:8].js`
      return `js/${
        language
          ? `monaco-language-${language.toLowerCase()}`
          : isProduction
          ? `[name].[contenthash:8]`
          : `[name]`
      }.js`
    }
  },
  optimization: {
    // runtimeChunk: true,
    // removeAvailableModules: false,
    // removeEmptyChunks: false,
    // splitChunks: false,
    minimize: isProduction || isBeta,
    minimizer: [
      new EsbuildPlugin({
        target: 'es2015',
        minify: isProduction || isBeta,
        css: true,
        legalComments: 'none',
        drop: isProduction ? ['console'] : []
      })
    ],
    splitChunks: {
      chunks: 'async',
      name: (_, chunks) => chunks.map((item) => item.name).join('-')
    },
    moduleIds: 'named',
    chunkIds: 'named'
  },
  externals: {
    // jquery: '$'
    // react: 'React',
    // 'react-dom': 'ReactDOM'
  },
  module: {
    rules: [
      {
        test: /\.svelte$/,
        include: [
          fileURLToPath(new URL('web_src', import.meta.url)),
          fileURLToPath(new URL('node_modules', import.meta.url))
        ],
        use: {
          loader: 'svelte-loader',
          options: {
            preprocess: sveltePreprocess({
              typescript: true,
              postcss: true
            }),
            compilerOptions: {
              dev: !isProduction
            },
            emitCss: isProduction,
            hotReload: !isProduction
          }
        }
      },
      {
        // required to prevent errors from Svelte on Webpack 5+, omit on Webpack 4
        test: /node_modules\/svelte\/.*\.mjs$/,
        include: [
          fileURLToPath(new URL('web_src', import.meta.url)),
          fileURLToPath(new URL('node_modules', import.meta.url))
        ],
        resolve: {
          fullySpecified: false
        }
      },
      {
        test: /\.worker\.js$/,
        include: [fileURLToPath(new URL('web_src', import.meta.url))],
        use: [
          {
            loader: 'worker-loader',
            options: {
              inline: 'no-fallback'
            }
          }
        ]
      },
      {
        test: /\.[jt]s$/,
        include: [fileURLToPath(new URL('web_src', import.meta.url))],
        use: [
          {
            loader: 'esbuild-loader',
            options: {
              loader: 'ts',
              target: 'es2015'
            }
          }
        ]
      },
      {
        test: /.css$/i,
        include: [
          fileURLToPath(new URL('web_src', import.meta.url)),
          fileURLToPath(new URL('node_modules', import.meta.url))
        ],
        use: [
          MiniCssExtractPlugin.loader,
          {
            loader: 'css-loader',
            options: {
              sourceMap: isBeta || isDevelopment,
              url: { filter: filterCssImport },
              import: { filter: filterCssImport }
            }
          },
          'postcss-loader',
          {
            loader: 'esbuild-loader',
            options: {
              loader: 'css',
              target: 'es2015'
            }
          }
        ]
      },
      {
        test: /\.svg$/,
        include: [fileURLToPath(new URL('public/img/svg', import.meta.url))],
        type: 'asset/source'
      },
      {
        test: /\.(ttf|woff2?)$/,
        type: 'asset/resource',
        generator: {
          filename: 'fonts/[name].[contenthash:8][ext]'
        }
      },
      {
        test: /\.png$/i,
        type: 'asset/resource',
        generator: {
          filename: 'img/webpack/[name].[contenthash:8][ext]'
        }
      }
    ]
  },
  plugins: [
    new MiniCssExtractPlugin({
      filename: 'css/[name].css',
      chunkFilename: `css/${
        isProduction ? `[name].[contenthash:8]` : `[name]`
      }.css`
    }),
    new webpack.SourceMapDevToolPlugin({
      filename: `${isProduction ? `[file].[contenthash:8]` : `[file]`}.map`,
      include: ['js/index.js', 'css/index.css']
    }),
    new MonacoWebpackPlugin({
      filename: `js/monaco-${
        isProduction ? `[name].[contenthash:8]` : `[name]`
      }.worker.js`
    }),
    new webpack.ProvidePlugin({
      process: 'process/browser.js'
    })
  ],
  performance: {
    hints: false,
    maxEntrypointSize: Infinity,
    maxAssetSize: Infinity
  },
  resolve: {
    symlinks: false,
    // see below for an explanation
    alias: {
      svelte: path.resolve('node_modules', 'svelte'),
      mermaid: path.resolve('node_modules', 'mermaid'),
      stylis: path.resolve('node_modules', 'stylis'),
      '@node_modules': path.resolve('node_modules'),
      'js-yaml': path.resolve('node_modules', 'js-yaml')
    },
    extensions: ['.mjs', '.js', '.ts', '.svelte'],
    mainFields: ['svelte', 'browser', 'module', 'main'],
    conditionNames: ['svelte']
  },
  watchOptions: {
    ignored: ['node_modules/**', 'public/js/**'],
    aggregateTimeout: 300,
    poll: 1000 // NOTE: this will trigger unstopped compile and eat lots of cpu
  },
  stats: {
    assetsSort: 'name',
    assetsSpace: Infinity,
    cached: false,
    cachedModules: false,
    children: false,
    chunkModules: false,
    chunkOrigins: false,
    chunksSort: 'name',
    colors: true,
    entrypoints: false,
    excludeAssets: [
      /^js\/monaco-language-.+\.js$/,
      !isProduction && /^js\/licenses.txt$/
    ].filter((item) => !!item),
    groupAssetsByChunk: false,
    groupAssetsByEmitStatus: false,
    groupAssetsByInfo: false,
    groupModulesByAttributes: false,
    modules: false,
    reasons: false,
    runtimeModules: false
  },
  cache: true
})
