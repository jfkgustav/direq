/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.html", "./**/*.templ", "./**/*.go","./view/**/*.{templ, html}", "./view/*.{templ, html}"],
  theme: {
    extend: {},
  },
  variants: {
    extend: {
      borderColor: ['focus'],
      outline: ['focus'],
    }
  },
  plugins: [
    require('tailwindcss'),
  ],
}

