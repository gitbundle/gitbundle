const AddAssetPlugin = require('add-asset-webpack-plugin')
const MiniCssExtractPlugin = require('mini-css-extract-plugin')
const MonacoWebpackPlugin = require('monaco-editor-webpack-plugin')
const { EsbuildPlugin } = require('esbuild-loader')
const path = require('path')
const webpack = require('webpack')
const WatchExternalFilesPlugin = require('webpack-watch-external-files-plugin')
const sveltePreprocess = require('svelte-preprocess')
// const { WebpackManifestPlugin } = require('webpack-manifest-plugin')
const glob = require('glob')
const SpeedMeasurePlugin = require('speed-measure-webpack-plugin')

const smp = new SpeedMeasurePlugin({
  disable: true,
  outputFormat: 'humanVerbose',
  loaderTopFiles: 10
  // compareLoadersBuild: {
  //   filePath: './build.json'
  // }
})

const themes = {}
// for (const path of glob("web_src/less/themes/*.less")) {
//   themes[parse(path).name] = [path];
// }

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

  if (
    cssFile.includes('font-awesome') &&
    /(eot|ttf|otf|woff|svg)$/.test(importedFile)
  ) {
    return false
  }

  return true
}

const isProduction = process.env.NODE_ENV !== 'development'
console.log('isProduction:', isProduction)

const entry = glob.sync('./web_src/{js,css}/*.{js,css}').reduce((acc, path) => {
  if (!isProduction) {
    const name = path
      .replace(/\.\/web_src\/(css|js)\//, '')
      .replace(/\.(css|js)$/, '')
    if (!(name in acc)) acc[name] = []
    acc[name].push(path)
    return acc
  }

  const name = 'index'
  if (!(name in acc)) acc[name] = []
  acc[name].push(path)
  return acc
}, {})
console.log(entry)

module.exports = smp.wrap({
  mode: isProduction ? 'production' : 'development',
  // TODO: optimise the bundle.js for enterprice??
  // entry: {
  //   index: [
  //     // fileURLToPath(new URL("web_src/js/jquery.js", import.meta.url)),
  //     './web_src/js/index.js',
  //     // './node_modules/easymde/dist/easymde.min.css',
  //     './web_src/css/index.css'
  //   ],
  //   swagger: ['./web_src/js/standalone/swagger.js', './web_src/css/standalone/swagger.css'],
  //   serviceworker: ['./web_src/js/serviceworker.js'],
  //   'eventsource.sharedworker': ['./web_src/js/features/eventsource.sharedworker.js'],
  //   ...themes
  // },
  entry,
  // TODO: disable for real production
  devtool: isProduction ? 'source-map' : 'eval-cheap-module-source-map',
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
    pathinfo: false,
    path: __dirname + '/public',
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
    removeAvailableModules: false,
    removeEmptyChunks: false,
    splitChunks: false,
    // minimize: isProduction,
    minimizer: [
      new EsbuildPlugin({
        target: 'es2015',
        minify: isProduction,
        css: true,
        legalComments: 'none'
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
    jquery: '$',
    react: 'React',
    'react-dom': 'ReactDOM'
  },
  module: {
    rules: [
      {
        test: /\.svelte$/,
        include: [__dirname + '/web_src', __dirname + '/node_modules'],
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
        include: [__dirname + '/web_src', __dirname + '/node_modules'],
        resolve: {
          fullySpecified: false
        }
      },
      {
        test: /\.worker\.js$/,
        include: path.resolve(__dirname, './web_src'),
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
        include: path.resolve(__dirname, './web_src'),
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
        include: [__dirname + '/web_src', __dirname + '/node_modules'],
        use: [
          MiniCssExtractPlugin.loader,
          {
            loader: 'css-loader',
            options: {
              sourceMap: true,
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
        include: __dirname + '/public/img/svg',
        type: 'asset/source'
      },
      // {
      //   test: /\.svg$/,
      //   issuer: /\.[jt]sx?$/,
      //   exclude: fileURLToPath(new URL("public/img/svg", import.meta.url)),
      //   use: [{ loader: "@svgr/webpack", options: { icon: true } }],
      // },
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
    // new WebpackManifestPlugin({
    //   basePath: 'public',
    //   publicPath: './public/'
    // }),
    // new WatchExternalFilesPlugin({
    //   // files: ['templates/**/*.html', 'web_src/**/*.js']
    //   files: ['templates/**/*.html']
    // }),
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
    // new AddAssetPlugin(
    //   'js/licenses.txt',
    //   `Licenses will be filled in the future.` //TODO: will be filled in the future
    // )
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
      stylis: path.resolve('node_modules', 'stylis')
    },
    extensions: ['.mjs', '.js', '.ts', '.svelte'],
    mainFields: ['svelte', 'browser', 'module', 'main'],
    conditionNames: ['svelte']
  },
  watchOptions: {
    ignored: ['node_modules/**'],
    aggregateTimeout: 300,
    poll: 1000
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
