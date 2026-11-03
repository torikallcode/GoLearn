import typography from '@tailwindcss/typography'

/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{js,jsx}'],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        brand: {
          50: '#eefdf8',
          100: '#d6f9ec',
          200: '#b0f1db',
          300: '#7be5c6',
          400: '#3fd0ab',
          500: '#16b591',
          600: '#0a9278',
          700: '#087563',
          800: '#0a5d50',
          900: '#0a4c43',
        },
      },
      fontFamily: {
        sans: ['Inter', 'ui-sans-serif', 'system-ui', 'sans-serif'],
        mono: ['JetBrains Mono', 'ui-monospace', 'SFMono-Regular', 'monospace'],
      },
    },
  },
  plugins: [typography],
}
