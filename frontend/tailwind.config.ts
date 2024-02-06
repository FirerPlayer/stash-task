import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
      },
      colors: {
        'primary': {
          50: '#fdf2e7',
          100: '#fce4cf',
          200: '#f8c9a0',
          300: '#f5ae70',
          400: '#f29340',
          500: '#ee7811',
          600: '#bf600d',
          700: '#8f480a',
          800: '#5f3007',
          900: '#301803',
          950: '#180c02',
        }
      }
    },
  },
  darkMode: "class",
  plugins: [],
};
export default config;
