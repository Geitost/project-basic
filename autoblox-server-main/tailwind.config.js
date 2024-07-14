/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.hbs", "./views/*.hbs"],
  theme: {
    extend: {
      colors: {
        primary: {
          DEFAULT: "#ea5455",
          50: "#ffffff",
          100: "#fbdddd",
          200: "#f7bbbb",
          300: "#f29899",
          400: "#ee7677",
          500: "#ea5455",
          600: "#bb4344",
          700: "#8c3233",
          800: "#5e2222",
          900: "#2f1111",
          950: "#000000",
        },
      },
      fontFamily: {
        sans: ["'Kanit'", "sans-serif"],
      },
    },
  },
  plugins: [],
};
