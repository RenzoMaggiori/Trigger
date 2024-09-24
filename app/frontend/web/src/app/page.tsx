import WordFadeIn from "@/components/magicui/word-fade-in";
import { Button } from "@/components/ui/button";
import Image from "next/image";
import Link from "next/link";
import { FcGoogle } from "react-icons/fc";

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
        <div className="max-w-md mx-auto flex flex-col md:flex-row gap-x-7">
          <Button className="w-full rounded-full text-lg py-6 px-12 mt-5 bg-orange-600 hover:bg-orange-700 text-white hover:text-white" variant="outline">
              Start with Email
          </Button>
          <Button className="w-full rounded-full border-black bg-white text-lg p-6 mt-5" variant="outline" asChild>
            <Link href="http://localhost:8000/api/auth/gmail/provider" className="flex items-center">
              <FcGoogle className="mr-2 text-2xl" /> Start with Google
            </Link>
          </Button>
        </div>

      </div>

    </div>
  );
}
