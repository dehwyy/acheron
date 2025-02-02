import { heroui } from "@heroui/theme"

/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./components/**/*.{js,ts,jsx,tsx,mdx}",
        "./app/**/*.{js,ts,jsx,tsx,mdx}",
        "./node_modules/@heroui/theme/dist/**/*.{js,ts,jsx,tsx}", // as monorepo
        "../../node_modules/@heroui/theme/dist/**/*.{js,ts,jsx,tsx}" // as workspace
    ],
    theme: {
        extend: {
            colors: {
                "purple-light": "hsl(235, 86.1%,77.5%)"
            },
            fontFamily: {
                sans: ["var(--font-sans)"],
                mono: ["var(--font-mono)"],
                curly: ["var(--font-sacramento)"],
                common: ["var(--font-common)"]
            },
            borderRadius: {
                lg: "var(--radius)",
                md: "calc(var(--radius) - 2px)",
                sm: "calc(var(--radius) - 4px)"
            }
        }
    },
    darkMode: ["class"],
    plugins: [
        heroui({
            themes: {
                "purple-dark": {
                    extend: "dark",
                    colors: {
                        background: "#0D001A",
                        foreground: "#ffffff",
                        primary: {
                            50: "#3B096C",
                            100: "#520F83",
                            200: "#7318A2",
                            300: "#9823C2",
                            400: "#c031e2",
                            500: "#DD62ED",
                            600: "#F182F6",
                            700: "#FCADF9",
                            800: "#FDD5F9",
                            900: "#FEECFE",
                            DEFAULT: "#DD62ED",
                            foreground: "#ffffff"
                        },
                        focus: "#F182F6"
                    },
                    layout: {
                        disabledOpacity: "0.3",
                        radius: {
                            small: "4px",
                            medium: "6px",
                            large: "8px"
                        },
                        borderWidth: {
                            small: "1px",
                            medium: "2px",
                            large: "3px"
                        }
                    }
                }
            }
        }),
        require("tailwindcss-animate")
    ]
}
