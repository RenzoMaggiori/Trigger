"use client"
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'
import { FcGoogle } from "react-icons/fc";
import { FiGithub } from "react-icons/fi";
import { PiMicrosoftOutlookLogo } from "react-icons/pi";

import React from 'react'
import { useSearchParams } from 'next/navigation'

const page = () => {
    // TODO: add href to the api for the tab buttons
    // TODO: add form for email login and sign up
    const searchParams = useSearchParams()
    const type = searchParams.get('type')
    const defaultValue = type === 'signup' ? 'Sign Up' : 'Log In';
    
    const tabs = [
        {
            trigger: "Log In",
            fields: [{ name: "Email", label: "example@example.com" }, { name: "Password", label: "password" }],
            buttons: [
                { icon: <FcGoogle className='mr-2 ' />, name: "Continue with Google", variant: "outline" },
                { icon: <FiGithub className='mr-2 ' />, name: "Continue with Github", variant: "outline", className: "bg-zinc-800 text-white hover:bg-zinc-950 hover:text-white" },
                { icon: <PiMicrosoftOutlookLogo className='mr-2 ' />, name: "Continue with Outlook", variant: "outline", className: "bg-sky-600 text-white hover:bg-sky-700 hover:text-white" },
            ]
        },
        {
            trigger: "Sign Up",
            fields: [{ name: "Email", label: "example@example.com" }, { name: "Password", label: "password" }],
            buttons: [
                { icon: <FcGoogle className='mr-2 ' />, name: "Start with Google", variant: "outline" },
                { icon: <FiGithub className='mr-2 ' />, name: "Start with Github", variant: "outline", className: "bg-zinc-800 text-white hover:bg-zinc-950 hover:text-white" },
                { icon: <PiMicrosoftOutlookLogo className='mr-2 ' />, name: "Start with Outlook", variant: "outline", className: "bg-sky-600 text-white hover:bg-sky-700 hover:text-white" },
            ]
        },
    ]
    return (
        <div className='flex justify-center items-center h-full'>
            <Tabs defaultValue={defaultValue} className="w-3/4 md:w-1/2 lg:w-1/3">
                <TabsList className="grid w-full grid-cols-2">
                    {tabs.map((item, index) => (
                        <TabsTrigger key={index} value={item.trigger}>{item.trigger}</TabsTrigger>
                    ))}
                </TabsList>
                {tabs.map((item, index) => (
                    <TabsContent key={index} value={item.trigger}>
                        <Card>
                            <CardHeader className='text-xl font-bold'>{item.trigger}</CardHeader>
                            <CardContent className="space-y-5 text-xl">
                                {item.fields.map((item, index) => (
                                    <div key={index} className="space-y-1">
                                        <Label htmlFor={item.name}>{item.name}</Label>
                                        <Input id={item.name} defaultValue={item.label} />
                                    </div>
                                ))}
                                <Button className='flex w-full items-center justify-center rounded-full bg-orange-600 hover:bg-orange-700'>{item.trigger}</Button>
                                <p className='flex items-center justify-center font-bold text-lg py-2'>or</p>
                                <div className='flex flex-col w-full justify-center items-center text-center gap-5'>
                                    {item.buttons.map((button, key) => (
                                        <Button key={key} variant="outline" className={`${button.className} w-2/3 rounded-full py-5 text-sm md:text-lg`}>{button.icon}{button.name}</Button>
                                    ))}
                                </div>
                            </CardContent>
                        </Card>
                    </TabsContent>
                ))}
            </Tabs>
        </div>
    )
}

export default page