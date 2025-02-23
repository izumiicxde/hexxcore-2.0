"use client";
import Link from "next/link";
import React, { useState } from "react";
import { ThemeToggle } from "./theme-toggler";
import { Menu, X } from "lucide-react";

export const Navbar = () => {
  const [isOpen, setIsOpen] = useState(false);

  const links = [
    { name: "dashboard", link: "/dashboard" },
    { name: "today", link: "/dashboard/today" },
    { name: "mark attendance", link: "/dashboard/subjects" },
  ];

  return (
    <div className="p-3 flex items-center justify-between w-full bg-white dark:bg-gray-900 shadow-md">
      {/* Logo */}
      <div className="text-xl font-bold uppercase">Zen0</div>

      {/* Desktop Navigation */}
      <nav className="hidden md:flex gap-6">
        {links.map((link) => (
          <Link
            key={link.name}
            href={link.link}
            className="capitalize hover:underline transition"
          >
            {link.name}
          </Link>
        ))}
      </nav>

      {/* Theme Toggle */}
      <div className="hidden md:block">
        <ThemeToggle />
      </div>

      {/* Mobile Menu Button */}
      <button className="md:hidden" onClick={() => setIsOpen(!isOpen)}>
        {isOpen ? <X size={24} /> : <Menu size={24} />}
      </button>

      {/* Mobile Navigation */}
      {isOpen && (
        <div className="absolute top-16 right-4 w-48 bg-white dark:bg-gray-800 shadow-lg rounded-lg p-4 flex flex-col gap-3 md:hidden">
          {links.map((link) => (
            <Link
              key={link.name}
              href={link.link}
              className="capitalize hover:underline transition"
              onClick={() => setIsOpen(false)}
            >
              {link.name}
            </Link>
          ))}
          <ThemeToggle />
        </div>
      )}
    </div>
  );
};
