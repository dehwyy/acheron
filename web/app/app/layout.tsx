import type { Metadata } from "next"

import { fontCommon, fontSacramento, fontSans } from "@/config/fonts"
import { siteConfig } from "@/config/site"

import clsx from "clsx"
import { Providers } from "./providers"
import "@/styles/globals.css"

export const metadata: Metadata = {
    title: {
        default: siteConfig.name,
        template: `%s - ${siteConfig.name}`
    },
    description: siteConfig.description,
    icons: {
        icon: "/favicon.ico"
    }
}

export default function RootLayout({ children }: { children: React.ReactNode }) {
    return (
        <html
            suppressHydrationWarning
            lang="en"
        >
            <head />
            <body
                className={clsx(
                    "min-h-screen bg-background font-sans antialiased",
                    fontSans.variable,
                    fontSacramento.variable,
                    fontCommon.variable
                )}
            >
                <Providers themeProps={{ attribute: "class", defaultTheme: "dark", enableSystem: true }}>
                    <div className="relative flex flex-col h-screen font-common">{children}</div>
                </Providers>
            </body>
        </html>
    )
}
