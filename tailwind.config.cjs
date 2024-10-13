/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['web_src/**/*.{js,jsx,ts,tsx,svelte}', 'templates/**/*.html'],
  theme: {
    extend: {
      colors: require('./colors.config.cjs'),
      screens: {
        '3xl': '1792px',
        'cu-w': '2560px'
      },
      // tabSize: {
      //   2: '0.5rem',
      //   4: '1rem'
      // },
      gridTemplateColumns: {
        13: 'repeat(13, minmax(0, 1fr))',
        14: 'repeat(14, minmax(0, 1fr))',
        15: 'repeat(15, minmax(0, 1fr))',
        16: 'repeat(16, minmax(0, 1fr))'
      },
      gridColumn: {
        'span-13': 'span 13 / span 13',
        'span-14': 'span 14 / span 14',
        'span-15': 'span 15 / span 15',
        'span-16': 'span 16 / span 16'
      }
      // gridAutoColumns: {
      //   '2fr': 'minmax(0, 2fr)'
      // }
    }
  },
  plugins: [
    require('@tailwindcss/typography'),
    // require('@tailwindcss/line-clamp'),//NOTE: tw3.3 has include this by default
    // require('@tailwindcss/forms'),
    require('daisyui'),
    require('tailwindcss-animate'),
    require('tailwindcss-animated')
  ],
  daisyui: {
    styled: true,
    base: true,
    utils: true,
    logs: true,
    rtl: false,
    darkTheme: 'geniusDark',
    themes: [
      {
        geniusDark: {
          primary: '#2563eb', //#09244B
          secondary: '#155e75',
          accent: '#6b21a8',
          neutral: '#334155',
          info: '#1e40af',
          success: '#166534',
          warning: '#eab308',
          error: '#991b1b',
          'base-100': '#2a303c',
          'base-200': '#262b36',
          'base-300': '#222731',
          'base-content': '#A6ADBB',
          'primary-content': '#ffffff',
          'secondary-content': '#E9EBEE',
        }
      },
      'dark',
      'light'
    ]
  }
}
