import { AuthAtom } from "@/lib/store/global"
import { Avatar, Dropdown, DropdownItem, DropdownMenu, DropdownTrigger } from "@heroui/react"
import { useAtomValue } from "jotai"
import { useMemo } from "react"

export function Navigation() {
    const { userImage, username, userId } = useAtomValue(AuthAtom)

    const dropdownItems = useMemo(() => {
        if (!userId) {
            return [
                { label: "Sign in", href: "/auth/login" },
                { label: "Development", href: "/dev" },
            ]
        }

        return [
            { label: "Profile", href: `/profile/${username}` },
            { label: "Settings", href: "/settings" },
            { label: "Following", href: `/profile/${username}/following` },
            { label: "Logout", href: "/auth/logout" }
        ]
    }, [userId, username])

    return (
        <Dropdown>
            <DropdownTrigger>
                <Avatar

                    isBordered
                    as="button"
                    className="select-none"
                    src={userImage!} // || Dev.Img
                />
            </DropdownTrigger>
            <DropdownMenu>
                {dropdownItems.map(v => (
                    <DropdownItem
                        key={v.label}
                        href={v.href}
                        variant="faded"
                        className="py-2"
                    >
                        {v.label}
                    </DropdownItem>
                ))}
            </DropdownMenu>
        </Dropdown>
    )
}
