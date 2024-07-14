import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          DEFAULT: "#EA5455",
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
    },
  },
  plugins: [],
};
export default config;
