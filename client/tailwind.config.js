const defaultTheme = require("tailwindcss/defaultTheme");
module.exports = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
    screens: {
      xsm: { max: "460px" },
      ...defaultTheme.screens,
    },
  },
  plugins: [require("daisyui"), "macros"],
  format: "auto",
  daisyui: {
    themes: false,
  },
};
