package com.coderconsole.tdd.generic;

import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.TestInstance;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.HashMap;
import java.util.Map;
import java.util.stream.Stream;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.assertThrows;

@TestInstance(TestInstance.Lifecycle.PER_CLASS)
class RewardCoinsParameterisedTest {

    private RewardCoins rewardsCoins;

    @BeforeAll
    void setUp() {
        Map<String, Long> couponCoins = new HashMap<>();
        couponCoins.put("IND2024", 10L);
        couponCoins.put("IND4040", 40L);
        couponCoins.put("IN9090", 50L);
        rewardsCoins = new RewardCoins(couponCoins);
    }

    @ParameterizedTest
    @MethodSource("validCouponCoins")
    @DisplayName("given valid coupon code return the reward coins for the customer")
    void rewardCoinsBasedOnCouponCode(String couponCode, Long expectedCoins) throws RewardCoins.InValidCouponCode {
        Long rewardCoins = rewardsCoins.rewardCoins(couponCode);
        assertThat(rewardCoins).isEqualTo(expectedCoins);
    }

    @ParameterizedTest
    @MethodSource("invalidCouponCoins")
    @DisplayName("given invalid coupon code throw InValidCouponCode")
    void throwExceptionForInvalidCouponCode(String couponCode) {
        assertThrows(RewardCoins.InValidCouponCode.class, () ->
                rewardsCoins.rewardCoins(couponCode));
    }

    private static Stream<Arguments> validCouponCoins() {
        return Stream.of(Arguments.of("IND2024", 10L),
                Arguments.of("IN9090", 50L));
    }

    private static Stream<Arguments> invalidCouponCoins() {
        return Stream.of(Arguments.of("2023"),
                Arguments.of("200"),
                Arguments.of("100"));
    }
}