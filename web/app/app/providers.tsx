"use client"

import type { ThemeProviderProps } from "next-themes"
import { HeroUIProvider } from "@heroui/system"
import { ThemeProvider as NextThemesProvider } from "next-themes"
import { useRouter } from "next/navigation"

interface ProvidersProps {
    children: React.ReactNode
    themeProps?: ThemeProviderProps
}

export function Providers({ children, themeProps }: ProvidersProps) {
    const router = useRouter()

    return (
        <HeroUIProvider navigate={router.push}>
            <NextThemesProvider {...themeProps}>{children}</NextThemesProvider>
        </HeroUIProvider>
    )
}
