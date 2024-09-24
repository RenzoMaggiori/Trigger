"use client"
import React, { useState } from 'react'
import { Button } from './button'
import Link from 'next/link'
import { LogoIcon } from './logoIcon'
import { IoMenu } from "react-icons/io5";

export function Navbar() {
    const [loggedIn, setLoggedIn] = useState(false);

    const navbarItems = [
        { name: "Home", href: "/home" },
        { name: "Community", href: "/community" },
        { name: "Triggers", href: "/triggers" },
    ]

    return (
        <nav className="flex bg-white border-gray-500 dark:bg-zinc-950 h-16">
            <div className="w-full flex flex-nowrap items-center p-4">
                <a href="/" className="flex items-center space-x-3 rtl:space-x-reverse absolute">
                    <LogoIcon className="h-12 w-[200px] dark:fill-white" />
                </a>
                <Button type='button' className="absolute right-2 md:hidden px-2" variant="ghost">
                    <IoMenu className='h-7 w-7'/>
                </Button>
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
                    <Button className='rounded-full border-black text-lg' variant="outline">
                        Log In
                    </Button>
                    <Button className='rounded-full bg-orange-600 hover:bg-orange-700 text-lg'>
                        Sign Up
                    </Button>
                </div>
            </div>
        </nav>
    )
}
