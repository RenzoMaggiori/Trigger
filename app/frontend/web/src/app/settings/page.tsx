"use client"
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { ScrollArea } from '@/components/ui/scroll-area';
import { Switch } from '@/components/ui/switch';
import React from 'react'
import { FcGoogle } from "react-icons/fc";

import { FaDiscord } from "react-icons/fa";
import { IoLogoGithub } from "react-icons/io";
import { FaSlack } from "react-icons/fa6";
import { PiMicrosoftOutlookLogo } from "react-icons/pi";

import { FaCircle } from "react-icons/fa6";

const services = [
    {
        name: "Google",
        connected: false,
        connect: "/",
        disconect: "/",
        icon: <FcGoogle className='w-5 h-5' />,
        fields: [{ name: "Show on Profile", connected: true }, { name: "Connection", connected: false }]
    },
    {
        name: "Discord",
        connected: true,
        connect: "/", disconect: "/",
        icon: <FaDiscord className='w-5 h-5 text-blue-500' />,
        fields: [{ name: "Show on Profile", connected: true }, { name: "Connection", connected: false }]
    },
    {
        name: "Slack",
        connected: false,
        connect: "/",
        disconect: "/",
        icon: <FaSlack className='w-5 h-5' />,
        fields: [{ name: "Show on Profile", connected: true }, { name: "Connection", connected: false }]
    },
    {
        name: "Outlook",
        connected: false,
        connect: "/",
        disconect: "/",
        icon: <PiMicrosoftOutlookLogo className='w-5 h-5 text-black' />,
        fields: [{ name: "Show on Profile", connected: true }, { name: "Connection", connected: false }]
    },
    {
        name: "Github",
        connected: false,
        connect: "/",
        disconect: "/",
        icon: <IoLogoGithub className='w-5 h-5' />,
        fields: [{ name: "Show on Profile", connected: true }, { name: "Connection", connected: false }]
    },
];

const page = () => {
    const [serviceList, setServiceList] = React.useState(services);

    const handleSwitchChange = (serviceIndex: number, fieldIndex: number) => {
        const updatedServices = [...serviceList];
        const targetField = updatedServices[serviceIndex].fields[fieldIndex];

        targetField.connected = !targetField.connected;

        setServiceList(updatedServices);
    };
    return (
        <div className='flex flex-col w-full h-full items-center justify-center gap-5'>
            <ScrollArea className='w-full md:w-2/3 lg:w-1/2 max-h-[80vh] items-center justify-center border p-5 rounded-md'>
                <div className='flex flex-col items-center justify-center gap-5 w-full'>
                    {services.map((item, key) => (
                        <Card className='flex flex-col w-full h-auto' key={key}>
                            <CardHeader className={`rounded-t-md border-b`}>
                                <CardTitle className='flex items-center justify-between text-xl text-start font-bold'>
                                    <div className='flex items-center gap-x-2'>
                                        {item.icon}
                                        {item.name}
                                    </div>
                                    <div className={`flex items-center ${item.connected ? 'text-green-500' : 'text-red-500'}`}>
                                        <div className='hidden md:block'>
                                            {item.connected ? 'Connected' : 'Disconnected'}
                                        </div>
                                        <FaCircle className={`ml-2 ${item.connected ? 'text-green-500' : 'text-red-500'}`} />
                                    </div>
                                </CardTitle>
                            </CardHeader>
                            <CardContent className={`flex w-full h-full rounded-b-md items-end justify-center`}>
                                <div className='w-full flex flex-col text-lg items-start justify-start gap-y-3 mt-5'>
                                    {item.fields?.map((field, index) => (
                                        <div key={index} className='flex flex-row w-full items-center justify-between text-black font-bold'>
                                            {field.name}
                                            <Switch
                                                className='data-[state=checked]:bg-green-400 data-[state=unchecked]:bg-red-500'
                                                checked={field.connected}
                                                onClick={() => handleSwitchChange(key, index)}
                                            />
                                        </div>
                                    ))}
                                </div>
                            </CardContent>
                        </Card>
                    ))}
                </div>
            </ScrollArea>
        </div>
    )
}

export default page