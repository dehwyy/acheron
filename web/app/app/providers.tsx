"use client"

import type { ThemeProviderProps } from "next-themes"
import { NextUIProvider } from "@nextui-org/system"
import { ThemeProvider as NextThemesProvider } from "next-themes"
import { useRouter } from "next/navigation"

interface ProvidersProps {
    children: React.ReactNode
    themeProps?: ThemeProviderProps
}

export function Providers({ children, themeProps }: ProvidersProps) {
    const router = useRouter()

    return (
        <NextUIProvider navigate={router.push}>
            <NextThemesProvider {...themeProps}>{children}</NextThemesProvider>
        </NextUIProvider>
    )
}
