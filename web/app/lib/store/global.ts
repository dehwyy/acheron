import { atomWithStorage } from "jotai/utils"

enum LocalStorageKey {
    Navbar = "navbar",
    Auth = "auth"
}

export const NavbarAtom = atomWithStorage(LocalStorageKey.Navbar, {
    isExpanded: true
})

export const AuthAtom = atomWithStorage(LocalStorageKey.Auth, {
    userId: null as string | null,
    username: null as string | null,
    userImage: null as string | null
})
