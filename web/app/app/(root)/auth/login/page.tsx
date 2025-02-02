"use client"
import { Box } from "@/components/@layout/builder/essential"
import { IconDiscord } from "@/components/icons/Discord"
import { IconGoogle } from "@/components/icons/Google"
import { Button, Divider, Form, Input } from "@heroui/react"
import { useState } from "react"

export default function Page() {
    const [isVisiblePassword, setIsVisiblePassword] = useState(false)
    return (
        <Box variant="gradient" w="350px" alignSelf="center" className="pt-7 pb-5 px-7">
            <h1 className="text-center text-2xl font-semibol">Acheron</h1>
            <Form className="flex flex-col gap-y-5 py-5">
                <Input variant="underlined" label="Username" name="username" autoComplete="off" />
                <Input
                    variant="underlined"
                    label="Password"
                    name="password"
                    type="password"
                    // TODO
                    endContent={(
                        <button
                            className="focus:outline-none"
                            type="button"
                            onClick={() => setIsVisiblePassword(v => !v)}
                        >
                            {isVisiblePassword ? <span>💀</span> : <span>😇</span>}
                        </button>
                    )}
                />
                <Button variant="shadow" color="secondary" className="w-full">
                    Log In
                </Button>
            </Form>
            <div className="relative py-3">
                <Divider />
                <p className="absolute top-1/2 -translate-y-1/2 left-1/2 -translate-x-1/2">
                    <span className="text-sm font-common">Or</span>
                </p>
            </div>
            <div className="flex flex-col gap-y-3 py-5">
                <Button variant="shadow" startContent={<IconDiscord className="min-h-[40px] max-h-[40px] min-w-[40px] max-w-[40px] p-2" />}>Continue with Discord</Button>
                <Button variant="shadow" startContent={<IconGoogle className="min-h-[40px] max-h-[40px] min-w-[40px] max-w-[40px] p-2" />}>Continue with Google</Button>
            </div>
        </Box>
    )
}
