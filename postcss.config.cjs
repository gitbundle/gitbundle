module.exports = {
  plugins: {
    // '@marketing-relevance/postcss-font-awesome': {},
    'postcss-import': {},
    'tailwindcss/nesting': {},
    tailwindcss: {},
    'postcss-preset-env': {
      features: { 'nesting-rules': true }
    },
    autoprefixer: {}
  }
}
