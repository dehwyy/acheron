"use client"
import type { PageProps } from "@/types"
import { Box, Container } from "@/components/@layout/builder/essential"
import { RTCVideoPlayer } from "@/components/@layout/streaming/rtc/player"
import { IconChevronUp } from "@/components/icons/ChevronUp"
import { IconPlus } from "@/components/icons/Plus"
import { IconTrash } from "@/components/icons/Trash"
import { Button, Dropdown, DropdownItem, DropdownMenu, DropdownTrigger, Image, Link } from "@heroui/react"
import clsx from "clsx"
import { useMemo, useState } from "react"

interface Params {
    username: string
}

const imageSrc = "https://app.requestly.io/delay/3000/https://sun9-25.userapi.com/impg/h4MrpyrOltsN8xlZVgist3rhcfiR4wCMXSycpw/sHpyNf87uW8.jpg?size=960x384&quality=95&crop=0,25,2560,1024&sign=66294a4472243ee76c8ae35c6cf6f385&c_uniq_tag=gGOsuEKM10XUUnb6Pca3LTpyWFJ5h75Z6SPEidTFvcE&type=helpers"
const avatarSrc = "https://app.requestly.io/delay/2000/https://sun9-46.userapi.com/impg/MJT6IVmTL0KWnPrSVYxlOaCxQtCn-dIhr-sN5Q/8IQk5a_8JgQ.jpg?size=1280x1280&quality=95&sign=52c696ca9a1bf4051dc9e038b42ddcaa&type=album"
const userStatus = "<Optional description>"

const userPanelHeight = "64px"
const dropdownActions = [
    { label: "Add friend", action: () => { }, icon: <IconPlus className="h-[20px]" /> },
    { label: "Block", action: () => { }, icon: <IconTrash className="h-[20px]" /> },
]

export default function Page({ params }: PageProps<Params>) {
    const [isDropdownOpen, setIsDropdownOpen] = useState(false)
    const editUrl = useMemo(() => `/edit/${params.username}`, [params.username])

    return (
        <div className="max-w-full" style={{ width: 700 }}>
            <section className="relative w-full">
                <Image isZoomed src={imageSrc} alt="user background" height={250} width={700} className="object-cover object-center" />
                <div className="absolute -bottom-3 w-full z-10">
                    <Container className="overflow-y-visible gap-x-10">
                        <Box variant="gradient" h={userPanelHeight} w="300px" className="overflow-y-visible flex-row items-center">
                            <div className="relative w-[150px] h-[40px]">
                                <div className="absolute left-0 bottom-0 z-20">
                                    <Image src={avatarSrc} alt="user avatar" width={150} height={150} radius="full" className="p-0.5 bg-gradient-to-br from-default-100/30 to-default-100/10 shadow-sm shadow-default-100" />
                                </div>
                            </div>
                            <div>{params.username}</div>
                        </Box>
                        <Box variant="gradient" h={userPanelHeight} grow className="flex-row gap-x-5">
                            <Button variant="shadow" color="primary" className="flex-1" href={editUrl} as={Link}> Edit profile</Button>
                            <Dropdown onOpenChange={state => setIsDropdownOpen(state)} placement="bottom-end" showArrow offset={20} backdrop="transparent">
                                <DropdownTrigger>
                                    <Button variant="shadow" color="danger" endContent={<IconChevronUp className={clsx(isDropdownOpen || "rotate-180", "transition-transform")} />}>More</Button>
                                </DropdownTrigger>
                                <DropdownMenu color="primary" variant="flat">
                                    {dropdownActions.map(({ label, action, icon }) => <DropdownItem key={label} onPress={action} startContent={icon}>{label}</DropdownItem>)}
                                </DropdownMenu>
                            </Dropdown>
                        </Box>
                    </Container>
                    {/* <Box variant="gradient" h="60px" className="flex-row items-center overflow-visible gap-x-5">
                        </div>
                        <div className="grow">
                            <div>
                                {params.username}
                            </div>
                            <div>

                            </div>
                        </div>
                    </Box> */}
                </div>
            </section>
            <section className="h-[1000px]">
                <RTCVideoPlayer streamName={params.username} />
            </section>
        </div>
    )
}
