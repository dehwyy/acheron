import AppHeader from "@/components/@layout/builder/essential/AppHeader"

interface AppShellProps {
    children: React.ReactNode
    withHeader?: boolean
}

export default function AppShell({ children, withHeader }: AppShellProps) {
    return (
        <div className="flex justify-center min-h-screen">
            <div className="flex flex-col w-full relative">
                {withHeader && <AppHeader />}
                <main className="flex-1 overflow-hidden">
                    <div className="w-full flex gap-x-3 px-5 py-3 transition-all h-screen max-h-full">{children}</div>
                </main>
            </div>
        </div>
    )
}
