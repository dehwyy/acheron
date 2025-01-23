import { Card } from "@nextui-org/react"
import clsx from "clsx"
import { useMemo } from "react"

export type Variant = "gradient" | "gradientPurple" | "unstyled"
type Shadow = "sm" | "md" | "lg" | "none"

const defaultVariant: Variant = "unstyled"

function newVariant(className: string, shadow?: Shadow, isBlured?: boolean) {
    return {
        className,
        shadow,
        isBlured
    }
}

const variants = {
    gradient: newVariant("bg-background/60 dark:bg-gradient-to-br from-default-100/30 to-default-100/10", "sm", true),
    gradientPurple: newVariant("bg-purple-600/60 dark:bg-gradient-to-br from-default-100/30 to-purple-400/30", "sm", true),
    unstyled: newVariant("bg-transparent shadow-none")
} as {
    [key in Variant]: ReturnType<typeof newVariant>
}

interface CardPresetedProps {
    children: React.ReactNode
    className?: string
    variant?: Variant
}

export function CardPreseted({ children, className, variant }: CardPresetedProps) {
    const v = useMemo(() => {
        return variants[variant || defaultVariant]
    }, [variant])

    return (
        <Card isBlurred={v.isBlured} shadow={v.shadow} className={clsx(className, v.className)}>
            {children}
        </Card>
    )
}
