"use client"

import { Logo, Navigation, VisibilityToggler } from "$layout/essential/navbar"
import { TransformTranslate } from "@/lib/const"
import { NavbarAtom } from "@/lib/store/global"
import {
    Navbar as HeroUINavbar,
    NavbarBrand,
    NavbarContent,
    NavbarItem
} from "@heroui/react"

import clsx from "clsx"
import { useAtomValue } from "jotai"
import { useMemo } from "react"

export default function Navbar() {
    const { isExpanded } = useAtomValue(NavbarAtom)

    const shift = useMemo(() => {
        return `${isExpanded ? 0 : -TransformTranslate.NavbarHide}px`
    }, [isExpanded])

    return (
        <HeroUINavbar
            shouldHideOnScroll
            isBordered
            maxWidth="full"
            style={{
                translate: `0 ${shift}`,
                marginTop: shift
            }}
            className="transition-all"
        >
            <div className="absolute left-1/2 -translate-x-1/2 -bottom-3.5">
                <VisibilityToggler />
            </div>
            <NavbarBrand className={clsx(isExpanded || "-translate-y-4", "transition-all")}>
                <Logo />
            </NavbarBrand>
            <NavbarContent
                justify="end"
                className={clsx(isExpanded || "-translate-y-4", "transition-all")}
            >
                <NavbarItem>
                    <Navigation />
                </NavbarItem>
            </NavbarContent>
        </HeroUINavbar>
    )
}
