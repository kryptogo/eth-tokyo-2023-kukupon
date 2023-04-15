import { useQuery } from "@tanstack/react-query";

interface Campaign {
  id: string;
  required_condition: string[];
  image: string;
  sponsor: string;
}

const getCampagins = async (): Promise<Campaign[]> =>
  fetch(`${process.env.NEXT_PUBLIC_API_URL}/campaigns`).then((res) =>
    res.json()
  );

export const useCampaigns = (options?: any) =>
  useQuery(["campaigns"], getCampagins, options);
