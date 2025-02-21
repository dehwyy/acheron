"use client"
import type { PageProps } from "@/types"
import { Box } from "@/components/@layout/builder/essential"
import { Button } from "@heroui/button"
import { Input } from "@heroui/input"
import { Snippet } from "@heroui/snippet"

interface Params {
    username: string
}

const token = "eufghf=31dmaml"
export default function Page({ params }: PageProps<Params>) {
    return (
        <Box w="700px" className="gap-y-5">
            <Box variant="gradient" className="px-7 py-5 gap-y-5">
                <h2 className="text-center font-bold text-xl">General</h2>
                <Field label="Nickname">
                    <Input value={params.username} className="flex-1" variant="bordered" type="text" color="secondary" />
                </Field>
                <Field label="User image">
                    <Input type="file" variant="bordered" color="secondary" className="flex-1" />
                </Field>
                <Field label="Profile background">
                    <Input type="file" variant="bordered" color="secondary" className="flex-1" />
                </Field>
                <Field label="Stream token">
                    <Snippet symbol="" codeString={token} color="secondary" disableTooltip>{`<Click to copy>`}</Snippet>
                </Field>
                <div className="ml-auto flex gap-x-3 pt-5">
                    <Button color="danger" variant="shadow" className="px-10">Discard</Button>
                    <Button color="secondary" variant="shadow" className="px-10">Apply</Button>
                </div>
            </Box>
            <Box variant="gradient" className="px-7 py-5 gap-y-5">
                <h2 className="text-center font-bold text-xl">Integrations</h2>

            </Box>
        </Box>
    )
}

function Field({ children, label }: { children: React.ReactNode, label: string }) {
    return (
        <div className="flex items-center gap-x-7 w-full">
            <p className="w-[150px]">
                {`${label}: `}
            </p>
            {children}
        </div>
    )
}
