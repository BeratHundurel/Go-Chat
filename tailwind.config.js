import { fontFamily } from "tailwindcss/defaultTheme";

module.exports = {
    content: ["./**/*.html", "./**/*.templ", "./**/*.go",],
    safelist: [],
    theme: {
        container: {
            center: true,
            padding: "2rem",
            screens: {
                "2xl": "1400px"
            }
        },
        extend: {
            colors: {
                border: "hsl(var(--border) / <alpha-value>)",
                input: "hsl(var(--input) / <alpha-value>)",
                ring: "hsl(var(--ring) / <alpha-value>)",
                background: "hsl(var(--background) / <alpha-value>)",
                foreground: "hsl(var(--foreground) / <alpha-value>)",
                primary: {
                    DEFAULT: "hsl(var(--primary) / <alpha-value>)",
                    foreground: "hsl(var(--primary-foreground) / <alpha-value>)"
                },
                secondary: {
                    DEFAULT: "hsl(var(--secondary) / <alpha-value>)",
                    foreground: "hsl(var(--secondary-foreground) / <alpha-value>)"
                },
                destructive: {
                    DEFAULT: "hsl(var(--destructive) / <alpha-value>)",
                    foreground: "hsl(var(--destructive-foreground) / <alpha-value>)"
                },
                muted: {
                    DEFAULT: "hsl(var(--muted) / <alpha-value>)",
                    foreground: "hsl(var(--muted-foreground) / <alpha-value>)"
                },
                accent: {
                    DEFAULT: "hsl(var(--accent) / <alpha-value>)",
                    foreground: "hsl(var(--accent-foreground) / <alpha-value>)"
                },
                popover: {
                    DEFAULT: "hsl(var(--popover) / <alpha-value>)",
                    foreground: "hsl(var(--popover-foreground) / <alpha-value>)"
                },
                card: {
                    DEFAULT: "hsl(var(--card) / <alpha-value>)",
                    foreground: "hsl(var(--card-foreground) / <alpha-value>)"
                },
                'ruby': {
                    '50': '#fef1f8',
                    '100': '#fde6f3',
                    '200': '#fecce8',
                    '300': '#fea3d4',
                    '400': '#fc6ab6',
                    '500': '#f63e99',
                    '600': '#e61c76',
                    '700': '#c80e5b',
                    '800': '#a90f4d',
                    '900': '#8a1142',
                    '950': '#550223',
                },
                'night': {
                    '50': '#f6f7f9',
                    '100': '#eceff2',
                    '200': '#d4dbe3',
                    '300': '#afbcca',
                    '400': '#8499ac',
                    '500': '#647c93',
                    '600': '#506479',
                    '700': '#415163',
                    '800': '#394653',
                    '900': '#333d47',
                    '950': '#111418',
                },
                'smoke': {
                    '50': '#f3f6f8',
                    '100': '#e1eaec',
                    '200': '#c7d6da',
                    '300': '#a0b8c0',
                    '400': '#72939e',
                    '500': '#577883',
                    '600': '#4a6470',
                    '700': '#41545d',
                    '800': '#3b484f',
                    '900': '#343e45',
                    '950': '#0c0f11',
                },
            },
            borderRadius: {
                lg: "var(--radius)",
                md: "calc(var(--radius) - 2px)",
                sm: "calc(var(--radius) - 4px)"
            },
            fontFamily: {
                sans: [...fontFamily.sans]
            }
        }
    },
};
