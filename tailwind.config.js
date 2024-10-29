/** @type {import('tailwindcss').Config} */
const colors = require("tailwindcss/colors");

module.exports = {
  content: ["./view/**/*.{templ,html,js,ts,go}"],
  safelist: [
    {
      pattern: /basis-/,
      variants: ["sm", "md", "lg", "xl", "2xl"],
    },
    {
      pattern: /bg-/,
      variants: ["sm", "md", "lg", "xl", "2xl"],
    },
    {
      pattern: /static|fixed|relative|absolute/,
    },
    {
      pattern: /rounded-/,
      variants: ["sm", "md", "lg", "xl", "2xl"],
    },
  ],
  theme: {
    container: {
      center: true,
      padding: "2rem",
      screens: {
        "2xl": "1400px",
      },
    },
    colors: {
      transparent: "transparent",
      scred: `var(--scred)`,
      scdeepred: `var(--scdeepred)`,
      scultraviolet: `var(--scultraviolet)`,
      scpurple: `var(--scpurple)`,
      scsky: `var(--scsky)`,
      scpink: `var(--scpink)`,
      scblue: `var(--scblue)`,
      scblack: `var(--scblack)`,
      scwhite: `var(--scwhite)`,
      scgray: {
        46: `var(--scgray-46)`,
        76: `var(--scgray-76)`,
        84: `var(--scgray-84)`,
        90: `var(--scgray-90)`,
        95: `var(--scgray-95)`,
        98: `var(--scgray-98)`,
      },
      black: `var(--scblack)`,
      blue: colors.blue,
      purple: "#7e5bef",
      pink: "#ff49db",
      green: colors.green,
      yellow: "#ffc82c",
      gray: colors.gray,
      red: colors.red,
      white: `var(--scwhite)`,
      sky: colors.sky,
      slate: colors.slate,
      border: `var(--scpurple)`,
      input: `var(--scgray-46)`,
      ring: `var(--scblue)`,
      background: `var(--scgray-95)`,
      foreground: `var(--scblack)`,
      primary: {
        DEFAULT: `var(--scultraviolet)`,
        foreground: `var(--scwhite)`,
      },
      secondary: {
        DEFAULT: "#eb001a",
        foreground: `var(--scwhite)`,
      },
      destructive: {
        DEFAULT: `var(--scred)`,
        foreground: "#f8fafc",
      },
      muted: {
        DEFAULT: `var(--scgray-95)`,
        foreground: `var(--scgray-46)`,
      },
      accent: {
        DEFAULT: "var(--scgray-96)",
        foreground: `var(--scblue)`,
      },
      popover: {
        DEFAULT: `var(--scblack)`,
        foreground: `var(--scsky)`,
      },
      card: {
        DEFAULT: `var(--scblack)`,
        foreground: `var(--scgray-90)`,
      },
    },
    borderRadius: {
      full: "9999px",
      xl: `calc(var(--radius) * 2)`,
      lg: `var(--radius)`,
      md: `calc(var(--radius) - 2px)`,
      sm: "calc(var(--radius) - 4px)",
    },
    screens: {
      sm: "481px",
      md: "681px",
      lg: "769px",
      xl: "1025px",
      "2xl": "1281px",
    },
    extend: {
      backgroundImage: {
        "gradient-sc":
          "linear-gradient(-45deg, #1d08bd, #CC0017, #1d08bd, #EB001A)",
        "gradient-main":
          "linear-gradient(180deg, var(--scgray-90) 0%, var(--scgray-98) 18%)",
        "gradient-subtle":
          "linear-gradient(51deg, rgb(222, 219, 255) -10%, rgb(249, 249, 249) 40%, rgb(249, 249, 249) 70%, rgb(255, 207, 207) 120%)",
      },
      backgroundSize: {
        xl: "400% 400%",
      },
      animation: {
        gradient: "gradient 15s ease infinite",
      },
      fontFamily: {
        primary: ["DM Sans", "sans-serif"],
        sans: ["DM Sans", "sans-serif"],
      },
      zIndex: {
        top: "2100000012", // value is needed to overcome "sc-frame" element on the page
      },
      mask: {
        search:
          "url(\"data:image/svg+xml,%3Csvg xmlnsi='http://www.w3.org/2000/svg' viewBox='0 0 24 24'%3E%3Cpath d='M23.384,21.619,16.855,15.09a9.284,9.284,0,1,0-1.768,1.768l6.529,6.529a1.266,1.266,0,0,0,1.768,0A1.251,1.251,0,0,0,23.384,21.619ZM2.75,9.5a6.75,6.75,0,1,1,6.75,6.75A6.758,6.758,0,0,1,2.75,9.5Z'/%3E%3C/svg%3E \") center center/3rem 3rem no-repeat",
      },
    },
  },
  plugins: [require("@tailwindcss/typography"), require("tailwindcss-animate")],
};
