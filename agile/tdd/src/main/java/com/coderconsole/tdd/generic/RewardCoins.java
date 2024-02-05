package com.coderconsole.tdd.generic;

import java.util.Map;

public class RewardCoins {

    private final Map<String, Long> couponCoins;
    RewardCoins(Map<String, Long> couponCoins){
        this.couponCoins = couponCoins;
    }

    public Long rewardCoins(String couponCode) throws InValidCouponCode {
        validateCouponCode(couponCode);

        if (couponCoins.containsKey(couponCode))
            return couponCoins.get(couponCode);

        throw new InValidCouponCode("Please enter valid coupon code");
    }

    private void validateCouponCode(String couponCode) throws InValidCouponCode {
        if (couponCode == null)
            throw new InValidCouponCode("Please enter valid coupon code");
    }


    static class InValidCouponCode extends Exception {
        public InValidCouponCode(String message) {
            super(message);
        }
    }

}
