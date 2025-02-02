import type { LinkProps } from "@heroui/link"
import { Link } from "@heroui/link"

interface LinkSmoothProps extends LinkProps {
    anchorId: string
    children: React.ReactNode
}

export function LinkSmooth({ anchorId, children, onPress, ...props }: LinkSmoothProps) {
    return (
        <Link
            onPress={(e) => {
                document.getElementById(anchorId)?.scrollIntoView({ behavior: "smooth" })
                onPress?.(e)
            }}
            href={`#${anchorId}`}
            {...props}
        >
            {children}
        </Link>
    )
}
