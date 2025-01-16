import antfu from "@antfu/eslint-config"

export default antfu({
    type: "app",
    toml: false,
    markdown: false,
    yaml: {
        overrides: {
            "yaml/indent": 2
        }
    },
    react: {
        overrides: {
            "react-hooks/exhaustive-deps": "off",
            "react-dom/no-missing-button-type": "off"
        }
    },
    typescript: {
        overrides: {
            "@typescript-eslint/no-unused-vars": "warn",
            "unused-imports/no-unused-imports": "warn",
            "unused-imports/no-unused-vars": "warn"
        }
    },
    rules: {
        "style/comma-dangle": "off"
    },
    stylistic: {
        quotes: "double",
        semi: false,
        jsx: true,
        indent: 4
    }
})
