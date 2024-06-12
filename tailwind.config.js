/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/**/*.{templ,html,js,ts}"],
  safelist: [
    {
      pattern: /basis-/,
      variants: ["sm", "md", "lg", "xl", "2xl"],
    },
    {
      pattern: /static|fixed|relative|absolute/,
    },
  ],
  theme: {
    screens: {
      sm: "640px",
      md: "768px",
      lg: "1024px",
      xl: "1280px",
      "2xl": "1536px",
    },
    extend: {
      zIndex: {
        top: "2100000012", // value is needed to overcome "sc-frame" element on the page
      },
    },
  },
  plugins: [require("@tailwindcss/typography")],
};
