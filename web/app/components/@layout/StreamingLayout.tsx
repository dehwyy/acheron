import { AppShell, Box, Container } from "$layout/essential"

interface Props {
    children: React.ReactNode
}

export function StreamingLayout({ children }: Props) {
    return (
        <AppShell withHeader>
            <Container grow>
                {children}
            </Container>
            <Container w="200px" flexHorizontal>
                <Box variant="gradient">
                    Some sheesh
                </Box>
            </Container>
        </AppShell>
    )
}
