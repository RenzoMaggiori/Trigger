import React from 'react'
import { LogoIcon } from './LogoIcon'
import { NavigationMenu, NavigationMenuContent, NavigationMenuItem, NavigationMenuLink, NavigationMenuList, NavigationMenuTrigger, navigationMenuTriggerStyle } from "@/components/ui/navigation-menu"
import Link from 'next/link'
import { cn } from '@/lib/utils'

const ProductContent = () => {
    const listItems = [
        {title: "Introduction", href: "/docs", content: "Discover how Trigger streamlines app integrations and automates workflows with ease. Get started quickly!"},
        {title: "Getting Started", href: "/docs/installation", content: "Follow this guide to set up Trigger and connect your first API integration in just a few steps."},
        {title: "API Connections", href: "/docs/connections", content: "Understand the different types of API connections you can create and how to manage them effectively."},
    ]

    return (
        <ul className="grid gap-3 p-6 md:w-[400px] lg:w-[500px] lg:grid-cols-[.75fr_1fr]">
            <li className="row-span-3">
                <NavigationMenuLink asChild>
                    <a
                        className="flex h-full w-full select-none flex-col justify-end rounded-md bg-gradient-to-b from-muted/50 to-muted p-6 no-underline outline-none focus:shadow-md"
                        href="/"
                    >
                        <LogoIcon className="h-12 w-[200px] dark:fill-white" />
                        <div className="mb-2 mt-4 text-lg font-medium">
                        </div>
                        <p className="text-sm leading-tight text-muted-foreground">
                            Effortlessly connect your favorite apps with Trigger. Create seamless workflows that automate tasks and simplify processes. User-friendly. Flexible. Open Source.
                        </p>
                    </a>
                </NavigationMenuLink>
            </li>
            {listItems.map((item, key) => (
                <ListItem key={key} href={item.href} title={item.title}>
                    {item.content}
                </ListItem>
            ))}
        </ul>
    );
}

const ConnectionsContent = () => {
    const components = [
        { title: "Email", href: "/email", description: "Integrate your email service to automate personalized email notifications and responses." },
        { title: "Slack", href: "/slack", description: "Connect Slack to streamline team communication with automated alerts and updates." },
        { title: "Google Sheets", href: "/sheets", description: "Link Google Sheets to automate data entry and synchronize information across your workflows." },
        { title: "Trello", href: "/trello", description: "Integrate Trello boards to manage tasks effortlessly and automate project updates." },
        { title: "Webhook", href: "/webhook", description: "Set up webhooks to receive real-time data from various sources and trigger workflows instantly." },
        { title: "Calendar", href: "/calendar", description: "Connect your calendar to automate event scheduling and receive reminders for important dates." }
    ];
    
    return (
        <ul className="grid w-[400px] gap-3 p-4 md:w-[500px] md:grid-cols-2 lg:w-[600px] ">
              {components.map((component) => (
                <ListItem
                  key={component.title}
                  title={component.title}
                  href={component.href}
                >
                  {component.description}
                </ListItem>
              ))}
            </ul>
    )
}

export const Navbar = () => {

    const navbarItems = [
        { name: "Home", href: "/", type: "link" },
        { name: "Product", content: <ProductContent /> },
        { name: "Connections", content: <ConnectionsContent />},
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
                        <NavigationMenu>
                            <NavigationMenuList>
                                {navbarItems.map((item, key) => (
                                    <div key={key}>
                                        {item.type === "link" ? (
                                            <Link href={item.href} legacyBehavior passHref>
                                                <NavigationMenuLink className={`${navigationMenuTriggerStyle()} text-xl text-black dark:text-white`}                                                >
                                                    {item.name}
                                                </NavigationMenuLink>
                                            </Link>
                                        ) : (
                                            <NavigationMenuItem>
                                                <NavigationMenuTrigger className='text-black dark:text-white text-xl'>{item.name}</NavigationMenuTrigger>
                                                <NavigationMenuContent>
                                                    {item.content}
                                                </NavigationMenuContent>
                                            </NavigationMenuItem>
                                        )}
                                    </div>
                                ))}
                            </NavigationMenuList>
                        </NavigationMenu>
                    </div>
                </div>
            </div>
        </nav>
    )
}

const ListItem = React.forwardRef<
    React.ElementRef<"a">,
    React.ComponentPropsWithoutRef<"a">
>(({ className, title, children, ...props }, ref) => {
    return (
        <li>
            <NavigationMenuLink asChild>
                <a
                    ref={ref}
                    className={cn(
                        "block select-none space-y-1 rounded-md p-3 leading-none no-underline outline-none transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground",
                        className
                    )}
                    {...props}
                >
                    <div className="text-sm font-medium leading-none">{title}</div>
                    <p className="line-clamp-2 text-sm leading-snug text-muted-foreground">
                        {children}
                    </p>
                </a>
            </NavigationMenuLink>
        </li>
    )
})
ListItem.displayName = "ListItem"
