package com.coderconsole.tdd.checkout;

import com.coderconsole.tdd.checkout.PaymentMode;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.assertj.core.api.AssertionsForInterfaceTypes.assertThat;

class PaymentModeTest {


    PaymentMode paymentMode;

    @BeforeEach
    void setUp(){
        paymentMode = new PaymentMode();
    }

    @Test
    @DisplayName("given payment mode as paytm return true")
    void paytm(){
        boolean isValidPayment = paymentMode.triggerPayment("paytm");
        assertThat(isValidPayment).isTrue();
    }

    @Test
    @DisplayName("given payment mode as payu return true")
    void payu(){
        boolean isValidPayment = paymentMode.triggerPayment("payu");
        assertThat(isValidPayment).isTrue();
    }

    @Test
    @DisplayName("given payment mode as credit card return true")
    void credicard(){
        boolean isValidPayment = paymentMode.triggerPayment("cc");
        assertThat(isValidPayment).isTrue();
    }


}
