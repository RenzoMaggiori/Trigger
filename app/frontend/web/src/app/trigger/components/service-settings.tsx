import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';
import React from 'react'


const OutlookSettings = () => {
    const inputs = [
        { label: "Destination", placeholder: "example@example.com" },
        { label: "Title", placeholder: "Example title..." },
        { label: "Subject", placeholder: "Example subject..." },
    ]
    return (
        <div className='flex flex-col gap-y-4'>
            {inputs.map((item, key) => (
                <div key={key}>
                    <Label>{item.label}</Label>
                    <Input placeholder={item.placeholder} />
                </div>

            ))}
            <div>
                <Label>Email body</Label>
                <Textarea placeholder='Example body...' className='resize-none h-[200px]' />
            </div>
        </div>
    );
}

export { OutlookSettings }