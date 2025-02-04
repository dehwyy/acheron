import AppHeader from "@/components/@layout/builder/essential/AppHeader"
import { Hoshi } from "./background/hoshi"

interface AppShellProps {
    children: React.ReactNode
    withHeader?: boolean
}

export function AppShell({ children, withHeader }: AppShellProps) {
    return (
        <Hoshi>
            <div className="flex justify-center min-h-screen">
                <div className="flex flex-col w-full relative">
                    {withHeader && <AppHeader />}
                    <main className="flex-1 overflow-hidden">
                        <div className="w-full flex gap-x-3 px-5 my-3 transition-all h-screen max-h-full">{children}</div>
                    </main>
                </div>
            </div>
        </Hoshi>
    )
}
