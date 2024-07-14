/* eslint-disable react/no-unescaped-entities */
import { useState } from "react";
import cx from "classnames";
import { m } from "framer-motion";

const faqs = [
  {
    title: "What are the Supported Games for AutoBlox?",
    content: `At the moment, AutoBlox exclusively supports Bloxburg due to its recent
        launch. However, we have exciting plans to introduce additional
        automations for various games in the future. If you have any specific
        games in mind that you'd like to see supported, please let us know on
        our discord server.`,
  },
  {
    title: "How Does Autoblox Work?",
    content: `Autoblox utilizes a combination of machine learning and user input to
    efficiently automate tasks. By mimicking mouse movements and clicks, it
    seamlessly automates gameplay with an impressive 99.9% undetectable
    success rate. Crucially, it operates without injecting any scripts into
    the game, ensuring a secure and reliable process.`,
  },
  {
    title: "Is Autoblox Safe To Use?",
    content: `While we have taken precautions to ensure safety, it's essential to
    understand that using third-party tools like autoblox may violate
    Roblox's Terms of Service. Users should use caution and consider the
    potential risks.`,
  },
];

const variants = {
  visible: { opacity: 1, x: 0 },
  hidden: { opacity: 0, x: -25 },
};

function Faq() {
  const [openIdx, setOpenIdx] = useState(-1);

  return (
    <div className="text-center px-4 mx-auto w-full max-w-5xl">
      <h3 className="text-3xl md:text-4xl font-semibold">
        Frequently Asked Questions
      </h3>
      <p className="text-gray-300 mt-2 mx-auto max-w-xl text-sm md:text-base">
        Here are some frequently asked questions; if your question isn't
        answered here, don't be afraid to ask someone in our{" "}
        <a
          href={`${process.env.NEXT_PUBLIC_API_URL}/discord`}
          target="_blank"
          className="font-medium border-b border-b-gray-300 hover:text-gray-100 hover:border-b-gray-100 transition">
          discord
        </a>
        .
      </p>
      <div className="flex flex-col mt-12 text-left">
        {faqs.map((faq, idx) => (
          <m.div
            initial="hidden"
            whileInView="visible"
            viewport={{ once: true, margin: "-25px" }}
            transition={{ delay: (100 * (idx + 1)) / 1000 }}
            variants={variants}
            key={`faq-${idx}`}
            className="hover:bg-zinc-800 border-b border-b-zinc-800 cursor-pointer"
            onClick={() => {
              setOpenIdx(openIdx === idx ? -1 : idx);
            }}>
            <div className="accordion-header flex space-x-5 px-5 items-center h-16">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 20 20"
                fill="currentColor"
                className={cx(
                  "w-5 h-5 duration-300 transition-transform transform",
                  openIdx === idx ? "rotate-90" : "rotate-0"
                )}>
                <path
                  fillRule="evenodd"
                  d="M7.21 14.77a.75.75 0 01.02-1.06L11.168 10 7.23 6.29a.75.75 0 111.04-1.08l4.5 4.25a.75.75 0 010 1.08l-4.5 4.25a.75.75 0 01-1.06-.02z"
                  clipRule="evenodd"
                />
              </svg>
              <h3>{faq.title}</h3>
            </div>
            <div
              className={cx(
                "px-5 pt-0 overflow-hidden ease-in-out duration-500 transition-all",
                openIdx === idx ? "max-h-96 opacity-100" : "max-h-0 opacity-0"
              )}>
              <p className="leading-6 text-start pb-4">{faq.content}</p>
            </div>
          </m.div>
        ))}
      </div>
    </div>
  );
}

export default Faq;
