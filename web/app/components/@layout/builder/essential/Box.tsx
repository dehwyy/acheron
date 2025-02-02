import type { Variant } from "@/components/reusable/CardPreseted"
import { CardPreseted } from "@/components/reusable/CardPreseted"
import { CardBody, ScrollShadow } from "@heroui/react"
import clsx from "clsx"

interface WrapperProps {
    children: React.ReactNode
    h?: string
    w?: string
    scrollable?: boolean
    grow?: boolean
    alignSelf?: string
}

interface BoxProps extends WrapperProps {
    className?: string
    variant?: Variant
}

export function Box({ children, className, ...props }: BoxProps) {
    return (
        <Wrapper {...props}>
            <CardPreseted variant={props.variant} className="min-h-full overflow-visible">
                <CardBody className={clsx("flex flex-col gap-y-2 overflow-y-auto", className)}>{children}</CardBody>
            </CardPreseted>
        </Wrapper>
    )
}

function Wrapper({ children, w, h, scrollable, grow, alignSelf }: WrapperProps) {
    return scrollable
        ? (
                <ScrollShadow style={{ height: h ?? "100%" }} className={clsx("pr-1", grow && "flex-1")}>
                    {children}
                </ScrollShadow>
            )
        : (
                <div
                    style={{ height: h, minHeight: h, maxHeight: h, width: w, alignSelf }}
                    className={clsx(grow && "flex-1")}
                >
                    {children}
                </div>
            )
}
