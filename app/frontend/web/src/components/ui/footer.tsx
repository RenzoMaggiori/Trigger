import React from 'react'
import { LogoIcon } from './logoIcon'
import { FaFacebookF } from "react-icons/fa";
import { FaLinkedinIn } from "react-icons/fa";
import { FiGithub } from "react-icons/fi";
import { Button } from './button';
import Link from 'next/link';

export const Footer = () => {
    const footer = [
        { title: "Contact Information", fields: ["Email: trigger@gmail.com", "Phone: +123 456 7890", "Address: 123 Main St, City, Country"] },
        { title: "Our Team", fields: ["John Doe - CEO", "Jane Smith - CTO", "Michael Brown - COO"] },
        { title: "Available Services", fields: ["Gmail", "Github", "Outlook", "Discord", "Slack"] },
    ];
    const footerButtons = [
        {icon: <FaFacebookF />, hover: "hover:bg-sky-700 hover:text-zinc-300", href: "https://www.facebook.com/"},
        {icon: <FaLinkedinIn />, hover: "hover:bg-sky-600 hover:text-zinc-300", href: "https://www.linkedin.com/"},
        {icon: <FiGithub />, hover: "hover:bg-zinc-950 hover:text-zinc-300", href: "https://www.github.com/"},
    ];
    return (
        <div className="flex flex-col w-full bg-zinc-300 dark:bg-zinc-950 text-zinc-800 dark:text-zinc-400 py-10">
            <div className="flex flex-col md:flex-row gap-10 justify-between w-full px-12 md:px-32">
                {footer.map((item, index) => (
                    <div key={index} className="text-start">
                        <h3 className="text-lg font-semibold mb-4">{item.title}</h3>
                        {item.fields.map((field, fieldIndex) => (
                            <p key={fieldIndex} className="text-sm mb-2 hover:underline cursor-pointer">{field}</p>
                        ))}
                    </div>
                ))}
            </div>

            <div className="flex flex-col md:flex-row justify-between items-center w-full mt-10 px-12 md:px-32 border-t border-gray-400 pt-6">
                <div className="flex items-center">
                    <LogoIcon className="h-8 w-auto dark:fill-white" />
                </div>

                <div className="mt-4 md:mt-0 text-center">
                    <span>© 2024 Trigger. All rights reserved.</span>
                </div>

                <div className="flex space-x-4 mt-4 md:mt-0">
                    {footerButtons.map((item, index) => (
                        <Button key={index} variant="ghost" className={`${item.hover} rounded-full bg-zinc-500 text-zinc-300`} asChild>
                            <Link href={`${item.href}`}>
                                {item.icon}
                            </Link>
                        </Button>
                    ))}
                </div>
            </div>
        </div>
    )
}
