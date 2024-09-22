import WordFadeIn from "@/components/magicui/word-fade-in";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardTitle } from "@/components/ui/card";
import Link from "next/link";
import { FcGoogle } from "react-icons/fc";
import email_arrived from "/public/email_arrived.png"
import email_sent from "/public/email_sent.png"
import Image from "next/image";

export default function Home() {
  return (
    <div className="flex flex-1 w-full bg-gray-200 dark:bg-zinc-900 justify-center">
      <div className="flex flex-col items-center justify-start text-center text-black dark:text-white pt-20 gap-y-5">
        <div className="text-5xl font-bold mb-4">
          <WordFadeIn words="Connect and Automate Effortlessly" as="h1" />
        </div>
        <div className="text-xl font-bold max-w-2xl">
          <WordFadeIn as="p" words="Trigger empowers you to connect services seamlessly. Automate tasks and enhance productivity by turning your ideas into efficient workflows." />
        </div>
        <Button className="rounded-full border-black bg-white text-lg p-6 mt-5" variant="outline" asChild>
          <Link href="http://localhost:8000/api/auth/gmail/provider">
            <FcGoogle className="mr-2" /> Sign in with Google
          </Link>
        </Button>
        <div className="flex flex-row items-center mt-10 gap-5">
          <Card className="max-w-[400px]">
            <CardTitle>
              <Image
                src={email_arrived}
                alt="email_arrived"
                className="rounded-t-md"
              />
            </CardTitle>
            <CardContent className="text-xl font-bold py-5">
              Email Received: Trigger has successfully captured this email and is ready to process it as part of your automated workflow.
            </CardContent>
          </Card>
          <Card className="max-w-[400px]">
            <CardTitle>
              <Image
                src={email_sent}
                alt="email_sent"
                className="rounded-t-md"
              />
            </CardTitle>
            <CardContent className="text-xl font-bold py-5">
              New Email Sent: Trigger has successfully dispatched your email, seamlessly integrating it into your automated workflow.
            </CardContent>
          </Card>
        </div>
      </div>

    </div>
  );
}
