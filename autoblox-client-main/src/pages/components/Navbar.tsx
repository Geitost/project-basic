import Image from "next/image";
import { Divide as Hamburger } from "hamburger-react";
import { useState } from "react";
import cn from "classnames";

import { m } from "framer-motion";

type NavLinkProps = {
  children: string;
  href?: string;
  delay?: number;
};

function NavLink({ children, href, delay }: NavLinkProps) {
  return (
    <m.a
      initial={{ opacity: 0, y: -10 }}
      animate={{ opacity: 1, y: 0 }}
      whileHover={{ color: "#d1d5db", scale: 1.1 }}
      transition={{ delay }}
      href={href || "#"}
      target="_blank"
      className="inline-block">
      {children}
    </m.a>
  );
}

function Navbar() {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <nav className="flex fixed w-full h-20 px-8 justify-between items-center z-50 backdrop-blur bg-zinc-950/20">
      <div className="text-primary flex gap-4 items-center">
        <Image src="/logo.svg" alt="AutoBlox logo" width={36} height={36} />
        <h1 className="text-2xl font-semibold">AutoBlox</h1>
      </div>
      <div>
        <ul className="gap-6 hidden md:flex">
          <li>
            <NavLink href={`${process.env.NEXT_PUBLIC_API_URL}/download`}>
              Download
            </NavLink>
          </li>
          <li>
            <NavLink
              href={`${process.env.NEXT_PUBLIC_API_URL}/key`}
              delay={50 / 1000}>
              Get Key
            </NavLink>
          </li>
          <li>
            <NavLink
              href={`${process.env.NEXT_PUBLIC_API_URL}/discord`}
              delay={100 / 1000}>
              Join Discord
            </NavLink>
          </li>
          <li>
            <m.a
              initial={{ opacity: 0, y: -10 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ delay: 150 / 1000 }}
              href={`${process.env.NEXT_PUBLIC_API_URL}/pro`}
              target="_blank"
              className="px-8 py-2 font-medium bg-primary rounded-lg hover:bg-primary-400 transition-colors">
              Go Pro
            </m.a>
          </li>
        </ul>
        <div className="block md:hidden">
          <Hamburger size={28} toggled={isOpen} toggle={setIsOpen} />

          {/* Mobile Menu */}
          <m.div
            initial={{ height: 0 }}
            animate={{ height: isOpen ? "auto" : 0 }}
            className="absolute overflow-clip left-0 top-20 w-full backdrop-blur-sm bg-zinc-950/20">
            <ul className="gap-4 p-6 justify-stretch flex flex-col">
              <li>
                <NavLink href={`${process.env.NEXT_PUBLIC_API_URL}/download`}>
                  Download
                </NavLink>
              </li>
              <li>
                <NavLink href={`${process.env.NEXT_PUBLIC_API_URL}/pro`}>
                  Get Key
                </NavLink>
              </li>
              <li>
                <NavLink href={`${process.env.NEXT_PUBLIC_API_URL}/discord`}>
                  Join Discord
                </NavLink>
              </li>
              <li className="mt-4">
                <a
                  target="_blank"
                  href={`${process.env.NEXT_PUBLIC_API_URL}/pro`}>
                  <div className="max-w-md mx-auto text-center py-2 font-medium bg-primary rounded-lg hover:bg-primary-400 transition-colors">
                    Go Pro
                  </div>
                </a>
              </li>
            </ul>
          </m.div>
        </div>
      </div>
    </nav>
  );
}

export default Navbar;
