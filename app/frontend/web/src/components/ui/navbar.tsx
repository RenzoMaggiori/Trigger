"use client"
import React from 'react'
import { Button } from './button'
import Link from 'next/link'
import { LogoIcon } from './logoIcon'
import { IoMenu } from "react-icons/io5";
import { Sheet, SheetContent, SheetDescription, SheetHeader, SheetTitle, SheetTrigger } from './sheet'

export function Navbar() {
    const [loggedIn, setLoggedIn] = React.useState(false);

    const navbarItems = [
        { name: "Home", href: "/home" },
        { name: "Community", href: "/community" },
        { name: "Settings", href: "/settings" },
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
                            {navbarItems.map((item, key) => (
                                <div key={key}>
                                    <Button asChild variant="outline" className='flex items-center justify-center text-xl rounded-full border-black'>
                                        <Link href={item.href}>
                                            {item.name}
                                        </Link>
                                    </Button>
                                </div>
                            ))}
                        </div>
                        {authButtons.map((item, key) => (
                            <Button key={key} className={item.className} variant={item.variant as "outline" | "default" | "link" | "destructive" | "secondary" | "ghost"} asChild>
                                <Link href={item.href}>
                                    {item.name}
                                </Link>
                            </Button>
                        ))}
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
