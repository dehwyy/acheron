import type { ButtonProps, TooltipProps } from "@heroui/react"
import { useHover } from "@custom-react-hooks/all"
import { Button, Tooltip } from "@heroui/react"
import clsx from "clsx"
import { useState } from "react"

interface ButtonTooltipedProps extends ButtonProps {
    children: React.ReactNode
    tooltip: TooltipProps

    className?: string
}

export function ButtonTooltiped({ children, className, ...props }: ButtonTooltipedProps) {
    const [isTooltipOpen, setTooltipOpen] = useState(false)
    const { isHovered, ref } = useHover<HTMLButtonElement>()

    return (
        <Tooltip
            isDisabled={isTooltipOpen}
            placement={props.tooltip.placement || "bottom"}
            delay={200}
            closeDelay={0}
            showArrow={props.tooltip.showArrow}
            content={props.tooltip.content}
            offset={props.tooltip.offset}
            className={props.tooltip.className}
        >
            <Button
                ref={ref}
                aria-expanded={props.isIconOnly}
                disableAnimation={props.isIconOnly}
                disableRipple={props.isIconOnly}
                className={clsx(

                    className,
                    props.isIconOnly
                        ? `bg-transparent hover:!opacity-100${isHovered || isTooltipOpen ? " stroke-default-800" : " stroke-default-500"}`
                        : "",
                )}
                {...props}
            >
                {children}
            </Button>
        </Tooltip>
    )
}
