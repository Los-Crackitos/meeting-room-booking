module.exports = {
  theme: {
    extend: {
      fill: theme => ({
        'current': 'currentColor',
        'gray-600': theme('colors.gray.600'),
        'gray-800': theme('colors.gray.800'),
      }),
    },
  },
  variants: {
    display: ['children', 'default', 'children-first', 'children-last', 'children-odd', 'children-even', 'children-not-first', 'children-hover', 'hover', 'children-focus', 'focus', 'children-focus-within', 'focus-within', 'children-active', 'active', 'children-visited', 'visited', 'children-disabled', 'disabled', 'responsive'],
  },
  plugins: [
    require('tailwindcss-children'),
  ],
}
