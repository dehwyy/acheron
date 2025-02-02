import { AppShell, Box, Container } from "$layout/essential"

interface Props {
    children: React.ReactNode
}

export function StreamingLayout({ children }: Props) {
    return (
        <AppShell withHeader>
            <Container w="200px" flexHorizontal>
                <Box variant="gradient" w="100%">
                    Some sheesh
                </Box>
            </Container>
            <Container grow className="pr-3">
                <div className="flex flex-1 items-start flex-col">
                    <div className="rounded-sm overflow-hidden min-w-full">
                        {children}
                    </div>
                </div>
                <Box grow variant="gradient">Some sheesh 3</Box>
            </Container>
            <Container w="140px" flexHorizontal>
                <Box variant="gradient" w="100%">
                    Some sheesh
                </Box>
            </Container>
        </AppShell>
    )
}
