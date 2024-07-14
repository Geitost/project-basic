import { Html, Head, Main, NextScript } from "next/document";

export default function Document() {
  return (
    <Html lang="en">
      <Head>
        {/*eslint-disable-next-line @next/next/no-title-in-document-head*/}
        <title>AutoBlox - Roblox Automation Tools</title>
        <meta
          name="description"
          content="Welcome to AutoBlox, efficient Roblox Automation solutions that takes subtlety and safety to heart. No scripts or injectors needed here, just a smooth and worry-free experience. The best part? It's completely free to get started! Head over to autoblox.xyz"
        />
      </Head>
      <body>
        <Main />
        <NextScript />
      </body>
    </Html>
  );
}
