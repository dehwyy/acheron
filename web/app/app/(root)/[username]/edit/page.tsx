"use client"
import type { PageProps } from "@/types"
import { Box } from "@/components/@layout/builder/essential"
import { ButtonFancy } from "@/components/reusable/ButtonFancy"

interface Params {
    username: string
}

export default function Page({ params }: PageProps<Params>) {
    return (
        <Box w="700px">
            <Box variant="gradient" className="flex-row items-center gap-x-7 w-full">
                <p>Nickname: </p>
                <p className="flex-1">{params.username}</p>
                <ButtonFancy>
                    Edit
                </ButtonFancy>
            </Box>
        </Box>
    )
}
