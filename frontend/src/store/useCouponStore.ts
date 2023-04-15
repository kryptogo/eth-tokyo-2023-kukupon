import { create } from "zustand";
import { devtools, persist } from "zustand/middleware";

interface CouponState {
  coupons: Map<string, string[]>;
  addCoupons: (id: string, codes: string[]) => void;
}

const useCouponStore = create<CouponState>()(
  devtools(
    persist(
      (set) => ({
        coupons: new Map(),

        addCoupons: (id: string, codes: string[]) =>
          set((state) => ({
            ...state,
            coupons: new Map(state.coupons).set(id, codes),
          })),
      }),
      {
        name: "coupon-storage",
      }
    )
  )
);

export default useCouponStore;
