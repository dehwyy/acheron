import { IconChevronUp } from "@/components/icons/ChevronUp"
import { NavbarAtom } from "@/lib/store/global"
import { Button } from "@heroui/button"
import clsx from "clsx"
import { useAtom } from "jotai"
import { useCallback } from "react"

export function VisibilityToggler() {
    const [{ isExpanded }, setStore] = useAtom(NavbarAtom)

    const toggleNavbar = useCallback(() => {
        setStore(v => ({ ...v, isExpanded: !v.isExpanded }))
    }, [])

    return (
        <Button
            onPress={toggleNavbar}
            isIconOnly
            disableAnimation
            className="h-[48px] px-10 items-end bg-transparent focus-visible:outline-none outline-none"
        >
            <IconChevronUp className={clsx(isExpanded || "rotate-180", "transition-all stroke-default-300 min-w-[16px]")} />
        </Button>
    )
}
