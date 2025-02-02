"use client"
import { AppShell, Container } from "$layout/essential"

interface Props {
    children: React.ReactNode
}

export function UserLayout({ children }: Props) {
    return (
        <AppShell withHeader>
            <Container className="justify-center">
                {children}
            </Container>
        </AppShell>
    )
}
