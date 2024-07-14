import Image from "next/image";
import { Kanit } from "next/font/google";
import Navbar from "./components/Navbar";
import Hero from "./components/Hero";
import StatsSection from "./components/StatsSection";
import BenefitsSection from "./components/BenefitsSection";
import Faq from "./components/Faq";
import CallToAction from "./components/CallToAction";
import Footer from "./components/Footer";

const kanit = Kanit({
  weight: ["100", "200", "300", "400", "500", "600", "700", "800", "900"],
  subsets: ["latin"],
});

export default function Home() {
  return (
    <div className={`${kanit.className}`}>
      <Navbar />
      <Hero />
      <StatsSection />
      <BenefitsSection />
      <section className="py-16 relative">
        <Faq />
        <CallToAction />
      </section>
      <Footer />
    </div>
  );
}
