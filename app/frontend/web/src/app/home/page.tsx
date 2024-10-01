"use client"
import { Sidebar, SidebarBody, SidebarLink } from '@/components/ui/sidebar'
import React from 'react'
import Image from 'next/image'
import { IconArrowLeft, IconBrandTabler, IconSettings, IconUserBolt } from '@tabler/icons-react'
import { MdAddBox } from "react-icons/md";
import { GrDocumentImage } from "react-icons/gr";
import { Button } from '@/components/ui/button'
import Link from 'next/link'
import { Input } from '@/components/ui/input'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { ScrollArea } from '@/components/ui/scroll-area'
import { FocusCards } from '@/components/ui/focus-cards'

const page = () => {
    const [open, setOpen] = React.useState(false);
    const links = [
        {
            href: "/",
            name: "New Trigger",
            icon: <MdAddBox className="text-neutral-700 dark:text-neutral-200 flex-shrink-0 mr-2" />,
        },
        {
            name: "Templates",
            href: "/",
            icon: <GrDocumentImage className="text-neutral-700 dark:text-neutral-200 flex-shrink-0 mr-2" />,
        },
        {
            name: "Settings",
            href: "#",
            icon: (
                <IconSettings className="text-neutral-700 dark:text-neutral-200 h-5 w-5 flex-shrink-0" />
            ),
        },
        {
            label: "Logout",
            href: "#",
            icon: (
                <IconArrowLeft className="text-neutral-700 dark:text-neutral-200 h-5 w-5 flex-shrink-0" />
            ),
        },
    ];

    const triggers = [
        { title: "Card one", href: "/", img: "https://fakeimg.pl/300x200" },
        { title: "Card one", href: "/", img: "https://fakeimg.pl/300x200" },
        { title: "Card one", href: "/", img: "https://fakeimg.pl/300x200" },
        { title: "Card one", href: "/", img: "https://fakeimg.pl/300x200" },
        { title: "Card one", href: "/", img: "https://fakeimg.pl/300x200" },
        { title: "Card one", href: "/", img: "https://fakeimg.pl/300x200" },
        { title: "Card one", href: "/", img: "https://fakeimg.pl/300x200" },
        { title: "Card one", href: "/", img: "https://fakeimg.pl/300x200" },
        { title: "Card one", href: "/", img: "https://fakeimg.pl/300x200" },
        { title: "Card one", href: "/", img: "https://fakeimg.pl/300x200" },
        { title: "Card one", href: "/", img: "https://fakeimg.pl/300x200" },
        { title: "Card one", href: "/", img: "https://fakeimg.pl/300x200" },
        { title: "Card one", href: "/", img: "https://fakeimg.pl/300x200" },
        { title: "Card one", href: "/", img: "https://fakeimg.pl/300x200" },
        { title: "Card one", href: "/", img: "https://fakeimg.pl/300x200" },
        { title: "Card one", href: "/", img: "https://fakeimg.pl/300x200" },
    ]
    React.useEffect(() => {
        const handleResize = () => {
            if (window.innerWidth >= 768) {
                setOpen(true);
            } else {
                setOpen(false)
            }
        };

        handleResize();
        window.addEventListener('resize', handleResize);

        return () => window.removeEventListener('resize', handleResize);
    }, [setOpen]);
    return (
        <div className='flex h-screen w-full overflow-hidden'>
            <Sidebar open={open}>
                <SidebarBody className='bg-white'>
                    <div className="mt-8 flex flex-col gap-2">
                        {links.map((link, idx) => (
                            <Button key={idx} variant="ghost" className='flex items-center justify-start text-center text-xl' asChild>
                                <Link href={link.href}>
                                    {link.icon}
                                    {link.name}
                                </Link>
                            </Button>
                        ))}
                    </div>

                </SidebarBody>
            </Sidebar>
            <div className='flex flex-col w-full'>
                <p className='text-3xl font-bold p-5'>Your Triggers</p>

                <ScrollArea>
                    <div className='flex flex-row flex-wrap gap-4 p-5 overflow-x-auto'>
                        {triggers.map((trigger, index) => (
                            <div key={index}>
                                <Link href={trigger.href}>
                                    <Card className='flex flex-col bg-zinc-100 shadow-md rounded-lg w-[200px]' key={index}>
                                        <CardHeader className='p-4 border-b'>
                                            <CardTitle className='text-xl font-bold'>
                                                <img src={trigger.img} alt={trigger.title} />
                                            </CardTitle>
                                        </CardHeader>
                                    </Card>
                                    <p className='font-bold text-md text-start p-1'>{trigger.title}</p>
                                </Link>
                            </div>
                        ))}
                    </div>
                </ScrollArea>
            </div>
        </div>
    )
}

export default page