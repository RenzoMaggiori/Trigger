"use client"
import { Card, CardContent } from '@/components/ui/card'
import React from 'react'
import { IoLogoGithub } from "react-icons/io";
import { Service, TriggerDraggable } from '../components/trigger-draggable';
import { Button } from '@/components/ui/button';

const Page = () => {
    const services: Service[] = [
        { icon: <IoLogoGithub className='w-5 h-5 mr-2' />, name: "Github", settings: [
            {type: "input", label: "text"}
        ] },
    ];
    const [droppedServices, setDroppedServices] = React.useState<Service[]>([]);
    const [settings, setSettings] = React.useState<Service["settings"]>([]);
    const [draggedIndex, setDraggedIndex] = React.useState<number | null>(null);

    const handleDragStart = (e: React.DragEvent<HTMLDivElement>, service: Service, index?: number): void => {
        const serviceData = { name: service.name, index };
        e.dataTransfer.setData("service", JSON.stringify(serviceData));
        if (index !== undefined) setDraggedIndex(index);
    };

    const handleDragOver = (e: React.DragEvent<HTMLDivElement>): void => {
        e.preventDefault();
    };

    const handleDrop = (e: React.DragEvent<HTMLDivElement>): void => {
        e.preventDefault();
        const droppedService = JSON.parse(e.dataTransfer.getData("service")) as { name: string; index?: number };
        if (droppedService.index !== undefined && draggedIndex !== null) {
            const reorderedItems = [...droppedServices];
            const [movedItem] = reorderedItems.splice(droppedService.index, 1);
            reorderedItems.splice(draggedIndex, 0, movedItem);
            setDroppedServices(reorderedItems);
        } else {
            const newService = services.find((service) => service.name === droppedService.name);
            if (newService && !droppedServices.includes(newService)) {
                setDroppedServices((prev) => [...prev, newService]);
            }
        }
        setDraggedIndex(null);
    };

    return (
        <div className='flex h-screen w-full overflow-hidden'>
            <div className='w-auto p-5'>
                <Card className='h-full'>
                    <p className='font-bold text-2xl p-3'>Services</p>
                    <CardContent className='flex flex-col items-center justify-start h-full py-5 gap-4'>
                        {services.map((item, key) => (
                            <div
                                key={key}
                                draggable
                                onDragStart={(e) => handleDragStart(e, item)}
                                className='cursor-move'
                            >
                                <TriggerDraggable service={item} />
                            </div>
                        ))}
                    </CardContent>
                </Card>
            </div>

            <div className='w-full p-5'>
                <Card className='w-full h-full'>
                    <CardContent
                        className='flex flex-row flex-wrap py-5 gap-x-4 h-full'
                        onDragOver={handleDragOver}
                        onDrop={handleDrop}
                    >
                        {droppedServices.length === 0 ? (
                            <p className='text-gray-500 w-full h-full'>Drop services here</p>
                        ) : (
                            droppedServices.map((item, index) => (
                                <div
                                    key={index}
                                    draggable
                                    onDragStart={(e) => handleDragStart(e, item, index)}
                                    className='cursor-move mb-2'
                                >
                                    <TriggerDraggable service={item} className="cursor-pointer" onClick={() => setSettings(item.settings)} />
                                </div>
                            ))
                        )}
                    </CardContent>
                </Card>
            </div>
            <div className='p-5'>
                {settings.length > 0 &&
                    <Card>
                        <CardContent>
                            text
                        </CardContent>
                    </Card>
                }
            </div>
        </div>
    );
};

export default Page;
