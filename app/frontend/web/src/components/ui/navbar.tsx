"use client"
import React, { useState } from 'react'
import { Button } from './button'
import Link from 'next/link'
import { LogoIcon } from './logoIcon'

export function Navbar() {
    const [loggedIn, setLoggedIn] = useState(false);

    const navbarItems = [
        { name: "Home", href: "/home" },
        { name: "Community", href: "/community" },
        { name: "Triggers", href: "/triggers" },
    ]

    return (
        <nav className="flex bg-white border-gray-500 dark:bg-zinc-950">
            <div className="w-full flex flex-nowrap items-center p-4">
                <a href="/" className="flex items-center space-x-3 rtl:space-x-reverse absolute">
                    <LogoIcon className="h-12 w-[200px] dark:fill-white" />
                </a>
                <button data-collapse-toggle="navbar-default" type="button" className="inline-flex items-center p-2 w-10 h-10 justify-center text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600" aria-controls="navbar-default" aria-expanded="false">
                    <span className="sr-only">Open main menu</span>
                    <svg className="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 17 14">
                        <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M1 1h15M1 7h15M1 13h15" />
                    </svg>
                </button>
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
            </div>
        </nav>
    )
}
