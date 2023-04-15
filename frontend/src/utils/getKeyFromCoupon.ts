import { SHA256 } from "crypto-js";

export function getKeyFromCoupon(coupon: string): string {
  const hash = SHA256(coupon).toString();
  const key = `0x${hash.slice(0, 64)}`;
  return key;
}
