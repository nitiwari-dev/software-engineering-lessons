package com.coderconsole.tdd;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static org.assertj.core.api.AssertionsForInterfaceTypes.assertThat;

public class PaymentMode {

    public PaymentMode paymentMode;

    @BeforeEach
    void setup() {
        paymentMode = new PaymentMode();
    }

    @Test
    void givenPaymentModeTriggerPayment() {
        boolean isPaymentSuccessful = paymentMode.triggerPayment("paytm");
        assertThat(isPaymentSuccessful).isTrue();
    }
}
