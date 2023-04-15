import "@/styles/globals.css";
import {
  getDefaultWallets,
  KryptogoKitProvider,
  connectorsForWallets,
  wallet,
} from "@kryptogo/kryptogokit";
import "@kryptogo/kryptogokit/styles.css";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import type { AppProps } from "next/app";
import { configureChains, createClient, WagmiConfig } from "wagmi";
import { polygonMumbai } from "wagmi/chains";
import { publicProvider } from "wagmi/providers/public";

const { chains, provider } = configureChains(
  [polygonMumbai],
  [publicProvider()]
);
const appName = "kukupon";
const { wallets } = getDefaultWallets({
  chains,
});

const connectors = connectorsForWallets([
  ...wallets,
  {
    groupName: "Other",
    wallets: [
      wallet.rainbow({ chains }),
      wallet.trust({ chains }),
      wallet.coinbase({ appName, chains }),
      wallet.argent({ chains }),
      wallet.ledger({ chains }),
    ],
  },
]);

const wagmiClient = createClient({
  autoConnect: true,
  connectors,
  provider,
});

const queryClient = new QueryClient();

export default function App({ Component, pageProps }: AppProps) {
  return (
    <QueryClientProvider client={queryClient}>
      <WagmiConfig client={wagmiClient}>
        <KryptogoKitProvider chains={chains}>
          <Component {...pageProps} />
        </KryptogoKitProvider>
      </WagmiConfig>
      {/* <ReactQueryDevtools initialIsOpen={false} /> */}
    </QueryClientProvider>
  );
}
