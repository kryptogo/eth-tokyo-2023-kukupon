/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx}",
    "./src/components/**/*.{js,ts,jsx,tsx}",
    "./src/app/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        cartoon: ['"M PLUS Rounded 1c"', "sans-serif"],
      },
      colors: {
        primary: "#73D252",
        primaryPink: "#FAF2E7",
        secondaryPink: "#CE6363",
      },
      backgroundImage: {
        anime: "url('/images/anime.png')",
        animeTwice: "url('/images/anime_twice.png')",
      },
      boxShadow: {
        primary: "0px 1px 35px -13px #93EF74",
      },
    },
  },
  plugins: [],
};
