"use client"
import { AppShell, Container } from "$layout/essential"

export function AuthLayout({ children }: { children: React.ReactNode }) {
    return (
        <AppShell withHeader>
            <Container className="justify-center">
                {children}
            </Container>
        </AppShell>
    )
}
