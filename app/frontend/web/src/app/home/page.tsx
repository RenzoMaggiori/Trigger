"use client"
import React from 'react'
import { MdAddBox } from "react-icons/md";
import { GrDocumentImage } from "react-icons/gr";
import { Button } from '@/components/ui/button'
import Link from 'next/link'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { ScrollArea } from '@/components/ui/scroll-area'
import { LogoIcon } from '@/components/ui/logoIcon'
import { SiGooglegemini } from "react-icons/si";

const page = () => {
    const links = [
        {
            href: "/",
            name: "Create Trigger",
            className: "bg-gradient-to-r from-blue-500 via-violet-500 to-fuchsia-500 animate-gradient text-white",
            icon: <MdAddBox className="text-white flex-shrink-0 mr-2" />,
        },
        {
            name: "Templates",
            href: "/",
            icon: <GrDocumentImage className="text-neutral-700 dark:text-neutral-200 flex-shrink-0 mr-2" />,
        },
        {
            name: "Triggers",
            href: "/",
            icon: <SiGooglegemini className="text-neutral-700 dark:text-neutral-200 flex-shrink-0 mr-2" />,
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


    return (
        <div className='flex h-screen w-full overflow-hidden'>
            <div className="flex flex-col h-full p-7">
                <div className="p-4">
                    <h2 className="text-2xl font-bold text-gray-800 dark:text-white mb-4">Dashboard</h2>
                </div>
                <div className='flex flex-col items-center justify-center gap-3'>
                    {links.map((item, key) => (
                        <Button key={key} className={`bg-white hover:bg-zinc-100 text-black ${item.className} w-full justify-start rounded-md`} asChild>
                            <Link href={item.href}>
                                {item.icon}
                                <p className='text-xl'>{item.name}</p>
                            </Link>
                        </Button>
                    ))}
                </div>
                <div className="mt-auto p-4">
                    <Button variant="ghost" className="w-full rounded-md justify-start text-red-600 hover:text-red-700 hover:bg-red-100 dark:hover:bg-red-900">
                        <p>Logout</p>
                    </Button>
                </div>
            </div>
            <div className='flex flex-col w-full p-5 overflow-x-auto'>
                <Card className='py-6'>
                    <CardContent>
                        <ScrollArea>
                        <div className='px-8 py-4 w-full'>

                            <Card className='w-full h-[200px] bg-gradient-to-tr from-blue-500 via-violet-500 to-fuchsia-500 p-5 rounded-lg shadow-lg animate-gradient'>
                                <CardContent className='flex flex-col h-full items-center justify-center text-white text-center font-bold text-lg lg:text-3xl gap-y-3'>
                                    <LogoIcon className='w-[150px] fill-yellow-500' />
                                    <p>Try Trigger for 30 days free</p>
                                    <Button className='bg-zinc-200 text-black p-5 hover:bg-zinc-100'>Start free trial</Button>
                                </CardContent>
                            </Card>
                        </div>

                        <p className='text-3xl font-bold p-5'>Your Triggers</p>

                            <div className='flex flex-row flex-wrap gap-4 p-5 items-center justify-center'>
                                {triggers.concat(triggers).map((trigger, index) => (
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
                    </CardContent>
                </Card>
            </div>
        </div>
    )
}

export default page