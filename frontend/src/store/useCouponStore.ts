import { create } from "zustand";
import { devtools, persist } from "zustand/middleware";

interface CouponState {
  coupons: Record<string, string[]>;
  addCoupons: (id: string, codes: string[]) => void;
}

const useCouponStore = create<CouponState>()(
  devtools(
    persist(
      (set) => ({
        coupons: {},

        addCoupons: (title: string, codes: string[]) =>
          set((state) => ({
            ...state,
            coupons: {
              ...state.coupons,
              [title]: codes,
            },
          })),
      }),
      {
        name: "coupon-storage",
      }
    )
  )
);

export default useCouponStore;
