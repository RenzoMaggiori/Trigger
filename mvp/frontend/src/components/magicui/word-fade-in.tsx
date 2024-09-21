"use client";

import { motion, Variants } from "framer-motion";
import { cn } from "@/lib/utils";

const defaultVariants: Variants = {
  hidden: { opacity: 0, y: 20 },
  visible: (i: number) => ({
    opacity: 1,
    y: 0,
    transition: { delay: i * 0.15 },
  }),
};

interface WordFadeInProps {
  words: string;
  className?: string;
  delay?: number;
  variants?: Variants;
  as?: 'h1' | 'p';
}

export default function WordFadeIn({
  words,
  delay = 0.15,
  className,
  variants = defaultVariants,
  as: Element = 'p',
}: WordFadeInProps) {
  const _words = words.split(" ");

  const MotionElement = Element === 'h1' ? motion.h1 : motion.p;

  return (
    <MotionElement
      variants={variants}
      initial="hidden"
      animate="visible"
      className={cn(
        "font-display text-center tracking-[-0.02em] text-black drop-shadow-sm dark:text-white",
        className,
      )}
    >
      {_words.map((word, i) => (
        <motion.span key={word} variants={variants} custom={i}>
          {word}{" "}
        </motion.span>
      ))}
    </MotionElement>
  );
}
