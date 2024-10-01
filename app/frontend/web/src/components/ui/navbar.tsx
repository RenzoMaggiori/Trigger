"use client"
import React from 'react'
import { Button } from './button'
import Link from 'next/link'
import { LogoIcon } from './logoIcon'
import { IoMenu } from "react-icons/io5";
import { Sheet, SheetContent, SheetDescription, SheetHeader, SheetTitle, SheetTrigger } from './sheet'
import { MdAddBox } from 'react-icons/md'
import { GrDocumentImage } from 'react-icons/gr'
import { SiGooglegemini } from 'react-icons/si'
import { IoSettingsOutline } from "react-icons/io5";
import { cn } from '@/lib/utils'

export function Navbar() {
    const [loggedIn, setLoggedIn] = React.useState(false);
    const [isHomePage, setIsHomePage] = React.useState(false);

    React.useEffect(() => {
        setIsHomePage(window.location.pathname === '/home');
    }, [window.location.pathname]);

    const navbarItems = [
        { name: "Home", href: "/home" },
        { name: "Community", href: "/community" },
        { name: "Settings", href: "/settings" },
    ]

    const homeItems = [
        {
            href: "/",
            name: "Create Trigger",
            className: "bg-gradient-to-r from-blue-500 via-violet-500 to-fuchsia-500 hover:bg-gradient-to-r hover:from-blue-600 hover:via-violet-600 hover:to-fuchsia-600 animate-gradient border-0 text-white",
            icon: <MdAddBox className="text-white w-5 h-5" />,
        },
        {
            name: "Templates",
            href: "/",
            icon: <GrDocumentImage className="text-black dark:text-white w-5 h-5" />,
        },
        {
            name: "Triggers",
            href: "/home",
            icon: <SiGooglegemini className="text-black dark:text-white w-5 h-5" />,
        },
        {
            name: "Settings",
            href: "/settings",
            icon: <IoSettingsOutline className="text-black dark:text-white w-5 h-5" />,
        },
    ]

    const authButtons = [
        { name: "Log In", href: "/auth?type=login", className: "rounded-full border-black text-lg", variant: "outline" },
        { name: "Sign Up", href: "/auth?type=signup", className: "rounded-full bg-orange-600 hover:bg-orange-700 text-lg", variant: "default" },
    ]

    return (
        <nav className="flex bg-white border-gray-500 dark:bg-zinc-950 min-h-16">
            <div className="w-full flex flex-nowrap items-center p-4">
                <a href="/" className="flex items-center space-x-3 rtl:space-x-reverse absolute">
                    <LogoIcon className="h-12 w-[200px] dark:fill-white" />
                </a>

                <Sheet>
                    <SheetTrigger className="absolute right-2 md:hidden px-2">
                        <IoMenu className='h-7 w-7' />
                    </SheetTrigger>
                    <SheetContent className='w-full flex flex-col gap-5' side="right">
                        <SheetHeader>
                            <SheetTitle><LogoIcon className="w-1/2 dark:fill-white" /></SheetTitle>
                            <SheetDescription className='text-start'>All reactions have a Trigger</SheetDescription>
                        </SheetHeader>
                        <div className='flex flex-col gap-5'>
                            {isHomePage && homeItems.map((item, key) => (
                                <div key={key}>
                                    <Button asChild className={cn(`flex bg-white border dark:bg-zinc-900 dark:hover:bg-zinc-950 dark:text-white border-zinc-700 hover:bg-zinc-100 text-black items-center justify-center text-xl rounded-full`, item.className)}>
                                        <Link href={item.href} className='gap-x-3'>
                                            {item.icon}
                                            <span className='mx-auto'>{item.name}</span>
                                        </Link>
                                    </Button>
                                </div>
                            ))}
                            {!isHomePage && loggedIn && navbarItems.map((item, key) => (
                                <div key={key}>
                                    <Button asChild variant="outline" className='flex items-center justify-center text-xl rounded-full border-black'>
                                        <Link href={item.href}>
                                            {item.name}
                                        </Link>
                                    </Button>
                                </div>
                            ))}
                        </div>
                        <div className='w-full flex flex-col mt-auto gap-5'>

                            {authButtons.map((item, key) => (
                                <Button key={key} className={item.className} variant={item.variant as "outline" | "default" | "link" | "destructive" | "secondary" | "ghost"} asChild>
                                    <Link href={item.href}>
                                        {item.name}
                                    </Link>
                                </Button>
                            ))}
                        </div>
                    </SheetContent>
                </Sheet>
                <div className="hidden w-full md:block md:w-auto mx-auto">
                    <div className='flex flex-row'>
                        {navbarItems.map((item, key) => (
                            <div key={key}>
                                <Button asChild variant="ghost" className='text-xl'>
                                    <Link href={item.href}>
                                        {item.name}
                                    </Link>
                                </Button>
                            </div>
                        ))}
                    </div>
                </div>
                <div className='absolute gap-x-4 right-6 hidden md:flex'>
                    {authButtons.map((item, key) => (
                        <Button key={key} className={item.className} variant={item.variant as "outline" | "default" | "link" | "destructive" | "secondary" | "ghost"} asChild>
                            <Link href={item.href}>
                                {item.name}
                            </Link>
                        </Button>
                    ))}
                </div>
            </div>
        </nav>
    )
}
