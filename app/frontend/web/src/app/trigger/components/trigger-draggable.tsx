import { Card, CardContent } from '@/components/ui/card'
import React from 'react'
import { Service } from '@/app/trigger/lib/types';

interface TriggerDraggableProps extends React.HTMLAttributes<HTMLDivElement> {
    service: Service;
}

export function TriggerDraggable({ service, className, ...props }: TriggerDraggableProps) {
    return (
        <Card className={className} {...props}>
            <CardContent className='flex flex-row items-center justify-center text-lg font-bold p-4'>
                {service.icon}
                <p>{service.name}</p>
            </CardContent>
        </Card>
    )
}
