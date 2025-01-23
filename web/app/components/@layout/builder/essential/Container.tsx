import clsx from "clsx"

interface AsideProps {
    w?: string
    h?: string
    flexHorizontal?: boolean
    grow?: boolean
    children?: React.ReactNode[] | React.ReactNode
    className?: string
}

export function Container({ w, h, flexHorizontal, grow, children, className }: AsideProps) {
    return (
        <section
            style={{
                height: h,
                minHeight: h,
                width: w ?? "100%"
            }}
            className={clsx(
                "flex gap-y-3 gap-x-1 overflow-y-auto",
                className,
                grow && "flex-1",
                flexHorizontal ? "flex-row" : "flex-col"
            )}
        >
            {children}
        </section>
    )
}
