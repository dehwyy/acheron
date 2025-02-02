import type { SVGProps } from "react"

export type IconSvgProps = SVGProps<SVGSVGElement> & {
    size?: number
}

export interface IconProps {
    className?: string
}

export interface LayoutProps<T = any> {
    children: React.ReactNode
}

export interface PageProps<P> {
    params: P
}
