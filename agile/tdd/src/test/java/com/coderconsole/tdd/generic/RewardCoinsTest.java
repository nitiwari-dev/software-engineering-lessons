package com.coderconsole.tdd.generic;


import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.TestInstance;

import java.util.HashMap;
import java.util.Map;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.assertThrows;

@TestInstance(TestInstance.Lifecycle.PER_CLASS)
class RewardCoinsTest {

    private RewardCoins rewardsCoins;

    @BeforeAll
    void setUp() {
        Map<String, Long> couponCoins = new HashMap<>();
        couponCoins.put("IND2024", 10L);
        couponCoins.put("IND4040", 40L);
        couponCoins.put("IN9090", 50L);
        rewardsCoins = new RewardCoins(couponCoins);
    }

    @Test
    @DisplayName("given valid coupon code return the reward coins for the customer")
    void rewardCoinsBasedOnCouponCode() throws RewardCoins.InValidCouponCode {
        String couponCode = "IND2024";
        Long rewardCoins = rewardsCoins.rewardCoins(couponCode);
        assertThat(rewardCoins).isEqualTo(10L);
    }

    @Test
    @DisplayName("given invalid coupon code throw InValidCouponCode")
    void throwExceptionForInvalidCouponCode() {
        String couponCode = "2023";

        assertThrows(RewardCoins.InValidCouponCode.class, () -> {
            rewardsCoins.rewardCoins(couponCode);
        });
    }


}
