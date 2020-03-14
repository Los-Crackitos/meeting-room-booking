module.exports = {
  theme: {
    extend: {
      fill: (theme) => ({
        current: 'currentColor',
        'gray-500': theme('colors.gray.500'),
        'gray-600': theme('colors.gray.600'),
        'gray-800': theme('colors.gray.800'),
      }),
    },
  },
  variants: {},
  plugins: [],
}
